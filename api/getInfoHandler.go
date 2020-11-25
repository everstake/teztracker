package api

import (
	"fmt"
	"github.com/everstake/teztracker/api/render"
	info "github.com/everstake/teztracker/gen/restapi/operations/app_info"
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/repos"
	"github.com/everstake/teztracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"time"
)

type getInfoHandler struct {
	provider   models.MarketDataProvider
	dbProvider DbProvider
	cache      *cache.Cache
}

const (
	stakingRatioCacheKey = "staking_ratio_%s"
	cacheTTL             = 2 * time.Minute
)

// Handle serves the Get Info request.
func (h *getInfoHandler) Handle(params info.GetInfoParams) middleware.Responder {
	md, err := h.provider.GetTezosMarketData(*params.Currency)
	if err != nil {
		logrus.Errorf("failed to get market data: %s", err.Error())
		return info.NewGetInfoInternalServerError()
	}
	net, err := ToNetwork(params.Network)
	if err != nil {
		return info.NewGetInfoBadRequest()
	}
	db, err := h.dbProvider.GetDb(net)
	if err != nil {
		return info.NewGetInfoInternalServerError()
	}

	service := services.New(repos.New(db), net)

	ratio, isFound := h.cache.Get(fmt.Sprintf(stakingRatioCacheKey, net))
	if !isFound {
		ratio, err = service.GetStakingRatio()
		if err != nil {
			logrus.Errorf("failed to get staking ratio: %s", err.Error())
		} else {
			h.cache.Set(fmt.Sprintf(stakingRatioCacheKey, net), ratio, cacheTTL)
		}
	}

	return info.NewGetInfoOK().WithPayload(render.Info(*params.Currency, md, ratio.(float64), service.BlocksInCycle()))
}
