package api

import (
	"github.com/everstake/teztracker/api/render"
	"github.com/everstake/teztracker/gen/restapi/operations/accounts"
	"github.com/everstake/teztracker/repos"
	"github.com/everstake/teztracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

type getBakersStakeChange struct {
	provider DbProvider
}

// Handle serves the Get Period Info request.
func (h *getBakersStakeChange) Handle(params accounts.GetBakersStakeChangeParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetBakersStakeChangeBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return accounts.NewGetBakersStakeChangeBadRequest()
	}

	service := services.New(repos.New(db), net)
	data, err := service.GetBakersStakeChange()
	if err != nil {
		logrus.Errorf("GetBakersStakeChange: %s", err.Error())
		return accounts.NewGetBakersStakeChangeInternalServerError()
	}

	return accounts.NewGetBakersStakeChangeOK().WithPayload(render.BakersDelegators(data))
}
