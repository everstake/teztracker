package api

import (
	"github.com/everstake/teztracker/api/render"
	"github.com/everstake/teztracker/gen/restapi/operations/blocks"
	"github.com/everstake/teztracker/repos"
	"github.com/everstake/teztracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

type getBakingRightsHandler struct {
	provider DbProvider
}

// Handle serves the Get Block List request.
func (h *getBakingRightsHandler) Handle(params blocks.GetBakingRightsParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return blocks.NewGetBakingRightsBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return blocks.NewGetBakingRightsNotFound()
	}
	service := services.New(repos.New(db), net)
	limiter := NewLimiter(params.Limit, params.Offset)
	priorityTo := 0
	if params.PrioritiesTo != nil {
		priorityTo = int(*params.PrioritiesTo)
	}
	count, bs, err := service.BakingRightsList(params.BlockID, priorityTo, limiter)
	if err != nil {
		logrus.Errorf("failed to get baking rights: %s", err.Error())
		return blocks.NewGetBakingRightsNotFound()
	}

	return blocks.NewGetBakingRightsOK().WithPayload(render.BlocksBakingRights(bs)).WithXTotalCount(count)
}
