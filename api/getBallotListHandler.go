package api

import (
	"github.com/everstake/teztracker/api/render"
	vt "github.com/everstake/teztracker/gen/restapi/operations/voting"
	"github.com/everstake/teztracker/repos"
	"github.com/everstake/teztracker/services"
	"github.com/go-openapi/runtime/middleware"
	"strconv"
)

type getBallotsHandler struct {
	provider DbProvider
}

// Handle serves the Get Ballots List request.
func (h *getBallotsHandler) Handle(params vt.GetBallotsByPeriodIDParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return vt.NewGetBallotsByPeriodIDNotFound()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return vt.NewGetBallotsByPeriodIDNotFound()
	}
	service := services.New(repos.New(db), net)

	limiter := NewLimiter(params.Limit, params.Offset)

	id, err := strconv.ParseInt(params.ID, 10, 64)
	if err != nil {
		return vt.NewGetBallotsByPeriodIDNotFound()
	}

	votes, _, err := service.GetBallotVoters(id, limiter)
	if err != nil {
		return vt.NewGetBallotsByPeriodIDNotFound()
	}

	return vt.NewGetBallotsByPeriodIDOK().WithPayload(render.BallotVoters(votes))
}
