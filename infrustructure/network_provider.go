package infrustructure

import (
	"errors"
	"fmt"
	"github.com/everstake/teztracker/config"
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/services/rpc_client/client"
	"github.com/jinzhu/gorm"
	"github.com/rubenv/sql-migrate"
	"os"
	"path/filepath"
)

const defaultMigrationsDir = "repos/migrations"

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

		//Migration
		migrationsDir := defaultMigrationsDir
		if migrations := os.Getenv("POSTGRESQL_MIGRATIONS_PATH"); migrations != "" {
			migrationsDir = migrations
		}

		err = Migrate(db, migrationsDir)
		if err != nil {
			return nil, err
		}

		provider.networks[k] = NetworkContext{
			Db:           db,
			ClientConfig: v.NodeRpc,
		}
	}
	return provider, nil
}

func Migrate(db *gorm.DB, migrationsDir string) (err error) {
	ex, err := os.Executable()
	if err != nil {
		return err
	}

	dir := filepath.Join(filepath.Dir(ex), migrationsDir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		dir = migrationsDir
		if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
			return errors.New("Migrations dir does not exist: " + dir)
		}
	}

	migrations := &migrate.FileMigrationSource{
		Dir: dir,
	}

	migrate.SetTable("migrations")

	_, err = migrate.Exec(db.DB(), "postgres", migrations, migrate.Up)
	if err != nil {
		return err
	}

	return nil
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
