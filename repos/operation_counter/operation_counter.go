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
