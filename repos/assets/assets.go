package assets

import (
	"github.com/everstake/teztracker/models"
	"github.com/jinzhu/gorm"
)

type (
	// Repository is the account repo implementation.
	Repository struct {
		db *gorm.DB
	}

	Repo interface {
		GetTokensList() (int64, []models.AssetInfo, error)
		GetTokenInfo(tokenID string) (models.AssetInfo, error)
		GetTokenHolders(tokenID string) ([]models.AssetHolder, error)
		GetAssetOperations(tokenIDs, operationTypes, accountIDs []string, limit, offset uint) (count int64, info []models.AssetOperationReport, err error)
		GetUnprocessedAssetTxs(tokenID string) ([]models.Operation, error)
		CreateAssetOperations(models.AssetOperation) error
	}
)

// New creates an instance of repository using the provided db.
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

const (
	assetsTable     = "tezos.asset_info"
	assetOperations = "tezos.asset_operations"
)

func (r *Repository) CreateAssetOperations(operation models.AssetOperation) error {
	// Create creates a AssetOperation.
	return r.db.Model(&operation).Create(&operation).Error
}

func (r *Repository) GetTokensList() (count int64, tokens []models.AssetInfo, err error) {
	db := r.db.Table(assetsTable)

	err = db.Count(&count).Error
	if err != nil {
		return 0, nil, err
	}

	err = db.Select("*").Find(&tokens).Error
	if err != nil {
		return 0, nil, err
	}

	return count, tokens, nil
}

func (r *Repository) GetTokenInfo(tokenID string) (info models.AssetInfo, err error) {
	db := r.db.Select("*").Table(assetsTable).
		Where("account_id = ?", tokenID)

	err = db.Find(&info).Error

	return info, err
}

func (r *Repository) GetTokenHolders(tokenID string) (holders []models.AssetHolder, err error) {

	return holders, err
}

func (r *Repository) GetUnprocessedAssetTxs(tokenID string) (ops []models.Operation, err error) {
	db := r.db.Select("operations.operation_group_hash").
		Model(&models.Operation{}).
		Joins("LEFT OUTER JOIN tezos.asset_operations on operations.operation_group_hash = asset_operations.operation_group_hash").
		Where("source = ? OR destination = ?", tokenID, tokenID).
		Where("status = ?", "applied").
		Where("asset_operations.operation_group_hash IS NULL").
		Group("operations.operation_group_hash")

	err = db.Find(&ops).Error
	return ops, nil
}

func (r *Repository) GetAssetOperations(tokenIDs, operationTypes, accountIDs []string, limit, offset uint) (count int64, info []models.AssetOperationReport, err error) {

	db := r.db.Select("*").Table(assetOperations)

	if len(tokenIDs) > 0 {
		db = db.Joins("LEFT JOIN tezos.registered_tokens on id = token_id")
		db = db.Where("account_id IN (?)", tokenIDs)
	}

	if len(operationTypes) == 1 {
		switch operationTypes[0] {
		case "transfer":
			db = db.Where("type = ?", "transfer")
		case "other":
			db = db.Where("type != ?", "transfer")
		}
	}

	if len(accountIDs) > 0 {
		db = db.Where("asset_operations.sender IN (?) OR asset_operations.receiver IN (?)", accountIDs, accountIDs)
	}

	err = db.Count(&count).Error
	if err != nil {
		return 0, nil, err
	}

	db = db.Joins("LEFT JOIN tezos.operations on (asset_operations.operation_group_hash=operations.operation_group_hash and internal is not TRUE)").
		Order("asset_operations.timestamp desc").Limit(limit).Offset(offset)

	err = db.Find(&info).Error
	if err != nil {
		return count, info, err
	}

	return count, info, nil
}
