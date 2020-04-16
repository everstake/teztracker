package future_rights

import (
	"context"
	"github.com/everstake/teztracker/repos/account"
	"github.com/everstake/teztracker/repos/future_endorsement_rights"

	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/repos/block"
	"github.com/everstake/teztracker/repos/future_baking_rights"
)

type BlocksRepo interface {
	Last() (block models.Block, err error)
}

type RightsRepo interface {
	List(limit, offset uint, filter models.RightFilter) (rights []models.FutureBakingRight, err error)
	Create(right models.FutureBakingRight) error
}

type RightsProvider interface {
	RightsFor(ctx context.Context, blockFrom, blockTo, currentHead int64) ([]models.FutureBakingRight, error)
	EndorsementRightsFor(ctx context.Context, blockFrom, blockTo, currentHead int64) ([]models.FutureEndorsementRight, error)
	BlocksInCycle() int64
}

type UnitOfWork interface {
	Start(ctx context.Context)
	RollbackUnlessCommitted()
	Commit() error
	GetBlock() block.Repo
	GetFutureBakingRight() future_baking_rights.Repo
	GetFutureEndorsementRight() future_endorsement_rights.Repo
	GetAccount() account.Repo
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

	err = unit.GetAccount().RefreshView()
	if err != nil {
		return 0, err
	}

	return count, nil
}

func SaveFutureRightsForBlockRange(ctx context.Context, blockFrom, blockTo, headLevel int64, unit UnitOfWork, provider RightsProvider) (int, error) {
	rights, err := provider.RightsFor(ctx, blockFrom, blockTo, headLevel)
	if err != nil {
		return 0, err
	}

	for i := range rights {
		rights[i].Cycle = (rights[i].Level - 1) / provider.BlocksInCycle()
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

func SaveNewEndorsementRights(ctx context.Context, unit UnitOfWork, provider RightsProvider) (count int, err error) {
	blocks := unit.GetBlock()
	lastBlock, err := blocks.Last()
	if err != nil {
		return 0, err
	}
	lastCycle := lastBlock.MetaCycle
	lastKnownRightsBlock := (lastCycle + 6) * provider.BlocksInCycle()

	rightsRepo := unit.GetFutureEndorsementRight()
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
		cnt, err := SaveFutureEndorsementRightsForBlockRange(ctx, nextBlockToScan, endRange, lastBlock.MetaLevel, unit, provider)
		if err != nil {
			return 0, err
		}
		count += cnt
		nextBlockToScan = endRange + 1
	}
	return count, nil
}

func SaveFutureEndorsementRightsForBlockRange(ctx context.Context, blockFrom, blockTo, headLevel int64, unit UnitOfWork, provider RightsProvider) (int, error) {
	rights, err := provider.EndorsementRightsFor(ctx, blockFrom, blockTo, headLevel)
	if err != nil {
		return 0, err
	}

	for i := range rights {
		rights[i].Cycle = (rights[i].Level - 1) / provider.BlocksInCycle()
	}

	unit.Start(ctx)
	defer unit.RollbackUnlessCommitted()
	rightsRepo := unit.GetFutureEndorsementRight()

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
