package services

import (
	"github.com/everstake/teztracker/models"
)

func (t *TezTracker) GetAccountRewardsList(accountID string, limits Limiter) (count int64, rewards []models.AccountReward, err error) {
	lastBlock, err := t.repoProvider.GetBlock().Last()
	if err != nil {
		return 0, rewards, err
	}

	count, rewards, err = t.repoProvider.GetAccount().RewardsList(accountID, limits.Limit(), limits.Offset())
	if err != nil {
		return 0, nil, err
	}

	var blockReward, endorsementReward int64
	for i := range rewards {
		//Use FutureBakingRewards for future cycles
		if rewards[i].Cycle > lastBlock.MetaCycle {
			rewards[i].BakingRewards = rewards[i].FutureBakingCount * BlockReward
		}

		rewards[i].Status = getRewardStatus(rewards[i].Cycle, lastBlock.MetaCycle)

		blockReward = BlockReward
		endorsementReward = EndorsementReward

		if rewards[i].Cycle < CarthageCycle {
			blockReward = BabylonBlockReward
			endorsementReward = BabylonEndorsementRewards
		}

		rewards[i].Losses += rewards[i].MissedBaking*blockReward + rewards[i].MissedEndorsements*endorsementReward
	}

	return count, rewards, nil
}
