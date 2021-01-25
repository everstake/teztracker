package api

import (
	"github.com/everstake/teztracker/api/render"
	"github.com/everstake/teztracker/gen/restapi/operations/assets"
	"github.com/everstake/teztracker/repos"
	"github.com/everstake/teztracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

type getAssetsListHandler struct {
	provider DbProvider
}

func (h *getAssetsListHandler) Handle(params assets.GetAssetsListParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return assets.NewGetAssetsListBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return assets.NewGetAssetsListBadRequest()
	}
	service := services.New(repos.New(db), net)

	total, tokens, err := service.TokensList(NewLimiter(params.Limit, params.Offset))
	if err != nil {
		logrus.Errorf("failed to get tokens list: %s", err.Error())
		return assets.NewGetAssetsListNotFound()
	}

	return assets.NewGetAssetsListOK().WithXTotalCount(total).WithPayload(render.AssetsList(tokens))
}

type getAssetInfoHandler struct {
	provider DbProvider
}

func (h *getAssetInfoHandler) Handle(params assets.GetAssetTokenInfoParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return assets.NewGetAssetTokenInfoBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return assets.NewGetAssetTokenInfoBadRequest()
	}
	service := services.New(repos.New(db), net)

	total, err := service.TokenInfo(params.AssetID)
	if err != nil {
		logrus.Errorf("failed to get token info: %s", err.Error())
		return assets.NewGetAssetTokenInfoNotFound()
	}

	return assets.NewGetAssetTokenInfoOK().WithPayload(render.AssetInfo(total))
}

type getAssetOperationListHandler struct {
	provider DbProvider
}

func (h *getAssetOperationListHandler) Handle(params assets.GetAssetOperationsListParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return assets.NewGetAssetOperationsListBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return assets.NewGetAssetOperationsListBadRequest()
	}
	service := services.New(repos.New(db), net)

	count, ops, err := service.TokenOperations(params.AssetID, params.Type, params.AccountID, params.BlockLevel, NewLimiter(params.Limit, params.Offset))
	if err != nil {
		logrus.Errorf("failed to get token operations: %s", err.Error())
		return assets.NewGetAssetOperationsListNotFound()
	}

	return assets.NewGetAssetOperationsListOK().WithXTotalCount(count).WithPayload(render.AssetOperations(ops))
}

type getAssetHoldersHandler struct {
	provider DbProvider
}

func (h *getAssetHoldersHandler) Handle(params assets.GetAssetTokenHoldersListParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return assets.NewGetAssetTokenHoldersListBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return assets.NewGetAssetTokenHoldersListBadRequest()
	}
	service := services.New(repos.New(db), net)

	total, err := service.TokenHolders(params.AssetID)
	if err != nil {
		logrus.Errorf("failed to get token holders: %s", err.Error())
		return assets.NewGetAssetTokenHoldersListNotFound()
	}

	return assets.NewGetAssetTokenHoldersListOK().WithPayload(render.AssetHolders(total))
}
