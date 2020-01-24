package render

import (
	genModels "github.com/everstake/teztracker/gen/models"
	"github.com/everstake/teztracker/models"
)

// Snapshot renders an app level model to a gennerated OpenAPI model.
func Snapshot(b models.Snapshot) *genModels.Snapshots {
	return &genModels.Snapshots{
		Cycle:         b.Cycle,
		Rolls:         b.Rolls,
		SnapshotBlock: b.BlockLevel,
	}

}

// Snapshots renders a slice of app level Snapshots into a slice of OpenAPI models.
func Snapshots(bs []models.Snapshot) []*genModels.Snapshots {
	blocks := make([]*genModels.Snapshots, len(bs))
	for i := range bs {
		blocks[i] = Snapshot(bs[i])
	}
	return blocks
}
