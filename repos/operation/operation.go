package operation

import (
	"fmt"
	"github.com/everstake/teztracker/models"
	"github.com/go-openapi/validate"
	"github.com/jinzhu/gorm"
)

//go:generate mockgen -source ./operation.go -destination ./mock_operation/main.go Repo
type (
	// Repository is the operation repo implementation.
	Repository struct {
		db *gorm.DB
	}

	Repo interface {
		List(ids, kinds []string, inBlocks, accountIDs []string, limit, offset uint, since int64, operationsIDs []int64) (operations []models.Operation, err error)
		ListAsc(kinds []string, limit, offset uint, after int64) (operations []models.Operation, err error)
		Count(ids, kinds, inBlocks, accountIDs []string, maxOperationID int64) (count int64, err error)
		EndorsementsFor(blockLevel int64) (operations []models.Operation, err error)
		Last() (operation models.Operation, err error)
		ListDoubleEndorsementsWithoutLevel(limit, offset uint) (operations []models.Operation, err error)
		UpdateLevel(operation models.Operation) error
		AccountOperationCount(string) ([]models.OperationCount, error)
		AccountEndorsements(accountID string, cycle int64, limit uint, offset uint) (count int64, operations []models.Operation, err error)
	}
)

const (
	endorsementKind = "endorsement"
	delegationKind  = "delegation"
	activationKind  = "activate_account"
)

// New creates an instance of repository using the provided db.
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Count counts a number of operations sutisfying the filter.
func (r *Repository) Count(ids, kinds, inBlocks, accountIDs []string, maxOperationID int64) (count int64, err error) {
	db := r.getFilteredDB(ids, kinds, inBlocks, accountIDs, nil, true)
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

//TODO refactor to preload for all types
func (r *Repository) getFilteredDB(hashes, kinds, inBlocks, accountIDs []string, operationIDs []int64, count bool) (db *gorm.DB) {
	selectQ := "*"
	db = r.db.Model(&models.Operation{})

	if len(hashes) > 0 || len(operationIDs) > 0 {
		//Join for extend info
		if !count && len(kinds) == 0 {

			kindMap := map[string]bool{}
			//Preload operations by hashes

			//Preload operations
			preloadDb := r.db
			if len(operationIDs) == 0 {
				preloadDb = preloadDb.Where("operations.operation_group_hash IN (?)", hashes)
			} else {
				preloadDb = preloadDb.Where("operations.operation_id IN (?)", operationIDs)
			}
			op := []models.Operation{}
			err := preloadDb.Model(&models.Operation{}).Find(&op).Error
			if err != nil {
				return
			}

			var operationIds []int64
			for key := range op {
				kindMap[op[key].Kind.String] = true
				operationIDs = append(operationIds, op[key].OperationID.Int64)
			}

			db = db.Where("operations.operation_id IN (?)", operationIDs)

			for kind := range kindMap {
				switch kind {
				case delegationKind:
					db = db.Joins("left join tezos.accounts_history as ah on (ah.block_level=operations.block_level and account_id=source and operations.kind='delegation')")
				case endorsementKind:
					selectQ = fmt.Sprintf("%s, %s", selectQ, "bur.change endorsement_reward, bud.change endorsement_deposit")
					db = db.Joins("left join tezos.balance_updates as bur on (operations.operation_group_hash = bur.operation_group_hash and bur.category='rewards')").
						Joins("left join tezos.balance_updates as bud on (operations.operation_group_hash = bud.operation_group_hash and bud.category='deposits')")
				case activationKind:
					selectQ = fmt.Sprintf("%s, %s", selectQ, "bua.change claimed_amount")
					db = db.Joins("left join tezos.balance_updates as bua on (operations.operation_group_hash = bua.operation_group_hash and bua.kind='contract')")
				}
			}
		} else {
			db = db.Where("operations.operation_group_hash IN (?)", hashes)
		}
	}

	if len(kinds) > 0 {
		if !count && validate.Enum("", "", delegationKind, kinds) == nil {
			db = db.Joins("left join tezos.accounts_history as ah on (ah.block_level=operations.block_level and account_id=source)")
		}

		db = db.Where("operations.kind IN (?)", kinds)
	}

	if len(inBlocks) > 0 {
		db = db.Where("operations.block_hash IN (?)", inBlocks)
	}

	if len(accountIDs) > 0 {
		if len(kinds) == 1 && kinds[0] == "transaction" {
			db = db.Where("operations.source IN (?) OR operations.destination IN (?)", accountIDs, accountIDs)
		} else {
			db = db.Where("operations.delegate IN (?) OR operations.pkh IN (?) OR operations.source IN (?) OR operations.public_key IN (?) OR operations.destination IN (?) OR operations.originated_contracts IN (?)", accountIDs, accountIDs, accountIDs, accountIDs, accountIDs, accountIDs)
		}
	}
	//Join Aliases
	if !count {
		selectQ = fmt.Sprintf("%s, %s", selectQ, "des.baker_name as destination_name, s.baker_name as source_name, del.baker_name as delegate_name")
		db = db.Joins("left join tezos.public_bakers as des on operations.destination = des.delegate").
			Joins("left join tezos.public_bakers as s on operations.source = s.delegate").
			Joins("left join tezos.public_bakers as del on operations.delegate = del.delegate")
	}

	db = db.Select(selectQ)
	return db
}

// List returns a list of operations from the newest to oldest.
// limit defines the limit for the maximum number of operations returned.
// since is used to paginate results based on the operation id.
// As the result is ordered descendingly the operations with operation_id < since will be returned.
func (r *Repository) List(hashes, kinds []string, inBlocks, accountIDs []string, limit, offset uint, since int64, operationsIDs []int64) (operations []models.Operation, err error) {
	db := r.getFilteredDB(hashes, kinds, inBlocks, accountIDs, operationsIDs, false)

	if since > 0 {
		db = db.Where("operations.operation_id < ?", since)
	}

	db = db.Order("operations.operation_id desc").
		Limit(limit).
		Offset(offset)

	//TODO Join with baker_endorsements
	if len(inBlocks) == 1 && len(kinds) == 1 && kinds[0] == endorsementKind {
		db = r.db.Raw("SELECT * from (?) as s left join tezos.balance_updates on (s.operation_group_hash = balance_updates.operation_group_hash and category = 'rewards')", db.SubQuery())
	}

	err = db.Find(&operations).Error

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
	db := r.getFilteredDB(nil, kinds, nil, nil, nil, false)

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
		Where("operations.block_level = ?", blockLevel+1).
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
	db := r.getFilteredDB(nil, nil, nil, []string{acc}, nil, true)

	err = db.Select("kind,count(1)").
		Group("kind").
		Scan(&counts).Error
	if err != nil {
		return counts, err
	}

	return counts, nil
}

func (r *Repository) AccountEndorsements(accountID string, cycle int64, limit uint, offset uint) (count int64, operations []models.Operation, err error) {
	db := r.db.Model(&models.Operation{}).
		Where("delegate = ?", accountID).
		Where("kind = ?", endorsementKind).
		Where("cycle = ?", cycle)

	err = db.Count(&count).Error
	if err != nil {
		return count, operations, err
	}

	db = db.Order("operation_id desc").
		Limit(limit).
		Offset(offset)

	//TODO Join with baker_endorsements
	db = r.db.Raw("SELECT * from (?) as s left join tezos.balance_updates on (s.operation_group_hash = balance_updates.operation_group_hash and category = 'rewards')", db.SubQuery())
	err = db.Find(&operations).Error

	return count, operations, nil
}
