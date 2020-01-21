package api

import (
	"github.com/bullblock-io/tezTracker/api/render"
	"github.com/bullblock-io/tezTracker/gen/restapi/operations/accounts"
	"github.com/bullblock-io/tezTracker/repos"
	"github.com/bullblock-io/tezTracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

type getAccountHandler struct {
	provider DbProvider
}

// Handle serves the Get Account request.
func (h *getAccountHandler) Handle(params accounts.GetAccountParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetAccountBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return accounts.NewGetAccountInternalServerError()
	}
	service := services.New(repos.New(db), net)

	acc, err := service.GetAccount(params.AccountID)

	if err != nil {
		if err == services.ErrNotFound {
			return accounts.NewGetAccountNotFound()
		}
		logrus.Errorf("failed to get acc: %s", err.Error())
		return accounts.NewGetAccountInternalServerError()
	}

	return accounts.NewGetAccountOK().WithPayload(render.Account(acc))
}
