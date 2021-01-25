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
			Alias:       in[i].Alias,
			Address:     in[i].Address,
			Tag:         in[i].Address,
			Description: in[i].Description,
		}
	}
	return out
}

func UserNotesWithBalance(in []models.UserNoteWithBalance) []*genModels.UserNoteWithBalance {
	out := make([]*genModels.UserNoteWithBalance, len(in))
	for i := range in {
		out[i] = &genModels.UserNoteWithBalance{
			Alias:       &in[i].Alias,
			Address:     &in[i].Address,
			Tag:         &in[i].Address,
			Description: &in[i].Description,
			Balance:     &in[i].Balance.Int64,
		}
	}
	return out
}
