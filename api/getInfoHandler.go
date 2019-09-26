package api

import (
	"github.com/bullblock-io/tezTracker/api/render"
	info "github.com/bullblock-io/tezTracker/gen/restapi/operations/app_info"
	"github.com/bullblock-io/tezTracker/models"
	"github.com/bullblock-io/tezTracker/repos"
	"github.com/bullblock-io/tezTracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// MarketDataProvider is an interface for getting actual price and price changes.
type MarketDataProvider interface {
	GetTezosMarketData() (md models.MarketInfo, err error)
}

type getInfoHandler struct {
	provider MarketDataProvider
	db       *gorm.DB
}

// Handle serves the Get Info request.
func (h *getInfoHandler) Handle(params info.GetInfoParams) middleware.Responder {
	md, err := h.provider.GetTezosMarketData()
	if err != nil {
		logrus.Errorf("failed to get market data: %s", err.Error())
		return info.NewGetInfoInternalServerError()
	}
	service := services.New(repos.New(h.db))
	ratio, err := service.GetStakingRatio()
	if err != nil {
		logrus.Errorf("failed to get staking ratio: %s", err.Error())
	}
	return info.NewGetInfoOK().WithPayload(render.Info(md, ratio))
}
