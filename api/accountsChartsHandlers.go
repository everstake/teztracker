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

	filter := models.AggTimeFilter{Period: params.Period}
	if params.From != nil && *params.From > 0 {
		filter.From = time.Unix(*params.From, 0)
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

type getAccountsTotalAggCount struct {
	provider DbProvider
}

func (h *getAccountsTotalAggCount) Handle(params accounts.GetAccountsTotalAggCountParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetAccountsTotalAggCountBadRequest()
	}

	db, err := h.provider.GetDb(net)
	if err != nil {
		return accounts.NewGetAccountsTotalAggCountNotFound()
	}

	service := services.New(repos.New(db), net)

	filter := models.AggTimeFilter{Period: params.Period}
	if params.From != nil && *params.From > 0 {
		filter.From = time.Unix(*params.From, 0)
	}
	if params.To != nil && *params.To > 0 {
		filter.To = time.Unix(*params.To, 0)
	}
	resp, err := service.GetTotalAccountCountByPeriod(filter)
	if err != nil {
		log.Errorf("GetTotalAccountCountByPeriod error: %s", err)
		return accounts.NewGetAccountsTotalAggCountInternalServerError()
	}

	return accounts.NewGetAccountsTotalAggCountOK().WithPayload(render.AggTimeInt(resp))
}

type getContractsAggCount struct {
	provider DbProvider
}

func (h *getContractsAggCount) Handle(params accounts.GetContractsAggCountParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetContractsAggCountBadRequest()
	}

	db, err := h.provider.GetDb(net)
	if err != nil {
		return accounts.NewGetContractsAggCountNotFound()
	}

	service := services.New(repos.New(db), net)

	filter := models.AggTimeFilter{Period: params.Period}
	if params.From != nil && *params.From > 0 {
		filter.From = time.Unix(*params.From, 0)
	}
	if params.To != nil && *params.To > 0 {
		filter.To = time.Unix(*params.To, 0)
	}
	resp, err := service.GetContractCountByPeriod(filter)
	if err != nil {
		log.Errorf("GetContractCountByPeriod error: %s", err)
		return accounts.NewGetContractsAggCountInternalServerError()
	}

	return accounts.NewGetContractsAggCountOK().WithPayload(render.AggTimeInt(resp))
}

type getContractsTotalAggCount struct {
	provider DbProvider
}

func (h *getContractsTotalAggCount) Handle(params accounts.GetContractsTotalAggCountParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetContractsTotalAggCountBadRequest()
	}

	db, err := h.provider.GetDb(net)
	if err != nil {
		return accounts.NewGetContractsTotalAggCountNotFound()
	}

	service := services.New(repos.New(db), net)

	filter := models.AggTimeFilter{Period: params.Period}
	if params.From != nil && *params.From > 0 {
		filter.From = time.Unix(*params.From, 0)
	}
	if params.To != nil && *params.To > 0 {
		filter.To = time.Unix(*params.To, 0)
	}
	resp, err := service.GetTotalContractCountByPeriod(filter)
	if err != nil {
		log.Errorf("GetTotalContractCountByPeriod error: %s", err)
		return accounts.NewGetContractsTotalAggCountInternalServerError()
	}

	return accounts.NewGetContractsTotalAggCountOK().WithPayload(render.AggTimeInt(resp))
}

type getActiveAccountsAggCount struct {
	provider DbProvider
}

func (h *getActiveAccountsAggCount) Handle(params accounts.GetActiveAccountsAggCountParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetActiveAccountsAggCountBadRequest()
	}

	db, err := h.provider.GetDb(net)
	if err != nil {
		return accounts.NewGetActiveAccountsAggCountNotFound()
	}

	service := services.New(repos.New(db), net)
	resp, err := service.GetActiveAccounts(params.Period)
	if err != nil {
		log.Errorf("GetActiveAccounts error: %s", err)
		return accounts.NewGetActiveAccountsAggCountInternalServerError()
	}

	return accounts.NewGetActiveAccountsAggCountOK().WithPayload(render.AggTimeInt(resp))
}


type getInactiveAccountsAggCount struct {
	provider DbProvider
}

func (h *getInactiveAccountsAggCount) Handle(params accounts.GetInactiveAccountsAggCountParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetInactiveAccountsAggCountBadRequest()
	}

	db, err := h.provider.GetDb(net)
	if err != nil {
		return accounts.NewGetInactiveAccountsAggCountNotFound()
	}

	filter := models.AggTimeFilter{Period: params.Period}
	if params.From != nil && *params.From > 0 {
		filter.From = time.Unix(*params.From, 0)
	}
	if params.To != nil && *params.To > 0 {
		filter.To = time.Unix(*params.To, 0)
	}
	service := services.New(repos.New(db), net)
	resp, err := service.GetInactiveAccounts(filter)
	if err != nil {
		log.Errorf("GetInactiveAccounts error: %s", err)
		return accounts.NewGetInactiveAccountsAggCountInternalServerError()
	}

	return accounts.NewGetInactiveAccountsAggCountOK().WithPayload(render.AggTimeInt(resp))
}


type getLowBalanceTotalAggCount struct {
	provider DbProvider
}

func (h *getLowBalanceTotalAggCount) Handle(params accounts.GetLowBalanceTotalAggCountParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetLowBalanceTotalAggCountBadRequest()
	}

	db, err := h.provider.GetDb(net)
	if err != nil {
		return accounts.NewGetLowBalanceTotalAggCountNotFound()
	}

	filter := models.AggTimeFilter{Period: params.Period}
	if params.From != nil && *params.From > 0 {
		filter.From = time.Unix(*params.From, 0)
	}
	if params.To != nil && *params.To > 0 {
		filter.To = time.Unix(*params.To, 0)
	}
	service := services.New(repos.New(db), net)
	resp, err := service.GetAccountsWithLowBalance(filter)
	if err != nil {
		log.Errorf("GetAccountsWithLowBalance error: %s", err)
		return accounts.NewGetLowBalanceTotalAggCountInternalServerError()
	}

	return accounts.NewGetLowBalanceTotalAggCountOK().WithPayload(render.AggTimeInt(resp))
}


type getBakersHolding struct {
	provider DbProvider
}

func (h *getBakersHolding) Handle(params accounts.GetBakersHoldingParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetBakersHoldingBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return accounts.NewGetBakersHoldingNotFound()
	}
	service := services.New(repos.New(db), net)
	resp, err := service.GetHoldingPoints()
	if err != nil {
		log.Errorf("GetHoldingPoints error: %s", err)
		return accounts.NewGetBakersHoldingInternalServerError()
	}

	return accounts.NewGetBakersHoldingOK().WithPayload(render.BakersHoldings(resp))
}
