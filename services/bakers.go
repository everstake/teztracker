package services

import (
	"fmt"
	"github.com/everstake/teztracker/models"
)

const (
	PreservedCycles            = 5
	Precision                  = 6
	XTZ                        = 1000000
	BlockSecurityDeposit       = 512 * XTZ
	EndorsementSecurityDeposit = 64 * XTZ
	BlockReward                = 40 * XTZ
	LowPriorityBlockReward     = 6 * XTZ
	BabylonBlockReward         = 24 * XTZ
	EndorsementReward          = 1.25 * XTZ
	BabylonEndorsementRewards  = 1.75 * XTZ
	CarthageCycle              = 208
	BlockEndorsers             = 32
	TokensPerRoll              = 8000
	TotalLocked                = (BlockSecurityDeposit + EndorsementSecurityDeposit*BlockEndorsers) * BlocksInMainnetCycle * (PreservedCycles + 1)
	BlockLockEstimate          = BlockReward + BlockSecurityDeposit + BlockEndorsers*(EndorsementReward+EndorsementSecurityDeposit)
)

// BakerList retrives up to limit of bakers after the specified id.
func (t *TezTracker) PublicBakerList(limits Limiter) (bakers []models.Baker, count int64, err error) {
	r := t.repoProvider.GetBaker()
	count, err = r.PublicBakersCount()
	if err != nil {
		return nil, 0, err
	}

	bakers, err = r.PublicBakersList(limits.Limit(), limits.Offset())
	if err != nil {
		return nil, 0, err
	}

	block, err := t.repoProvider.GetBlock().Last()
	if err != nil {
		return nil, 0, err
	}

	//Get last snapshot
	_, snap, err := t.repoProvider.GetSnapshots().Find(block.MetaCycle - PreservedCycles)
	if err != nil {
		return nil, 0, err
	}

	for i := range bakers {
		bakers[i].StakingCapacity = t.calcBakerCapacity(bakers[i], snap.Rolls)
	}

	return bakers, count, nil
}

func (t *TezTracker) PublicBakersSearchList() (list []models.PublicBakerSearch, err error) {
	list, err = t.repoProvider.GetBaker().PublicBakersSearchList()
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (t *TezTracker) BakerList(limits Limiter) (bakers []models.Baker, count int64, err error) {
	r := t.repoProvider.GetBaker()
	count, err = r.Count()
	if err != nil {
		return nil, 0, err
	}

	bakers, err = r.List(limits.Limit(), limits.Offset())
	return bakers, count, err
}

//Used BakingBad capacity formula
func (t *TezTracker) calcBakerCapacity(bi models.Baker, totalRolls int64) int64 {
	bakerBalanceF := float64(bi.Balance - bi.FrozenEndorsementRewards - bi.FrozenBakingRewards)
	totalRollsF := float64(totalRolls)

	bakerShare := bakerBalanceF / float64(TotalLocked)

	bakerRollsCapacity := totalRollsF * bakerShare
	return int64(bakerRollsCapacity * float64(TokensPerRoll) * float64(XTZ))
}

func (t *TezTracker) GetCurrentCycle() (int64, error) {
	r := t.repoProvider.GetBlock()
	lastBlock, err := r.Last()
	if err != nil {
		return 0, err
	}
	return lastBlock.MetaCycle, nil
}

func getFirstPreservedBlock(currentCycle, blocksInCycle int64) (fpb int64) {
	fpc := currentCycle - PreservedCycles

	if fpc > 0 {
		fpb = fpc*blocksInCycle + 1
	}
	return fpb
}

//TODO change this method
func (t *TezTracker) GetBakerInfo(accountID string) (bi *models.Baker, err error) {
	r := t.repoProvider.GetBaker()
	found, baker, err := r.Find(accountID)
	if err != nil {
		return bi, err
	}
	if !found {
		return nil, nil
	}

	err = t.calcDepositRewards(&baker.BakerStats, baker.AccountID)
	if err != nil {
		return bi, err
	}

	block, err := t.repoProvider.GetBlock().Last()
	if err != nil {
		return nil, err
	}

	//Get last snapshot
	_, snap, err := t.repoProvider.GetSnapshots().Find(block.MetaCycle - PreservedCycles - 2)
	if err != nil {
		return nil, err
	}

	baker.StakingCapacity = t.calcBakerCapacity(baker, snap.Rolls)

	return &baker, nil
}

func (t *TezTracker) calcDepositRewards(bi *models.BakerStats, accountID string) (err error) {

	bi.BakingDeposits = bi.BakingCount * BlockSecurityDeposit
	bi.BakingRewards = bi.FrozenBakingRewards

	bi.EndorsementDeposits = bi.EndorsementCount * EndorsementSecurityDeposit
	bi.EndorsementRewards = bi.FrozenEndorsementRewards

	return nil
}

func (t *TezTracker) getLockedBalance() (int64, error) {
	blockR := t.repoProvider.GetBlock()
	lastBlock, err := blockR.Last()
	if err != nil {
		return 0, err
	}
	curCycle := lastBlock.MetaCycle
	fpb := getFirstPreservedBlock(curCycle, t.BlocksInCycle())
	lockedBlocks := lastBlock.Level.Int64 - fpb
	lockedBalanceEstimate := lockedBlocks * BlockLockEstimate

	return lockedBalanceEstimate, nil
}

// GetStakingRatio calculates the rough ratio of staked balance to the total supply.
func (t *TezTracker) GetStakingRatio() (float64, error) {
	lockedBalanceEstimate, err := t.getLockedBalance()
	if err != nil {
		return 0, err
	}

	ar := t.repoProvider.GetAccount()
	liquidBalance, err := ar.TotalBalance()
	if err != nil {
		return 0, err
	}

	br := t.repoProvider.GetBaker()
	stakedBalance, err := br.TotalStakingBalance()
	if err != nil {
		return 0, err
	}

	supply := liquidBalance + lockedBalanceEstimate
	if supply == 0 {
		return 0, nil
	}

	lastBlock, err := t.repoProvider.GetBlock().Last()
	if err != nil {
		return 0, nil
	}

	bakingRewards, err := br.TotalBakingRewards("", lastBlock.MetaCycle-PreservedCycles, lastBlock.MetaCycle)
	if err != nil {
		return 0, nil
	}

	endorsementRewards, err := br.TotalEndorsementRewards("", lastBlock.MetaCycle-PreservedCycles, lastBlock.MetaCycle)
	if err != nil {
		return 0, nil
	}

	stakedBalance = stakedBalance - bakingRewards - endorsementRewards
	ratio := float64(stakedBalance) / float64(supply)

	return ratio, nil
}

func (t *TezTracker) GetThirdPartyBakers() (bakers []models.ThirdPartyBakerAgg, err error) {
	tpbRepo := t.repoProvider.GetThirdPartyBakers()
	return tpbRepo.GetAggregatedBakers()
}

func (t *TezTracker) GetNumberOfDelegators() (items []models.BakerDelegators, err error) {
	block, err := t.repoProvider.GetBlock().Last()
	if err != nil {
		return nil, fmt.Errorf("get LastBlock: %s", err.Error())
	}
	items, err = t.repoProvider.GetBaker().NumberOfDelegators(uint64(block.MetaCycle))
	if err != nil {
		return nil, fmt.Errorf("NumberOfDelegators: %s", err.Error())
	}
	return items, nil
}

func (t *TezTracker) GetBakersStakeChange() (items []models.BakerDelegators, err error) {
	block, err := t.repoProvider.GetBlock().Last()
	if err != nil {
		return nil, fmt.Errorf("get LastBlock: %s", err.Error())
	}
	prevStakes, err := t.repoProvider.GetBaker().GetBakersStake(uint64(block.MetaCycle - 1))
	if err != nil {
		return nil, fmt.Errorf("BakersStake: %s", err.Error())
	}
	prevStakesMap := make(map[string]int64)
	for _, stake := range prevStakes {
		prevStakesMap[stake.Address] = stake.Value
	}
	lastStakes, err := t.repoProvider.GetBaker().GetBakersStake(uint64(block.MetaCycle))
	if err != nil {
		return nil, fmt.Errorf("BakersStake: %s", err.Error())
	}
	items = make([]models.BakerDelegators, len(lastStakes))
	for i := range lastStakes {
		var diff int64
		p, ok := prevStakesMap[lastStakes[i].Address]
		if ok {
			diff = lastStakes[i].Value - p
		}
		items[i] = models.BakerDelegators{
			Baker:   lastStakes[i].Baker,
			Address: lastStakes[i].Address,
			Value:   diff,
		}
	}
	return items, nil
}

func (t *TezTracker) GetBakersVoting() (voting models.BakersVoting, err error) {
	bakers, err := t.repoProvider.GetBaker().GetBakersVoting()
	if err != nil {
		return voting, fmt.Errorf("GetBakersVoting: %s", err.Error())
	}
	count, err := t.repoProvider.GetVotingPeriod().ProposalsCount()
	if err != nil {
		return voting, fmt.Errorf("ProposalsCount: %s", err.Error())
	}
	return models.BakersVoting{
		ProposalsCount: count,
		Bakers:         bakers,
	}, nil
}
