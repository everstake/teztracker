package services

import (
	"github.com/everstake/teztracker/models"
)

func (t *TezTracker) TokensList(limiter Limiter) (assets []models.AssetInfo, err error) {
	r := t.repoProvider.GetAssets()

	assets, err = r.GetTokensList()
	if err != nil {
		return assets, err
	}

	return assets, nil
}

func (t *TezTracker) TokenInfo(assetID string) (info models.AssetInfo, err error) {

	r := t.repoProvider.GetAssets()

	info, err = r.GetTokenInfo(assetID)
	if err != nil {
		return info, err
	}

	return info, nil
}

func (t *TezTracker) TokenOperations(assetID string, operationsType string, limits Limiter) (operations []models.AssetOperationReport, err error) {

	r := t.repoProvider.GetAssets()

	info, err := r.GetTokenInfo(assetID)
	if err != nil {
		return operations, err
	}

	var isTransfer bool
	if operationsType == "transfer" {
		isTransfer = true
	}

	operations, err = r.GetAssetOperations(info.ID, isTransfer, limits.Limit(), limits.Offset())
	if err != nil {
		return operations, err
	}

	return operations, nil
}

func (t *TezTracker) TokenHolders(assetID string) (holders []models.AssetHolder, err error) {

	r := t.repoProvider.GetAssets()
	holders, err = r.GetTokenHolders(assetID)
	if err != nil {
		return nil, err
	}

	return holders, nil
}
