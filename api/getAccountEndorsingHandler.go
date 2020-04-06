package api

import (
	"github.com/everstake/teztracker/api/render"
	"github.com/everstake/teztracker/gen/restapi/operations/accounts"
	"github.com/everstake/teztracker/repos"
	"github.com/everstake/teztracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

type getAccountEndorsingListHandler struct {
	provider DbProvider
}

// Handle serves the Get Account Baking List request.
func (h *getAccountEndorsingListHandler) Handle(params accounts.GetAccountEndorsingListParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetAccountEndorsingListBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return accounts.NewGetAccountEndorsingListBadRequest()
	}
	service := services.New(repos.New(db), net)
	limiter := NewLimiter(params.Limit, params.Offset)

	count, accs, err := service.GetAccountEndorsingList(params.AccountID, limiter)
	if err != nil {
		logrus.Errorf("failed to get account baking: %s", err.Error())
		return accounts.NewGetAccountBakingListNotFound()
	}
	return accounts.NewGetAccountEndorsingListOK().WithPayload(render.AccountEndorsingList(accs)).WithXTotalCount(count)
}

type getAccountEndorsementsHandler struct {
	provider DbProvider
}

// Handle serves the Get Account baked blocks List request.
func (h *getAccountEndorsementsHandler) Handle(params accounts.GetAccountEndorsementsByCycleListParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetAccountEndorsementsByCycleListBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return accounts.NewGetAccountEndorsementsByCycleListBadRequest()
	}
	service := services.New(repos.New(db), net)
	limiter := NewLimiter(params.Limit, params.Offset)

	count, accs, err := service.GetAccountEndorsementsList(params.AccountID, params.CycleID, limiter)
	if err != nil {
		logrus.Errorf("failed to get accounts: %s", err.Error())
		return accounts.NewGetAccountEndorsementsByCycleListNotFound()
	}
	return accounts.NewGetAccountEndorsementsByCycleListOK().WithPayload(render.Operations(accs)).WithXTotalCount(count)
}

type getAccountTotalEndorsingHandler struct {
	provider DbProvider
}

func (h *getAccountTotalEndorsingHandler) Handle(params accounts.GetAccountTotalEndorsingParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetAccountTotalEndorsingBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return accounts.NewGetAccountTotalEndorsingBadRequest()
	}
	service := services.New(repos.New(db), net)

	total, err := service.GetAccountEndorsingTotal(params.AccountID)
	if err != nil {
		logrus.Errorf("failed to get accounts: %s", err.Error())
		return accounts.NewGetAccountTotalEndorsingNotFound()
	}

	return accounts.NewGetAccountTotalEndorsingOK().WithPayload(render.AccountEndorsing(total))
}
