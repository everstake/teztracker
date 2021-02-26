package api

import (
	"github.com/everstake/teztracker/api/render"
	"github.com/everstake/teztracker/gen/restapi/operations/accounts"
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/repos"
	"github.com/everstake/teztracker/services"
	"github.com/go-openapi/runtime/middleware"
	log "github.com/sirupsen/logrus"
	"time"
)

type getAccountsAggCount struct {
	provider DbProvider
}

func (h *getAccountsAggCount) Handle(params accounts.GetAccountsAggCountParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetAccountsAggCountBadRequest()
	}

	db, err := h.provider.GetDb(net)
	if err != nil {
		return accounts.NewGetAccountsAggCountNotFound()
	}

	service := services.New(repos.New(db), net)

	filter := models.AggTimeFilter{
		From:   time.Unix(params.From, 0),
		Period: params.Period,
	}
	if params.To != nil && *params.To > 0 {
		filter.To = time.Unix(*params.To, 0)
	}
	resp, err := service.GetAccountCountByPeriod(filter)
	if err != nil {
		log.Errorf("GetAccountCountByPeriod error: %s", err)
		return accounts.NewGetAccountsAggCountInternalServerError()
	}

	return accounts.NewGetAccountsAggCountOK().WithPayload(render.AggTimeInt(resp))
}
