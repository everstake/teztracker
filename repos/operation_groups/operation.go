package operation_groups

import (
	"github.com/everstake/teztracker/models"
	"github.com/jinzhu/gorm"
)

//go:generate mockgen -source ./operation.go -destination ./mock_operation_group/main.go Repo
type (
	// Repository is the operation groups repo implementation.
	Repository struct {
		db *gorm.DB
	}

	Repo interface {
		GetGroupsFor(block models.Block) (og []*models.OperationGroup, err error)
	}
)

// New creates an instance of repository using the provided db.
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Last returns the last added block.
func (r *Repository) GetGroupsFor(block models.Block) (og []*models.OperationGroup, err error) {
	filter := models.OperationGroup{BlockID: block.Hash}
	err = r.db.Model(&filter).Where(&filter).Find(&og).Error
	return og, err
}
