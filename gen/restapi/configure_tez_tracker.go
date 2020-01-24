// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/rs/cors"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/everstake/teztracker/gen/restapi/operations"
	"github.com/everstake/teztracker/gen/restapi/operations/accounts"
	"github.com/everstake/teztracker/gen/restapi/operations/app_info"
	"github.com/everstake/teztracker/gen/restapi/operations/blocks"
	"github.com/everstake/teztracker/gen/restapi/operations/fees"
	"github.com/everstake/teztracker/gen/restapi/operations/operation_groups"
	"github.com/everstake/teztracker/gen/restapi/operations/operations_list"
)

//go:generate swagger generate server --target ../../gen --name TezTracker --spec ../../swagger/swagger.yml --exclude-main

func configureFlags(api *operations.TezTrackerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TezTrackerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.AccountsGetAccountHandler == nil {
		api.AccountsGetAccountHandler = accounts.GetAccountHandlerFunc(func(params accounts.GetAccountParams) middleware.Responder {
			return middleware.NotImplemented("operation accounts.GetAccount has not yet been implemented")
		})
	}
	if api.AccountsGetAccountsListHandler == nil {
		api.AccountsGetAccountsListHandler = accounts.GetAccountsListHandlerFunc(func(params accounts.GetAccountsListParams) middleware.Responder {
			return middleware.NotImplemented("operation accounts.GetAccountsList has not yet been implemented")
		})
	}
	if api.FeesGetAvgFeesHandler == nil {
		api.FeesGetAvgFeesHandler = fees.GetAvgFeesHandlerFunc(func(params fees.GetAvgFeesParams) middleware.Responder {
			return middleware.NotImplemented("operation fees.GetAvgFees has not yet been implemented")
		})
	}
	if api.BlocksGetBlockHandler == nil {
		api.BlocksGetBlockHandler = blocks.GetBlockHandlerFunc(func(params blocks.GetBlockParams) middleware.Responder {
			return middleware.NotImplemented("operation blocks.GetBlock has not yet been implemented")
		})
	}
	if api.BlocksGetBlocksHeadHandler == nil {
		api.BlocksGetBlocksHeadHandler = blocks.GetBlocksHeadHandlerFunc(func(params blocks.GetBlocksHeadParams) middleware.Responder {
			return middleware.NotImplemented("operation blocks.GetBlocksHead has not yet been implemented")
		})
	}
	if api.BlocksGetBlocksListHandler == nil {
		api.BlocksGetBlocksListHandler = blocks.GetBlocksListHandlerFunc(func(params blocks.GetBlocksListParams) middleware.Responder {
			return middleware.NotImplemented("operation blocks.GetBlocksList has not yet been implemented")
		})
	}
	if api.AppInfoGetInfoHandler == nil {
		api.AppInfoGetInfoHandler = app_info.GetInfoHandlerFunc(func(params app_info.GetInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation app_info.GetInfo has not yet been implemented")
		})
	}
	if api.OperationGroupsGetOperationGroupHandler == nil {
		api.OperationGroupsGetOperationGroupHandler = operation_groups.GetOperationGroupHandlerFunc(func(params operation_groups.GetOperationGroupParams) middleware.Responder {
			return middleware.NotImplemented("operation operation_groups.GetOperationGroup has not yet been implemented")
		})
	}
	if api.OperationGroupsGetOperationGroupsHandler == nil {
		api.OperationGroupsGetOperationGroupsHandler = operation_groups.GetOperationGroupsHandlerFunc(func(params operation_groups.GetOperationGroupsParams) middleware.Responder {
			return middleware.NotImplemented("operation operation_groups.GetOperationGroups has not yet been implemented")
		})
	}
	if api.OperationsListGetOperationsListHandler == nil {
		api.OperationsListGetOperationsListHandler = operations_list.GetOperationsListHandlerFunc(func(params operations_list.GetOperationsListParams) middleware.Responder {
			return middleware.NotImplemented("operation operations_list.GetOperationsList has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	corsHandler := cors.New(cors.Options{
		AllowedHeaders: []string{"*"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{},
		ExposedHeaders: []string{"X-Total-Count"},
		MaxAge:         60,
	})
	return corsHandler.Handler(handler)
}
