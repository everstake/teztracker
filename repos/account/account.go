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
		BakingTotal(string) (models.AccountBaking, error)
		BakingList(accountID string, limit uint, offset uint) (int64, []models.AccountBaking, error)
		FutureBakingList(accountID string) ([]models.AccountBaking, error)
		EndorsingTotal(string) (models.AccountEndorsing, error)
		EndorsingList(accountID string, limit uint, offset uint) (int64, []models.AccountEndorsing, error)
		RewardsList(accountID string, limit uint, offset uint) (count int64, rewards []models.AccountReward, err error)
		PrevBalance(string, time.Time) (bool, models.AccountBalance, error)
		RefreshView() error
		RefreshAccountFutureBakingView() error
		RefreshAccountBakingView() error
	}
)

const (
	accountMaterializedView = "tezos.account_materialized_view"
	futureBakingView        = "tezos.future_baking_rights_materialized_view"
	bakingView              = "tezos.baking_materialized_view"
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

func (r *Repository) BakingTotal(accountID string) (total models.AccountBaking, err error) {
	db := r.db.Select("avg(avg_priority) avg_priority, sum(reward) reward, sum(count) count, sum(missed) missed, sum(stolen) stolen").
		Table("tezos.baking_materialized_view").
		Model(&models.AccountBaking{}).
		Where("delegate = ?", accountID)

	err = db.Find(&total).Error

	return total, err
}

func (r *Repository) BakingList(accountID string, limit uint, offset uint) (count int64, baking []models.AccountBaking, err error) {
	db := r.db.Table("tezos.baking_materialized_view").
		Model(&models.AccountBaking{}).
		Where("delegate = ?", accountID)

	err = db.Count(&count).Error
	if err != nil {
		return 0, baking, err
	}

	err = db.Order("cycle desc").Limit(limit).
		Offset(offset).
		Find(&baking).Error

	return count, baking, err
}

func (r *Repository) EndorsingTotal(accountID string) (total models.AccountEndorsing, err error) {
	db := r.db.Select("sum(reward) reward, count(1) count, sum(missed) missed").
		Table("tezos.baker_endorsements").
		Model(&models.AccountEndorsing{}).
		Where("delegate = ?", accountID)

	err = db.Find(&total).Error

	return total, err
}

func (r *Repository) EndorsingList(accountID string, limit uint, offset uint) (count int64, endorsing []models.AccountEndorsing, err error) {
	db := r.db.Table("tezos.baker_cycle_endorsements_view").
		Model(&models.AccountEndorsing{}).
		Where("delegate = ?", accountID)

	err = db.Count(&count).Error
	if err != nil {
		return 0, endorsing, err
	}

	err = db.Order("cycle desc").Limit(limit).
		Offset(offset).
		Find(&endorsing).Error

	return count, endorsing, err
}

func (r *Repository) RewardsList(accountID string, limit uint, offset uint) (count int64, rewards []models.AccountReward, err error) {
	db := r.db.Select("br.*, bmv.reward as baking_rewards, bmv.missed missed_baking, fbrmv.count as future_baking_count, emv.reward endorsement_rewards ,emv.missed missed_endorsements").
		Table("tezos.baking_rewards as br").
		Joins("left join tezos.future_baking_rights_materialized_view fbrmv on br.baker = fbrmv.delegate and br.cycle = fbrmv.cycle").
		Joins(" left join tezos.baking_materialized_view bmv on br.baker = bmv.delegate and br.cycle = bmv.cycle").
		Joins("left join tezos.endorsements_materialized_view emv on br.baker = emv.delegate and br.cycle = emv.cycle").
		Model(&models.AccountReward{}).
		Where("baker = ?", accountID)

	err = db.Count(&count).Error
	if err != nil {
		return 0, rewards, err
	}

	err = db.Order("cycle desc").Limit(limit).
		Offset(offset).
		Find(&rewards).Error

	return count, rewards, err
}

func (r *Repository) FutureBakingList(accountID string) (baking []models.AccountBaking, err error) {
	db := r.db.Table("tezos.future_baking_rights_materialized_view").
		Model(&models.AccountBaking{}).
		Where("delegate = ?", accountID)

	err = db.Order("cycle desc").Find(&baking).Error

	return baking, err
}

func (r *Repository) RefreshAccountFutureBakingView() (err error) {
	err = r.db.Exec(fmt.Sprint("REFRESH MATERIALIZED VIEW ", futureBakingView)).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) RefreshAccountBakingView() (err error) {
	err = r.db.Exec(fmt.Sprint("REFRESH MATERIALIZED VIEW ", bakingView)).Error
	if err != nil {
		return err
	}
	return nil
}
