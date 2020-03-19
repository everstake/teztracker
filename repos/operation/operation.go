package operation

import (
	"github.com/everstake/teztracker/models"
	"github.com/jinzhu/gorm"
)

//go:generate mockgen -source ./operation.go -destination ./mock_operation/main.go Repo
type (
	// Repository is the operation repo implementation.
	Repository struct {
		db *gorm.DB
	}

	Repo interface {
		List(ids, kinds []string, inBlocks, accountIDs []string, limit, offset uint, since int64) (operations []models.Operation, err error)
		ListAsc(kinds []string, limit, offset uint, after int64) (operations []models.Operation, err error)
		Count(ids, kinds, inBlocks, accountIDs []string, maxOperationID int64) (count int64, err error)
		EndorsementsFor(blockLevel int64) (operations []models.Operation, err error)
		Last() (operation models.Operation, err error)
		ListDoubleEndorsementsWithoutLevel(limit, offset uint) (operations []models.Operation, err error)
		UpdateLevel(operation models.Operation) error
		AccountOperationCount(string) ([]models.OperationCount, error)
	}
)

const endorsementKind = "endorsement"

// New creates an instance of repository using the provided db.
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Count counts a number of operations sutisfying the filter.
func (r *Repository) Count(ids, kinds, inBlocks, accountIDs []string, maxOperationID int64) (count int64, err error) {
	db := r.getFilteredDB(ids, kinds, inBlocks, accountIDs)
	if maxOperationID > 0 {
		db = db.Where("operation_id <= ?", maxOperationID)
	}
	snapshotCount := int64(0)
	if len(ids) == 0 && len(inBlocks) == 0 && len(accountIDs) == 0 && len(kinds) == 1 {
		counter := models.OperationCounter{}
		lastCounterDb := r.db.Model(&counter).Where("cnt_operation_type = ?", kinds[0])
		if err := lastCounterDb.Last(&counter).Error; err == nil {
			db = db.Where("operation_id > ?", counter.LastOperationID)
			snapshotCount = counter.Count
		}

	}
	err = db.Count(&count).Error
	return count + snapshotCount, err
}

func (r *Repository) getFilteredDB(ids, kinds []string, inBlocks, accountIDs []string) *gorm.DB {
	db := r.db.Model(&models.Operation{})
	if len(ids) > 0 {
		db = db.Where("operation_group_hash IN (?)", ids)
	}
	if len(kinds) > 0 {
		db = db.Where("kind IN (?)", kinds)
	}

	if len(inBlocks) > 0 {
		db = db.Where("block_hash IN (?)", inBlocks)
	}
	if len(accountIDs) > 0 {
		if len(kinds) == 1 && kinds[0] == "transaction" {
			db = db.Where("source IN (?) OR destination IN (?)", accountIDs, accountIDs)
		} else {
			db = db.Where("delegate IN (?) OR pkh IN (?) OR source IN (?) OR public_key IN (?) OR destination IN (?)", accountIDs, accountIDs, accountIDs, accountIDs, accountIDs)
		}
	}
	return db
}

// List returns a list of operations from the newest to oldest.
// limit defines the limit for the maximum number of operations returned.
// since is used to paginate results based on the operation id.
// As the result is ordered descendingly the operations with operation_id < since will be returned.
func (r *Repository) List(ids, kinds []string, inBlocks, accountIDs []string, limit, offset uint, since int64) (operations []models.Operation, err error) {
	db := r.getFilteredDB(ids, kinds, inBlocks, accountIDs)

	if since > 0 {
		db = db.Where("operation_id < ?", since)
	}
	err = db.Order("operation_id desc").
		Limit(limit).
		Offset(offset).
		Find(&operations).Error
	return operations, err
}

func (r *Repository) ListDoubleEndorsementsWithoutLevel(limit, offset uint) (operations []models.Operation, err error) {
	db := r.db.Model(&models.Operation{}).Where("kind IN (?)", []string{"double_endorsement_evidence"}).Where("level is null")
	err = db.Order("operation_id asc").
		Limit(limit).
		Offset(offset).
		Find(&operations).Error
	return operations, err
}

func (r *Repository) UpdateLevel(operation models.Operation) error {
	return r.db.Select("level").Save(&operation).Error
}

func (r *Repository) ListAsc(kinds []string, limit, offset uint, after int64) (operations []models.Operation, err error) {
	db := r.getFilteredDB(nil, kinds, nil, nil)

	if after > 0 {
		db = db.Where("operation_id > ?", after)
	}
	err = db.Order("operation_id asc").
		Limit(limit).
		Offset(offset).
		Find(&operations).Error
	return operations, err
}

// EndorsementsFor returns a list of endorsement operations for the provided block level.
func (r *Repository) EndorsementsFor(blockLevel int64) (operations []models.Operation, err error) {
	err = r.db.Select("*").Model(&models.Operation{}).
		Joins("left join tezos.balance_updates on (operations.operation_group_hash = balance_updates.operation_group_hash and category = 'rewards')").
		Where("operations.kind = ?", endorsementKind).
		// the endorsements of the block with blockLevel can only be in a block with level (blockLevel + 1)
		Where("block_level = ?", blockLevel+1).
		Order("operation_id DESC").
		Find(&operations).Error
	return operations, err
}

// Last returns the last known operation.
func (r *Repository) Last() (operation models.Operation, err error) {
	db := r.db.Model(&operation)
	err = db.Last(&operation).Error
	return operation, err
}

func (r *Repository) AccountOperationCount(acc string) (counts []models.OperationCount, err error) {
	db := r.getFilteredDB(nil, nil, nil, []string{acc})

	err = db.Select("kind,count(1)").
		Group("kind").
		Scan(&counts).Error
	if err != nil {
		return counts, err
	}

	return counts, nil
}
