package render

import (
	genModels "github.com/everstake/teztracker/gen/models"
	"github.com/everstake/teztracker/models"
)

func NFTContracts(с []models.NFTContract) []*genModels.NFTContractRow {
	contracts := make([]*genModels.NFTContractRow, len(с))
	for i := range с {
		contracts[i] = NFTContract(с[i])
	}

	return contracts
}

func NFTContract(c models.NFTContract) *genModels.NFTContractRow {
	return &genModels.NFTContractRow{
		Address:          c.AccountId,
		Description:      c.Description,
		Name:             c.Name,
		NftsNumber:       c.NFTsNumber,
		OperationsNumber: c.OperationsNum,
	}
}

func NFTTokens(tns []models.NFTToken) []*genModels.NFTTokenRow {
	tokens := make([]*genModels.NFTTokenRow, len(tns))
	for i := range tns {
		tokens[i] = NFTToken(tns[i])
	}

	return tokens
}

func NFTToken(t models.NFTToken) *genModels.NFTTokenRow {
	return &genModels.NFTTokenRow{
		Amount:       t.Amount,
		Category:     t.Category,
		Description:  t.Description,
		Decimals:     &t.Decimals,
		CreatedAt:    t.CreatedAt.Unix(),
		IpfsSource:   t.IpfsSource,
		IsForSale:    t.IsForSale,
		IssuedBy:     t.IssuedBy,
		LastPrice:    t.LastPrice,
		LastActiveAt: t.LastActiveAt.Unix(),
		Name:         t.Name,
		TokenID:      int64(t.ID),
	}
}

func NFTDistribution(d models.NFTDistribution) *genModels.NFTContractDistribution {

	return &genModels.NFTContractDistribution{
		Distribution:     AssetHolders(d.Holders),
		TotalTokenNum:    &d.TokenNum,
		UniqueHoldersNum: &d.UniqueHoldersNum,
	}
}

func NFTOwnership(d models.NFTContractOwnership) *genModels.NFTContractOwnership {

	return &genModels.NFTContractOwnership{
		UniqueHoldersNum: &d.UniqueHoldersNum,
		MultiOwners:      &d.MultiTokenHolders,
		SingleOwners:     &d.SingleTokenHolders,
		WhalesCount:      &d.WhaleTokenHolders,
	}
}
