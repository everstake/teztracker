package api

import (
	"github.com/everstake/teztracker/api/render"
	"github.com/everstake/teztracker/gen/restapi/operations/app_info"
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/repos/thirdparty_bakers"
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
	repo := thirdparty_bakers.New(db)
	bakers, err := repo.GetAll()
	if err != nil {
		logrus.Errorf("getThirdPartyBakersHandler: repo.GetAll: %s", err.Error())
		return app_info.NewGetThirdPartyBakersHandlerInternalServerError()
	}
	return app_info.NewGetThirdPartyBakersHandlerOK().WithPayload(render.ThirdPartyBakers(bakers))
}
