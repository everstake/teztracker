package api

import (
	"github.com/everstake/teztracker/gen/restapi/operations/accounts"
	"github.com/everstake/teztracker/repos"
	"github.com/everstake/teztracker/services"
	"github.com/go-openapi/runtime/middleware"
	log "github.com/sirupsen/logrus"
)

type getAccountReportHandler struct {
	provider DbProvider
}

func (h *getAccountReportHandler) Handle(params accounts.GetAccountReportParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetAccountReportBadRequest()
	}

	db, err := h.provider.GetDb(net)
	if err != nil {
		return accounts.NewGetAccountReportNotFound()
	}

	service := services.New(repos.New(db), net)

	resp, err := service.GetAccountReport(params.AccountID, params.From, params.To, params.OperationType)
	if err != nil {
		log.Errorf("GetAccountReport error: %s", err)
		return accounts.NewGetAccountReportInternalServerError()
	}

	return accounts.NewGetAccountReportOK().WithPayload(resp)
}
