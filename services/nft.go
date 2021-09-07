package services

import (
	"github.com/everstake/teztracker/models"
)

// GetBlockEndorsements finds a block and returns endorsements for it.
func (t *TezTracker) GetNFTContracts(limits Limiter) (contracts []models.NFTContract, count int64, err error) {

	contracts, count, err = t.repoProvider.GetNFT().NTFContractsList("", limits.Limit(), limits.Offset())
	if err != nil {
		return contracts, 0, err
	}

	return contracts, count, nil
}

func (t *TezTracker) GetNFTContractOperations(contractID string, limits Limiter) (operations []models.Operation, count int64, err error) {
	repo := t.repoProvider.GetOperation()

	operations, count, err = repo.ContractOperationsList(contractID, []string{"transaction"}, []string{"mint", "buy", "list_token", "transfer"}, 0, limits.Offset(), limits.Limit(), models.SortDesc)
	if err != nil {
		return operations, 0, err
	}

	return operations, 0, err
}

func (t *TezTracker) GetNFTContractOperationsChart(contractID, period string, from, to int64) (chart []models.ChartData, err error) {
	nftContracts, _, err := t.repoProvider.GetNFT().NTFContractsList(contractID, 1, 0)
	if err != nil {
		return nil, err
	}

	if len(nftContracts) != 1 {
		return nil, ErrNotFound
	}

	chart, err = t.repoProvider.GetChart().OperationsNumber(from, to, period, nftContracts[0].AccountId, []string{"transaction"}, []string{"buy", "list_token", "delist_token", "register_auction", "end_auction", "swap"})
	if err != nil {
		return nil, err
	}

	return chart, nil
}

func (t *TezTracker) GetNFTContract(contractID string) (contract models.NFTContract, err error) {

	contracts, _, err := t.repoProvider.GetNFT().NTFContractsList(contractID, 1, 0)
	if err != nil {
		return contract, err
	}

	if len(contracts) == 0 {
		return contract, ErrNotFound
	}

	return contracts[0], nil
}

func (t *TezTracker) GetNFTContractTokens(contractID string, limits Limiter) (tokens []models.NFTToken, count int64, err error) {
	nftRepo := t.repoProvider.GetNFT()
	contracts, _, err := nftRepo.NTFContractsList(contractID, 1, 0)
	if err != nil {
		return tokens, 0, err
	}

	if len(contracts) == 0 {
		return tokens, 0, ErrNotFound
	}

	tokens, count, err = nftRepo.TokensList(contracts[0].ID, nil, limits.Limit(), limits.Offset())
	if err != nil {
		return tokens, 0, err
	}

	return tokens, count, nil
}

func (t *TezTracker) GetNFTContractDistribution(contractID string) (distribution models.NFTDistribution, err error) {
	nftRepo := t.repoProvider.GetNFT()
	contracts, _, err := nftRepo.NTFContractsList(contractID, 1, 0)
	if err != nil {
		return distribution, err
	}

	if len(contracts) == 0 {
		return distribution, ErrNotFound
	}

	holders, totalCount, err := nftRepo.ContractTokenHolders(contracts[0].LedgerBigMap, 100)
	if err != nil {
		return distribution, err
	}

	_, tokensCount, err := nftRepo.TokensList(contracts[0].ID, nil, 1, 0)
	if err != nil {
		return distribution, err
	}

	distribution = models.NFTDistribution{
		Holders:          holders,
		UniqueHoldersNum: totalCount,
		TokenNum:         tokensCount,
	}

	return distribution, nil
}

func (t *TezTracker) GetNFTContractOwnership(contractID string) (ownership models.NFTContractOwnership, err error) {
	nftRepo := t.repoProvider.GetNFT()

	contracts, _, err := nftRepo.NTFContractsList(contractID, 1, 0)
	if err != nil {
		return ownership, err
	}

	if len(contracts) == 0 {
		return ownership, ErrNotFound
	}

	_, totalHoldersCount, err := nftRepo.ContractTokenHolders(contracts[0].LedgerBigMap, 1)
	if err != nil {
		return ownership, err
	}

	multiTokenHolders, err := nftRepo.TokenHoldersCount(contracts[0].LedgerBigMap, 1, false)
	if err != nil {
		return ownership, err
	}

	whaleTokenHolders, err := nftRepo.TokenHoldersCount(contracts[0].LedgerBigMap, 10, false)
	if err != nil {
		return ownership, err
	}

	singleTokenHolders, err := nftRepo.TokenHoldersCount(contracts[0].LedgerBigMap, 1, true)
	if err != nil {
		return ownership, err
	}

	return models.NFTContractOwnership{
		UniqueHoldersNum:   totalHoldersCount,
		SingleTokenHolders: singleTokenHolders,
		MultiTokenHolders:  multiTokenHolders,
		WhaleTokenHolders:  whaleTokenHolders,
	}, nil
}

func (t *TezTracker) GetNFTContractToken(contractID string, tokenID int64) (token models.NFTToken, err error) {
	nftRepo := t.repoProvider.GetNFT()
	contracts, _, err := nftRepo.NTFContractsList(contractID, 1, 0)
	if err != nil {
		return token, err
	}

	if len(contracts) == 0 {
		return token, ErrNotFound
	}

	tokens, _, err := nftRepo.TokensList(contracts[0].ID, &tokenID, 1, 0)
	if err != nil {
		return token, err
	}

	if len(tokens) != 1 {
		return token, ErrNotFound
	}

	return tokens[0], nil
}

func (t *TezTracker) GetNFTHolders(contractID string, tokenID int64, limiter Limiter) (holders []models.AssetHolder, count int64, err error) {
	nftRepo := t.repoProvider.GetNFT()

	contracts, _, err := nftRepo.NTFContractsList(contractID, 1, 0)
	if err != nil {
		return holders, 0, err
	}

	if len(contracts) == 0 {
		return holders, 0, ErrNotFound
	}

	holders, count, err = nftRepo.TokenHoldersList(contracts[0].LedgerBigMap, &tokenID, limiter.Limit(), limiter.Offset())
	if err != nil {
		return holders, 0, err
	}

	return orderAndTruncateHoldersList(holders), count, nil
}
