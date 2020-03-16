package api

import (
	"github.com/everstake/teztracker/api/render"
	vt "github.com/everstake/teztracker/gen/restapi/operations/voting"
	"github.com/everstake/teztracker/repos"
	"github.com/everstake/teztracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

type getPeriodListHandler struct {
	provider DbProvider
}

// Handle serves the Get Proposal List request.
func (h *getPeriodListHandler) Handle(params vt.GetPeriodsListParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return vt.NewGetProposalsByPeriodIDNotFound()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return vt.NewGetProposalsByPeriodIDNotFound()
	}
	service := services.New(repos.New(db), net)

	periods, err := service.VotingPeriodsList()
	if err != nil {
		logrus.Errorf("failed to get voting periods list: %s", err.Error())
		return vt.NewGetProposalsByPeriodIDNotFound()
	}

	return vt.NewGetPeriodsListOK().WithPayload(render.Periods(periods))
}
