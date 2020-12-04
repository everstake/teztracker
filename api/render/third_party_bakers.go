package render

import (
	genModels "github.com/everstake/teztracker/gen/models"
	"github.com/everstake/teztracker/models"
)

func ThirdPartyBakers(bakers []models.ThirdPartyBakerAgg) (result []*genModels.ThirdPartyBakers) {
	for _, b := range bakers {
		var providers []*genModels.ThirdPartyBakersProvidersItems0
		for _, p := range b.Providers {
			providers = append(providers, &genModels.ThirdPartyBakersProvidersItems0{
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
			})
		}
		result = append(result, &genModels.ThirdPartyBakers{
			Baker:     b.Address,
			Providers: providers,
		})
	}
	return result
}
