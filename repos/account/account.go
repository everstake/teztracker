package account

import (
	"github.com/everstake/teztracker/models"
	"github.com/jinzhu/gorm"
)

//go:generate mockgen -source ./account.go -destination ./mock_account/main.go Repo
type (
	// Repository is the account repo implementation.
	Repository struct {
		db *gorm.DB
	}

	Repo interface {
		List(limit uint, offset uint, filter models.AccountFilter) (int64, []models.Account, error)
		Filter(filter models.Account, limit, offset uint) (accounts []models.Account, err error)
		Count(filter models.Account) (int64, error)
		Find(filter models.Account) (found bool, acc models.Account, err error)
		TotalBalance() (int64, error)
	}
)

// New creates an instance of repository using the provided db.
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) getDb(filter models.AccountFilter) *gorm.DB {
	db := r.db.Select("accounts.*, asof").Model(&models.Account{}).
		Joins("natural join tezos.account_create_time_view")
	if filter.After != "" {
		db = db.Where("account_id > ?", filter.After)
	}

	if filter.Type == models.AccountTypeAccount {
		db = db.Where("account_id like 'tz%'")
	} else if filter.Type == models.AccountTypeContract {
		db = db.Where("account_id like 'KT1%'")
	}
	return db
}

// List returns a list of accounts from the newest to oldest.
// limit defines the limit for the maximum number of accounts returned.
// before is used to paginate results based on the level. As the result is ordered descendingly the accounts with level < before will be returned.
func (r *Repository) List(limit, offset uint, filter models.AccountFilter) (count int64, accounts []models.Account, err error) {
	db := r.getDb(filter)
	if err := db.Count(&count).Error; err != nil {
		return 0, nil, err
	}
	err = db.Order("account_id asc").
		Limit(limit).
		Offset(offset).
		Find(&accounts).Error
	return count, accounts, err
}

// Count counts a number of accounts sutisfying the filter.
func (r *Repository) Count(filter models.Account) (count int64, err error) {
	err = r.db.Model(&filter).
		Where(&filter).Count(&count).Error
	return count, err
}

// Filter returns a list of accounts that sutisfies the filter.
func (r *Repository) Filter(filter models.Account, limit, offset uint) (accounts []models.Account, err error) {
	err = r.db.Select("accounts.*, asof").Model(&filter).
		Where(&filter).
		Joins("natural join tezos.account_create_time_view").
		Order("account_id asc").
		Limit(limit).
		Offset(offset).
		Find(&accounts).Error
	return accounts, err
}

// Find looks up for an account with filter.
func (r *Repository) Find(filter models.Account) (found bool, acc models.Account, err error) {
	if res := r.db.Select("accounts.*, asof").Model(&filter).Joins("natural join tezos.account_create_time_view").Where(&filter).Find(&acc); res.Error != nil {
		if res.RecordNotFound() {
			return false, acc, nil
		}
		return false, acc, res.Error
	}
	return true, acc, nil
}

// TotalBalance gets the total available balance of all accounts.
func (r *Repository) TotalBalance() (b int64, err error) {
	bal := struct {
		Balance int64 `json:"balance"`
	}{}
	err = r.db.Table("tezos.accounts").Select("SUM(balance) balance").First(&bal).Error
	if err != nil {
		return 0, err
	}
	return bal.Balance, nil
}
