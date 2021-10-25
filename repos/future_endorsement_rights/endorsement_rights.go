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
		List(filter models.RightFilter, limit, offset uint) (rights []models.FutureEndorsementRight, err error)
		Count(models.RightFilter) (count int64, err error)
		CreateBulk(rights []models.EndorsementRight) error
	}
)

// New creates an instance of repository using the provided db.
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db.Model(&models.FutureEndorsementRight{}),
	}
}

func (r *Repository) getDb(filter models.RightFilter) *gorm.DB {
	db := r.db.Select("fer.*, baker_name as delegate_name").Table("tezos.future_endorsement_rights_view as fer").
		Joins("left join tezos.public_bakers as pb on fer.delegate = pb.delegate")
	if len(filter.BlockLevels) != 0 {
		db = db.Where("block_level IN (?)", filter.BlockLevels)
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
func (r *Repository) List(filter models.RightFilter, limit, offset uint) (rights []models.FutureEndorsementRight, err error) {
	db := r.getDb(filter)
	err = db.Order("block_level asc").
		Offset(offset).
		Limit(limit).
		Find(&rights).Error

	return rights, err
}

func (r *Repository) Last() (found bool, right models.FutureEndorsementRight, err error) {
	db := r.getDb(models.RightFilter{})
	if res := db.Order("block_level desc").Take(&right); res.Error != nil {
		if res.RecordNotFound() {
			return false, right, nil
		}
		return false, right, res.Error
	}
	return true, right, nil
}

// Find looks up for rights with filter.
func (r *Repository) Find(filter models.RightFilter) (found bool, right models.FutureBakingRight, err error) {
	if res := r.getDb(filter).Find(&right); res.Error != nil {
		if res.RecordNotFound() {
			return false, right, nil
		}
		return false, right, res.Error
	}
	return true, right, nil
}

func (r *Repository) Count(filter models.RightFilter) (count int64, err error) {
	err = r.getDb(filter).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *Repository) CreateBulk(rights []models.EndorsementRight) error {
	insertRecords := make([]interface{}, len(rights))
	for i := range rights {
		insertRecords[i] = rights[i]
	}

	db := r.db.Table("tezos.endorsing_rights")

	return gormbulk.BulkInsert(db, insertRecords, 2000)
}
