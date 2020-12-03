package render

import (
	genModels "github.com/everstake/teztracker/gen/models"
	"github.com/everstake/teztracker/models"
	"sort"
)

func ThirdPartyBakers(items []models.ThirdPartyBaker) (result []*genModels.ThirdPartyBakers) {
	type baker struct {
		address        string
		stakingBalance int64
		providers      []models.ThirdPartyBaker
	}
	mp := make(map[string]baker)
	for _, item := range items {
		p := mp[item.Address].providers
		balance := item.StakingBalance
		if balance > 0 {
			balance = item.StakingBalance
		}
		mp[item.Address] = baker{
			address:        item.Address,
			stakingBalance: balance,
			providers:      append(p, item),
		}
	}
	var bakers []baker
	for _, item := range mp {
		bakers = append(bakers, baker{
			address:        item.address,
			stakingBalance: item.stakingBalance,
			providers:      item.providers,
		})
	}
	sort.Slice(bakers, func(i, j int) bool {
		return bakers[i].stakingBalance > bakers[j].stakingBalance
	})
	for _, b := range bakers {
		var providers []*genModels.ThirdPartyBakersProvidersItems0
		for _, p := range b.providers {
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
			Baker:     b.address,
			Providers: providers,
		})
	}
	return result
}
