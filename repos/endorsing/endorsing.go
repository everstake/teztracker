package endorsing

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
		FutureEndorsingList(accountID string) ([]models.AccountEndorsing, error)
		EndorsingTotal(string) (models.AccountEndorsing, error)
		EndorsingList(accountID string, limit uint, offset uint) (int64, []models.AccountEndorsing, error)
	}
)

// New creates an instance of repository using the provided db.
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
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

func (r *Repository) FutureEndorsingList(accountID string) (endorsing []models.AccountEndorsing, err error) {
	db := r.db.Table("tezos.baker_future_endorsement_view").
		Model(&models.AccountEndorsing{}).
		Where("delegate = ?", accountID)

	err = db.Order("cycle desc").Find(&endorsing).Error

	return endorsing, err
}
