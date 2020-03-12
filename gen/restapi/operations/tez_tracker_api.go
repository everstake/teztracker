// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/everstake/teztracker/gen/restapi/operations/accounts"
	"github.com/everstake/teztracker/gen/restapi/operations/app_info"
	"github.com/everstake/teztracker/gen/restapi/operations/blocks"
	"github.com/everstake/teztracker/gen/restapi/operations/fees"
	"github.com/everstake/teztracker/gen/restapi/operations/operation_groups"
	"github.com/everstake/teztracker/gen/restapi/operations/operations_list"
	"github.com/everstake/teztracker/gen/restapi/operations/voting"
)

// NewTezTrackerAPI creates a new TezTracker instance
func NewTezTrackerAPI(spec *loads.Document) *TezTrackerAPI {
	return &TezTrackerAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		AccountsGetAccountHandler: accounts.GetAccountHandlerFunc(func(params accounts.GetAccountParams) middleware.Responder {
			return middleware.NotImplemented("operation AccountsGetAccount has not yet been implemented")
		}),
		AccountsGetAccountDelegatorsHandler: accounts.GetAccountDelegatorsHandlerFunc(func(params accounts.GetAccountDelegatorsParams) middleware.Responder {
			return middleware.NotImplemented("operation AccountsGetAccountDelegators has not yet been implemented")
		}),
		AccountsGetAccountsListHandler: accounts.GetAccountsListHandlerFunc(func(params accounts.GetAccountsListParams) middleware.Responder {
			return middleware.NotImplemented("operation AccountsGetAccountsList has not yet been implemented")
		}),
		FeesGetAvgFeesHandler: fees.GetAvgFeesHandlerFunc(func(params fees.GetAvgFeesParams) middleware.Responder {
			return middleware.NotImplemented("operation FeesGetAvgFees has not yet been implemented")
		}),
		AccountsGetBakersListHandler: accounts.GetBakersListHandlerFunc(func(params accounts.GetBakersListParams) middleware.Responder {
			return middleware.NotImplemented("operation AccountsGetBakersList has not yet been implemented")
		}),
		BlocksGetBakingRightsHandler: blocks.GetBakingRightsHandlerFunc(func(params blocks.GetBakingRightsParams) middleware.Responder {
			return middleware.NotImplemented("operation BlocksGetBakingRights has not yet been implemented")
		}),
		VotingGetBallotsByPeriodIDHandler: voting.GetBallotsByPeriodIDHandlerFunc(func(params voting.GetBallotsByPeriodIDParams) middleware.Responder {
			return middleware.NotImplemented("operation VotingGetBallotsByPeriodID has not yet been implemented")
		}),
		BlocksGetBlockHandler: blocks.GetBlockHandlerFunc(func(params blocks.GetBlockParams) middleware.Responder {
			return middleware.NotImplemented("operation BlocksGetBlock has not yet been implemented")
		}),
		BlocksGetBlockBakingRightsHandler: blocks.GetBlockBakingRightsHandlerFunc(func(params blocks.GetBlockBakingRightsParams) middleware.Responder {
			return middleware.NotImplemented("operation BlocksGetBlockBakingRights has not yet been implemented")
		}),
		BlocksGetBlockEndorsementsHandler: blocks.GetBlockEndorsementsHandlerFunc(func(params blocks.GetBlockEndorsementsParams) middleware.Responder {
			return middleware.NotImplemented("operation BlocksGetBlockEndorsements has not yet been implemented")
		}),
		BlocksGetBlocksHeadHandler: blocks.GetBlocksHeadHandlerFunc(func(params blocks.GetBlocksHeadParams) middleware.Responder {
			return middleware.NotImplemented("operation BlocksGetBlocksHead has not yet been implemented")
		}),
		BlocksGetBlocksListHandler: blocks.GetBlocksListHandlerFunc(func(params blocks.GetBlocksListParams) middleware.Responder {
			return middleware.NotImplemented("operation BlocksGetBlocksList has not yet been implemented")
		}),
		AccountsGetContractsListHandler: accounts.GetContractsListHandlerFunc(func(params accounts.GetContractsListParams) middleware.Responder {
			return middleware.NotImplemented("operation AccountsGetContractsList has not yet been implemented")
		}),
		OperationsListGetDoubleBakingsListHandler: operations_list.GetDoubleBakingsListHandlerFunc(func(params operations_list.GetDoubleBakingsListParams) middleware.Responder {
			return middleware.NotImplemented("operation OperationsListGetDoubleBakingsList has not yet been implemented")
		}),
		BlocksGetFutureBakingRightsHandler: blocks.GetFutureBakingRightsHandlerFunc(func(params blocks.GetFutureBakingRightsParams) middleware.Responder {
			return middleware.NotImplemented("operation BlocksGetFutureBakingRights has not yet been implemented")
		}),
		AppInfoGetInfoHandler: app_info.GetInfoHandlerFunc(func(params app_info.GetInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation AppInfoGetInfo has not yet been implemented")
		}),
		VotingGetNonVotersByPeriodIDHandler: voting.GetNonVotersByPeriodIDHandlerFunc(func(params voting.GetNonVotersByPeriodIDParams) middleware.Responder {
			return middleware.NotImplemented("operation VotingGetNonVotersByPeriodID has not yet been implemented")
		}),
		OperationGroupsGetOperationGroupHandler: operation_groups.GetOperationGroupHandlerFunc(func(params operation_groups.GetOperationGroupParams) middleware.Responder {
			return middleware.NotImplemented("operation OperationGroupsGetOperationGroup has not yet been implemented")
		}),
		OperationGroupsGetOperationGroupsHandler: operation_groups.GetOperationGroupsHandlerFunc(func(params operation_groups.GetOperationGroupsParams) middleware.Responder {
			return middleware.NotImplemented("operation OperationGroupsGetOperationGroups has not yet been implemented")
		}),
		OperationsListGetOperationsListHandler: operations_list.GetOperationsListHandlerFunc(func(params operations_list.GetOperationsListParams) middleware.Responder {
			return middleware.NotImplemented("operation OperationsListGetOperationsList has not yet been implemented")
		}),
		VotingGetPeriodHandler: voting.GetPeriodHandlerFunc(func(params voting.GetPeriodParams) middleware.Responder {
			return middleware.NotImplemented("operation VotingGetPeriod has not yet been implemented")
		}),
		VotingGetPeriodsListHandler: voting.GetPeriodsListHandlerFunc(func(params voting.GetPeriodsListParams) middleware.Responder {
			return middleware.NotImplemented("operation VotingGetPeriodsList has not yet been implemented")
		}),
		VotingGetProposalVotesListHandler: voting.GetProposalVotesListHandlerFunc(func(params voting.GetProposalVotesListParams) middleware.Responder {
			return middleware.NotImplemented("operation VotingGetProposalVotesList has not yet been implemented")
		}),
		VotingGetProposalsByPeriodIDHandler: voting.GetProposalsByPeriodIDHandlerFunc(func(params voting.GetProposalsByPeriodIDParams) middleware.Responder {
			return middleware.NotImplemented("operation VotingGetProposalsByPeriodID has not yet been implemented")
		}),
		VotingGetProtocolsListHandler: voting.GetProtocolsListHandlerFunc(func(params voting.GetProtocolsListParams) middleware.Responder {
			return middleware.NotImplemented("operation VotingGetProtocolsList has not yet been implemented")
		}),
		GetSnapshotsHandler: GetSnapshotsHandlerFunc(func(params GetSnapshotsParams) middleware.Responder {
			return middleware.NotImplemented("operation GetSnapshots has not yet been implemented")
		}),
	}
}

/*TezTrackerAPI the tez tracker API */
type TezTrackerAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// AccountsGetAccountHandler sets the operation handler for the get account operation
	AccountsGetAccountHandler accounts.GetAccountHandler
	// AccountsGetAccountDelegatorsHandler sets the operation handler for the get account delegators operation
	AccountsGetAccountDelegatorsHandler accounts.GetAccountDelegatorsHandler
	// AccountsGetAccountsListHandler sets the operation handler for the get accounts list operation
	AccountsGetAccountsListHandler accounts.GetAccountsListHandler
	// FeesGetAvgFeesHandler sets the operation handler for the get avg fees operation
	FeesGetAvgFeesHandler fees.GetAvgFeesHandler
	// AccountsGetBakersListHandler sets the operation handler for the get bakers list operation
	AccountsGetBakersListHandler accounts.GetBakersListHandler
	// BlocksGetBakingRightsHandler sets the operation handler for the get baking rights operation
	BlocksGetBakingRightsHandler blocks.GetBakingRightsHandler
	// VotingGetBallotsByPeriodIDHandler sets the operation handler for the get ballots by period ID operation
	VotingGetBallotsByPeriodIDHandler voting.GetBallotsByPeriodIDHandler
	// BlocksGetBlockHandler sets the operation handler for the get block operation
	BlocksGetBlockHandler blocks.GetBlockHandler
	// BlocksGetBlockBakingRightsHandler sets the operation handler for the get block baking rights operation
	BlocksGetBlockBakingRightsHandler blocks.GetBlockBakingRightsHandler
	// BlocksGetBlockEndorsementsHandler sets the operation handler for the get block endorsements operation
	BlocksGetBlockEndorsementsHandler blocks.GetBlockEndorsementsHandler
	// BlocksGetBlocksHeadHandler sets the operation handler for the get blocks head operation
	BlocksGetBlocksHeadHandler blocks.GetBlocksHeadHandler
	// BlocksGetBlocksListHandler sets the operation handler for the get blocks list operation
	BlocksGetBlocksListHandler blocks.GetBlocksListHandler
	// AccountsGetContractsListHandler sets the operation handler for the get contracts list operation
	AccountsGetContractsListHandler accounts.GetContractsListHandler
	// OperationsListGetDoubleBakingsListHandler sets the operation handler for the get double bakings list operation
	OperationsListGetDoubleBakingsListHandler operations_list.GetDoubleBakingsListHandler
	// BlocksGetFutureBakingRightsHandler sets the operation handler for the get future baking rights operation
	BlocksGetFutureBakingRightsHandler blocks.GetFutureBakingRightsHandler
	// AppInfoGetInfoHandler sets the operation handler for the get info operation
	AppInfoGetInfoHandler app_info.GetInfoHandler
	// VotingGetNonVotersByPeriodIDHandler sets the operation handler for the get non voters by period ID operation
	VotingGetNonVotersByPeriodIDHandler voting.GetNonVotersByPeriodIDHandler
	// OperationGroupsGetOperationGroupHandler sets the operation handler for the get operation group operation
	OperationGroupsGetOperationGroupHandler operation_groups.GetOperationGroupHandler
	// OperationGroupsGetOperationGroupsHandler sets the operation handler for the get operation groups operation
	OperationGroupsGetOperationGroupsHandler operation_groups.GetOperationGroupsHandler
	// OperationsListGetOperationsListHandler sets the operation handler for the get operations list operation
	OperationsListGetOperationsListHandler operations_list.GetOperationsListHandler
	// VotingGetPeriodHandler sets the operation handler for the get period operation
	VotingGetPeriodHandler voting.GetPeriodHandler
	// VotingGetPeriodsListHandler sets the operation handler for the get periods list operation
	VotingGetPeriodsListHandler voting.GetPeriodsListHandler
	// VotingGetProposalVotesListHandler sets the operation handler for the get proposal votes list operation
	VotingGetProposalVotesListHandler voting.GetProposalVotesListHandler
	// VotingGetProposalsByPeriodIDHandler sets the operation handler for the get proposals by period ID operation
	VotingGetProposalsByPeriodIDHandler voting.GetProposalsByPeriodIDHandler
	// VotingGetProtocolsListHandler sets the operation handler for the get protocols list operation
	VotingGetProtocolsListHandler voting.GetProtocolsListHandler
	// GetSnapshotsHandler sets the operation handler for the get snapshots operation
	GetSnapshotsHandler GetSnapshotsHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *TezTrackerAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *TezTrackerAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *TezTrackerAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *TezTrackerAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *TezTrackerAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *TezTrackerAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *TezTrackerAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the TezTrackerAPI
func (o *TezTrackerAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.AccountsGetAccountHandler == nil {
		unregistered = append(unregistered, "accounts.GetAccountHandler")
	}

	if o.AccountsGetAccountDelegatorsHandler == nil {
		unregistered = append(unregistered, "accounts.GetAccountDelegatorsHandler")
	}

	if o.AccountsGetAccountsListHandler == nil {
		unregistered = append(unregistered, "accounts.GetAccountsListHandler")
	}

	if o.FeesGetAvgFeesHandler == nil {
		unregistered = append(unregistered, "fees.GetAvgFeesHandler")
	}

	if o.AccountsGetBakersListHandler == nil {
		unregistered = append(unregistered, "accounts.GetBakersListHandler")
	}

	if o.BlocksGetBakingRightsHandler == nil {
		unregistered = append(unregistered, "blocks.GetBakingRightsHandler")
	}

	if o.VotingGetBallotsByPeriodIDHandler == nil {
		unregistered = append(unregistered, "voting.GetBallotsByPeriodIDHandler")
	}

	if o.BlocksGetBlockHandler == nil {
		unregistered = append(unregistered, "blocks.GetBlockHandler")
	}

	if o.BlocksGetBlockBakingRightsHandler == nil {
		unregistered = append(unregistered, "blocks.GetBlockBakingRightsHandler")
	}

	if o.BlocksGetBlockEndorsementsHandler == nil {
		unregistered = append(unregistered, "blocks.GetBlockEndorsementsHandler")
	}

	if o.BlocksGetBlocksHeadHandler == nil {
		unregistered = append(unregistered, "blocks.GetBlocksHeadHandler")
	}

	if o.BlocksGetBlocksListHandler == nil {
		unregistered = append(unregistered, "blocks.GetBlocksListHandler")
	}

	if o.AccountsGetContractsListHandler == nil {
		unregistered = append(unregistered, "accounts.GetContractsListHandler")
	}

	if o.OperationsListGetDoubleBakingsListHandler == nil {
		unregistered = append(unregistered, "operations_list.GetDoubleBakingsListHandler")
	}

	if o.BlocksGetFutureBakingRightsHandler == nil {
		unregistered = append(unregistered, "blocks.GetFutureBakingRightsHandler")
	}

	if o.AppInfoGetInfoHandler == nil {
		unregistered = append(unregistered, "app_info.GetInfoHandler")
	}

	if o.VotingGetNonVotersByPeriodIDHandler == nil {
		unregistered = append(unregistered, "voting.GetNonVotersByPeriodIDHandler")
	}

	if o.OperationGroupsGetOperationGroupHandler == nil {
		unregistered = append(unregistered, "operation_groups.GetOperationGroupHandler")
	}

	if o.OperationGroupsGetOperationGroupsHandler == nil {
		unregistered = append(unregistered, "operation_groups.GetOperationGroupsHandler")
	}

	if o.OperationsListGetOperationsListHandler == nil {
		unregistered = append(unregistered, "operations_list.GetOperationsListHandler")
	}

	if o.VotingGetPeriodHandler == nil {
		unregistered = append(unregistered, "voting.GetPeriodHandler")
	}

	if o.VotingGetPeriodsListHandler == nil {
		unregistered = append(unregistered, "voting.GetPeriodsListHandler")
	}

	if o.VotingGetProposalVotesListHandler == nil {
		unregistered = append(unregistered, "voting.GetProposalVotesListHandler")
	}

	if o.VotingGetProposalsByPeriodIDHandler == nil {
		unregistered = append(unregistered, "voting.GetProposalsByPeriodIDHandler")
	}

	if o.VotingGetProtocolsListHandler == nil {
		unregistered = append(unregistered, "voting.GetProtocolsListHandler")
	}

	if o.GetSnapshotsHandler == nil {
		unregistered = append(unregistered, "GetSnapshotsHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *TezTrackerAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *TezTrackerAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *TezTrackerAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *TezTrackerAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *TezTrackerAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *TezTrackerAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the tez tracker API
func (o *TezTrackerAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *TezTrackerAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{platform}/{network}/accounts/{accountId}"] = accounts.NewGetAccount(o.context, o.AccountsGetAccountHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{platform}/{network}/accounts/{accountId}/delegators"] = accounts.NewGetAccountDelegators(o.context, o.AccountsGetAccountDelegatorsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{platform}/{network}/accounts"] = accounts.NewGetAccountsList(o.context, o.AccountsGetAccountsListHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{platform}/{network}/operations/avgFees"] = fees.NewGetAvgFees(o.context, o.FeesGetAvgFeesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{platform}/{network}/bakers"] = accounts.NewGetBakersList(o.context, o.AccountsGetBakersListHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{platform}/{network}/baking_rights"] = blocks.NewGetBakingRights(o.context, o.BlocksGetBakingRightsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{network}/ballots/{id}"] = voting.NewGetBallotsByPeriodID(o.context, o.VotingGetBallotsByPeriodIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{platform}/{network}/blocks/{hash}"] = blocks.NewGetBlock(o.context, o.BlocksGetBlockHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{platform}/{network}/blocks/{hash}/baking_rights"] = blocks.NewGetBlockBakingRights(o.context, o.BlocksGetBlockBakingRightsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{platform}/{network}/blocks/{hash}/endorsements"] = blocks.NewGetBlockEndorsements(o.context, o.BlocksGetBlockEndorsementsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{platform}/{network}/blocks/head"] = blocks.NewGetBlocksHead(o.context, o.BlocksGetBlocksHeadHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{platform}/{network}/blocks"] = blocks.NewGetBlocksList(o.context, o.BlocksGetBlocksListHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{platform}/{network}/contracts"] = accounts.NewGetContractsList(o.context, o.AccountsGetContractsListHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{platform}/{network}/double_bakings"] = operations_list.NewGetDoubleBakingsList(o.context, o.OperationsListGetDoubleBakingsListHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{platform}/{network}/future_baking_rights"] = blocks.NewGetFutureBakingRights(o.context, o.BlocksGetFutureBakingRightsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{platform}/{network}/info"] = app_info.NewGetInfo(o.context, o.AppInfoGetInfoHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{network}/non_voters/{id}"] = voting.NewGetNonVotersByPeriodID(o.context, o.VotingGetNonVotersByPeriodIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{platform}/{network}/operation_groups/{operationGroupId}"] = operation_groups.NewGetOperationGroup(o.context, o.OperationGroupsGetOperationGroupHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{platform}/{network}/operation_groups"] = operation_groups.NewGetOperationGroups(o.context, o.OperationGroupsGetOperationGroupsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{platform}/{network}/operations"] = operations_list.NewGetOperationsList(o.context, o.OperationsListGetOperationsListHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{network}/period"] = voting.NewGetPeriod(o.context, o.VotingGetPeriodHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{network}/periods"] = voting.NewGetPeriodsList(o.context, o.VotingGetPeriodsListHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{network}/proposal_votes/{id}"] = voting.NewGetProposalVotesList(o.context, o.VotingGetProposalVotesListHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{network}/proposals"] = voting.NewGetProposalsByPeriodID(o.context, o.VotingGetProposalsByPeriodIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{network}/protocols"] = voting.NewGetProtocolsList(o.context, o.VotingGetProtocolsListHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/data/{platform}/{network}/snapshots"] = NewGetSnapshots(o.context, o.GetSnapshotsHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *TezTrackerAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *TezTrackerAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *TezTrackerAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *TezTrackerAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
