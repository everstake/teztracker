package api

import (
	"github.com/everstake/teztracker/api/render"
	"github.com/everstake/teztracker/gen/restapi/operations/accounts"
	"github.com/everstake/teztracker/services/whales"
	"github.com/go-openapi/runtime/middleware"
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

type getWhaleTransfersHandler struct{}

func (h *getWhaleTransfersHandler) Handle(params accounts.GetWhaleTranfersParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetWhaleTranfersBadRequest()
	}
	data := whales.Service.GetData(net)
	return accounts.NewGetWhaleTranfersOK().WithPayload(render.WhaleTransfers(data))
}
