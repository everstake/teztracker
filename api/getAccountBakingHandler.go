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

	count, accs, err := service.GetAccountBakedBlocksList(params.AccountID, limiter)
	if err != nil {
		logrus.Errorf("failed to get accounts: %s", err.Error())
		return accounts.NewGetAccountBakedBlocksListNotFound()
	}
	return accounts.NewGetAccountBakedBlocksListOK().WithPayload(render.Blocks(accs)).WithXTotalCount(count)
}

type getAccountTotalBakingHandler struct {
	provider DbProvider
}

func (h *getAccountTotalBakingHandler) Handle(params accounts.GetAccountTotalBakingListParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetAccountTotalBakingListBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return accounts.NewGetAccountTotalBakingListBadRequest()
	}
	service := services.New(repos.New(db), net)

	total, err := service.GetAccountBakingTotal(params.AccountID)
	if err != nil {
		logrus.Errorf("failed to get accounts: %s", err.Error())
		return accounts.NewGetAccountTotalBakingListNotFound()
	}

	return accounts.NewGetAccountTotalBakingListOK().WithPayload(render.AccountBaking(total))
}
