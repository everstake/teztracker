package double_baking

import (
	"github.com/bullblock-io/tezTracker/models"
	"github.com/jinzhu/gorm"
)

type (
	// Repository is the baking evidences repo implementation.
	Repository struct {
		db *gorm.DB
	}

	Repo interface {
		List(options models.DoubleBakingEvidenceQueryOptions) (count int64, evidences []models.DoubleBakingEvidence, err error)
		Last() (found bool, evidence models.DoubleBakingEvidence, err error)
		Create(evidence models.DoubleBakingEvidence) error
	}
)

// New creates an instance of repository using the provided db.
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db.Model(&models.DoubleBakingEvidence{}),
	}
}

func (r *Repository) getDb(options models.DoubleBakingEvidenceQueryOptions) *gorm.DB {
	db := r.db.Model(&models.DoubleBakingEvidence{})
	if options.LoadOperation {
		db = db.Preload("Operation")
	}

	if len(options.BlockIDs) != 0 {
		db = db.Where("dbe_block_hash IN (?)", options.BlockIDs)
	}
	if len(options.OperationHashes) != 0 {
		db = db.Joins("natural join operations")
		db = db.Where("operations.operation_group_hash in (?)", options.OperationHashes)
	}
	return db
}

// List returns a list of evidences from the newest to oldest.
func (r *Repository) List(options models.DoubleBakingEvidenceQueryOptions) (count int64, evidences []models.DoubleBakingEvidence, err error) {
	db := r.getDb(options)
	if err := db.Count(&count).Error; err != nil {
		return 0, nil, err
	}

	if options.Limit > 0 {
		db = db.Limit(options.Limit)
	}
	if options.Offset > 0 {
		db = db.Offset(options.Offset)
	}

	err = db.Order("operation_id desc").
		Find(&evidences).Error
	return count, evidences, err
}

func (r *Repository) Last() (found bool, evidence models.DoubleBakingEvidence, err error) {
	db := r.db.Model(&evidence)
	if res := db.Order("operation_id desc").Take(&evidence); res.Error != nil {
		if res.RecordNotFound() {
			return false, evidence, nil
		}
		return false, evidence, res.Error
	}
	return true, evidence, nil
}

// Create creates a DoubleBakingEvidence.
func (r *Repository) Create(evidence models.DoubleBakingEvidence) error {
	return r.db.Model(&evidence).Create(&evidence).Error
}
