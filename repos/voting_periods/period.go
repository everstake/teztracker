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
		GetCurrentPeriodId() (int64, error)
		Info(id int64) (models.PeriodInfo, error)
		Ballots(id int64) ([]models.PeriodBallot, error)
		StatsByKind(periodType string) ([]models.PeriodInfo, error)
	}
)

// New creates an instance of repository using the provided db.
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetCurrentPeriodId() (id int64, err error) {
	period := struct {
		ID int64
	}{}

	err = r.db.Select("max(id) as id").Table("tezos.voting_period").
		Find(&period).Error
	if err != nil {
		return 0, err
	}

	return period.ID, nil
}

func (r *Repository) Info(id int64) (periodInfo models.PeriodInfo, err error) {
	err = r.db.Select("vp.*, psw.*, ptsv.total_rolls, ptsv.total_bakers").Table("tezos.voting_period as vp").
		Joins("left join tezos.period_stat_view as psw on id = psw.period").
		Joins("left join tezos.period_total_stat_view as ptsv on id = ptsv.period").
		Where("id = ?", id).
		Find(&periodInfo).Error
	if err != nil {
		return periodInfo, err
	}

	return periodInfo, nil
}

func (r *Repository) StatsByKind(periodKind string) (periods []models.PeriodInfo, err error) {
	err = r.db.Select("psv.*, ptsv.total_rolls, ptsv.total_bakers").
		Table("tezos.period_stat_view as psv").
		Joins("left join tezos.period_total_stat_view as ptsv on psv.period = ptsv.period").
		Where("kind = ?", periodKind).
		Order("psv.period asc").Scan(&periods).Error
	if err != nil {
		return periods, err
	}

	return periods, nil
}

func (r *Repository) Ballots(id int64) (periodBallots []models.PeriodBallot, err error) {
	err = r.db.Select("*").Table("tezos.proposal_stat_view").
		Where("period = ?", id).Scan(&periodBallots).Error
	if err != nil {
		return periodBallots, err
	}

	return periodBallots, nil
}
