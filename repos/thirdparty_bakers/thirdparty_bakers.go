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
		DeleteAll() error
		Create(bakers []models.ThirdPartyBaker) error
		GetAggregatedBakers() (bakers []models.ThirdPartyBakerAgg, err error)
	}
)

// New creates an instance of repository using the provided db.
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) getDb() *gorm.DB {
	db := r.db.
		Model(&models.ThirdPartyBaker{})
	return db
}

// Get all aggregated third party bakers
func (r *Repository) GetAggregatedBakers() (bakers []models.ThirdPartyBakerAgg, err error) {
	err = r.db.Select("address,max(staking_balance) as staking_balance,json_agg(json_build_object('provider',provider,'number',number,'name',name,'address',address,'yield',yield,'staking_balance',staking_balance,'fee',fee,'available_capacity',available_capacity,'efficiency',efficiency,'payout_accuracy',payout_accuracy)) as providers").
		Model(models.ThirdPartyBakerAgg{}).
		Table("tezos.third_party_bakers").
		Order("staking_balance DESC").
		Group("address").
		Find(&bakers).
		Error
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
