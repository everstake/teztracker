package services

import (
	"github.com/everstake/teztracker/models"
)

func (t *TezTracker) GetAccountBakingList(accountID string, limits Limiter) (count int64, list []models.AccountBaking, err error) {
	lastBlock, err := t.repoProvider.GetBlock().Last()
	if err != nil {
		return 0, nil, err
	}

	count, list, err = t.repoProvider.GetBaking().BakingList(accountID, limits.Limit(), limits.Offset())
	if err != nil {
		return 0, nil, err
	}

	for i := range list {
		list[i].Status = getRewardStatus(list[i].Cycle, lastBlock.MetaCycle)
		list[i].TotalDeposit = (list[i].Stolen + list[i].Count) * getBlockSecurityDepositByCycle(list[i].Cycle)
	}

	return count, list, nil
}

func (t *TezTracker) GetAccountFutureBakingList(accountID string) (list []models.AccountBaking, err error) {
	lastBlock, err := t.repoProvider.GetBlock().Last()
	if err != nil {
		return list, err
	}

	list, err = t.repoProvider.GetBaking().FutureBakingList(accountID)
	if err != nil {
		return nil, err
	}

	for i := range list {
		list[i].Status = getRewardStatus(list[i].Cycle, lastBlock.MetaCycle)
		list[i].TotalDeposit = list[i].Count * getBlockSecurityDepositByCycle(list[i].Cycle)
		list[i].Reward = list[i].Count * getBlockRewardByCycle(list[i].Cycle, 0)
	}

	return list, nil
}

func (t *TezTracker) GetAccountBakedBlocksList(accountID string, cycle int64, limits Limiter) (count int64, list []models.Block, err error) {
	count, list, err = t.repoProvider.GetBlock().BakedBlocksList(accountID, cycle, limits.Limit(), limits.Offset())
	if err != nil {
		return 0, nil, err
	}

	for i := range list {
		list[i].Deposit = getBlockSecurityDepositByCycle(list[i].MetaCycle)
	}

	return count, list, nil
}

func (t *TezTracker) GetAccountBakingTotal(accountID string) (total models.AccountBaking, err error) {
	total, err = t.repoProvider.GetBaking().BakingTotal(accountID)
	if err != nil {
		return total, err
	}

	total.TotalDeposit = (total.Stolen + total.Count) * getBlockSecurityDepositByCycle(total.Cycle)

	return total, nil
}

func getBlockSecurityDepositByCycle(cycle int64) int64 {

	if cycle < GranadaCycle {
		return FlorenceBlockSecurityDeposit
	}

	return GranadaBlockSecurityDeposit
}

func getBlockRewardByCycle(cycle, priority int64) int64 {

	if priority == 0 {
		if cycle < GranadaCycle {
			return FlorenceBlockReward
		}
		return GranadaBlockReward
	}

	if cycle < GranadaCycle {
		return FlorenceLowPriorityBlockReward
	}

	return GradanaLowPriorityBlockReward
}

func getEndorsementSecurityDepositByCycle(cycle int64) int64 {

	if cycle < GranadaCycle {
		return FlorenceEndorsementSecurityDeposit
	}

	return GranadaEndorsementSecurityDeposit
}

func getEndorsementRewardByCycle(cycle int64) int64 {

	if cycle < GranadaCycle {
		return FlorenceEndorsementReward
	}

	return GranadaEndorsementReward
}
