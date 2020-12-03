package snapshots

import (
	"context"
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/repos/block"
	"github.com/everstake/teztracker/repos/rolls"
	"github.com/everstake/teztracker/repos/snapshots"
)

type BlocksRepo interface {
	Last() (block models.Block, err error)
}

type RightsRepo interface {
	List(limit, offset uint) (count int64, snaps []models.Snapshot, err error)
	Create(right models.Snapshot) error
}

type SnapshotProvider interface {
	SnapshotForCycle(ctx context.Context, cycle int64, useHead bool) (snap models.Snapshot, err error)
	RollsForBlock(ctx context.Context, blockLevel int64) (roll []models.Roll, err error)
}

type UnitOfWork interface {
	Start(ctx context.Context)
	RollbackUnlessCommitted()
	Commit() error
	GetBlock() block.Repo
	GetSnapshots() snapshots.Repo
	GetRolls() rolls.Repo
}

const CyclesInAdvance = 6
const FirstCycleWithSnapshots = 7

func SaveNewSnapshots(ctx context.Context, unit UnitOfWork, provider SnapshotProvider) (count int, err error) {
	blocks := unit.GetBlock()
	lastBlock, err := blocks.Last()
	if err != nil {
		return 0, err
	}
	lastCycle := lastBlock.MetaCycle + CyclesInAdvance

	snapRepo := unit.GetSnapshots()
	_, list, err := snapRepo.List(1, 0)
	if err != nil {
		return 0, err
	}

	nextCycleToScan := int64(FirstCycleWithSnapshots)
	if len(list) > 0 {
		if list[0].Snapshot.Cycle >= lastCycle {
			return 0, nil
		}
		nextCycleToScan = list[0].Snapshot.Cycle + 1
	}

	//Last cycle not finished yet so will be selected later
	for ; nextCycleToScan < lastCycle; nextCycleToScan++ {
		err := SaveSnapshotForCycle(ctx, nextCycleToScan, unit, provider, nextCycleToScan >= lastBlock.MetaCycle)
		if err != nil {
			return 0, err
		}
		count++
	}

	return count, nil
}

func SaveSnapshotForCycle(ctx context.Context, cycle int64, unit UnitOfWork, provider SnapshotProvider, isFromFuture bool) error {
	snap, err := provider.SnapshotForCycle(ctx, cycle, isFromFuture)
	if err != nil {
		return err
	}

	if snap.Rolls == 0 {

		rolls, _, err := unit.GetRolls().RollsAndBakersInBlock(snap.BlockLevel)
		if err != nil {
			return err
		}

		snap.Rolls = rolls
	}

	err = unit.GetSnapshots().Create(snap)
	if err != nil {
		return err
	}

	return nil
}
