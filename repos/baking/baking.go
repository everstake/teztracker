package baking

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
		BakingTotal(string) (models.AccountBaking, error)
		BakingList(accountID string, limit uint, offset uint) (int64, []models.AccountBaking, error)
		FutureBakingList(accountID string) ([]models.AccountBaking, error)
	}
)

// New creates an instance of repository using the provided db.
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) BakingTotal(accountID string) (total models.AccountBaking, err error) {
	db := r.db.Select("avg(avg_priority) avg_priority, sum(reward) reward, sum(count) count, sum(missed) missed, sum(stolen) stolen").
		Table("tezos.baker_cycle_bakings_view").
		Model(&models.AccountBaking{}).
		Where("delegate = ?", accountID)

	err = db.Find(&total).Error

	return total, err
}

func (r *Repository) BakingList(accountID string, limit uint, offset uint) (count int64, baking []models.AccountBaking, err error) {
	db := r.db.Table("tezos.baker_cycle_bakings_view").
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

func (r *Repository) FutureBakingList(accountID string) (baking []models.AccountBaking, err error) {
	db := r.db.Table("tezos.baker_future_baking_rights_view").
		Model(&models.AccountBaking{}).
		Where("delegate = ?", accountID)

	err = db.Order("cycle desc").Find(&baking).Error

	return baking, err
}
