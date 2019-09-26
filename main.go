package main

import (
	"os"
	"strings"

	"github.com/bullblock-io/tezTracker/api"
	"github.com/bullblock-io/tezTracker/config"
	"github.com/bullblock-io/tezTracker/gen/restapi"
	"github.com/bullblock-io/tezTracker/gen/restapi/operations"
	"github.com/go-openapi/loads"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	cfg := config.Parse()

	db, err := gorm.Open("postgres", cfg.SqlConnectionString)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer db.Close()
	db.SetLogger(&config.DbLogger{})

	// Enable log mode only on trace level. It's safe to set it to true always, but that'll be a little slower.
	db.LogMode(strings.EqualFold(cfg.LogLevel, log.TraceLevel.String()))

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}
	apiServer := operations.NewTezTrackerAPI(swaggerSpec)

	// pass services instance to API handlers
	api.SetHandlers(apiServer, db)

	server := restapi.NewServer(apiServer)
	server.ConfigureAPI()

	defer func() {
		if err := server.Shutdown(); err != nil {
			log.Fatalln(err)
		}
	}()

	server.Port = cfg.Port
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
