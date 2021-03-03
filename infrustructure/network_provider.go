package infrustructure

import (
	"fmt"
	"github.com/everstake/teztracker/config"
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/repos"
	"github.com/everstake/teztracker/services"
	"github.com/everstake/teztracker/services/daily_stats"
	"github.com/everstake/teztracker/services/mailer"
	"github.com/everstake/teztracker/services/mempool"
	"github.com/everstake/teztracker/services/rpc_client/client"
	"github.com/everstake/teztracker/services/watcher"
	"github.com/everstake/teztracker/ws"
	"github.com/jinzhu/gorm"
	"strings"
)

type NetworkContext struct {
	Db           *gorm.DB
	Mempool      *mempool.Mempool
	WS           *ws.Hub
	ClientConfig client.TransportConfig
}

type Provider struct {
	networks map[models.Network]NetworkContext
}

func New(configs map[models.Network]config.NetworkConfig, cfg config.Config) (*Provider, error) {
	provider := &Provider{
		networks: make(map[models.Network]NetworkContext),
	}
	var serviceProviders []services.Provider
	for k, v := range configs {
		db, err := gorm.Open("postgres", v.SqlConnectionString)
		if err != nil {
			return nil, err
		}
		db.SetLogger(&config.DbLogger{})
		gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
			if strings.Contains(defaultTableName, "views") {
				defaultTableName = defaultTableName[:len(defaultTableName)-1]
			}
			return "tezos." + defaultTableName
		}
		serviceProviders = append(serviceProviders, repos.New(db))

		hub := ws.NewHub()
		//Start hub
		go hub.Run()

		m, err := mempool.NewMempool(v, hub)
		if err != nil {
			return nil, err
		}

		go m.MonitorMempool()

		var mail mailer.Mail
		mail = mailer.NewFakeMailer()
		if models.NetworkMain == k && cfg.Production {
			mail = mailer.New(cfg.SmtpHost, cfg.SmtpPort, cfg.SmtpUser, cfg.SmtpPassword)
		}

		//TODO make graceful stop
		w := watcher.NewWatcher(v.SqlConnectionString, hub, repos.New(db), mail)

		//Start watcher
		go w.Start()

		provider.networks[k] = NetworkContext{
			Db:           db,
			Mempool:      m,
			WS:           hub,
			ClientConfig: v.NodeRpc,
		}
	}
	stats := daily_stats.NewDailyStats(serviceProviders)
	go stats.Run()
	return provider, nil
}

func (p *Provider) EnableTraceLevel() {
	for _, v := range p.networks {
		v.Db.LogMode(true)

	}
}
func (p *Provider) Close() {
	for _, v := range p.networks {
		v.Db.Close()
		v.Mempool.Cancel()
	}
}

func (p *Provider) GetDb(net models.Network) (*gorm.DB, error) {
	if netcont, ok := p.networks[net]; ok {
		return netcont.Db, nil
	}
	return nil, fmt.Errorf("not enabled network")
}

func (p *Provider) GetMempool(net models.Network) (*mempool.Mempool, error) {
	if netcont, ok := p.networks[net]; ok {
		return netcont.Mempool, nil
	}
	return nil, fmt.Errorf("not enabled network")
}

func (p *Provider) GetWS(net models.Network) (*ws.Hub, error) {
	if netcont, ok := p.networks[net]; ok {
		return netcont.WS, nil
	}
	return nil, fmt.Errorf("not enabled network")
}

func (p *Provider) GetRpcConfig(net models.Network) (cfg client.TransportConfig, err error) {
	if netcont, ok := p.networks[net]; ok {
		return netcont.ClientConfig, nil
	}
	return cfg, fmt.Errorf("not enabled network")
}
