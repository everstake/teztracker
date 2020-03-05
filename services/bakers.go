package services

import (
	"github.com/everstake/teztracker/models"
)

const (
	PreservedCycles            = 5
	XTZ                        = 1000000
	BlockSecurityDeposit       = 512 * XTZ
	EndorsementSecurityDeposit = 64 * XTZ
	BlockReward                = 40 * XTZ
	EndorsementReward          = 1.25 * XTZ
	BlockEndorsers             = 32
	BlockLockEstimate          = BlockReward + BlockSecurityDeposit + BlockEndorsers*(EndorsementReward+EndorsementSecurityDeposit)
)

// BakerList retrives up to limit of bakers after the specified id.
func (t *TezTracker) BakerList(limits Limiter) (bakers []models.Baker, count int64, err error) {
	r := t.repoProvider.GetBaker()
	count, err = r.Count()
	if err != nil {
		return nil, 0, err
	}

	bakers, err = r.List(limits.Limit(), limits.Offset())
	return bakers, count, err
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

func (t *TezTracker) GetBakerInfo(accountID string) (bi *models.BakerInfo, err error) {
	r := t.repoProvider.GetBaker()
	found, delegate, err := r.Find(accountID)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, nil
	}
	bi = &models.BakerInfo{Delegate: delegate}
	curCycle, err := t.GetCurrentCycle()
	if err != nil {
		return bi, err
	}

	fpb := getFirstPreservedBlock(curCycle, t.BlocksInCycle())
	counter, err := r.BlocksCountBakedBy([]string{accountID}, fpb)
	if err != nil {
		return bi, err
	}
	var blocksCount int64
	if len(counter) == 1 {
		blocksCount = counter[0].Count
	}
	bi.BakingDeposits = blocksCount * BlockSecurityDeposit
	bi.BakingRewards = blocksCount * BlockReward

	endCounter, err := r.EndorsementsCountBy([]string{accountID}, fpb)
	if err != nil {
		return bi, err
	}
	var endorsementCount int64
	var endorsementWeight float64
	if len(endCounter) == 1 {
		endorsementCount = endCounter[0].Count
		endorsementWeight = endCounter[0].Weight
	}
	bi.EndorsementDeposits = endorsementCount * EndorsementSecurityDeposit
	bi.EndorsementRewards = int64(endorsementWeight * EndorsementReward)

	return bi, nil
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

	ratio := float64(stakedBalance) / float64(supply)

	return ratio, nil
}
