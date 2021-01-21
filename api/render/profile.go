package render

import (
	genModels "github.com/everstake/teztracker/gen/models"
	"github.com/everstake/teztracker/models"
)

func UserAddresses(in []models.UserAddressWithBalance) []*genModels.UserAddressWithBalance {
	out := make([]*genModels.UserAddressWithBalance, len(in))
	for i := range in {
		out[i] = &genModels.UserAddressWithBalance{
			Address:             &in[i].Address,
			DelegationsEnabled:  &in[i].DelegationsEnabled,
			InTransfersEnabled:  &in[i].InTransfersEnabled,
			OutTransfersEnabled: &in[i].OutTransfersEnabled,
			Balance:             &in[i].Balance.Int64,
		}
	}
	return out
}

func UserNotes(in []models.UserNote) []*genModels.UserNote {
	out := make([]*genModels.UserNote, len(in))
	for i := range in {
		out[i] = &genModels.UserNote{
			Alias: in[i].Alias,
			Text:  in[i].Text,
		}
	}
	return out
}
