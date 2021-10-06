package operation_counter

import (
	"github.com/everstake/teztracker/models"
	"github.com/jinzhu/gorm"
)

type (
	// Repository is the operation counter repo implementation.
	Repository struct {
		db *gorm.DB
	}

	Repo interface {
		Create(cntr models.OperationCounter) (id int64, err error)
		OperationsCount(kinds []string) (count int64, err error)
	}
)

// New creates an instance of repository using the provided db.
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Last returns the last added block.
func (r *Repository) Create(cntr models.OperationCounter) (id int64, err error) {
	err = r.db.Create(&cntr).Error
	return cntr.ID, err
}

func (r *Repository) OperationsCount(kinds []string) (count int64, err error) {

	s := struct {
		S int64
	}{}

	err = r.db.Select("sum(cnt_count) s").
		Table("tezos.last_operation_counters").
		Where("cnt_operation_type IN (?)", kinds).
		Find(&s).Error
	if err != nil {
		return count, err
	}

	return s.S, nil
}
