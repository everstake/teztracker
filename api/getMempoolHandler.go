package api

import (
	"github.com/everstake/teztracker/api/render"
	"github.com/everstake/teztracker/gen/restapi/operations/mempool"
	"github.com/go-openapi/runtime/middleware"
)

type getMempoolHandler struct {
	provider MempoolProvider
}

func (h *getMempoolHandler) Handle(params mempool.GetMempoolOperationsParams) middleware.Responder {

	net, err := ToNetwork(params.Network)
	if err != nil {
		return mempool.NewGetMempoolOperationsBadRequest()
	}

	mem, err := h.provider.GetMempool(net)
	if err != nil {
		return mempool.NewGetMempoolOperationsBadRequest()
	}

	op, err := mem.GetMempool()
	if err != nil {
		return mempool.NewGetMempoolOperationsInternalServerError()
	}

	return mempool.NewGetMempoolOperationsOK().WithPayload(render.MempoolOperationsList(op))
}
