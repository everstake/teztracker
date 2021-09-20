package services

import (
	"fmt"
	"strconv"

	"github.com/everstake/teztracker/models"
	"github.com/guregu/null"
)

func (t *TezTracker) BakingRightsList(blockLevelOrHash []string, priorityTo int, limiter Limiter) (count int64, blocksWithRights []models.Block, err error) {
	filter := models.RightFilter{
		PriorityTo: priorityTo,
		IsFuture:   false,
	}
	count = int64(len(blockLevelOrHash))

	blockRepo := t.repoProvider.GetBlock()
	if count == 0 {
		last, err := blockRepo.Last()
		if err != nil {
			return 0, nil, err
		}
		lastLevel := last.Level.Int64
		rangeEnd := lastLevel - int64(limiter.Offset())
		if rangeEnd < 0 {
			return 0, nil, fmt.Errorf("out of range")
		}
		rangeStart := rangeEnd - int64(limiter.Limit())
		if rangeStart < 0 {
			rangeStart = 0
		}
		// we have a block with number 0
		count = lastLevel + 1
		for ; rangeStart <= rangeEnd; rangeStart++ {
			filter.BlockLevels = append(filter.BlockLevels, rangeStart)
		}
	} else {
		for i := range blockLevelOrHash {
			if level, e := strconv.ParseInt(blockLevelOrHash[i], 10, 64); e == nil {
				filter.BlockLevels = append(filter.BlockLevels, level)
			} else {
				filter.BlockHashes = append(filter.BlockHashes, blockLevelOrHash[i])
			}
		}
	}
	blocks, err := blockRepo.Filter(filter.BlockFilter)
	if err != nil {
		return count, nil, err
	}
	if len(filter.BlockHashes) > 0 {
		filter.BlockLevels = make([]int64, len(blocks))
		for i := range blocks {
			filter.BlockLevels[i] = blocks[i].Level.Int64
		}
	}
	r := t.repoProvider.GetFutureBakingRight()
	rights, err := r.ListDesc(filter)
	if err != nil {
		return count, nil, err
	}

	blockMap := map[int64]*models.Block{}
	for i := range blocks {
		blockMap[blocks[i].Level.Int64] = &blocks[i]
	}
	for i := range rights {
		blockMap[rights[i].BlockLevel].BakingRights = append(blockMap[rights[i].BlockLevel].BakingRights, rights[i])
	}
	return count, blocks, nil
}

func (t *TezTracker) FutureBakingRightsList(priorityTo int, limiter Limiter) (count int64, blocksWithRights []models.FutureBlockBakingRight, err error) {
	blockRepo := t.repoProvider.GetBlock()
	lastBlock, err := blockRepo.Last()
	if err != nil {
		return 0, nil, err
	}
	lastCycle := lastBlock.MetaCycle
	lastLevel := lastBlock.Level.Int64
	lastKnownRightsBlock := (lastCycle + 6) * t.BlocksInCycle()
	count = lastKnownRightsBlock - lastLevel

	rangeStart := lastLevel + 1 + int64(limiter.Offset())
	if rangeStart > lastKnownRightsBlock {
		return 0, nil, fmt.Errorf("out of range")
	}
	rangeEnd := rangeStart + int64(limiter.Limit())
	if rangeEnd > lastKnownRightsBlock {
		rangeEnd = lastKnownRightsBlock
	}
	filter := models.RightFilter{
		PriorityTo: priorityTo,
		IsFuture:   true,
	}
	for ; rangeStart <= rangeEnd; rangeStart++ {
		filter.BlockLevels = append(filter.BlockLevels, rangeStart)
	}
	r := t.repoProvider.GetFutureBakingRight()
	//Not use limit here
	rights, err := r.List(filter, 0, 0)
	if err != nil {
		return count, nil, err
	}
	curBlock := int64(-1)
	for i := range rights {
		if curBlock < rights[i].BlockLevel {
			curBlock = rights[i].BlockLevel
			blockRights := models.FutureBlockBakingRight{
				Level: curBlock,
			}
			blocksWithRights = append(blocksWithRights, blockRights)
		}
		blocksWithRights[len(blocksWithRights)-1].Rights = append(blocksWithRights[len(blocksWithRights)-1].Rights, rights[i])

	}

	return count, blocksWithRights, nil
}

// GetBlockEndorsements finds a block and returns endorsements for it.
func (t *TezTracker) GetBlockBakingRights(hashOrLevel string) (rights []models.FutureBakingRight, count int64, err error) {
	var level int64
	if i, e := strconv.ParseInt(hashOrLevel, 10, 64); e == nil {
		level = i
	} else {
		r := t.repoProvider.GetBlock()
		var filter models.Block
		filter.Hash = null.StringFrom(hashOrLevel)
		found, block, err := r.Find(filter)
		if err != nil {
			return nil, 0, err
		}
		if !found {
			return nil, 0, ErrNotFound
		}
		level = block.Level.Int64
	}
	filter := models.RightFilter{
		IsFuture: true,
	}
	filter.BlockLevels = []int64{level}
	repo := t.repoProvider.GetFutureBakingRight()
	rights, err = repo.ListDesc(filter)
	return rights, int64(len(rights)), err
}

func (t *TezTracker) GetAccountFutureBakingRights(accountID string, cycle int64, limits Limiter) (count int64, futureRights []models.FutureBakingRight, err error) {
	lastBlock, err := t.repoProvider.GetBlock().Last()
	if err != nil {
		return 0, nil, err
	}

	cycleFirstBlock := cycle*t.BlocksInCycle() + 1
	//Return future part of active cycle
	if lastBlock.MetaCycle == cycle {
		cycleFirstBlock = lastBlock.MetaLevel + 1
	}

	repo := t.repoProvider.GetFutureBakingRight()
	filter := models.RightFilter{
		BlockFilter: models.BlockFilter{
			FromID: null.IntFrom(cycleFirstBlock),
			ToID:   null.IntFrom(cycleFirstBlock + t.BlocksInCycle()),
		},
		PriorityTo: 10,
		Delegates:  []string{accountID},
		IsFuture:   true,
	}

	count, err = repo.Count(filter)
	if err != nil {
		return 0, nil, err
	}

	futureRights, err = repo.List(filter, limits.Limit(), limits.Offset())
	if err != nil {
		return 0, nil, err
	}

	for i := range futureRights {
		futureRights[i].Reward = getBlockRewardByCycle(futureRights[i].Cycle.Int64, int64(futureRights[i].Priority))
		futureRights[i].Deposit = getBlockSecurityDepositByCycle(futureRights[i].Cycle.Int64)
	}

	return count, futureRights, nil
}
