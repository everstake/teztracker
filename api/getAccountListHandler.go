package api

import (
	"github.com/bullblock-io/tezTracker/api/render"
	"github.com/bullblock-io/tezTracker/gen/restapi/operations/accounts"
	"github.com/bullblock-io/tezTracker/repos"
	"github.com/bullblock-io/tezTracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

type getAccountListHandler struct {
	provider DbProvider
}

// Handle serves the Get Account List request.
func (h *getAccountListHandler) Handle(params accounts.GetAccountsListParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetAccountsListBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return accounts.NewGetAccountsListNotFound()
	}
	service := services.New(repos.New(db), net)
	limiter := NewLimiter(params.Limit, params.Offset)
	before := ""
	if params.AfterID != nil {
		before = *params.AfterID
	}
	accs, count, err := service.AccountList(before, limiter)
	if err != nil {
		logrus.Errorf("failed to get accounts: %s", err.Error())
		return accounts.NewGetAccountsListNotFound()
	}
	return accounts.NewGetAccountsListOK().WithPayload(render.Accounts(accs)).WithXTotalCount(count)
}
