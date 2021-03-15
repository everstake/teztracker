package services

import (
	"fmt"
	"strconv"
	"time"

	"github.com/everstake/teztracker/models"
	"github.com/guregu/null"
)

const lostRewardsCacheKey = "lost_rewards"

// ErrNotFound is an error returned when the requested entity doesn't exist in the repository.
var ErrNotFound = fmt.Errorf("not found")

// HeadBlock retrieves the last added block from the repository.
func (t *TezTracker) HeadBlock() (models.Block, error) {
	r := t.repoProvider.GetBlock()
	return r.Last()
}

// BlockList retrives up to limit of blocks before the specified level.
func (t *TezTracker) BlockList(beforeLevel uint64, limits Limiter) ([]models.Block, int64, error) {
	r := t.repoProvider.GetBlock()
	last, err := r.Last()
	if err != nil {
		return nil, 0, err
	}
	blocks, err := r.ListExtended(limits.Limit(), limits.Offset(), beforeLevel)
	return blocks, last.Level.Int64 + 1, err
}

// GetBlockWithOperationGroups retrieves a block by hash or by level. It loads OperationGroups into the found block.
func (t *TezTracker) GetBlockWithOperationGroups(hashOrLevel string) (block models.Block, err error) {
	r := t.repoProvider.GetBlock()

	var filter models.Block
	if i, e := strconv.ParseInt(hashOrLevel, 10, 64); e == nil {
		filter.Level = null.IntFrom(i)
	} else {
		filter.Hash = null.StringFrom(hashOrLevel)
	}
	found, block, err := r.FindExtended(filter)
	if err != nil {
		return block, err
	}
	if !found {
		return block, ErrNotFound
	}

	found, prevBlock, err := r.Find(models.Block{Level: null.IntFrom(block.Level.Int64 - 1)})
	if err != nil {
		return block, err
	}
	if found {
		block.BlockTime = int64(block.Timestamp.Sub(prevBlock.Timestamp).Seconds())
	}

	ogr := t.repoProvider.GetOperationGroup()
	ogs, err := ogr.GetGroupsFor(block)
	if err != nil {
		return block, err
	}
	block.OperationGroups = ogs
	return block, nil
}

func (t *TezTracker) SaveLostRewardsInCache() error {
	repo := t.repoProvider.GetBlock()
	for period, duration := range models.GetChartPeriods() {
		items, err := repo.GetLostRewardsAgg(models.AggTimeFilter{
			From:   time.Now().Add(-duration),
			Period: period,
		})
		if err != nil {
			return fmt.Errorf("repo.GetLostRewardsAgg: %s", err.Error())
		}
		storageKey := fmt.Sprintf("%s_%s", lostRewardsCacheKey, period)
		err = t.repoProvider.GetStorage().Set(storageKey, items)
		if err != nil {
			return fmt.Errorf("GetStorage: Set: %s", err.Error())
		}
	}
	return nil
}

func (t *TezTracker) GetLostRewards(period string) (items []models.AggTimeInt, err error) {
	err = models.ValidatePeriod(period)
	if err != nil {
		return items, fmt.Errorf("ValidatePeriod: %s", err.Error())
	}
	storageKey := fmt.Sprintf("%s_%s", lostRewardsCacheKey, period)
	err = t.repoProvider.GetStorage().Get(storageKey, &items)
	if err != nil {
		return items, fmt.Errorf("GetStorage: Set: %s", err.Error())
	}
	return items, nil
}

func (t *TezTracker) GetLostBlocksCountAgg(filter models.AggTimeFilter) (items []models.AggTimeInt, err error) {
	return t.repoProvider.GetBlock().GetLostBlocksCountAgg(filter)
}

func (t *TezTracker) GetLostEndorsingCountAgg(filter models.AggTimeFilter) (items []models.AggTimeInt, err error) {
	return t.repoProvider.GetEndorsing().GetLostEndorsingCountAgg(filter)
}
