package thirdparty_bakers

import (
	"context"
	"fmt"
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/repos/thirdparty_bakers"
	"github.com/everstake/teztracker/services/thirdparty_bakers/bakingbad"
	"github.com/everstake/teztracker/services/thirdparty_bakers/tezosnodes"
	"github.com/everstake/teztracker/services/thirdparty_bakers/tzstats"
)

const (
	BakingBadProvider    = "baking-bad"
	MyTezosBakerProvider = "mytezosbaker"
	TezosNodesProvider   = "tezos-nodes"
	TzstatsProvider      = "tzstats"
	mainProvider         = BakingBadProvider
)

type (
	BakersProvider interface {
		GetBakers() (bakers []models.ThirdPartyBaker, err error)
	}
	UnitOfWork interface {
		Start(ctx context.Context)
		Commit() error
		GetThirdPartyBakers() thirdparty_bakers.Repo
	}
)

type ThirdPartyBakers struct {
	bakersRepo thirdparty_bakers.Repo
	providers  map[string]BakersProvider
}

func getProviders() map[string]BakersProvider {
	return map[string]BakersProvider{
		BakingBadProvider:  bakingbad.New(),
		TezosNodesProvider: tezosnodes.New(),
		TzstatsProvider:    tzstats.New(),
	}
}

func UpdateBakers(ctx context.Context, unit UnitOfWork) error {
	bakersRepo := unit.GetThirdPartyBakers()
	var bakers []models.ThirdPartyBaker
	providers := getProviders()
	mainBakerProvider, ok := providers[mainProvider]
	if !ok {
		return fmt.Errorf("not found main provider")
	}
	addressesWhiteList := make(map[string]struct{})
	mainProviderBakers, err := mainBakerProvider.GetBakers()
	if err != nil {
		return fmt.Errorf("mainBakerProvider.GetBakers: %s", err.Error())
	}
	for _, baker := range mainProviderBakers {
		addressesWhiteList[baker.Address] = struct{}{}
	}
	for name, provider := range providers {
		items, err := provider.GetBakers()
		if err != nil {
			return fmt.Errorf("provider(%s): %s", name, err.Error())
		}
		var validItems []models.ThirdPartyBaker
		for key := range items {
			if _, ok := addressesWhiteList[items[key].Address]; !ok {
				continue
			}
			item := items[key]
			item.Provider = name
			validItems = append(validItems, item)
		}
		bakers = append(bakers, validItems...)
	}
	ctx, cancel := context.WithCancel(context.Background())
	unit.Start(ctx)
	defer cancel()
	err = bakersRepo.DeleteAll()
	if err != nil {
		return fmt.Errorf("bakersRepo.DeleteAll: %s", err.Error())
	}
	err = bakersRepo.Create(bakers)
	if err != nil {
		return fmt.Errorf("bakersRepo.Create: %s", err.Error())
	}
	err = unit.Commit()
	if err != nil {
		return fmt.Errorf("unit.Commit: %s", err.Error())
	}
	return nil
}
