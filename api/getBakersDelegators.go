package api

import (
	"github.com/everstake/teztracker/api/render"
	"github.com/everstake/teztracker/gen/restapi/operations/accounts"
	"github.com/everstake/teztracker/repos"
	"github.com/everstake/teztracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

type getBakersDelegators struct {
	provider DbProvider
}

// Handle serves the Get Period Info request.
func (h *getBakersDelegators) Handle(params accounts.GetBakersDelegatorsParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetBakersDelegatorsBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return accounts.NewGetBakersDelegatorsBadRequest()
	}

	service := services.New(repos.New(db), net)
	data, err := service.GetNumberOfDelegators()
	if err != nil {
		logrus.Errorf("getBakersDelegators: %s", err.Error())
		return accounts.NewGetBakersDelegatorsInternalServerError()
	}

	return accounts.NewGetBakersDelegatorsOK().WithPayload(render.BakersDelegators(data))
}
