package api

import (
	"github.com/everstake/teztracker/api/render"
	vt "github.com/everstake/teztracker/gen/restapi/operations/voting"
	"github.com/everstake/teztracker/repos"
	"github.com/everstake/teztracker/services"
	"github.com/go-openapi/runtime/middleware"
	"strconv"
)

type getProposalVotesHandler struct {
	provider DbProvider
}

// Handle serves the Get Proposal List request.
func (h *getProposalVotesHandler) Handle(params vt.GetProposalVotesListParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return vt.NewGetProposalVotesListNotFound()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return vt.NewGetProposalVotesListNotFound()
	}
	service := services.New(repos.New(db), net)

	limiter := NewLimiter(params.Limit, params.Offset)

	id, err := strconv.ParseInt(params.ID, 10, 64)
	if err != nil {
		return vt.NewGetProposalsByPeriodIDNotFound()
	}

	votes, _, err := service.GetProposalVoters(id, limiter)
	if err != nil {
		return vt.NewGetProposalsByPeriodIDNotFound()
	}

	return vt.NewGetProposalVotesListOK().WithPayload(render.ProposalVoters(votes))
}