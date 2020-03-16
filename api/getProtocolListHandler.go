package api

import (
	"github.com/everstake/teztracker/api/render"
	vt "github.com/everstake/teztracker/gen/restapi/operations/voting"
	"github.com/everstake/teztracker/repos"
	"github.com/everstake/teztracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

type getProtocolListHandler struct {
	provider DbProvider
}

// Handle serves the Get Proposal List request.
func (h *getProtocolListHandler) Handle(params vt.GetProtocolsListParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return vt.NewGetProtocolsListNotFound()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return vt.NewGetProtocolsListNotFound()
	}
	service := services.New(repos.New(db), net)

	limiter := NewLimiter(params.Limit, params.Offset)

	votes, count, err := service.GetProtocolsList(limiter)
	if err != nil {
		logrus.Errorf("failed to get proposal voters: %s", err.Error())
		return vt.NewGetProtocolsListNotFound()
	}

	return vt.NewGetProtocolsListOK().WithPayload(render.Protocols(votes)).WithXTotalCount(count)
}
