package render

import (
	genModels "github.com/everstake/teztracker/gen/models"
	"github.com/everstake/teztracker/models"
)

func AccountBaking(acb models.AccountBaking) *genModels.AccountBakingRow {
	return &genModels.AccountBakingRow{
		AvgPriority:  &acb.AvgPriority,
		Blocks:       &acb.Count,
		Cycle:        &acb.Cycle,
		CycleStart:   GetUnixFromNullTime(acb.CycleStart),
		CycleEnd:     GetUnixFromNullTime(acb.CycleEnd),
		Status:       string(acb.Status),
		Missed:       &acb.Missed,
		Rewards:      &acb.Reward,
		Stolen:       &acb.Stolen,
		TotalDeposit: &acb.TotalDeposit,
	}
}

func AccountBakingList(accb []models.AccountBaking) []*genModels.AccountBakingRow {
	accbs := make([]*genModels.AccountBakingRow, len(accb))
	for i := range accb {
		accbs[i] = AccountBaking(accb[i])
	}
	return accbs
}
