package render

import (
	genModels "github.com/bullblock-io/tezTracker/gen/models"
	"github.com/bullblock-io/tezTracker/models"
)

// Baker renders an app level model to a generated OpenAPI model.
func Baker(a models.Baker) *genModels.BakersRow {
	return &genModels.BakersRow{
		AccountID:      &a.AccountID,
		Blocks:         &a.Blocks,
		Endorsements:   &a.Endorsements,
		StakingBalance: &a.StakingBalance,
		Fees:           &a.Fees,
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

// BakerInfo renders a baker info details.
func BakerInfo(bi *models.BakerInfo) *genModels.BakerInfo {
	if bi == nil {
		return nil
	}
	return &genModels.BakerInfo{
		EvaluatedBalance:    bi.Balance,
		StakingBalance:      bi.StakingBalance,
		BakingDeposits:      bi.BakingDeposits,
		BakingRewards:       bi.BakingRewards,
		EndorsementDeposits: bi.EndorsementDeposits,
		EndorsementRewards:  bi.EndorsementRewards,
	}
}
