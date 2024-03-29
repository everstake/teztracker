package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"

	"github.com/everstake/teztracker/models"
	"github.com/shopspring/decimal"
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
	BabylonBlockReward  = 24 * XTZ

	GradanaLowPriorityBlockReward  = 3 * XTZ
	FlorenceLowPriorityBlockReward = 6 * XTZ

	GranadaEndorsementReward  = 0.078125 * XTZ
	FlorenceEndorsementReward = 1.25 * XTZ

	BabylonEndorsementRewards = 1.75 * XTZ

	CarthageCycle = 208
	GranadaCycle  = 388

	GranadaBlockEndorsers  = 256
	FlorenceBlockEndorsers = 32

	TokensPerRoll = 8000

	TotalLocked = (GranadaBlockSecurityDeposit + GranadaEndorsementSecurityDeposit*GranadaBlockEndorsers) * BlocksInMainnetCycle * (PreservedCycles + 1)

	BlockLockEstimate = GranadaBlockReward + GranadaBlockSecurityDeposit + GranadaBlockEndorsers*(GranadaEndorsementReward+GranadaEndorsementSecurityDeposit)

	bakerMediaSource       = "https://api.tzkt.io/v1/accounts/%s?metadata=true"
	holdingPointStorageKey = "holding_points"
)

// BakerList retrives up to limit of bakers after the specified id.
func (t *TezTracker) PublicBakerList(limits Limiter, favorites []string) (publicBakers []models.PublicBaker, count int64, err error) {
	r := t.repoProvider.GetBaker()
	count, err = r.PublicBakersCount()
	if err != nil {
		return nil, 0, err
	}

	bakers, err := r.PublicBakersList(limits.Limit(), limits.Offset(), favorites)
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

	// insert bakers changes
	var changes map[string]models.BakerChanges
	isFound, err := t.repoProvider.GetStorage().Get(models.BakersChangesStorageKey, &changes)
	if err != nil {
		return nil, 0, fmt.Errorf("GetStorage: Get: %s", err.Error())
	}

	if !isFound {
		return nil, 0, fmt.Errorf("Baker changes not found")
	}

	publicBakers = make([]models.PublicBaker, len(bakers))
	for i, baker := range bakers {
		pb := models.PublicBaker{
			Baker: baker,
		}
		change, ok := changes[baker.AccountID]
		if ok {
			pb.StakeChange = change.Balance
			pb.DelegatorsChange = change.Delegators
		}
		publicBakers[i] = pb
	}
	return publicBakers, count, nil
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
		fpb = (GranadaCycle)*blocksInCycle/2 + (fpc-GranadaCycle)*blocksInCycle + 1
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

func (t *TezTracker) SaveHoldingPoints() error {
	points := []float64{0.05, 0.33, 0.51, 0.8}
	bakers, err := t.repoProvider.GetBaker().List(100000, 0, nil)
	if err != nil {
		return fmt.Errorf("repoProvider.Baker: List: %s", err.Error())
	}
	sort.Slice(bakers, func(i, j int) bool {
		return bakers[i].StakingBalance > bakers[j].StakingBalance
	})
	var total int64
	for _, baker := range bakers {
		total += baker.StakingBalance
	}

	var holdingPoints []models.HoldingPoint
	for _, point := range points {
		var amount int64
		var count int64
		for _, baker := range bakers {
			amount += baker.StakingBalance
			count++
			p, _ := decimal.NewFromInt(amount).Div(decimal.NewFromInt(total)).Float64()
			if p >= point {
				break
			}
		}
		holdingPoints = append(holdingPoints, models.HoldingPoint{
			Percent: point,
			Amount:  amount,
			Count:   count,
		})
	}
	err = t.repoProvider.GetStorage().Set(holdingPointStorageKey, holdingPoints)
	if err != nil {
		return fmt.Errorf("GetStorage: Set: %s", err.Error())
	}
	return nil
}

func (t *TezTracker) GetHoldingPoints() (items []models.HoldingPoint, err error) {
	_, err = t.repoProvider.GetStorage().Get(holdingPointStorageKey, &items)
	return items, err
}
