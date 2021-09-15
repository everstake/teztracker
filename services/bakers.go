package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/everstake/teztracker/models"
)

const (
	PreservedCycles = 5
	XTZ             = 1000000

	GranadaBlockSecurityDeposit  = 640 * XTZ
	FlorenceBlockSecurityDeposit = 512 * XTZ

	GranadaEndorsementSecurityDeposit  = 2.5 * XTZ
	FlorenceEndorsementSecurityDeposit = 64 * XTZ

	GranadaBlockReward  = 20 * XTZ
	FlorenceBlockReward = 40 * XTZ

	GradanaLowPriorityBlockReward  = 3 * XTZ
	FlorenceLowPriorityBlockReward = 6 * XTZ

	BabylonBlockReward = 24 * XTZ

	GranadaEndorsementReward  = 1.25 * XTZ
	FlorenceEndorsementReward = 1.25 * XTZ
	BabylonEndorsementRewards = 1.75 * XTZ
	CarthageCycle             = 208
	GranadaCycle              = 388

	GranadaBlockEndorsers  = 256
	FlorenceBlockEndorsers = 32

	TokensPerRoll = 8000

	TotalLocked = (GranadaBlockSecurityDeposit + GranadaEndorsementSecurityDeposit*GranadaBlockEndorsers) * BlocksInMainnetCycle * (PreservedCycles + 1)

	BlockLockEstimate = GranadaBlockReward + GranadaBlockSecurityDeposit + GranadaBlockEndorsers*(GranadaEndorsementReward+GranadaEndorsementSecurityDeposit)

	bakerMediaSource = "https://api.tzkt.io/v1/accounts/%s?metadata=true"
)

// BakerList retrives up to limit of bakers after the specified id.
func (t *TezTracker) PublicBakerList(limits Limiter, favorites []string) (bakers []models.Baker, count int64, err error) {
	r := t.repoProvider.GetBaker()
	count, err = r.PublicBakersCount()
	if err != nil {
		return nil, 0, err
	}

	bakers, err = r.PublicBakersList(limits.Limit(), limits.Offset(), favorites)
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

func (t *TezTracker) BakerList(limits Limiter, favorites []string) (bakers []models.Baker, count int64, err error) {
	r := t.repoProvider.GetBaker()
	count, err = r.Count()
	if err != nil {
		return nil, 0, err
	}

	bakers, err = r.List(limits.Limit(), limits.Offset(), favorites)
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

	bi.BakingDeposits = bi.BakingCount * GranadaBlockSecurityDeposit
	bi.BakingRewards = bi.FrozenBakingRewards

	bi.EndorsementDeposits = bi.EndorsementCount * GranadaBlockSecurityDeposit
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

func (t *TezTracker) UpdateBakersSocialMedia() error {
	bakersRepo := t.repoProvider.GetBaker()
	bakers, err := bakersRepo.PublicBakersList(10000, 0, nil)
	if err != nil {
		return fmt.Errorf("bakersRepo.PublicBakersList: %s", err.Error())
	}
	for _, baker := range bakers {
		media, err := getBakerMediaData(baker.AccountID)
		if err != nil {
			return fmt.Errorf("getBakerMediaData: %s", err.Error())
		}
		baker.Media = string(media)
		err = bakersRepo.UpdateBaker(baker)
		if err != nil {
			return fmt.Errorf("bakersRepo.UpdateBaker: %s", err.Error())
		}
	}
	return nil
}

type bakerInfo struct {
	Metadata struct {
		Site     string `json:"site,omitempty"`
		Twitter  string `json:"twitter,omitempty"`
		Telegram string `json:"telegram,omitempty"`
		Reddit   string `json:"reddit,omitempty"`
	} `json:"metadata"`
}

func getBakerMediaData(address string) (media []byte, err error) {
	resp, err := http.Get(fmt.Sprintf(bakerMediaSource, address))
	if err != nil {
		return media, fmt.Errorf("http.Get: %s", err.Error())
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return media, fmt.Errorf("ioutil.ReadAll: %s", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		return media, fmt.Errorf("bad request status: %d", resp.StatusCode)
	}
	var info bakerInfo
	err = json.Unmarshal(data, &info)
	if err != nil {
		return media, fmt.Errorf("json.Unmarshal: %s", err.Error())
	}
	media, _ = json.Marshal(info.Metadata)
	return media, nil
}
