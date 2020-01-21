package api

import (
	"github.com/bullblock-io/tezTracker/api/render"
	"github.com/bullblock-io/tezTracker/gen/restapi/operations/accounts"
	"github.com/bullblock-io/tezTracker/repos"
	"github.com/bullblock-io/tezTracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

type getAccountDelegatorsHandler struct {
	provider DbProvider
}

// Handle serves the Get Account Delegators request.
func (h *getAccountDelegatorsHandler) Handle(params accounts.GetAccountDelegatorsParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetAccountDelegatorsBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return accounts.NewGetAccountDelegatorsInternalServerError()
	}

	service := services.New(repos.New(db), net)
	limiter := NewLimiter(params.Limit, params.Offset)
	accs, count, err := service.AccountDelegatorsList(params.AccountID, limiter)
	if err != nil {
		logrus.Errorf("failed to get account's delegators: %s", err.Error())
		return accounts.NewGetAccountDelegatorsInternalServerError()
	}
	return accounts.NewGetAccountDelegatorsOK().WithPayload(render.Accounts(accs)).WithXTotalCount(count)
}
