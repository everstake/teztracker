package infrustructure

import (
	"fmt"
	"github.com/bullblock-io/tezTracker/config"
	"github.com/bullblock-io/tezTracker/models"
	"github.com/bullblock-io/tezTracker/services/rpc_client/client"
	"github.com/jinzhu/gorm"
)

type NetworkContext struct {
	Db           *gorm.DB
	ClientConfig client.TransportConfig
}

type Provider struct {
	networks map[models.Network]NetworkContext
}

func New(configs map[models.Network]config.NetworkConfig) (*Provider, error) {
	provider := &Provider{
		networks: make(map[models.Network]NetworkContext),
	}
	for k, v := range configs {
		db, err := gorm.Open("postgres", v.SqlConnectionString)
		if err != nil {
			return nil, err
		}
		db.SetLogger(&config.DbLogger{})
		gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
			return "tezos." + defaultTableName
		}
		provider.networks[k] = NetworkContext{
			Db:           db,
			ClientConfig: v.NodeRpc,
		}
	}
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
	}
}
func (p *Provider) GetDb(net models.Network) (*gorm.DB, error) {
	if netcont, ok := p.networks[net]; ok {
		return netcont.Db, nil
	}
	return nil, fmt.Errorf("not enabled network")
}

func (p *Provider) GetRpcConfig(net models.Network) (cfg client.TransportConfig, err error) {
	if netcont, ok := p.networks[net]; ok {
		return netcont.ClientConfig, nil
	}
	return cfg, fmt.Errorf("not enabled network")
}
