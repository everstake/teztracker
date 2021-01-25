package assets

import (
	"fmt"
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
		GetAssetOperations(tokenIDs, operationTypes, accountIDs []string, blockLevels []int64, limit, offset uint) (count int64, info []models.AssetOperationReport, err error)
		GetAssetReport(tokenID uint64, params models.ReportFilter) ([]models.AssetReport, error)
		GetAccountAssetsBalances(hexAddress string) (holders []models.AccountAssetBalance, err error)
		GetUnprocessedAssetTxs(tokenID string) ([]models.Operation, error)
		CreateAssetOperations(models.AssetOperation) error
		FindOperations(operationIDs []int64, limit uint64) (operations []models.AssetOperation, err error)
		GetRegisteredToken(tokenID uint64) (token models.RegisteredToken, err error)
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
	err = r.db.
		Select("key as address, value balance").
		Table("tezos.big_map_contents").
		Joins("LEFT JOIN tezos.originated_account_maps oam on oam.big_map_id = big_map_contents.big_map_id").
		Joins("LEFT JOIN tezos.big_maps bms on oam.big_map_id = bms.big_map_id").
		Where("account_id = ?", tokenID).
		//Get only ledger bigmap
		Where("key_type IN ('address','bytes')").
		Where("value is not null").
		Find(&holders).Error
	return holders, err
}

func (r *Repository) GetAccountAssetsBalances(hexAddress string) (holders []models.AccountAssetBalance, err error) {

	err = r.db.
		Select("key as address, value balance, asset_info.*").
		Table("tezos.big_map_contents").
		Joins("LEFT JOIN tezos.originated_account_maps oam on oam.big_map_id = big_map_contents.big_map_id").
		Joins("LEFT JOIN tezos.asset_info on asset_info.account_id = oam.account_id").
		Where("key like ?", fmt.Sprint("%", hexAddress, "%")).
		Where("value is not null").
		Find(&holders).Error
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

func (r *Repository) GetAssetOperations(tokenIDs, operationTypes, accountIDs []string, blockLevels []int64, limit, offset uint) (count int64, info []models.AssetOperationReport, err error) {

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

	if len(blockLevels) > 0 {
		db = db.Where("asset_operations.block_level IN (?)", blockLevels)
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

func (r *Repository) FindOperations(operationIDs []int64, limit uint64) (operations []models.AssetOperation, err error) {
	q := r.db.Model(&models.AssetOperation{})
	if len(operationIDs) > 0 {
		q = q.Where("operation_id in (?)", operationIDs)
	}
	if limit != 0 {
		q = q.Limit(limit)
	}
	err = q.Find(&operations).Error
	return operations, err
}

func (r *Repository) GetRegisteredToken(tokenID uint64) (token models.RegisteredToken, err error) {
	err = r.db.Model(&models.RegisteredToken{}).Where("id = ?", tokenID).First(&token).Error
	return token, err
}

func (r *Repository) GetAssetReport(tokenID uint64, params models.ReportFilter) (report []models.AssetReport, err error) {
	err = r.db.
		Select("block_level, timestamp, type kind, operation_group_hash, ticker coin, amount :: decimal / (10 ^ scale) amount, 'applied' status, sender source, receiver destination, 0 fee").
		Table("tezos.asset_operations").
		Joins("left join tezos.registered_tokens on token_id = id").
		Where("token_id = ?", tokenID).
		Where("timestamp >= to_timestamp(?) :: timestamp without time zone", params.From).
		Where("timestamp <= to_timestamp(?) :: timestamp without time zone", params.To).
		Limit(params.Limit).
		Order("timestamp desc").
		Find(&report).Error
	if err != nil {
		return nil, err
	}

	return report, nil
}
