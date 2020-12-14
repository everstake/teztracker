package render

import (
	genModels "github.com/everstake/teztracker/gen/models"
	"github.com/everstake/teztracker/models"
)

func AssetsList(ash []models.AssetInfo) []*genModels.TokenAssetRow {
	ths := make([]*genModels.TokenAssetRow, len(ash))
	for i := range ash {
		ths[i] = AssetInfo(ash[i])
	}
	return ths
}

func AssetInfo(asi models.AssetInfo) *genModels.TokenAssetRow {
	return &genModels.TokenAssetRow{
		AccountID:   asi.AccountId,
		CreatedAt:   asi.Timestamp.Unix(),
		Manager:     asi.Source,
		Name:        asi.Name,
		Ticker:      asi.Ticker,
		Precision:   &asi.Scale,
		TotalSupply: asi.Balance,
	}
}

func AssetHolders(ash []models.AssetHolder) []*genModels.TokenHolderRow {
	ths := make([]*genModels.TokenHolderRow, len(ash))
	for i := range ash {
		ths[i] = AssetHolder(ash[i])
	}
	return ths
}

func AssetHolder(acb models.AssetHolder) *genModels.TokenHolderRow {
	bal := int64(acb.Balance)
	return &genModels.TokenHolderRow{
		AccountID: string(acb.Address),
		Balance:   &bal,
	}
}

func AccountAssetBalances(ash []models.AccountAssetBalance) []*genModels.AccountAssetBalanceRow {
	asb := make([]*genModels.AccountAssetBalanceRow, len(ash))
	for i := range ash {
		asb[i] = AccountAssetBalance(ash[i])
	}
	return asb
}

func AccountAssetBalance(acb models.AccountAssetBalance) *genModels.AccountAssetBalanceRow {
	scale := acb.AssetInfo.Scale
	return &genModels.AccountAssetBalanceRow{
		AccountID: string(acb.Address),
		Balance:   int64(acb.AssetHolder.Balance),
		TokenInfo: &genModels.TokenAssetRow{
			AccountID: acb.AssetInfo.AccountId,
			Name:      acb.AssetInfo.Name,
			Precision: &scale,
			Ticker:    acb.AssetInfo.Ticker,
		},
	}
}

func AssetOperations(ash []models.AssetOperationReport) []*genModels.AssetOperation {
	ths := make([]*genModels.AssetOperation, len(ash))
	for i := range ash {
		ths[i] = AssetOperation(ash[i])
	}
	return ths
}

func AssetOperation(acb models.AssetOperationReport) *genModels.AssetOperation {
	return &genModels.AssetOperation{
		OperationGroupHash: acb.OperationGroupHash,
		From:               acb.Sender,
		To:                 acb.Receiver,
		Type:               acb.Type,
		Amount:             &acb.Amount,
		Fee:                &acb.Fee,
		GasLimit:           &acb.GasLimit,
		StorageLimit:       &acb.StorageLimit,
		Timestamp:          acb.Timestamp.Unix(),
	}
}
