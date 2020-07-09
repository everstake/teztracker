package api

import (
	"github.com/everstake/teztracker/api/render"
	"github.com/everstake/teztracker/gen/restapi/operations/accounts"
	"github.com/everstake/teztracker/repos"
	"github.com/everstake/teztracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

type getBakerListHandler struct {
	provider DbProvider
}

// Handle serves the Get Baker List request.
func (h *getBakerListHandler) Handle(params accounts.GetBakersListParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetBakersListBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return accounts.NewGetBakersListNotFound()
	}
	service := services.New(repos.New(db), net)
	limiter := NewLimiter(params.Limit, params.Offset)

	accs, count, err := service.BakerList(limiter)
	if err != nil {
		logrus.Errorf("failed to get accounts: %s", err.Error())
		return accounts.NewGetBakersListNotFound()
	}
	return accounts.NewGetBakersListOK().WithPayload(render.Bakers(accs)).WithXTotalCount(count)
}

type getPublicBakerListHandler struct {
	provider DbProvider
}

// Handle serves the Get Baker List request.
func (h *getPublicBakerListHandler) Handle(params accounts.GetPublicBakersListParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetBakersListBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return accounts.NewGetBakersListNotFound()
	}
	service := services.New(repos.New(db), net)
	limiter := NewLimiter(params.Limit, params.Offset)

	accs, count, err := service.PublicBakerList(limiter)
	if err != nil {
		logrus.Errorf("failed to get accounts: %s", err.Error())
		return accounts.NewGetBakersListNotFound()
	}

	return accounts.NewGetPublicBakersListOK().WithPayload(render.PublicBakers(accs)).WithXTotalCount(count)
}

type getPublicBakerSearchListHandler struct {
	provider DbProvider
}

// Handle serves the Get Baker List request.
func (h *getPublicBakerSearchListHandler) Handle(params accounts.GetPublicBakersListForSearchParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetPublicBakersListForSearchBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return accounts.NewGetPublicBakersListForSearchBadRequest()
	}
	service := services.New(repos.New(db), net)

	accs, err := service.PublicBakersSearchList()
	if err != nil {
		logrus.Errorf("failed to get public bakers search list: %s", err.Error())
		return accounts.NewGetPublicBakersListForSearchNotFound()
	}

	return accounts.NewGetPublicBakersListForSearchOK().WithPayload(render.PublicBakersSearch(accs))
}

type getBakerSecurityDepositFutureListHandler struct {
	provider DbProvider
}

// Handle serves the Get Baker Security Deposit Future request.
func (h *getBakerSecurityDepositFutureListHandler) Handle(params accounts.GetAccountSecurityDepositListParams) middleware.Responder {

	net, err := ToNetwork(params.Network)
	if err != nil {
		return accounts.NewGetAccountSecurityDepositListBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return accounts.NewGetAccountSecurityDepositListBadRequest()
	}
	service := services.New(repos.New(db), net)

	accs, err := service.GetAccountSecurityDepositList(params.AccountID)
	if err != nil {
		logrus.Errorf("failed to get public bakers search list: %s", err.Error())
		return accounts.NewGetAccountSecurityDepositListNotFound()
	}

	return accounts.NewGetAccountSecurityDepositListOK().WithPayload(render.AccountSecurityDepositList(accs))

}
