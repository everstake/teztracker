package rolls

import (
	"context"
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/repos/block"
	"github.com/everstake/teztracker/repos/rolls"
	"github.com/everstake/teztracker/repos/snapshots"
	"github.com/everstake/teztracker/repos/voting_periods"
)

type RollProvider interface {
	RollsForBlock(ctx context.Context, blockLevel int64) (roll []models.Roll, err error)
}

type UnitOfWork interface {
	GetBlock() block.Repo
	GetSnapshots() snapshots.Repo
	GetRolls() rolls.Repo
	GetVotingPeriod() voting_periods.Repo
}

func SaveRolls(ctx context.Context, unit UnitOfWork, provider RollProvider) (count int, err error) {

	votingRepo := unit.GetVotingPeriod()
	rollsRepo := unit.GetRolls()
	periods, err := votingRepo.List()
	if err != nil {
		return count, err
	}

	lastProcessedPeriod, err := rollsRepo.LastVotingPeriod()
	if err != nil {
		return count, err
	}

	for i := range periods {

		if periods[i].ID < 10 || periods[i].ID <= lastProcessedPeriod {
			continue
		}

		//Skip testing
		if periods[i].Type == "testing" {
			continue
		}

		rolls, err := provider.RollsForBlock(ctx, periods[i].StartBlock)
		if err != nil {
			return count, err
		}

		for j := range rolls {
			rolls[j].VotingPeriod = periods[i].ID
			rolls[j].BlockLevel = periods[i].StartBlock
			rolls[j].Cycle = (periods[i].StartBlock - 1)
		}

		err = rollsRepo.CreateBulk(rolls)
		if err != nil {
			return 0, err
		}

		count += len(rolls)
	}

	return count, nil
}
