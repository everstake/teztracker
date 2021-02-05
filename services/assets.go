package services

import (
	"github.com/everstake/teztracker/models"
)

func (t *TezTracker) TokensList(limiter Limiter) (count int64, assets []models.AssetInfo, err error) {
	r := t.repoProvider.GetAssets()

	count, assets, err = r.GetTokensList()
	if err != nil {
		return count, assets, err
	}

	return count,
		assets, nil
}

func (t *TezTracker) TokenInfo(assetID string) (info models.AssetInfo, err error) {

	r := t.repoProvider.GetAssets()

	info, err = r.GetTokenInfo(assetID)
	if err != nil {
		return info, err
	}

	return info, nil
}

func (t *TezTracker) TokenOperations(assetID string, operationsType string, limits Limiter) (count int64, operations []models.AssetOperationReport, err error) {

	r := t.repoProvider.GetAssets()

	info, err := r.GetTokenInfo(assetID)
	if err != nil {
		return 0, operations, err
	}

	var isTransfer bool
	if operationsType == "transfer" {
		isTransfer = true
	}

	count, operations, err = r.GetAssetOperations(info.ID, isTransfer, limits.Limit(), limits.Offset())
	if err != nil {
		return 0, operations, err
	}

	return count, operations, nil
}

func (t *TezTracker) TokenHolders(assetID string) (holders []models.AssetHolder, err error) {

	r := t.repoProvider.GetAssets()
	holders, err = r.GetTokenHolders(assetID)
	if err != nil {
		return nil, err
	}

	return holders, nil
}

func (t *TezTracker) GetAssetReport(assetID string, from, to int64, operations []string) (resp []byte, err error) {

	return nil, nil
}
