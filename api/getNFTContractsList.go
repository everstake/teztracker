package api

import (
	"log"
	"strconv"

	"github.com/everstake/teztracker/api/render"
	nft "github.com/everstake/teztracker/gen/restapi/operations/n_f_t"
	"github.com/everstake/teztracker/repos"
	"github.com/everstake/teztracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

type getNFTContractsListHandler struct {
	provider DbProvider
}

func (h *getNFTContractsListHandler) Handle(params nft.GetNFTContractsListParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return nft.NewGetNFTContractsListBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return nft.NewGetNFTContractsListNotFound()
	}
	service := services.New(repos.New(db), net)
	limiter := NewLimiter(params.Limit, params.Offset)

	payload, count, err := service.GetNFTContracts(limiter)
	if err != nil {
		logrus.Errorf("failed to get NFT contracts: %s", err.Error())
		return nft.NewGetNFTContractsListInternalServerError()
	}

	return nft.NewGetNFTContractsListOK().WithPayload(render.NFTContracts(payload)).WithXTotalCount(count)
}

type getNFTContractHandler struct {
	provider DbProvider
}

func (h *getNFTContractHandler) Handle(params nft.GetNFTContractParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return nft.NewGetNFTContractBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return nft.NewGetNFTContractNotFound()
	}
	service := services.New(repos.New(db), net)

	payload, err := service.GetNFTContract(params.ContractID)
	if err != nil {
		if err == services.ErrNotFound {
			return nft.NewGetNFTContractNotFound()
		}
		logrus.Errorf("failed to get NFT contract: %s", err.Error())
		return nft.NewGetNFTContractInternalServerError()
	}

	return nft.NewGetNFTContractOK().WithPayload(render.NFTContract(payload))
}

type getNFTContractOperationsListHandler struct {
	provider DbProvider
}

func (h *getNFTContractOperationsListHandler) Handle(params nft.GetNFTContractOperationsParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return nft.NewGetNFTContractOperationsBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return nft.NewGetNFTContractOperationsNotFound()
	}

	service := services.New(repos.New(db), net)
	limiter := NewLimiter(params.Limit, params.Offset)

	payload, count, err := service.GetNFTContractOperations(params.ContractID, limiter)
	if err != nil {
		if err == services.ErrNotFound {
			return nft.NewGetNFTContractOperationsNotFound()
		}
		logrus.Errorf("failed to get NFT contract operations: %s", err.Error())
		return nft.NewGetNFTContractOperationsInternalServerError()
	}

	return nft.NewGetNFTContractOperationsOK().WithXTotalCount(count).WithPayload(render.Operations(payload))
}

type getNFTContractOperationsChartHandler struct {
	provider DbProvider
}

func (h *getNFTContractOperationsChartHandler) Handle(params nft.GetNFTContractOperationsChartParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return nft.NewGetNFTContractOperationsChartBadRequest()
	}

	db, err := h.provider.GetDb(net)
	if err != nil {
		return nft.NewGetNFTContractOperationsChartBadRequest()
	}

	service := services.New(repos.New(db), net)

	chart, err := service.GetNFTContractOperationsChart(params.ContractID, params.Period, params.From, params.To)
	if err != nil {
		logrus.Errorf("failed to get NFT contract operations chart: %s", err.Error())
		return nft.NewGetNFTContractOperationsChartInternalServerError()
	}

	return nft.NewGetNFTContractOperationsChartOK().WithPayload(render.ChartData(chart))
}

type getNFTContractDistributionHandler struct {
	provider DbProvider
}

func (h *getNFTContractDistributionHandler) Handle(params nft.GetNFTContractDistributionParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return nft.NewGetNFTContractDistributionBadRequest()
	}

	db, err := h.provider.GetDb(net)
	if err != nil {
		return nft.NewGetNFTContractDistributionNotFound()
	}

	service := services.New(repos.New(db), net)

	dist, err := service.GetNFTContractDistribution(params.ContractID)
	if err != nil {
		logrus.Errorf("failed to get NFT contract distribution: %s", err.Error())
		return nft.NewGetNFTContractDistributionInternalServerError()
	}

	return nft.NewGetNFTContractDistributionOK().WithPayload(render.NFTDistribution(dist))
}

type getNFTContractOwnershipHandler struct {
	provider DbProvider
}

func (h *getNFTContractOwnershipHandler) Handle(params nft.GetNFTContractOwnershipParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return nft.NewGetNFTContractOwnershipBadRequest()
	}

	db, err := h.provider.GetDb(net)
	if err != nil {
		return nft.NewGetNFTContractOwnershipNotFound()
	}

	service := services.New(repos.New(db), net)

	ownership, err := service.GetNFTContractOwnership(params.ContractID)
	if err != nil {
		logrus.Errorf("failed to get NFT contract ownership: %s", err.Error())
		return nft.NewGetNFTContractOwnershipInternalServerError()
	}

	log.Print("ownership ", ownership)

	return nft.NewGetNFTContractOwnershipOK().WithPayload(render.NFTOwnership(ownership))
}

type getNFTContractTokensListHandler struct {
	provider DbProvider
}

func (h *getNFTContractTokensListHandler) Handle(params nft.GetNFTContractTokensListParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return nft.NewGetNFTContractTokensListBadRequest()
	}

	db, err := h.provider.GetDb(net)
	if err != nil {
		return nft.NewGetNFTContractTokensListNotFound()
	}

	service := services.New(repos.New(db), net)
	limiter := NewLimiter(params.Limit, params.Offset)

	payload, count, err := service.GetNFTContractTokens(params.ContractID, limiter)
	if err != nil {
		logrus.Errorf("failed to get NFT contract tokens: %s", err.Error())
		return nft.NewGetNFTContractsListInternalServerError()
	}

	return nft.NewGetNFTContractTokensListOK().WithPayload(render.NFTTokens(payload)).WithXTotalCount(count)
}

type getNFTContractTokenHandler struct {
	provider DbProvider
}

func (h *getNFTContractTokenHandler) Handle(params nft.GetNFTContractTokenParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return nft.NewGetNFTContractTokenBadRequest()
	}

	db, err := h.provider.GetDb(net)
	if err != nil {
		return nft.NewGetNFTContractTokenNotFound()
	}

	tokenID, err := strconv.ParseInt(params.TokenID, 10, 64)
	if err != nil {
		return nft.NewGetNFTContractTokenBadRequest()
	}

	service := services.New(repos.New(db), net)

	payload, err := service.GetNFTContractToken(params.ContractID, tokenID)
	if err != nil {
		logrus.Errorf("failed to get NFT contract token: %s", err.Error())
		return nft.NewGetNFTContractTokenInternalServerError()
	}

	return nft.NewGetNFTContractTokenOK().WithPayload(render.NFTToken(payload))
}

type getNFTContractTokenHoldersHandler struct {
	provider DbProvider
}

func (h *getNFTContractTokenHoldersHandler) Handle(params nft.GetNFTContractTokenHoldersParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return nft.NewGetNFTContractTokenHoldersBadRequest()
	}

	db, err := h.provider.GetDb(net)
	if err != nil {
		return nft.NewGetNFTContractTokenHoldersNotFound()
	}

	tokenID, err := strconv.ParseInt(params.TokenID, 10, 64)
	if err != nil {
		return nft.NewGetNFTContractTokenHoldersBadRequest()
	}

	service := services.New(repos.New(db), net)

	limiter := NewLimiter(params.Limit, params.Offset)

	payload, count, err := service.GetNFTHolders(params.ContractID, tokenID, limiter)
	if err != nil {
		logrus.Errorf("failed to get NFT token holders: %s", err.Error())
		return nft.NewGetNFTContractTokenHoldersInternalServerError()
	}

	return nft.NewGetNFTContractTokenHoldersOK().WithPayload(render.AssetHolders(payload)).WithXTotalCount(count)
}
