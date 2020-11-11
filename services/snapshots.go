package services

import (
	"github.com/everstake/teztracker/models"
)

func (t *TezTracker) Snapshots(limiter Limiter) (count int64, snapshots []models.SnapshotsView, err error) {
	repo := t.repoProvider.GetSnapshots()
	return repo.List(limiter.Limit(), limiter.Offset())
}
