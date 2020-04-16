package future_endorsement_rights

import (
	"github.com/everstake/teztracker/models"
	"github.com/jinzhu/gorm"
	gormbulk "github.com/t-tiger/gorm-bulk-insert"
)

type (
	// Repository is the baking rights repo implementation.
	Repository struct {
		db *gorm.DB
	}

	Repo interface {
		Last() (found bool, right models.FutureEndorsementRight, err error)
		CreateBulk(rights []models.FutureEndorsementRight) error
	}
)

// New creates an instance of repository using the provided db.
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db.Model(&models.FutureEndorsementRight{}),
	}
}

func (r *Repository) getDb(filter models.BakingRightFilter) *gorm.DB {
	db := r.db.Select("fer.*, baker_name as delegate_name").Table("tezos.future_endorsement_rights as fer").
		Joins("left join tezos.public_bakers as pb on fer.delegate = pb.delegate")
	if len(filter.BlockLevels) != 0 {
		db = db.Where("level IN (?)", filter.BlockLevels)
	}
	if len(filter.Delegates) != 0 {
		db = db.Where("fer.delegate IN (?)", filter.Delegates)
	}
	if filter.PriorityFrom != 0 {
		db = db.Where("priority >= ?", filter.PriorityFrom)
	}
	if filter.PriorityTo != 0 {
		db = db.Where("priority <= ?", filter.PriorityTo)
	}
	return db
}

// List returns a list of rights from the oldest to the newest.
// limit defines the limit for the maximum number of rights returned.
// since is used to paginate results based on the level. As the result is ordered descendingly the rights with level < since will be returned.
func (r *Repository) List(filter models.BakingRightFilter) (rights []models.FutureBakingRight, err error) {
	db := r.getDb(filter)
	err = db.Order("level asc, priority asc").
		Find(&rights).Error
	return rights, err
}

// ListDesc returns a list of rights from the newest to the oldest.
// limit defines the limit for the maximum number of rights returned.
// since is used to paginate results based on the level. As the result is ordered descendingly the rights with level < since will be returned.
func (r *Repository) ListDesc(filter models.BakingRightFilter) (rights []models.FutureBakingRight, err error) {
	db := r.getDb(filter)
	err = db.Order("level desc, priority asc").
		Find(&rights).Error
	return rights, err
}

func (r *Repository) Last() (found bool, right models.FutureEndorsementRight, err error) {
	db := r.getDb(models.BakingRightFilter{})
	if res := db.Order("level desc").Take(&right); res.Error != nil {
		if res.RecordNotFound() {
			return false, right, nil
		}
		return false, right, res.Error
	}
	return true, right, nil
}

// Find looks up for rights with filter.
func (r *Repository) Find(filter models.BakingRightFilter) (found bool, right models.FutureBakingRight, err error) {
	if res := r.getDb(filter).Find(&right); res.Error != nil {
		if res.RecordNotFound() {
			return false, right, nil
		}
		return false, right, res.Error
	}
	return true, right, nil
}

// Create creates a FutureBakingRight.
func (r *Repository) Create(right models.FutureBakingRight) error {
	return r.db.Model(&right).Create(&right).Error
}

func (r *Repository) CreateBulk(rights []models.FutureEndorsementRight) error {
	insertRecords := make([]interface{}, len(rights))
	for i := range rights {
		insertRecords[i] = rights[i]
	}
	return gormbulk.BulkInsert(r.db, insertRecords, 2000)
}
