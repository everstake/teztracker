package render

import (
	genModels "github.com/everstake/teztracker/gen/models"
	"github.com/everstake/teztracker/models"
)

// Snapshot renders an app level model to a gennerated OpenAPI model.
func Snapshot(b models.SnapshotView) *genModels.Snapshots {
	return &genModels.Snapshots{
		Cycle:         b.Snapshot.Cycle,
		CycleStart:    GetUnixFromNullTime(b.CycleStart),
		CycleEnd:      GetUnixFromNullTime(b.CycleEnd),
		Rolls:         b.Rolls,
		SnapshotBlock: b.BlockLevel,
	}

}

// Snapshots renders a slice of app level Snapshots into a slice of OpenAPI models.
func Snapshots(bs []models.SnapshotView) []*genModels.Snapshots {
	blocks := make([]*genModels.Snapshots, len(bs))
	for i := range bs {
		blocks[i] = Snapshot(bs[i])
	}
	return blocks
}
