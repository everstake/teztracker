package api

import (
	"github.com/everstake/teztracker/api/render"
	ops "github.com/everstake/teztracker/gen/restapi/operations/operations_list"
	"github.com/everstake/teztracker/repos"
	"github.com/everstake/teztracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

type getDoubleBakingsListHandler struct {
	provider DbProvider
}

// Handle serves the Get Operations List request.
func (h *getDoubleBakingsListHandler) Handle(params ops.GetDoubleBakingsListParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return ops.NewGetDoubleBakingsListBadRequest()
	}

	db, err := h.provider.GetDb(net)
	if err != nil {
		return ops.NewGetDoubleBakingsListNotFound()
	}
	service := services.New(repos.New(db), net)

	limiter := NewLimiter(params.Limit, params.Offset)

	operations, count, err := service.GetDoubleBakings(params.OperationID, params.BlockID, limiter)
	if err != nil {
		logrus.Errorf("failed to get operations: %s", err.Error())
		return ops.NewGetDoubleBakingsListNotFound()

	}

	return ops.NewGetDoubleBakingsListOK().WithPayload(render.DoubleBakings(operations)).WithXTotalCount(count)
}
