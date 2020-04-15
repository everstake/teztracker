package services

import "github.com/everstake/teztracker/models"

func (t *TezTracker) GetAccountBakingList(accountID string, limits Limiter) (count int64, list []models.AccountBaking, err error) {
	count, list, err = t.repoProvider.GetAccount().BakingList(accountID, limits.Limit(), limits.Offset())
	if err != nil {
		return 0, nil, err
	}

	for i := range list {
		list[i].Status = models.StatusPending
		list[i].TotalDeposit = (list[i].Stolen + list[i].Count) * BlockSecurityDeposit
	}

	return count, list, nil
}

func (t *TezTracker) GetAccountFutureBakingList(accountID string) (list []models.AccountBaking, err error) {
	lastBlock, err := t.repoProvider.GetBlock().Last()
	if err != nil {
		return list, err
	}

	list, err = t.repoProvider.GetAccount().FutureBakingList(accountID)
	if err != nil {
		return nil, err
	}

	for i := range list {
		list[i].Status = getRewardStatus(list[i].Cycle, lastBlock.MetaCycle)
		list[i].TotalDeposit = (list[i].Stolen + list[i].Count) * BlockSecurityDeposit
	}

	return list, nil
}

func (t *TezTracker) GetAccountBakedBlocksList(accountID string, cycle int64, limits Limiter) (count int64, list []models.Block, err error) {
	count, list, err = t.repoProvider.GetBlock().BakedBlocksList(accountID, cycle, limits.Limit(), limits.Offset())
	if err != nil {
		return 0, nil, err
	}

	return count, list, nil
}

func (t *TezTracker) GetAccountBakingTotal(accountID string) (total models.AccountBaking, err error) {
	total, err = t.repoProvider.GetAccount().BakingTotal(accountID)
	if err != nil {
		return total, err
	}

	total.TotalDeposit = (total.Stolen + total.Count) * BlockSecurityDeposit

	return total, nil
}
