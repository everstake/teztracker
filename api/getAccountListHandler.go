package api

import (
	"github.com/everstake/teztracker/api/render"
	"github.com/everstake/teztracker/gen/restapi/operations/accounts"
	"github.com/everstake/teztracker/repos"
	"github.com/everstake/teztracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	"time"
)

type getAccountListHandler struct {
	provider DbProvider
}

// Handle serves the Get Account List request.
func (h *getAccountListHandler) Handle(params accounts.GetAccountsListParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetAccountsListBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return accounts.NewGetAccountsListNotFound()
	}
	service := services.New(repos.New(db), net)
	limiter := NewLimiter(params.Limit, params.Offset)
	before := ""
	if params.AfterID != nil {
		before = *params.AfterID
	}
	accs, count, err := service.AccountList(before, limiter)
	if err != nil {
		logrus.Errorf("failed to get accounts: %s", err.Error())
		return accounts.NewGetAccountsListNotFound()
	}
	return accounts.NewGetAccountsListOK().WithPayload(render.Accounts(accs)).WithXTotalCount(count)
}

type getAccountBalanceListHandler struct {
	provider DbProvider
}

// Handle serves the Get Account List request.
func (h *getAccountBalanceListHandler) Handle(params accounts.GetAccountBalanceListParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetAccountBalanceListBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return accounts.NewGetAccountBalanceListBadRequest()
	}
	service := services.New(repos.New(db), net)

	if params.From <= 0 || params.To <= 0 || params.To < params.From {
		return accounts.NewGetAccountBalanceListBadRequest()
	}

	from := time.Unix(params.From, 0)
	to := time.Unix(params.To, 0)

	if to.Sub(from) > 24*31*time.Hour {
		return accounts.NewGetAccountBalanceListBadRequest()
	}

	accs, err := service.GetAccountBalanceHistory(params.AccountID, from, to)
	if err != nil {
		logrus.Errorf("failed to get accounts: %s", err.Error())
		return accounts.NewGetAccountsListNotFound()
	}

	return accounts.NewGetAccountBalanceListOK().WithPayload(render.AccountBalances(accs))
}

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
