package api

import (
	"github.com/everstake/teztracker/api/render"
	"github.com/everstake/teztracker/gen/restapi/operations/voting"
	"github.com/everstake/teztracker/repos"
	"github.com/everstake/teztracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	"strconv"
)

type getPeriodInfoHandler struct {
	provider DbProvider
}

// Handle serves the Get Period Info request.
func (h *getPeriodInfoHandler) Handle(params voting.GetPeriodParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return voting.NewGetPeriodBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return voting.NewGetPeriodNotFound()
	}
	service := services.New(repos.New(db), net)
	var id *int64
	if params.ID != nil {
		parseID, err := strconv.ParseInt(*params.ID, 10, 64)
		if err == nil {
			id = &parseID
		}
	}

	period, err := service.VotingPeriod(id)
	if err != nil {
		logrus.Errorf("failed to get voting period: %s", err.Error())
		return voting.NewGetPeriodNotFound()
	}

	return voting.NewGetPeriodOK().WithPayload(render.Period(period))
}
