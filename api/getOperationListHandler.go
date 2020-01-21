package api

import (
	"github.com/bullblock-io/tezTracker/api/render"
	ops "github.com/bullblock-io/tezTracker/gen/restapi/operations/operations_list"
	"github.com/bullblock-io/tezTracker/repos"
	"github.com/bullblock-io/tezTracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

type getOperationListHandler struct {
	provider DbProvider
}

// Handle serves the Get Operations List request.
func (h *getOperationListHandler) Handle(params ops.GetOperationsListParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return ops.NewGetOperationsListBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return ops.NewGetOperationsListNotFound()
	}
	service := services.New(repos.New(db), net)

	limiter := NewLimiter(params.Limit, params.Offset)
	before := int64(0)
	if params.BeforeID != nil {
		before = *params.BeforeID
	}

	operations, count, err := service.GetOperations(params.OperationID, params.OperationKind, params.BlockID, params.AccountID, limiter, before)

	if err != nil {
		logrus.Errorf("failed to get operations: %s", err.Error())
		return ops.NewGetOperationsListNotFound()

	}

	return ops.NewGetOperationsListOK().WithPayload(render.Operations(operations)).WithXTotalCount(count)
}
