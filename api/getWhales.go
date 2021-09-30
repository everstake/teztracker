package api

import (
	"github.com/everstake/teztracker/api/render"
	"github.com/everstake/teztracker/gen/restapi/operations/accounts"
	"github.com/everstake/teztracker/gen/restapi/operations/operations_list"
	"github.com/everstake/teztracker/repos"
	"github.com/everstake/teztracker/services"
	"github.com/everstake/teztracker/services/whales"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

type getWhaleAccountsHandler struct{}

func (h *getWhaleAccountsHandler) Handle(params accounts.GetWhaleAccountsParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetWhaleAccountsBadRequest()
	}
	data := whales.Service.GetData(net)
	return accounts.NewGetWhaleAccountsOK().WithPayload(render.WhaleAccounts(data))
}

type getWhaleTransfersHandler struct {
	provider DbProvider
}

func (h *getWhaleTransfersHandler) Handle(params operations_list.GetWhaleTranfersParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return operations_list.NewGetWhaleTranfersBadRequest()
	}

	db, err := h.provider.GetDb(net)
	if err != nil {
		return operations_list.NewGetWhaleTranfersBadRequest()
	}
	service := services.New(repos.New(db), net)

	limiter := NewLimiter(params.Limit, params.Offset)

	ops, err := service.GetWhaleTransfers(limiter, *params.Period)
	if err != nil {
		logrus.Errorf("failed to get whale transfers: %s", err.Error())
		return operations_list.NewGetWhaleTranfersInternalServerError()
	}

	return operations_list.NewGetWhaleTranfersOK().WithPayload(render.Operations(ops))
}
