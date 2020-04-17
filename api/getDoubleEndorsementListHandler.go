package api

import (
	"github.com/everstake/teztracker/api/render"
	ops "github.com/everstake/teztracker/gen/restapi/operations/operations_list"
	"github.com/everstake/teztracker/repos"
	"github.com/everstake/teztracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

type getDoubleEndorsementListHandler struct {
	provider DbProvider
}

// Handle serves the Get Operations List request.
func (h *getDoubleEndorsementListHandler) Handle(params ops.GetDoubleEndorsementsListParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return ops.NewGetDoubleEndorsementsListBadRequest()
	}

	db, err := h.provider.GetDb(net)
	if err != nil {
		return ops.NewGetDoubleEndorsementsListNotFound()
	}
	service := services.New(repos.New(db), net)

	limiter := NewLimiter(params.Limit, params.Offset)

	operations, count, err := service.GetDoubleEndorsements(params.OperationID, params.BlockID, limiter)
	if err != nil {
		logrus.Errorf("failed to get operations: %s", err.Error())
		return ops.NewGetDoubleEndorsementsListNotFound()

	}

	return ops.NewGetDoubleEndorsementsListOK().WithPayload(render.DoubleOperations(operations)).WithXTotalCount(count)
}
