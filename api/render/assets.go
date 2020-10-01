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
		Balance:     &asi.Balance,
		CreatedAt:   asi.Timestamp.Unix(),
		Manager:     asi.Source,
		Name:        asi.Name,
		Precision:   &asi.Scale,
		TotalSupply: 0,
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
	return &genModels.TokenHolderRow{
		AccountID: acb.AccountID,
		Balance:   &acb.Balance,
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
