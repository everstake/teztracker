package main

import (
	"os"
	"strings"

	"github.com/bullblock-io/tezTracker/api"
	"github.com/bullblock-io/tezTracker/config"
	"github.com/bullblock-io/tezTracker/gen/restapi"
	"github.com/bullblock-io/tezTracker/gen/restapi/operations"
	"github.com/bullblock-io/tezTracker/infrustructure"
	"github.com/bullblock-io/tezTracker/models"
	"github.com/bullblock-io/tezTracker/services"
	"github.com/go-openapi/loads"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/roylee0704/gron"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	cfg := config.Parse()

	networks := make(map[models.Network]config.NetworkConfig)
	if cfg.Mainnet.SqlConnectionString != "" {
		networks[models.NetworkMain] = cfg.Mainnet
	}
	if cfg.Babylonnet.SqlConnectionString != "" {
		networks[models.NetworkBabylon] = cfg.Babylonnet
	}
	if len(networks) == 0 {
		log.Fatalln("no networks are configured")
	}

	provider, err := infrustructure.New(networks)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer provider.Close()
	// Enable log mode only on trace level. It's safe to set it to true always, but that'll be a little slower.
	if strings.EqualFold(cfg.LogLevel, log.TraceLevel.String()) {
		provider.EnableTraceLevel()
	}

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}
	apiServer := operations.NewTezTrackerAPI(swaggerSpec)

	// pass services instance to API handlers
	api.SetHandlers(apiServer, provider)

	server := restapi.NewServer(apiServer)
	server.ConfigureAPI()

	defer func() {
		if err := server.Shutdown(); err != nil {
			log.Fatalln(err)
		}
	}()
	cron := gron.New()
	for k := range networks {
		db, err := provider.GetDb(k)
		if err != nil {
			log.Fatalln(err)
		}
		rpc, err := provider.GetRpcConfig(k)
		if err != nil {
			log.Fatalln(err)
		}
		// Using models.NetworkMain instead of k due to stupid nodes configuration for babylonnet.
		// todo: if something is not workign for testnets, check this one.
		services.AddToCron(cron, cfg, db, rpc, models.NetworkMain, k == models.NetworkBabylon)
	}

	cron.Start()
	defer cron.Stop()

	server.Port = cfg.Port
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
