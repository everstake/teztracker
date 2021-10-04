package render

import (
	genModels "github.com/everstake/teztracker/gen/models"
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/services/whales"
)

func WhaleAccounts(data whales.Data) *genModels.WhaleAccounts {
	accounts := make([]*genModels.WhaleAccountsAccountsItems0, len(data.Accounts))
	for i := range data.Accounts {
		accounts[i] = &genModels.WhaleAccountsAccountsItems0{
			Address: data.Accounts[i].AccountID.String,
			Balance: data.Accounts[i].Balance.Int64,
		}
	}
	transfers := make([]*genModels.WhaleAccountsTransfersItems0, len(data.Transfers))
	for i := range data.Transfers {
		transfers[i] = &genModels.WhaleAccountsTransfersItems0{
			Amount: data.Transfers[i].Amount,
			From:   data.Transfers[i].Source,
			To:     data.Transfers[i].Destination,
		}
	}
	return &genModels.WhaleAccounts{
		Accounts:  accounts,
		Transfers: transfers,
	}
}

func WhaleAccountList(data []models.Account) (acs []*genModels.WhaleAccount) {
	acs = make([]*genModels.WhaleAccount, len(data))

	for i := range data {
		acs[i] = &genModels.WhaleAccount{
			AccountID: data[i].AccountID.String,
			Amount:    data[i].Balance.Int64,
		}

	}

	return acs
}
