package double_baking

import (
	"github.com/everstake/teztracker/models"
	"github.com/jinzhu/gorm"
)

type (
	// Repository is the baking evidences repo implementation.
	Repository struct {
		db *gorm.DB
	}

	Repo interface {
		List(options models.DoubleOperationEvidenceQueryOptions) (count int64, evidences []models.DoubleOperationEvidenceExtended, err error)
		Last() (found bool, evidence models.DoubleOperationEvidenceExtended, err error)
		Create(evidence models.DoubleOperationEvidence) error
	}
)

// New creates an instance of repository using the provided db.
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db.Model(&models.DoubleOperationEvidence{}),
	}
}

func (r *Repository) getDb(options models.DoubleOperationEvidenceQueryOptions) *gorm.DB {
	db := r.db.Select("*").Table("tezos.double_operation_evidences").
		Where("doe_type = ?", options.Type).
		Joins("left join tezos.public_bakers as off on doe_offender = off.delegate").
		Joins("left join tezos.public_bakers as evd on doe_evidence_baker = evd.delegate")
	if options.LoadOperation {
		db = db.Preload("Operation")
	}

	if len(options.BlockIDs) != 0 {
		db = db.Where("doe_block_hash IN (?)", options.BlockIDs)
	}
	if len(options.OperationHashes) != 0 {
		db = db.Joins("natural join tezos.operations")
		db = db.Where("operations.operation_group_hash in (?)", options.OperationHashes)
	}

	if len(options.OperationIDs) != 0 {
		db = db.Where("double_operation_evidences.operation_id in (?)", options.OperationIDs)
	}

	return db
}

// List returns a list of evidences from the newest to oldest.
func (r *Repository) List(options models.DoubleOperationEvidenceQueryOptions) (count int64, evidences []models.DoubleOperationEvidenceExtended, err error) {
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

func (r *Repository) Last() (found bool, evidence models.DoubleOperationEvidenceExtended, err error) {
	db := r.getDb(models.DoubleOperationEvidenceQueryOptions{Type: models.DoubleOperationTypeBaking})

	if res := db.Order("operation_id desc").Take(&evidence); res.Error != nil {
		if res.RecordNotFound() {
			return false, evidence, nil
		}
		return false, evidence, res.Error
	}
	return true, evidence, nil
}

// Create creates a DoubleOperationEvidence.
func (r *Repository) Create(evidence models.DoubleOperationEvidence) error {
	return r.db.Model(&evidence).Create(&evidence).Error
}
