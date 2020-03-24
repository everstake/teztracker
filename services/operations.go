package services

import (
	"strconv"

	"github.com/everstake/teztracker/models"
	"github.com/guregu/null"
)

// GetOperations gets the operations filtering by operation kinds and blocks wiht pagination.
func (t *TezTracker) GetOperations(ids, kinds, inBlocks, accountIDs []string, limits Limiter, before int64) (operations []models.Operation, count int64, err error) {
	r := t.repoProvider.GetOperation()
	count, err = r.Count(ids, kinds, inBlocks, accountIDs, 0)
	if err != nil {
		return nil, 0, err
	}

	operations, err = r.List(ids, kinds, inBlocks, accountIDs, limits.Limit(), limits.Offset(), before)
	return operations, count, err
}

// GetBlockEndorsements finds a block and returns endorsements for it.
func (t *TezTracker) GetBlockEndorsements(hashOrLevel string) (operations []models.Operation, count int64, err error) {
	r := t.repoProvider.GetBlock()
	var filter models.Block
	if i, e := strconv.ParseInt(hashOrLevel, 10, 64); e == nil {
		filter.Level = null.IntFrom(i)
	} else {
		filter.Hash = null.StringFrom(hashOrLevel)
	}
	found, block, err := r.Find(filter)
	if err != nil {
		return nil, 0, err
	}
	if !found {
		return nil, 0, ErrNotFound
	}
	or := t.repoProvider.GetOperation()
	operations, err = or.EndorsementsFor(block.Level.Int64)
	return operations, int64(len(operations)), err
}

// GetOperations gets the operations filtering by operation kinds and blocks wiht pagination.
func (t *TezTracker) GetDoubleBakings(ids, inBlocks []string, limits Limiter) (operations []models.DoubleBakingEvidence, count int64, err error) {
	r := t.repoProvider.GetDoubleBaking()
	options := models.DoubleBakingEvidenceQueryOptions{
		BlockIDs:        inBlocks,
		OperationHashes: ids,
		LoadOperation:   true,
		Limit:           limits.Limit(),
		Offset:          limits.Offset(),
	}
	count, operations, err = r.List(options)
	return operations, count, err
}
