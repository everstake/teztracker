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
		Find(accountID string) (bool, models.Baker, error)
		List(limit, offset uint) ([]models.Baker, error)
		Count() (int64, error)
		BlocksCountBakedBy(ids []string, startingLevel int64) (counter []BakerCounter, err error)
		EndorsementsCountBy(ids []string, startingLevel int64) (counter []BakerWeightedCounter, err error)
		TotalStakingBalance() (int64, error)
		RefreshView() error
		Balance(accountId string) (bal models.BakerBalance, err error)

		//New
		PublicBakersCount() (int64, error)
		PublicBakersList(limit, offset uint) (bakers []models.Baker, err error)
		BakerRegistryList() ([]models.BakerRegistry, error)
		SavePublicBaker(models.BakerRegistry) error
		PublicBakersSearchList() ([]models.PublicBakerSearch, error)

		TotalBakingRewards(accountId string, fromCycle, toCycle int64) (rewards int64, err error)
		TotalEndorsementRewards(accountId string, fromCycle, toCycle int64) (rewards int64, err error)
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

func (r *Repository) Find(accountID string) (found bool, baker models.Baker, err error) {
	if res := r.db.Select("tezos.baker_view.*, baker_name as name, (10000 - split)/100 as fee").Table(bakerMaterializedView).
		Joins("left join tezos.public_bakers on baker_view.account_id = public_bakers.delegate").
		Where("account_id = ?", accountID).
		Order("staking_balance desc").
		Find(&baker); res.Error != nil {
		if res.RecordNotFound() {
			return false, baker, nil
		}
		return false, baker, res.Error
	}
	return true, baker, nil
}

// List returns a list of bakers(accounts which have at least 1 endorsement operation) ordered by their staking balance.
// limit defines the limit for the maximum number of bakers returned,
// offset sets the offset for thenumber of rows returned.
func (r *Repository) List(limit, offset uint) (bakers []models.Baker, err error) {
	err = r.db.Select("tezos.baker_view.*,baker_name as name").Table(bakerMaterializedView).
		Joins("left join tezos.public_bakers on baker_view.account_id = public_bakers.delegate").
		Order("staking_balance desc").
		Limit(limit).
		Offset(offset).
		Find(&bakers).Error
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

func (r *Repository) Balance(accountId string) (bal models.BakerBalance, err error) {

	err = r.db.Table("tezos.bakers").
		Where("pkh = ?", accountId).
		Find(&bal).Error
	if err != nil {
		return bal, err
	}
	return bal, nil
}

// TotalStakingBalance gets the total staked balance of all delegates.
func (r *Repository) TotalStakingBalance() (b int64, err error) {
	bal := struct {
		Balance int64
	}{}
	err = r.db.Table("tezos.bakers").
		Select("SUM(staking_balance) balance").
		Where("deactivated is not true").
		Find(&bal).Error
	if err != nil {
		return 0, err
	}
	return bal.Balance, nil
}

func (r *Repository) TotalBakingRewards(accountId string, fromCycle, toCycle int64) (rewards int64, err error) {
	rew := struct {
		Rewards int64
	}{}

	db := r.db.Table("tezos.baker_cycle_bakings_view").
		Select("SUM(reward) rewards").
		Where("cycle >= ?", fromCycle).
		Where("cycle <= ?", toCycle)
	if accountId != "" {
		db = db.Where("delegate = ?", accountId)
	}

	err = db.Find(&rew).Error
	if err != nil {
		return 0, err
	}

	return rew.Rewards, nil
}

func (r *Repository) TotalEndorsementRewards(accountId string, fromCycle, toCycle int64) (rewards int64, err error) {
	rew := struct {
		Rewards int64
	}{}

	db := r.db.Table("tezos.baker_cycle_endorsements_view").
		Select("SUM(reward) rewards").
		Where("cycle >= ?", fromCycle).
		Where("cycle <= ?", toCycle)

	if accountId != "" {
		db = db.Where("delegate = ?", accountId)
	}

	err = db.Find(&rew).Error
	if err != nil {
		return 0, err
	}

	return rew.Rewards, nil
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

func (r *Repository) BakerRegistryList() (bakers []models.BakerRegistry, err error) {
	err = r.db.Model(&models.BakerRegistry{}).Scan(&bakers).Error
	if err != nil {
		return bakers, err
	}

	return bakers, nil
}

// Count counts a number of bakers sutisfying the filter.
func (r *Repository) PublicBakersCount() (count int64, err error) {
	err = r.db.Table("tezos.public_bakers").Where("is_hidden IS false").Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *Repository) PublicBakersList(limit, offset uint) (bakers []models.Baker, err error) {
	err = r.db.Select("pb.baker_name as name,delegate as account_id, bw.*, (10000 - split)/100 as fee ").Table("tezos.public_bakers as pb").
		Joins(fmt.Sprintf("left join %s as bw on bw.account_id = pb.delegate", bakerMaterializedView)).
		Where("is_hidden IS false").
		Order("COALESCE(staking_balance,0) desc").
		Limit(limit).
		Offset(offset).
		Find(&bakers).Error
	if err != nil {
		return nil, err
	}
	return bakers, nil
}

func (r *Repository) PublicBakersSearchList() (list []models.PublicBakerSearch, err error) {
	err = r.db.Table("tezos.public_bakers").Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}
func (r *Repository) SavePublicBaker(baker models.BakerRegistry) (err error) {
	if r.db.First(&models.BakerRegistry{}, "delegate = ?", baker.Delegate).RecordNotFound() {
		err = r.db.Create(&baker).Error
		return err
	}

	err = r.db.Save(baker).Error
	if err != nil {
		return err
	}

	return nil
}
