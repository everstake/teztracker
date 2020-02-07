package voting_periods

import (
	"github.com/everstake/teztracker/models"
	"github.com/jinzhu/gorm"
)

type (
	// Repository is the snapshots repo implementation.
	Repository struct {
		db *gorm.DB
	}

	Repo interface {
		Info(id string) (models.PeriodInfo, error)
	}
)

// New creates an instance of repository using the provided db.
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Info(id string) (periodInfo models.PeriodInfo, err error) {
	err = r.db.Select("*").Table("tezos.voting_period").
		Where("period = ?", id).Joins("left join tezos.period_stat_view on id = period").Find(&periodInfo).Error
	if err != nil {
		return periodInfo, err
	}

	return periodInfo, nil
}
