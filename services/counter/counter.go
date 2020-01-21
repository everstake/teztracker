package counter

import (
	"github.com/bullblock-io/tezTracker/models"
	"github.com/sirupsen/logrus"
)

type Counter struct {
}

func New() Counter {
	return Counter{}
}

type OpRepo interface {
	Count(ids, kinds, inBlocks, accountIDs []string, maxOperationID int64) (count int64, err error)
	Last() (operation models.Operation, err error)
}
type CounterRepo interface {
	Create(cntr models.OperationCounter) (id int64, err error)
}

func SaveCounterFor(kind string, repo OpRepo, cntRepo CounterRepo) error {
	lastOp, err := repo.Last()
	if err != nil {
		return err
	}
	cnt, err := repo.Count(nil, []string{kind}, nil, nil, lastOp.OperationID.Int64)
	if err != nil {
		return err
	}
	counter := models.OperationCounter{
		LastOperationID: lastOp.OperationID.Int64,
		OperationType:   kind,
		Count:           cnt,
	}
	_, err = cntRepo.Create(counter)
	return err
}
func SaveCounters(repo OpRepo, cntRepo CounterRepo) error {
	logrus.Tracef("Saving counters")
	kinds := []string{"endorsement", "proposals", "seed_nonce_revelation", "delegation", "transaction", "activate_account", "ballot", "origination", "reveal", "double_baking_evidence"}
	for i := range kinds {
		err := SaveCounterFor(kinds[i], repo, cntRepo)
		if err != nil {
			return err
		}
	}
	logrus.Tracef("Done saving counters")
	return nil
}
