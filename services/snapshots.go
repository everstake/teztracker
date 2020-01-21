package services

import (
	"github.com/bullblock-io/tezTracker/models"
)

func (t *TezTracker) Snapshots(limiter Limiter) (count int64, snapshots []models.Snapshot, err error) {
	repo := t.repoProvider.GetSnapshots()
	return repo.List(limiter.Limit(), limiter.Offset())
}
