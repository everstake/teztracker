package render

import (
	genModels "github.com/everstake/teztracker/gen/models"
	"github.com/everstake/teztracker/models"
)

// Baker renders an app level model to a generated OpenAPI model.
func Baker(a models.Baker) *genModels.BakersRow {
	return &genModels.BakersRow{
		AccountID: a.AccountID,
		BakerInfo: BakerInfo(&a),
	}
}

// Bakers renders a slice of app level Bakers into a slice of OpenAPI models.
func Bakers(ams []models.Baker) []*genModels.BakersRow {
	accs := make([]*genModels.BakersRow, len(ams))
	for i := range ams {
		accs[i] = Baker(ams[i])
	}
	return accs
}

func PublicBakers(ams []models.Baker) []genModels.PublicBaker {
	accs := make([]genModels.PublicBaker, len(ams))
	for i := range ams {
		bakerRow := Baker(ams[i])
		accs[i] = genModels.PublicBaker{BakersRow: *bakerRow}
	}
	return accs
}

// BakerInfo renders a baker info details.
func BakerInfo(bi *models.Baker) *genModels.BakerInfo {
	if bi == nil {
		return nil
	}

	return &genModels.BakerInfo{
		Name:                bi.Name,
		BakingSince:         bi.BakingSince,
		Rolls:               bi.Rolls,
		Fee:                 bi.Fee,
		Blocks:              bi.Blocks,
		Endorsements:        bi.Endorsements,
		ActiveDelegators:    bi.ActiveDelegations,
		StakingBalance:      bi.StakingBalance,
		EvaluatedBalance:    bi.Balance,
		FrozenBalance:       bi.FrozenBalance,
		BakingDeposits:      bi.BakingDeposits,
		BakingRewards:       bi.BakingRewards,
		EndorsementDeposits: bi.EndorsementDeposits,
		EndorsementRewards:  bi.EndorsementRewards,
		TotalPaidFees:       bi.TotalPaidFees,
	}
}
