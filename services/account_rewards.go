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
		rewards[i].BakingRewards += rewards[i].FutureBakingCount * getBlockRewardByCycle(rewards[i].Cycle, 0)
		rewards[i].EndorsementRewards += rewards[i].FutureEndorsementCount * getEndorsementRewardByCycle(rewards[i].Cycle)

		rewards[i].Status = getRewardStatus(rewards[i].Cycle, lastBlock.MetaCycle)

		blockReward = getBlockRewardByCycle(rewards[i].Cycle, 0)
		endorsementReward = getEndorsementRewardByCycle(rewards[i].Cycle)

		if rewards[i].Cycle < CarthageCycle {
			blockReward = BabylonBlockReward
			endorsementReward = BabylonEndorsementRewards
		}

		rewards[i].Losses += rewards[i].MissedBaking*blockReward + rewards[i].MissedEndorsements*endorsementReward
	}

	return count, rewards, nil
}

func (t *TezTracker) GetAccountSecurityDepositList(accountID string) (rewards []models.AccountRewardsCount, err error) {
	lastBlock, err := t.repoProvider.GetBlock().Last()
	if err != nil {
		return rewards, err
	}

	rewards, err = t.repoProvider.GetAccount().RewardsCountList(accountID, 11)
	if err != nil {
		return nil, err
	}

	bal, err := t.repoProvider.GetBaker().Balance(accountID)
	if err != nil {
		return nil, err
	}

	availableBond := bal.Balance - int64(bal.FrozenBalance)
	unfrozenCycle := PreservedCycles + 1
	var futureEndorsementDeposit, futureBakingDeposit int64
	cycles := PreservedCycles
	if cycles > len(rewards) {
		cycles = len(rewards)
	}
	//Start from active cycle
	for i := cycles; i >= 0; i-- {
		//Calc future deposit
		futureEndorsementDeposit = rewards[i].FutureEndorsementCount * getEndorsementSecurityDepositByCycle(rewards[i].Cycle)
		futureBakingDeposit = rewards[i].FutureBakingCount * getBlockSecurityDepositByCycle(rewards[i].Cycle)

		//Calc actual deposit
		rewards[i].ActualBakingSecurityDeposit = (rewards[i].BakingCount + rewards[i].StolenBaking) * getBlockSecurityDepositByCycle(rewards[i].Cycle)
		rewards[i].ActualEndorsementSecurityDeposit = rewards[i].EndorsementsCount * getEndorsementSecurityDepositByCycle(rewards[i].Cycle)

		//Calc expected deposit
		rewards[i].ExpectedBakingSecurityDeposit = futureBakingDeposit + rewards[i].ActualBakingSecurityDeposit
		rewards[i].ExpectedEndorsementSecurityDeposit = futureEndorsementDeposit + rewards[i].ActualEndorsementSecurityDeposit

		//Calc total deposit
		rewards[i].ActualTotalSecirityDeposit = rewards[i].ActualBakingSecurityDeposit + rewards[i].ActualEndorsementSecurityDeposit
		rewards[i].ExpectedTotalSecurityDeposit = rewards[i].ExpectedBakingSecurityDeposit + rewards[i].ExpectedEndorsementSecurityDeposit

		//Available bond = current amount - future deposit + unfroze deposit
		availableBond -= futureBakingDeposit + futureEndorsementDeposit

		//Unfroze deposit + unfroze rewards
		if len(rewards)-i > unfrozenCycle {
			availableBond += (rewards[i+unfrozenCycle].BakingCount+rewards[i+unfrozenCycle].StolenBaking)*getBlockSecurityDepositByCycle(rewards[i].Cycle) + rewards[i+unfrozenCycle].EndorsementsCount*getEndorsementSecurityDepositByCycle(rewards[i].Cycle)
			availableBond += rewards[i+unfrozenCycle].BakingReward + rewards[i+unfrozenCycle].EndorsementsReward
		}

		rewards[i].AvailableBond = availableBond
		rewards[i].Status = getRewardStatus(rewards[i].Cycle, lastBlock.MetaCycle)
	}

	return rewards[0:unfrozenCycle], nil
}
