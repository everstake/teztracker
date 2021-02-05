package thirdparty_bakers

import (
	"context"
	"fmt"
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/repos/thirdparty_bakers"
	"github.com/everstake/teztracker/services/thirdparty_bakers/bakingbad"
	"github.com/everstake/teztracker/services/thirdparty_bakers/mytezosbaker"
	"github.com/everstake/teztracker/services/thirdparty_bakers/tezosnodes"
)

const (
	BakingBadProvider    = "baking-bad"
	MyTezosBakerProvider = "mytezosbaker"
	TezosNodesProvider   = "tezos-nodes"
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
		BakingBadProvider:    bakingbad.New(),
		MyTezosBakerProvider: mytezosbaker.New(),
		TezosNodesProvider:   tezosnodes.New(),
	}
}

func UpdateBakers(ctx context.Context, unit UnitOfWork) error {
	bakersRepo := unit.GetThirdPartyBakers()
	var bakers []models.ThirdPartyBaker
	for name, provider := range getProviders() {
		items, err := provider.GetBakers()
		if err != nil {
			return fmt.Errorf("provider(%s): %s", name, err.Error())
		}
		for key := range items {
			items[key].Provider = name
		}
		bakers = append(bakers, items...)
	}
	ctx, cancel := context.WithCancel(context.Background())
	unit.Start(ctx)
	defer cancel()
	err := bakersRepo.DeleteAll()
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
