package api

import (
	"github.com/everstake/teztracker/api/render"
	"github.com/everstake/teztracker/gen/restapi/operations/app_info"
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/repos"
	"github.com/everstake/teztracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

type getThirdPartyBakersHandler struct {
	provider DbProvider
}

// Handle serves the Get Third Party Bakers request.
func (h *getThirdPartyBakersHandler) Handle(params app_info.GetThirdPartyBakersHandlerParams) middleware.Responder {
	db, err := h.provider.GetDb(models.NetworkMain)
	if err != nil {
		return app_info.NewGetThirdPartyBakersHandlerInternalServerError()
	}
	service := services.New(repos.New(db), models.NetworkMain)
	bakers, err := service.GetThirdPartyBakers()
	if err != nil {
		logrus.Errorf("service.GetThirdPartyBakers: %s", err.Error())
		return app_info.NewGetThirdPartyBakersHandlerInternalServerError()
	}
	return app_info.NewGetThirdPartyBakersHandlerOK().WithPayload(render.ThirdPartyBakers(bakers))
}
