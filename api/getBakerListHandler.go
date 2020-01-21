package api

import (
	"github.com/bullblock-io/tezTracker/api/render"
	"github.com/bullblock-io/tezTracker/gen/restapi/operations/accounts"
	"github.com/bullblock-io/tezTracker/repos"
	"github.com/bullblock-io/tezTracker/services"
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
