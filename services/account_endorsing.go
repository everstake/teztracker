package services

import "github.com/everstake/teztracker/models"

func (t *TezTracker) GetAccountEndorsingList(accountID string, limits Limiter) (count int64, list []models.AccountEndorsing, err error) {
	lastBlock, err := t.repoProvider.GetBlock().Last()
	if err != nil {
		return 0, list, err
	}

	count, list, err = t.repoProvider.GetEndorsing().EndorsingList(accountID, limits.Limit(), limits.Offset())
	if err != nil {
		return 0, nil, err
	}

	for i := range list {
		list[i].Status = getRewardStatus(list[i].Cycle, lastBlock.MetaCycle)
		list[i].TotalDeposit = list[i].Count * EndorsementSecurityDeposit
	}

	return count, list, nil
}

func (t *TezTracker) GetAccountEndorsingTotal(accountID string) (total models.AccountEndorsing, err error) {
	total, err = t.repoProvider.GetEndorsing().EndorsingTotal(accountID)
	if err != nil {
		return total, err
	}

	total.TotalDeposit = total.Count * EndorsementSecurityDeposit

	return total, nil
}

func (t *TezTracker) GetAccountEndorsementsList(accountID string, cycle int64, limits Limiter) (count int64, list []models.Operation, err error) {
	count, list, err = t.repoProvider.GetOperation().AccountEndorsements(accountID, cycle, limits.Limit(), limits.Offset())
	if err != nil {
		return 0, nil, err
	}

	for i := range list {
		list[i].EndorsementDeposit = EndorsementSecurityDeposit
	}

	return count, list, nil
}

func (t *TezTracker) GetAccountFutureEndorsementsList(accountID string) (list []models.AccountEndorsing, err error) {
	lastBlock, err := t.repoProvider.GetBlock().Last()
	if err != nil {
		return list, err
	}

	list, err = t.repoProvider.GetEndorsing().FutureEndorsingList(accountID)
	if err != nil {
		return nil, err
	}

	for i := range list {
		list[i].Status = getRewardStatus(list[i].Cycle, lastBlock.MetaCycle)
		list[i].TotalDeposit = list[i].Count * BlockSecurityDeposit
		list[i].Reward = list[i].Count * BlockReward
	}

	return list, nil
}

func getRewardStatus(cycle, currentCycle int64) (status models.RewardStatus) {
	switch {
	case cycle > currentCycle:
		status = models.StatusPending
	case cycle == currentCycle:
		status = models.StatusActive
	case cycle >= currentCycle-PreservedCycles:
		status = models.StatusFrozen
	default:
		status = models.StatusUnfrozen
	}

	return status
}
