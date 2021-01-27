package api

import (
	"github.com/everstake/teztracker/api/render"
	"github.com/everstake/teztracker/gen/restapi/operations/accounts"
	"github.com/everstake/teztracker/repos"
	"github.com/everstake/teztracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

type getBakersVoting struct {
	provider DbProvider
}

// Handle serves the Get Period Info request.
func (h *getBakersVoting) Handle(params accounts.GetBakersVotingParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetBakersVotingBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return accounts.NewGetBakersVotingBadRequest()
	}

	service := services.New(repos.New(db), net)
	data, err := service.GetBakersVoting()
	if err != nil {
		logrus.Errorf("GetBakersVoting: %s", err.Error())
		return accounts.NewGetBakersVotingInternalServerError()
	}

	return accounts.NewGetBakersVotingOK().WithPayload(render.BakersVoting(data))
}
