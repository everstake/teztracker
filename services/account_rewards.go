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
		rewards[i].BakingRewards += rewards[i].FutureBakingCount * BlockReward
		rewards[i].EndorsementRewards += rewards[i].FutureEndorsementCount * EndorsementReward

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
		futureEndorsementDeposit = rewards[i].FutureEndorsementCount * EndorsementSecurityDeposit
		futureBakingDeposit = rewards[i].FutureBakingCount * BlockSecurityDeposit

		//Calc actual deposit
		rewards[i].ActualBakingSecurityDeposit = (rewards[i].BakingCount + rewards[i].StolenBaking) * BlockSecurityDeposit
		rewards[i].ActualEndorsementSecurityDeposit = rewards[i].EndorsementsCount * EndorsementSecurityDeposit

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
			availableBond += (rewards[i+unfrozenCycle].BakingCount+rewards[i+unfrozenCycle].StolenBaking)*BlockSecurityDeposit + rewards[i+unfrozenCycle].EndorsementsCount*EndorsementSecurityDeposit
			availableBond += rewards[i+unfrozenCycle].BakingReward + rewards[i+unfrozenCycle].EndorsementsReward
		}

		rewards[i].AvailableBond = availableBond
		rewards[i].Status = getRewardStatus(rewards[i].Cycle, lastBlock.MetaCycle)
	}

	return rewards[0:unfrozenCycle], nil
}
