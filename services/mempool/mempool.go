package mempool

import (
	"context"
	"github.com/everstake/teztracker/config"
	gotez "github.com/goat-systems/go-tezos/v2"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"time"
)

const (
	monitoringUrl = "/chains/main/mempool/monitor_operations"
	mempoolKey    = "mempool"
	cacheTTL      = 1 * time.Minute
)

func NewMempool(cfg config.NetworkConfig) (*Mempool, error) {
	rpcURL := &url.URL{
		Scheme: cfg.NodeRpc.Schemes[0],
		Host:   cfg.NodeRpc.Host,
		Path:   cfg.NodeRpc.BasePath,
	}

	ctx, cancel := context.WithCancel(context.Background())

	return &Mempool{
		ctx:    ctx,
		Cancel: cancel,

		url:     rpcURL,
		cl:      http.DefaultClient,
		storage: cache.New(cacheTTL, cacheTTL),
	}, nil
}

type Mempool struct {
	url     *url.URL
	cl      *http.Client
	storage *cache.Cache
	ctx     context.Context
	Cancel  context.CancelFunc
}

func (m Mempool) MonitorMempool() {
	log.Info("Start mempool monitor")
	var err error
	ch := make(chan []gotez.Operations)
	for {
		select {
		case <-m.ctx.Done():
			log.Info("MonitorMempool done")
			return
		default:
			listenCtx, cancel := context.WithCancel(m.ctx)
			go m.listenResp(listenCtx, ch)
			err = m.monitorMempoolOperations(m.ctx, "", ch)
			if err != nil {
				log.Errorf("MonitorMempool error: %s", err)
			}

			//Close listen routine
			cancel()
			//Block closed, flush storage
			m.storage.Flush()
		}

	}
}

func (m Mempool) GetMempool() (op []gotez.Operations, err error) {
	stored, ok := m.storage.Get(mempoolKey)
	if !ok {
		return nil, nil
	}

	return stored.([]gotez.Operations), nil
}

func (m Mempool) monitorMempoolOperations(ctx context.Context, filter string, results chan []gotez.Operations) error {
	if filter == "" {
		filter = "applied=true"
	}

	return m.Do(m.ctx, http.MethodGet, monitoringUrl, filter, results)
}

func (m Mempool) listenResp(ctx context.Context, ch chan []gotez.Operations) {
	var stored interface{}
	for {
		select {
		case <-ctx.Done():
			//Close routine
			return
		case elem, ok := <-ch:
			//Chanel closed by request
			if !ok {
				return
			}

			stored, ok = m.storage.Get(mempoolKey)
			if !ok {
				m.storage.SetDefault(mempoolKey, elem)
				continue
			}

			m.storage.SetDefault(mempoolKey, append(stored.([]gotez.Operations), elem...))
		}
	}
}
