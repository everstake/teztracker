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
		GetTokensList() ([]models.AssetInfo, error)
		GetTokenInfo(tokenID string) (models.AssetInfo, error)
		GetTokenHolders(tokenID string) ([]models.AssetHolder, error)
		GetAssetOperations(tokenID uint64, isTransfer bool, limit, offset uint) (info []models.AssetOperationReport, err error)
		GetAssetTxs(tokenID string) ([]models.Operation, error)
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

func (r *Repository) GetTokensList() (tokens []models.AssetInfo, err error) {
	err = r.db.Select("*").Table(assetsTable).Find(&tokens).Error
	if err != nil {
		return nil, err
	}
	return tokens, nil
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

func (r *Repository) GetAssetTxs(tokenID string) (ops []models.Operation, err error) {
	db := r.db.Select("*").
		Model(&models.Operation{}).
		Where("source = ? OR destination = ?", tokenID, tokenID).
		Order("operation_id asc")

	err = db.Find(&ops).Error
	return ops, nil
}

func (r *Repository) GetAssetOperations(tokenID uint64, isTransfer bool, limit, offset uint) (info []models.AssetOperationReport, err error) {

	db := r.db.Select("*").Table(assetOperations).
		Joins("LEFT JOIN tezos.operations on (asset_operations.operation_group_hash=operations.operation_group_hash and internal is not TRUE)").
		Where("token_id = ?", tokenID)

	if isTransfer {
		db = db.Where("type = ?", "transfer")
	} else {
		db = db.Where("type != ?", "transfer")
	}

	err = db.Find(&info).Error

	return info, err
}
