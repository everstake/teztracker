package render

import (
	genModels "github.com/everstake/teztracker/gen/models"
	"github.com/everstake/teztracker/models"
)

func AccountEndorsing(acb models.AccountEndorsing) *genModels.AccountEndorsingRow {
	return &genModels.AccountEndorsingRow{
		Slots:        &acb.Count,
		Cycle:        &acb.Cycle,
		Status:       string(acb.Status),
		Missed:       &acb.Missed,
		Rewards:      &acb.Reward,
		TotalDeposit: &acb.TotalDeposit,
	}
}

func AccountEndorsingList(acce []models.AccountEndorsing) []*genModels.AccountEndorsingRow {
	accbs := make([]*genModels.AccountEndorsingRow, len(acce))
	for i := range acce {
		accbs[i] = AccountEndorsing(acce[i])
	}
	return accbs
}
