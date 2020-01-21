package future_rights

import (
	"context"

	"github.com/bullblock-io/tezTracker/models"
	"github.com/bullblock-io/tezTracker/repos/block"
	"github.com/bullblock-io/tezTracker/repos/future_baking_rights"
)

type BlocksRepo interface {
	Last() (block models.Block, err error)
}

type RightsRepo interface {
	List(limit, offset uint, filter models.BakingRightFilter) (rights []models.FutureBakingRight, err error)
	Find(filter models.BakingRightFilter) (found bool, right models.FutureBakingRight, err error)
	Create(right models.FutureBakingRight) error
}

type RightsProvider interface {
	RightsFor(ctx context.Context, blockFrom, blockTo, currentHead int64) ([]models.FutureBakingRight, error)
	BlocksInCycle() int64
}

type UnitOfWork interface {
	Start(ctx context.Context)
	RollbackUnlessCommitted()
	Commit() error
	GetBlock() block.Repo
	GetFutureBakingRight() future_baking_rights.Repo
}

const BlocksRangeSize = 256

func SaveNewBakingRights(ctx context.Context, unit UnitOfWork, provider RightsProvider) (count int, err error) {
	blocks := unit.GetBlock()
	lastBlock, err := blocks.Last()
	if err != nil {
		return 0, err
	}
	lastCycle := lastBlock.MetaCycle
	lastKnownRightsBlock := (lastCycle + 6) * provider.BlocksInCycle()

	rightsRepo := unit.GetFutureBakingRight()
	found, lastRight, err := rightsRepo.Last()
	if err != nil {
		return 0, err
	}
	var nextBlockToScan int64
	if !found {
		nextBlockToScan = 1
	} else {
		if lastRight.Level >= lastKnownRightsBlock {
			return 0, nil
		}
		nextBlockToScan = lastRight.Level + 1
	}
	for nextBlockToScan <= lastKnownRightsBlock {
		endRange := nextBlockToScan + BlocksRangeSize - 1
		if endRange > lastKnownRightsBlock {
			endRange = lastKnownRightsBlock
		}
		cnt, err := SaveFutureRightsForBlockRange(ctx, nextBlockToScan, endRange, lastBlock.MetaLevel, unit, provider)
		if err != nil {
			return 0, err
		}
		count += cnt
		nextBlockToScan = endRange + 1
	}
	return count, nil
}

func SaveFutureRightsForBlockRange(ctx context.Context, blockFrom, blockTo, headLevel int64, unit UnitOfWork, provider RightsProvider) (int, error) {
	rights, err := provider.RightsFor(ctx, blockFrom, blockTo, headLevel)
	if err != nil {
		return 0, err
	}

	unit.Start(ctx)
	defer unit.RollbackUnlessCommitted()
	rightsRepo := unit.GetFutureBakingRight()

	err = rightsRepo.CreateBulk(rights)
	if err != nil {
		return 0, err
	}

	err = unit.Commit()
	if err != nil {
		return 0, err
	}
	return len(rights), nil
}
