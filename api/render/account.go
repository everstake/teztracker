package render

import (
	genModels "github.com/everstake/teztracker/gen/models"
	"github.com/everstake/teztracker/models"
)

// Account renders an app level model to a gennerated OpenAPI model.
func Account(a models.Account) *genModels.AccountsRow {
	return &genModels.AccountsRow{
		AccountID:       a.AccountID.Ptr(),
		BlockID:         a.BlockID.Ptr(),
		Manager:         a.Manager.Ptr(),
		Spendable:       a.Spendable.Ptr(),
		DelegateSetable: a.DelegateSetable.Ptr(),
		DelegateValue:   a.DelegateValue,
		Counter:         a.Counter.Ptr(),
		Script:          a.Script,
		Storage:         a.Storage,
		Balance:         a.Balance.Ptr(),
		BlockLevel:      a.BlockLevel.Ptr(),
		BakerInfo:       BakerInfo(a.BakerInfo),
	}
}

// Accounts renders a slice of app level Accounts into a slice of OpenAPI models.
func Accounts(ams []models.Account) []*genModels.AccountsRow {
	accs := make([]*genModels.AccountsRow, len(ams))
	for i := range ams {
		accs[i] = Account(ams[i])
	}
	return accs
}
