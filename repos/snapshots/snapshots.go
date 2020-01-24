package snapshots

import (
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
		List(limit, offset uint) (count int64, snapshots []models.Snapshot, err error)
		RollsInBlock(block int64) (int64, error)
		Create(snapshot models.Snapshot) error
		CreateBulk(snapshots []models.Snapshot) error
	}
)

// New creates an instance of repository using the provided db.
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db.Model(&models.Snapshot{}),
	}
}

func (r *Repository) getDb() *gorm.DB {
	db := r.db.Model(&models.Snapshot{})

	return db
}

// List returns a list of snapshots from the newest to oldest.
// limit defines the limit for the maximum number of snapshots returned.
// since is used to paginate results based on the level. As the result is ordered descendingly the snapshots with level < since will be returned.
func (r *Repository) List(limit, offset uint) (count int64, snapshots []models.Snapshot, err error) {
	db := r.getDb()
	if err := db.Count(&count).Error; err != nil {
		return 0, nil, err
	}
	err = db.Order("snp_cycle desc").
		Limit(limit).
		Offset(offset).
		Find(&snapshots).Error
	return count, snapshots, err
}

// RollsInBlock returns the total number of rolls in a block.
func (r *Repository) RollsInBlock(block int64) (int64, error) {
	result := struct {
		S int64
	}{}

	err := r.db.Model(&models.Roll{}).
		Where("block_level = ?", block).
		Select("sum(rolls) as s").Scan(&result).Error
	return result.S, err
}

// Create creates a Snapshot.
func (r *Repository) Create(snapshot models.Snapshot) error {
	return r.db.Model(&snapshot).Create(&snapshot).Error
}

func (r *Repository) CreateBulk(snapshots []models.Snapshot) error {
	insertRecords := make([]interface{}, len(snapshots))
	for i := range snapshots {
		insertRecords[i] = snapshots[i]
	}
	return gormbulk.BulkInsert(r.db, insertRecords, 2000)
}
