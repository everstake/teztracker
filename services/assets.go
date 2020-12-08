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

	for i := range assets {

		assets[i].Balance, err = t.TokenTotalSupply(assets[i].AccountId)
		if err != nil {
			return count, assets, err
		}
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

	info.Balance, err = t.TokenTotalSupply(info.AccountId)
	if err != nil {
		return info, err
	}

	return info, nil
}

func (t *TezTracker) TokenTotalSupply(assetID string) (totalSupply int64, err error) {
	r := t.repoProvider.GetAssets()

	h, err := r.GetTokenHolders(assetID)
	if err != nil {
		return 0, err
	}

	for j := range h {
		totalSupply += int64(h[j].Balance)
	}

	return totalSupply, nil
}

func (t *TezTracker) TokenOperations(assetIDs, operationsTypes, accountIDs []string, limits Limiter) (count int64, operations []models.AssetOperationReport, err error) {

	count, operations, err = t.repoProvider.GetAssets().GetAssetOperations(assetIDs, operationsTypes, accountIDs, limits.Limit(), limits.Offset())
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
