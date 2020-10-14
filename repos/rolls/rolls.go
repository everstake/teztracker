package rolls

import (
	"fmt"
	"github.com/everstake/teztracker/models"
	"github.com/jinzhu/gorm"
	gormbulk "github.com/t-tiger/gorm-bulk-insert"
)

type (
	// Repository is the snapshots repo implementation.
	Repository struct {
		db *gorm.DB
	}

	Repo interface {
		RollsAndBakersInBlock(block int64) (int64, int64, error)
		CreateBulk(rolls []models.Roll) error
	}
)

// New creates an instance of repository using the provided db.
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db.Model(&models.Roll{}),
	}
}

// RollsInBlock returns the total number of rolls in a block.
func (r *Repository) RollsAndBakersInBlock(block int64) (int64, int64, error) {

	return 0, 0, fmt.Errorf("Not implemented now")

	result := struct {
		S int64
		C int64
	}{}

	err := r.db.Model(&models.Roll{}).
		Where("block_level = ?", block).
		Select("sum(rolls) as s, count(1) as c").Scan(&result).Error
	return result.S, result.C, err
}

func (r *Repository) CreateBulk(rolls []models.Roll) error {
	insertRecords := make([]interface{}, len(rolls))
	for i := range rolls {
		insertRecords[i] = rolls[i]
	}

	return gormbulk.BulkInsert(r.db, insertRecords, 2000)
}
