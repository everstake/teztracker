package api

import (
	"github.com/everstake/teztracker/api/render"
	vt "github.com/everstake/teztracker/gen/restapi/operations/voting"
	"github.com/everstake/teztracker/repos"
	"github.com/everstake/teztracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	"strconv"
)

type getNonVotersHandler struct {
	provider DbProvider
}

// Handle serves the Get Non Voters List request.
func (h *getNonVotersHandler) Handle(params vt.GetNonVotersByPeriodIDParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return vt.NewGetNonVotersByPeriodIDNotFound()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return vt.NewGetNonVotersByPeriodIDNotFound()
	}
	service := services.New(repos.New(db), net)

	limiter := NewLimiter(params.Limit, params.Offset)

	id, err := strconv.ParseInt(params.ID, 10, 64)
	if err != nil {
		return vt.NewGetNonVotersByPeriodIDNotFound()
	}

	votes, count, err := service.GetPeriodNonVoters(id, limiter)
	if err != nil {
		logrus.Errorf("failed to get period non voters list: %s", err.Error())
		return vt.NewGetNonVotersByPeriodIDNotFound()
	}

	return vt.NewGetNonVotersByPeriodIDOK().WithPayload(render.NonVoters(votes)).WithXTotalCount(count)
}
