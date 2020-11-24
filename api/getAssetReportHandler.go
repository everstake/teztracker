package api

import (
	"github.com/everstake/teztracker/gen/restapi/operations/assets"
	"github.com/everstake/teztracker/repos"
	"github.com/everstake/teztracker/services"
	"github.com/go-openapi/runtime/middleware"
	log "github.com/sirupsen/logrus"
)

type getAssetReportHandler struct {
	provider DbProvider
}

func (h *getAssetReportHandler) Handle(params assets.GetAssetReportParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return assets.NewGetAssetReportBadRequest()
	}

	db, err := h.provider.GetDb(net)
	if err != nil {
		return assets.NewGetAssetReportNotFound()
	}

	service := services.New(repos.New(db), net)

	resp, err := service.GetAssetReport(params.AssetID, params.From, params.To, params.OperationType)
	if err != nil {
		log.Errorf("GetAccountReport error: %s", err)
		return assets.NewGetAssetReportInternalServerError()
	}

	return assets.NewGetAssetReportOK().WithPayload(resp)
}
