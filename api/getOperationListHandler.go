package api

import (
	"fmt"
	"github.com/everstake/teztracker/api/render"
	ops "github.com/everstake/teztracker/gen/restapi/operations/operations_list"
	"github.com/everstake/teztracker/repos"
	"github.com/everstake/teztracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"
	"github.com/sirupsen/logrus"
	"strings"
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

	var kinds []string

	for key := range params.OperationKind {
		kinds = append(kinds, strings.Split(params.OperationKind[key], ",")...)
	}

	for key := range kinds {
		if err := validate.Enum(fmt.Sprintf("%s.%v", "operation_kind", key), "query", kinds[key], []interface{}{"endorsement", "endorsement_with_slot", "proposals", "seed_nonce_revelation", "delegation", "transaction", "activate_account", "ballot", "origination", "reveal", "double_baking_evidence", "double_endorsement_evidence"}); err != nil {
			return ops.NewGetOperationsListBadRequest()
		}
	}

	operations, count, err := service.GetOperations(params.OperationID, kinds, params.BlockID, params.AccountID, limiter, before)
	if err != nil {
		logrus.Errorf("failed to get operations: %s", err.Error())
		return ops.NewGetOperationsListNotFound()

	}

	return ops.NewGetOperationsListOK().WithPayload(render.Operations(operations)).WithXTotalCount(count)
}
