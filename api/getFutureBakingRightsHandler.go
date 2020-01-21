package api

import (
	"github.com/bullblock-io/tezTracker/api/render"
	"github.com/bullblock-io/tezTracker/gen/restapi/operations/blocks"
	"github.com/bullblock-io/tezTracker/repos"
	"github.com/bullblock-io/tezTracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

type getFutureBakingRightsHandler struct {
	provider DbProvider
}

// Handle serves the Get Block List request.
func (h *getFutureBakingRightsHandler) Handle(params blocks.GetFutureBakingRightsParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return blocks.NewGetFutureBakingRightsBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return blocks.NewGetFutureBakingRightsNotFound()
	}
	service := services.New(repos.New(db), net)

	limiter := NewLimiter(params.Limit, params.Offset)
	priorityTo := 0
	if params.PrioritiesTo != nil {
		priorityTo = int(*params.PrioritiesTo)
	}
	count, bs, err := service.FutureBakingRightsList(priorityTo, limiter)
	if err != nil {
		logrus.Errorf("failed to get future baking rights: %s", err.Error())
		return blocks.NewGetFutureBakingRightsNotFound()
	}

	return blocks.NewGetFutureBakingRightsOK().WithPayload(render.BlocksFutureBakingRights(bs)).WithXTotalCount(count)
}
