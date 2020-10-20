package api

import (
	"github.com/everstake/teztracker/gen/models"
	"github.com/everstake/teztracker/gen/restapi/operations/app_info"
	"github.com/everstake/teztracker/repos"
	"github.com/everstake/teztracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

type getHealthHandler struct {
	provider DbProvider
}

// Handle serves the Get Health request.
func (h *getHealthHandler) Handle(params app_info.GetHealthCheckInfoParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return app_info.NewGetHealthCheckInfoBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return app_info.NewGetHealthCheckInfoInternalServerError()
	}

	service := services.New(repos.New(db), net)

	err = service.Health()
	if err != nil {
		logrus.Errorf("failed to get Health: %s", err.Error())
		return app_info.NewGetHealthCheckInfoInternalServerError()
	}

	return app_info.NewGetHealthCheckInfoOK().WithPayload(&models.Health{Status: true})
}
