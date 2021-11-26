package api

import (
	"github.com/everstake/teztracker/api/render"
	"github.com/everstake/teztracker/gen/restapi/operations/accounts"
	"github.com/everstake/teztracker/gen/restapi/operations/blocks"
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/repos"
	"github.com/everstake/teztracker/services"
	"github.com/go-openapi/runtime/middleware"
	log "github.com/sirupsen/logrus"
	"time"
)


type getLostBlocksAggCount struct {
	provider DbProvider
}

func (h *getLostBlocksAggCount) Handle(params blocks.GetLostBlocksAggCountParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetLostBlocksAggCountBadRequest()
	}

	db, err := h.provider.GetDb(net)
	if err != nil {
		return accounts.NewGetLostBlocksAggCountNotFound()
	}

	filter := models.AggTimeFilter{Period: params.Period}
	if params.From != nil && *params.From > 0 {
		filter.From = time.Unix(*params.From, 0)
	}
	if params.To != nil && *params.To > 0 {
		filter.To = time.Unix(*params.To, 0)
	}
	service := services.New(repos.New(db), net)
	resp, err := service.GetLostBlocksCountAgg(filter)
	if err != nil {
		log.Errorf("GetLostBlocksCountAgg error: %s", err)
		return blocks.NewGetLostBlocksAggCountInternalServerError()
	}

	return blocks.NewGetLostBlocksAggCountOK().WithPayload(render.AggTimeInt(resp))
}

type getLostEndorsementsAggCount struct {
	provider DbProvider
}

func (h *getLostEndorsementsAggCount) Handle(params blocks.GetLostEndorsermentsAggCountParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetLostEndorsermentsAggCountBadRequest()
	}

	db, err := h.provider.GetDb(net)
	if err != nil {
		return accounts.NewGetLostEndorsermentsAggCountNotFound()
	}

	filter := models.AggTimeFilter{Period: params.Period}
	if params.From != nil && *params.From > 0 {
		filter.From = time.Unix(*params.From, 0)
	}
	if params.To != nil && *params.To > 0 {
		filter.To = time.Unix(*params.To, 0)
	}
	service := services.New(repos.New(db), net)
	resp, err := service.GetLostEndorsingCountAgg(filter)
	if err != nil {
		log.Errorf("GetLostEndorsingCountAgg error: %s", err)
		return blocks.NewGetLostEndorsermentsAggCountInternalServerError()
	}

	return blocks.NewGetLostEndorsermentsAggCountOK().WithPayload(render.AggTimeInt(resp))
}


type getLostRewardsAggCount struct {
	provider DbProvider
}

func (h *getLostRewardsAggCount) Handle(params blocks.GetLostRewardsAggParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetLostRewardsAggBadRequest()
	}

	db, err := h.provider.GetDb(net)
	if err != nil {
		return accounts.NewGetLostRewardsAggNotFound()
	}

	service := services.New(repos.New(db), net)
	resp, err := service.GetLostRewards(params.Period)
	if err != nil {
		log.Errorf("GetLostRewards error: %s", err)
		return blocks.NewGetLostRewardsAggInternalServerError()
	}

	return blocks.NewGetLostRewardsAggOK().WithPayload(render.AggTimeInt(resp))
}

