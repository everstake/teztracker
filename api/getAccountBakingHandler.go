package api

import (
	"github.com/everstake/teztracker/api/render"
	"github.com/everstake/teztracker/gen/restapi/operations/accounts"
	"github.com/everstake/teztracker/repos"
	"github.com/everstake/teztracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

type getAccountBakingListHandler struct {
	provider DbProvider
}

// Handle serves the Get Account Baking List request.
func (h *getAccountBakingListHandler) Handle(params accounts.GetAccountBakingListParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetAccountBakingListBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return accounts.NewGetAccountBakingListBadRequest()
	}
	service := services.New(repos.New(db), net)
	limiter := NewLimiter(params.Limit, params.Offset)

	count, accs, err := service.GetAccountBakingList(params.AccountID, limiter)
	if err != nil {
		logrus.Errorf("failed to get account baking: %s", err.Error())
		return accounts.NewGetAccountBakingListNotFound()
	}
	return accounts.NewGetAccountBakingListOK().WithPayload(render.AccountBakingList(accs)).WithXTotalCount(count)
}

type getAccountBakedBlocksListHandler struct {
	provider DbProvider
}

// Handle serves the Get Account baked blocks List request.
func (h *getAccountBakedBlocksListHandler) Handle(params accounts.GetAccountBakedBlocksListParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetAccountBakedBlocksListBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return accounts.NewGetAccountBakedBlocksListBadRequest()
	}
	service := services.New(repos.New(db), net)
	limiter := NewLimiter(params.Limit, params.Offset)

	count, accs, err := service.GetAccountBakedBlocksList(params.AccountID, params.CycleID, limiter)
	if err != nil {
		logrus.Errorf("failed to get accounts: %s", err.Error())
		return accounts.NewGetAccountBakedBlocksListNotFound()
	}
	return accounts.NewGetAccountBakedBlocksListOK().WithPayload(render.Blocks(accs)).WithXTotalCount(count)
}

type getAccountTotalBakingHandler struct {
	provider DbProvider
}

func (h *getAccountTotalBakingHandler) Handle(params accounts.GetAccountTotalBakingParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetAccountTotalBakingBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return accounts.NewGetAccountTotalBakingBadRequest()
	}
	service := services.New(repos.New(db), net)

	total, err := service.GetAccountBakingTotal(params.AccountID)
	if err != nil {
		logrus.Errorf("failed to get accounts: %s", err.Error())
		return accounts.NewGetAccountTotalBakingNotFound()
	}

	return accounts.NewGetAccountTotalBakingOK().WithPayload(render.AccountBaking(total))
}

type getAccountFutureBakingHandler struct {
	provider DbProvider
}

func (h *getAccountFutureBakingHandler) Handle(params accounts.GetAccountFutureBakingParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetAccountFutureBakingBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return accounts.NewGetAccountFutureBakingBadRequest()
	}
	service := services.New(repos.New(db), net)

	total, err := service.GetAccountFutureBakingList(params.AccountID)
	if err != nil {
		logrus.Errorf("failed to get accounts: %s", err.Error())
		return accounts.NewGetAccountFutureBakingNotFound()
	}

	return accounts.NewGetAccountFutureBakingOK().WithPayload(render.AccountBakingList(total))
}

type getAccountFutureBakingRightsHandler struct {
	provider DbProvider
}

func (h *getAccountFutureBakingRightsHandler) Handle(params accounts.GetAccountFutureBakingRightsByCycleParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetAccountFutureBakingRightsByCycleBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return accounts.NewGetAccountFutureBakingRightsByCycleBadRequest()
	}
	service := services.New(repos.New(db), net)
	limiter := NewLimiter(params.Limit, params.Offset)

	count, total, err := service.GetAccountFutureBakingRights(params.AccountID, params.CycleID, limiter)
	if err != nil {
		logrus.Errorf("failed to get future baking rights: %s", err.Error())
		return accounts.NewGetAccountFutureBakingNotFound()
	}

	return accounts.NewGetAccountFutureBakingRightsByCycleOK().WithPayload(render.BakingRights(total)).WithXTotalCount(count)
}
