package services

import (
	"github.com/everstake/teztracker/models"
)

const (
	cyclesOnVotingPeriod = 8
)

func (t *TezTracker) VotingPeriod(id string) (info models.PeriodInfo, err error) {
	repo := t.repoProvider.GetVotingPeriod()
	info, err = repo.Info(id)
	if err != nil {

	}

	info.TotalRolls, info.TotalBakers, err = t.repoProvider.GetSnapshots().RollsAndBakersInBlock(info.StartBlock)
	if err != nil {

	}

	return info, nil
}
