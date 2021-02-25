package render

import (
	"encoding/json"
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

func PublicBakersSearch(ams []models.PublicBakerSearch) []*genModels.PublicBakerSearch {
	accs := make([]*genModels.PublicBakerSearch, len(ams))
	for i := range ams {
		accs[i] = &genModels.PublicBakerSearch{AccountID: ams[i].Delegate, Name: ams[i].BakerName}
	}
	return accs
}

// BakerInfo renders a baker info details.
func BakerInfo(bi *models.Baker) *genModels.BakerInfo {
	if bi == nil {
		return nil
	}

	media := &genModels.BakerInfoMedia{}
	if len(bi.Media) != 0 {
		json.Unmarshal([]byte(bi.Media), media)
	}

	bakingSince := bi.BakingSince.Unix()

	return &genModels.BakerInfo{
		Name:                bi.Name,
		BakingSince:         &bakingSince,
		Rolls:               &bi.Rolls,
		Fee:                 &bi.Fee,
		Blocks:              &bi.Blocks,
		Endorsements:        &bi.Endorsements,
		ActiveDelegators:    &bi.ActiveDelegations,
		StakingBalance:      &bi.StakingBalance,
		StakingCapacity:     &bi.StakingCapacity,
		EvaluatedBalance:    &bi.Balance,
		FrozenBalance:       &bi.FrozenBalance,
		BakingDeposits:      &bi.BakingDeposits,
		BakingRewards:       &bi.BakingRewards,
		EndorsementDeposits: &bi.EndorsementDeposits,
		EndorsementRewards:  &bi.EndorsementRewards,
		TotalPaidFees:       bi.TotalPaidFees,
		Media:               media,
	}
}
