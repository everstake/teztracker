package account

import (
	"fmt"
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
		List(limit uint, offset uint, filter models.AccountFilter) (int64, []models.Account, error)
		Filter(filter models.Account, limit, offset uint) (accounts []models.Account, err error)
		Count(filter models.Account) (int64, error)
		Find(filter models.Account) (found bool, acc models.Account, err error)
		TotalBalance() (int64, error)
		Balances(string, time.Time, time.Time) ([]models.AccountBalance, error)
		PrevBalance(string, time.Time) (bool, models.AccountBalance, error)
		RewardsList(accountID string, limit uint, offset uint) (count int64, rewards []models.AccountReward, err error)
		RefreshView() error
	}
)

const (
	accountMaterializedView = "tezos.account_materialized_view"
)

// New creates an instance of repository using the provided db.
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

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
func (r *Repository) List(limit, offset uint, filter models.AccountFilter) (count int64, accounts []models.Account, err error) {
	db := r.getDb(filter)

	if err := db.Count(&count).Error; err != nil {
		return 0, nil, err
	}

	db = r.db.Select("accounts.*, created_at, last_active, account_name").
		Table("tezos.account_materialized_view as amv").
		Joins("inner join tezos.accounts on accounts.account_id = amv.account_id")

	if filter.Type == models.AccountTypeAccount {
		db = db.Where("amv.account_id like 'tz%'")
	} else if filter.Type == models.AccountTypeContract {
		db = db.Where("amv.account_id like 'KT1%'")
	}

	db = db.Order("created_at desc").
		Limit(limit).
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
func (r *Repository) Filter(filter models.Account, limit, offset uint) (accounts []models.Account, err error) {
	err = r.db.Select("accounts.*, created_at, last_active, account_name").Model(&filter).
		Where(&filter).
		Joins("natural join tezos.account_materialized_view").
		Order("account_id asc").
		Limit(limit).
		Offset(offset).
		Find(&accounts).Error
	return accounts, err
}

// Find looks up for an account with filter.
func (r *Repository) Find(filter models.Account) (found bool, acc models.Account, err error) {
	if res := r.db.Select("accounts.*, created_at, last_active").
		Model(&filter).
		Joins("natural join tezos.account_materialized_view").
		Where(&filter).Find(&acc); res.Error != nil {
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
		Where("account_id = ?", accountId).Scan(&bal).Error
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
		First(&balance); res.Error != nil {
		if res.RecordNotFound() {
			return false, balance, nil
		}
		return false, balance, res.Error
	}
	return true, balance, nil

}

func (r *Repository) RefreshView() (err error) {
	err = r.db.Exec(fmt.Sprint("REFRESH MATERIALIZED VIEW ", accountMaterializedView)).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) RewardsList(accountID string, limit uint, offset uint) (count int64, rewards []models.AccountReward, err error) {
	db := r.db.Table("tezos.baking_rewards as br").
		Where("baker = ?", accountID)

	err = db.Count(&count).Error
	if err != nil {
		return 0, rewards, err
	}

	err = db.Select("br.*, cbv.reward as baking_rewards, cbv.missed missed_baking, fbrv.count as future_baking_count, cev.reward endorsement_rewards ,cev.missed missed_endorsements, fev.count future_endorsement_count").
		Joins("left join tezos.baker_future_baking_rights_view fbrv on br.baker = fbrv.delegate and br.cycle = fbrv.cycle").
		Joins("left join tezos.baker_cycle_bakings_view cbv on br.baker = cbv.delegate and br.cycle = cbv.cycle").
		Joins("left join tezos.baker_cycle_endorsements_view cev on br.baker = cev.delegate and br.cycle = cev.cycle").
		Joins("left join tezos.baker_future_endorsement_view fev on br.baker = fev.delegate and br.cycle = fev.cycle").
		Order("cycle desc").Limit(limit).
		Offset(offset).
		Find(&rewards).Error

	return count, rewards, err
}
