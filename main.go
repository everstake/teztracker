package main

import (
	"flag"
	"github.com/everstake/teztracker/services/cmc"
	"os"
	"strings"

	"github.com/everstake/teztracker/api"
	"github.com/everstake/teztracker/config"
	"github.com/everstake/teztracker/gen/restapi"
	"github.com/everstake/teztracker/gen/restapi/operations"
	"github.com/everstake/teztracker/infrustructure"
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/services"
	"github.com/go-openapi/loads"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/roylee0704/gron"
	log "github.com/sirupsen/logrus"
)

var cronDisableFlag = flag.Bool("crondisable", false, "disable cron for api tests")

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	flag.Parse()

	cfg := config.Parse()

	networks := make(map[models.Network]config.NetworkConfig)
	if cfg.Mainnet.SqlConnectionString != "" {
		networks[models.NetworkMain] = cfg.Mainnet
	}
	if cfg.Carthagenet.SqlConnectionString != "" {
		networks[models.NetworkCarthage] = cfg.Carthagenet
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

	marketData := cmc.NewCoinGecko()
	// pass services instance to API handlers
	api.SetHandlers(apiServer, provider, marketData)

	server := restapi.NewServer(apiServer)
	server.ConfigureAPI()

	defer func() {
		if err := server.Shutdown(); err != nil {
			log.Fatalln(err)
		}
	}()

	if !*cronDisableFlag {
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

			ws, err := provider.GetWS(k)
			if err != nil {
				log.Fatalln(err)
			}

			// Using models.NetworkMain instead of k due to stupid nodes configuration for carthagenet.
			// todo: if something is not workign for testnets, check this one.
			services.AddToCron(cron, cfg, db, ws, marketData, rpc, models.NetworkMain, k == models.NetworkCarthage)
		}

		cron.Start()
		defer cron.Stop()
	}

	server.Port = cfg.Port
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
