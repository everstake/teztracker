package render

import (
	genModels "github.com/everstake/teztracker/gen/models"
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

func WhaleTransfers(data whales.Data) []*genModels.LargeTransfer {
	transfers := make([]*genModels.LargeTransfer, len(data.LargeTransfers))
	for i := range data.LargeTransfers {
		transfers[i] = &genModels.LargeTransfer{
			Amount: data.LargeTransfers[i].Amount,
			From:   data.LargeTransfers[i].Source,
			To:     data.LargeTransfers[i].Destination,
		}
	}
	return transfers
}
