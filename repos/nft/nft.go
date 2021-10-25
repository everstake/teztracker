package nft

import (
	"fmt"
	"time"

	"github.com/everstake/teztracker/models"
	"github.com/jinzhu/gorm"
	gormbulk "github.com/t-tiger/gorm-bulk-insert"
)

type (
	// Repository is the account repo implementation.
	Repository struct {
		db *gorm.DB
	}

	Repo interface {
		NTFContractsList(string, uint, uint) ([]models.NFTContract, int64, error)
		TokensList(int64, *int64, uint, uint) ([]models.NFTToken, int64, error)

		UpdateNTFContractLastHeight(contract models.NFTContract) error
		UpdateNTFContractLastOPHeight(contract models.NFTContract) error

		CreateBulk(rights []models.NFTToken) error
		UpdateNFTToken(contractID int64, tokenID uint64, isForSale bool, lastPrice *int64, lastActive time.Time) error

		ContractTokenHolders(mapID, limit int64) (tokens []models.AssetHolder, count int64, err error)
		TokenHoldersList(mapID int64, tokenID *int64, limit, offset uint) (tokens []models.AssetHolder, count int64, err error)
		TokenHoldersCount(mapID, tokenNum int64, isEqual bool) (int64, error)
	}
)

const (
	nftContractsViewTable = "tezos.nft_contracts_view"
)

// New creates an instance of repository using the provided db.
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) NTFContractsList(contractID string, limit, offset uint) (contracts []models.NFTContract, count int64, err error) {

	db := r.db.
		Table(nftContractsViewTable).
		Order("id desc").
		Limit(limit).
		Offset(offset)

	if contractID != "" {
		db = db.Where("account_id = ?", contractID)
	}

	err = db.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Find(&contracts).Error
	if err != nil {
		return nil, 0, err
	}

	return contracts, count, nil
}

func (r *Repository) TokensList(contractID int64, tokenID *int64, limit uint, offset uint) (tokens []models.NFTToken, count int64, err error) {

	db := r.db.
		Table("tezos.nft_tokens").
		Where("contract_id = ?", contractID)

	if tokenID != nil {
		db = db.Where("token_id = ?", tokenID)
	}

	err = db.Count(&count).Error
	if err != nil {
		return tokens, count, err
	}

	err = db.Order("token_id desc").
		Limit(limit).
		Offset(offset).
		Find(&tokens).
		Error
	if err != nil {
		return tokens, count, err
	}

	return tokens, count, nil
}

func (r *Repository) ContractTokenHolders(mapID, limit int64) (holders []models.AssetHolder, count int64, err error) {

	db := r.db.
		Select("address, sum(value :: integer) balance").
		Table("tezos.nft_tokens_ledger_view").
		Where("big_map_id = ?", mapID).
		Group("address")

	err = db.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Order("balance desc").
		Limit(limit).Find(&holders).Error
	if err != nil {
		return nil, 0, err
	}

	return
}

func (r *Repository) TokenHoldersList(mapID int64, tokenID *int64, limit uint, offset uint) (holders []models.AssetHolder, count int64, err error) {

	db := r.db.
		Select("key as address, value balance").
		Table("tezos.big_map_contents").
		Where("big_map_id = ?", mapID).
		Where("value <> '0'")

	if tokenID != nil {
		db = db.Where("key like (('Pair % '::text) || (?::text))", tokenID)
	}

	err = db.Count(&count).Error
	if err != nil {
		return holders, 0, err
	}

	err = db.Order("value desc").
		Find(&holders).Error
	if err != nil {
		return holders, 0, err
	}

	return holders, count, nil
}

func (r *Repository) TokenHoldersCount(mapID, tokenNum int64, isEqual bool) (count int64, err error) {

	db := r.db.Select("address, count(1) tokens_num").
		Table("tezos.nft_tokens_ledger_view").
		Where("big_map_id = ?", mapID).
		Group("address")

	eqSign := ">"
	if isEqual {
		eqSign = "="
	}

	db = r.db.Raw(fmt.Sprintf("SELECT count(1) num from ? s where tokens_num %s ?", eqSign), db.SubQuery(), tokenNum)

	err = db.Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *Repository) UpdateNTFContractLastHeight(contract models.NFTContract) error {
	return r.db.Table("tezos.nft_contracts").Where("id = ?", contract.ID).
		Updates(map[string]interface{}{
			"last_height":    contract.LastHeight,
			"operations_num": contract.OperationsNum,
		}).Error
}

func (r *Repository) UpdateNTFContractLastOPHeight(contract models.NFTContract) error {
	return r.db.Table("tezos.nft_contracts").Where("id = ?", contract.ID).
		Updates(map[string]interface{}{
			"last_update_height": contract.LastUpdateHeight,
		}).Error
}

func (r *Repository) UpdateNFTToken(contractID int64, tokenID uint64, isForSale bool, lastPrice *int64, lastActive time.Time) error {
	updates := map[string]interface{}{
		"is_for_sale":    isForSale,
		"last_active_at": lastActive,
	}

	if lastPrice != nil {
		updates["last_price"] = lastPrice
	}

	return r.db.Table("tezos.nft_tokens").
		Where("contract_id = ? and token_id = ?", contractID, tokenID).
		Updates(updates).Error
}

func (r *Repository) CreateBulk(rights []models.NFTToken) error {
	if len(rights) == 0 {
		return nil
	}

	insertRecords := make([]interface{}, len(rights))
	for i := range rights {
		insertRecords[i] = rights[i]
	}

	db := r.db.Table("tezos.nft_tokens")

	return gormbulk.BulkInsert(db, insertRecords, 2000)
}
