package account

import (
	"github.com/everstake/teztracker/models"
	"github.com/jinzhu/gorm"
	"time"
)

//go:generate mockgen -source ./account.go -destination ./mock_account/main.go Repo
type (
	// Repository is the account repo implementation.
	Repository struct {
		db *gorm.DB
	}

	Repo interface {
		List(limit uint, offset uint, filter models.AccountFilter) (int64, []models.AccountListView, error)
		Filter(filter models.Account, limit, offset uint) (accounts []models.AccountListView, err error)
		Count(filter models.Account) (int64, error)
		Find(filter models.Account) (found bool, acc models.AccountListView, err error)
		TotalBalance() (int64, error)
		Balances(string, time.Time, time.Time) ([]models.AccountBalance, error)
		PrevBalance(string, time.Time) (bool, models.AccountBalance, error)
		RewardsList(accountID string, limit uint, offset uint) (count int64, rewards []models.AccountReward, err error)
		RewardsCountList(accountID string, limit uint) (rewards []models.AccountRewardsCount, err error)
		CycleDelegatorsTotal(accountID string, cycleID int64) (reward models.AccountReward, err error)
		CycleDelegators(accountID string, cycle int64, limit uint, offset uint) (delegators []models.AccountDelegator, err error)
		GetReport(accountID string, params models.AccountReportFilter) (report []models.AccountReport, err error)
	}
)

const (
	accountsListView = "tezos.account_list_view"
)

// New creates an instance of repository using the provided db.
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

//Return clear table for count
func (r *Repository) getDb(filter models.AccountFilter) *gorm.DB {
	db := r.db.Model(&models.Account{})

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
func (r *Repository) List(limit, offset uint, filter models.AccountFilter) (count int64, accounts []models.AccountListView, err error) {
	db := r.getDb(filter)

	if err := db.Count(&count).Error; err != nil {
		return 0, nil, err
	}

	db = db.Model(models.AccountListView{})

	if filter.OrderBy == models.AccountOrderFieldCreatedAt {
		db = db.Order("created_at desc")
	} else if filter.OrderBy == models.AccountOrderFieldBalance {
		db = db.Order("balance desc")
	}

	db = db.Limit(limit).
		Offset(offset)

	err = db.Find(&accounts).Error
	return count, accounts, err
}

// Count counts a number of accounts sutisfying the filter.
func (r *Repository) Count(filter models.Account) (count int64, err error) {
	err = r.db.Model(&filter).
		Where(&filter).Count(&count).Error

	return count, err
}

// Filter returns a list of accounts that sutisfies the filter.
func (r *Repository) Filter(filter models.Account, limit, offset uint) (accounts []models.AccountListView, err error) {
	err = r.db.Model(models.AccountListView{}).
		Where(models.AccountListView{Account: filter}).
		Order("account_id asc").
		Limit(limit).
		Offset(offset).
		Find(&accounts).Error
	return accounts, err
}

// Find looks up for an account with filter.
func (r *Repository) Find(filter models.Account) (found bool, acc models.AccountListView, err error) {
	if res := r.db.Model(models.AccountListView{}).
		Where(models.AccountListView{Account: filter}).
		Find(&acc); res.Error != nil {
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

func (r *Repository) Balances(accountId string, from time.Time, to time.Time) (bal []models.AccountBalance, err error) {
	db := r.db.Table("tezos.accounts_history").
		Select("max(asof) as asof").
		Where("account_id = ?", accountId).
		Where("asof >= ?", from).
		Where("asof <= ?", to).
		Group("date_trunc('day', asof)")

	err = r.db.Table("tezos.accounts_history as ah").
		Select("ah.asof as time, balance").
		Joins("right join (?) as s on s.asof = ah.asof", db.QueryExpr()).
		Where("account_id = ?", accountId).
		Order("ah.asof asc").Scan(&bal).Error
	if err != nil {
		return bal, err
	}
	return bal, nil
}

func (r *Repository) PrevBalance(accountId string, from time.Time) (found bool, balance models.AccountBalance, err error) {
	if res := r.db.Table("tezos.accounts_history").
		Select("asof as time, balance").
		Where("account_id = ?", accountId).
		Where("asof < ?", from).
		Order("asof desc").
		First(&balance); res.Error != nil {
		if res.RecordNotFound() {
			return false, balance, nil
		}
		return false, balance, res.Error
	}
	return true, balance, nil

}

func (r *Repository) RewardsCountList(accountID string, limit uint) (rewards []models.AccountRewardsCount, err error) {

	err = r.db.Table("tezos.rewards_counter").
		Where("baker = ?", accountID).
		Limit(limit).
		Find(&rewards).Error

	return rewards, err
}

func (r *Repository) RewardsList(accountID string, limit uint, offset uint) (count int64, rewards []models.AccountReward, err error) {
	db := r.db.Table("tezos.baking_rewards as br").
		Where("baker = ?", accountID)

	err = db.Count(&count).Error
	if err != nil {
		return 0, rewards, err
	}

	err = db.Select("*").
		Table("tezos.rewards_counter").
		Limit(limit).
		Offset(offset).
		Find(&rewards).Error

	return count, rewards, err
}

func (r *Repository) CycleDelegatorsTotal(accountID string, cycleID int64) (reward models.AccountReward, err error) {
	err = r.db.Table("tezos.baking_rewards as br").
		Where("baker = ?", accountID).
		Where("cycle = ?", cycleID).
		Find(&reward).Error
	if err != nil {
		return reward, err
	}

	return reward, nil
}

func (r *Repository) CycleDelegators(accountID string, cycle int64, limit uint, offset uint) (delegators []models.AccountDelegator, err error) {
	err = r.db.Table("tezos.delegators_by_cycle").
		Where("delegate_value = ?", accountID).
		Where("cycle = ?", cycle).
		Order("balance desc").
		Limit(limit).
		Offset(offset).Find(&delegators).Error
	if err != nil {
		return delegators, err
	}

	return delegators, nil
}

func (r *Repository) GetReport(accountID string, params models.AccountReportFilter) (report []models.AccountReport, err error) {
	//Todo implement

	return nil, nil
}
