package api

import (
	"github.com/bullblock-io/tezTracker/api/render"
	ops "github.com/bullblock-io/tezTracker/gen/restapi/operations/operations_list"
	"github.com/bullblock-io/tezTracker/repos"
	"github.com/bullblock-io/tezTracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type getOperationListHandler struct {
	db *gorm.DB
}

// Handle serves the Get Operations List request.
func (h *getOperationListHandler) Handle(params ops.GetOperationsListParams) middleware.Responder {
	service := services.New(repos.New(h.db))
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
