package daily_stats

import (
	"fmt"
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
		GetDailyStats(key string, aggType string, filter models.AggTimeFilter) (items []models.AggTimeInt, err error)
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

func (r *Repository) GetDailyStats(key string, aggType string, filter models.AggTimeFilter) (items []models.AggTimeInt, err error) {
	err = filter.Validate()
	if err != nil {
		return nil, fmt.Errorf("filter.Validate: %s", err.Error())
	}
	q := r.db.Select(fmt.Sprintf("CAST(%s(value) AS INTEGER) as value, date_trunc('%s', date) as date", aggType, filter.Period)).
		Table("tezos.daily_stats").Where("key = ?", key).Group("date")
	if !filter.From.IsZero() {
		q = q.Where("date >= ?", filter.From)
	}
	if !filter.To.IsZero() {
		q = q.Where("date <= ?", filter.To)
	}
	err = q.Order("date").Find(&items).Error
	return items, err

}
