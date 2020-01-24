package render

import (
	genModels "github.com/everstake/teztracker/gen/models"
	"github.com/everstake/teztracker/models"
)

// OperationGroup renders an app level model to a gennerated OpenAPI model.
func OperationGroup(b models.OperationGroup) *genModels.OperationGroupsRow {
	return &genModels.OperationGroupsRow{
		Protocol:  b.Protocol.Ptr(),
		ChainID:   b.ChainID,
		Hash:      b.Hash.Ptr(),
		Branch:    b.Branch.Ptr(),
		Signature: b.Signature,
		BlockID:   b.BlockID.Ptr(),
	}
}
