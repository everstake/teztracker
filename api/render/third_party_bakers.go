package render

import (
	genModels "github.com/everstake/teztracker/gen/models"
	"github.com/everstake/teztracker/models"
)

func ThirdPartyBakers(bakers []models.ThirdPartyBakerAgg) (result []*genModels.ThirdPartyBakers) {
	result = make([]*genModels.ThirdPartyBakers, len(bakers))
	for i, b := range bakers {
		result[i] = &genModels.ThirdPartyBakers{
			Baker:     b.Address,
			Providers: ThirdPartyBakersProviders(b.Providers),
		}
	}
	return result
}

func ThirdPartyBakersProviders (tp models.ThirdPartyProviders) []*genModels.ThirdPartyBakersProvidersItems0 {
	providers := make([]*genModels.ThirdPartyBakersProvidersItems0, len(tp))
	for i, p := range tp {
		providers[i] = ThirdPartyBakersItem(p)
	}
	return providers
}

func ThirdPartyBakersItem (p models.ThirdPartyBaker) *genModels.ThirdPartyBakersProvidersItems0 {
	return &genModels.ThirdPartyBakersProvidersItems0{
		Address:           p.Address,
		AvailableCapacity: p.AvailableCapacity,
		Efficiency:        p.Efficiency,
		Fee:               p.Fee,
		Name:              p.Name,
		Number:            int64(p.Number),
		PayoutAccuracy:    p.PayoutAccuracy,
		Provider:          p.Provider,
		StakingBalance:    p.StakingBalance,
		Yield:             p.Yield,
	}
}


