package daily_stats

import (
	"github.com/everstake/teztracker/models"
	"github.com/jinzhu/gorm"
)

const DailyStatsTable = "tezos.daily_stats"

type (
	// Repository is the storage repo implementation.
	Repository struct {
		db *gorm.DB
	}

	Repo interface {
		Create(stat models.DailyStat) error
	}
)

// New creates an instance of repository using the provided db.
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) getDb() *gorm.DB {
	db := r.db.Table(DailyStatsTable)
	return db
}

func (r *Repository) Create(stat models.DailyStat) error {
	return r.getDb().Create(stat).Error
}

