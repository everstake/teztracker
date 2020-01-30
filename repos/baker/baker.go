package baker

import (
	"fmt"
	"github.com/everstake/teztracker/models"
	"github.com/jinzhu/gorm"
)

//go:generate mockgen -source ./baker.go -destination ./mock_baker/main.go Repo
type (
	// Repository is the baker repo implementation.
	Repository struct {
		db *gorm.DB
	}

	Repo interface {
		Find(accountID string) (bool, models.Delegate, error)
		List(limit, offset uint) ([]models.Baker, error)
		Count() (int64, error)
		BlocksCountBakedBy(ids []string, startingLevel int64) (counter []BakerCounter, err error)
		EndorsementsCountBy(ids []string, startingLevel int64) (counter []BakerWeightedCounter, err error)
		TotalStakingBalance() (int64, error)
		RefreshView() error
	}

	BakerCounter struct {
		Baker string
		Count int64
	}
	BakerWeightedCounter struct {
		BakerCounter
		Weight float64
	}
)

const (
	endorsementKind       = "endorsement"
	bakerMaterializedView = "tezos.baker_view"
	firstBlock            = 0
)

// New creates an instance of repository using the provided db.
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Find(accountID string) (found bool, delegate models.Delegate, err error) {
	filter := models.Delegate{Pkh: accountID}
	if res := r.db.Model(&filter).Where(&filter).Find(&delegate); res.Error != nil {
		if res.RecordNotFound() {
			return false, delegate, nil
		}
		return false, delegate, res.Error
	}
	return true, delegate, nil
}

// List returns a list of bakers(accounts which have at least 1 endorsement operation) ordered by their staking balance.
// limit defines the limit for the maximum number of bakers returned,
// offset sets the offset for thenumber of rows returned.
func (r *Repository) List(limit, offset uint) (bakers []models.Baker, err error) {
	err = r.db.Select("tezos.baker_view.*, name").Table(bakerMaterializedView).
		Joins("left join tezos.baker_alias on baker_view.account_id = baker_alias.address").
		Order("staking_balance desc").
		Limit(limit).
		Offset(offset).
		Find(&bakers).Debug().Error
	if err != nil {
		return nil, err
	}

	return bakers, err
}

// BlocksCountBakedBy returns a slice of block counters with the number of blocks baked by each baker among ids.
func (r *Repository) BlocksCountBakedBy(ids []string, startingLevel int64) (counter []BakerCounter, err error) {
	db := r.db.Model(&models.Block{}).
		Where("baker IN (?)", ids)
	if startingLevel > 0 {
		db = db.Where("level >= ?", startingLevel)
	}
	err = db.Select("baker, count(1) count").
		Group("baker").Scan(&counter).Error
	if err != nil {
		return nil, err
	}

	return counter, nil
}

// BlocksCountBakedBy returns a slice of block counters with the number of endorsements made by each baker among ids.
func (r *Repository) EndorsementsCountBy(ids []string, startingLevel int64) (counter []BakerWeightedCounter, err error) {
	db := r.db.Table("tezos.endorsements_view").
		Where("baker IN (?)", ids)
	if startingLevel > 0 {
		db = db.Where("block_level >= ?", startingLevel)
	}

	err = db.Select("SUM(count) as count, SUM(count*trunc(1/priority,6)) as weight, baker").
		Group("baker").Scan(&counter).Error
	if err != nil {
		return nil, err
	}

	return counter, nil
}

// TotalStakingBalance gets the total staked balance of all delegates.
func (r *Repository) TotalStakingBalance() (b int64, err error) {
	bal := struct {
		Balance int64
	}{}
	err = r.db.Table("tezos.delegates").Select("SUM(staking_balance) balance").Find(&bal).Error
	if err != nil {
		return 0, err
	}
	return bal.Balance, nil
}

// Count counts a number of bakers sutisfying the filter.
func (r *Repository) Count() (count int64, err error) {
	err = r.db.Table(bakerMaterializedView).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// RefreshView execute baker materialized view refresh
func (r *Repository) RefreshView() (err error) {

	err = r.db.Exec(fmt.Sprint("REFRESH MATERIALIZED VIEW CONCURRENTLY ", bakerMaterializedView)).Error
	if err != nil {
		return err
	}
	return nil
}
