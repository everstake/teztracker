package endorsing

import (
	"fmt"
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
		GetLostEndorsingCountAgg(filter models.AggTimeFilter) (items []models.AggTimeInt, err error)
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
	db := r.db.Table("tezos.baker_cycle_endorsements_view bce").
		Model(&models.AccountEndorsing{}).
		Where("delegate = ?", accountID)

	err = db.Count(&count).Error
	if err != nil {
		return 0, endorsing, err
	}

	err = db.Select("*").
		Joins("LEFT JOIN tezos.cycle_periods_view cp on bce.cycle = cp.cycle").
		Order("bce.cycle desc").Limit(limit).
		Offset(offset).
		Find(&endorsing).Error

	return count, endorsing, err
}

func (r *Repository) FutureEndorsingList(accountID string) (endorsing []models.AccountEndorsing, err error) {
	err = r.db.Select("*").
		Table("tezos.baker_future_endorsement_view bfe").
		Joins("LEFT JOIN tezos.cycle_periods_view cp on bfe.cycle = cp.cycle").
		Model(&models.AccountEndorsing{}).
		Where("delegate = ?", accountID).
		Order("bfe.cycle desc").Find(&endorsing).Error
	if err != nil {
		return nil, err
	}

	return endorsing, nil
}

func (r *Repository) GetLostEndorsingCountAgg(filter models.AggTimeFilter) (items []models.AggTimeInt, err error) {
	err = filter.Validate()
	if err != nil {
		return nil, fmt.Errorf("filter.Validate: %s", err.Error())
	}
	q := r.db.Select(fmt.Sprintf("count(*) as value, date_trunc('%s', blocks.timestamp) as date", filter.Period)).
		Table("tezos.baker_endorsements").
		Joins("left join tezos.blocks ON baker_endorsements.level = blocks.level").
		Where("baker_endorsements.missed != 0").
		Group("date")
	if !filter.From.IsZero() {
		q = q.Where("blocks.timestamp >= ?", filter.From)
	}
	if !filter.To.IsZero() {
		q = q.Where("blocks.timestamp <= ?", filter.To)
	}
	err = q.Order("date").Find(&items).Error
	return items, err
}
