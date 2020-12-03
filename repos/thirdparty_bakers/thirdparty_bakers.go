package thirdparty_bakers

import (
	"github.com/everstake/teztracker/models"
	"github.com/jinzhu/gorm"
	gormbulk "github.com/t-tiger/gorm-bulk-insert"
)

type (
	// Repository is the third party bakers repo implementation.
	Repository struct {
		db *gorm.DB
	}

	Repo interface {
		GetAll() (bakers []models.ThirdPartyBaker, err error)
		DeleteAll() error
		Create(bakers []models.ThirdPartyBaker) error
	}
)

// New creates an instance of repository using the provided db.
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) getDb() *gorm.DB {
	db := r.db.Select("*").
		Model(&models.ThirdPartyBaker{})
	return db
}

// Get all third party bakers
func (r *Repository) GetAll() (bakers []models.ThirdPartyBaker, err error) {
	err = r.getDb().Find(&bakers).Error
	return bakers, err
}

// Delete all third party bakers
func (r *Repository) DeleteAll() error {
	return r.getDb().Delete(&models.ThirdPartyBaker{}).Error
}

// Create third party bakers
func (r *Repository) Create(bakers []models.ThirdPartyBaker) error {
	insertRecords := make([]interface{}, len(bakers))
	for i := range bakers {
		insertRecords[i] = bakers[i]
	}
	return gormbulk.BulkInsert(r.db, insertRecords, 2000)
}
