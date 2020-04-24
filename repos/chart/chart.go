package chart

import (
	"fmt"
	"github.com/everstake/teztracker/models"
	"github.com/jinzhu/gorm"
)

type (
	// Repository is the block repo implementation.
	Repository struct {
		db *gorm.DB
	}

	Repo interface {
		BlocksNumber(from, to int64, period string) ([]models.ChartData, error)
		TransactionsVolume(from, to int64, period string) (data []models.ChartData, err error)
		OperationsNumber(from, to int64, period string) (data []models.ChartData, err error)
		FeesVolume(from, to int64, period string) (data []models.ChartData, err error)
		ActivationsNumber(from, to int64, period string) (data []models.ChartData, err error)
		AvgBlockDelay(from, to int64, period string) (data []models.ChartData, err error)
		DelegationVolume(from, to int64, period string) (data []models.ChartData, err error)
		Bakers(from, to int64, period string) (data []models.ChartData, err error)
		InsertWhaleAccounts(dayTime int64) error
		WhaleAccounts(from, to int64, period string) (data []models.ChartData, err error)
	}
)

// New creates an instance of repository using the provided db.
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) getDB(from, to int64, period string) *gorm.DB {
	db := r.db.
		Where("timestamp >= to_timestamp(?)", from).
		Where("timestamp <= to_timestamp(?)", to).
		Order("timestamp asc").
		Group(fmt.Sprintf("date_trunc('%s', timestamp)", period))

	return db
}

func (r *Repository) BlocksNumber(from, to int64, period string) (data []models.ChartData, err error) {
	db := r.getDB(from, to, period)

	err = db.Select(fmt.Sprintf("date_trunc('%s', timestamp) as timestamp, count(1) blocks", period)).
		Table("tezos.blocks").
		Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *Repository) TransactionsVolume(from, to int64, period string) (data []models.ChartData, err error) {
	db := r.getDB(from, to, period)

	err = db.Select(fmt.Sprintf("date_trunc('%s', timestamp) as timestamp, sum(amount) transaction_volume", period)).
		Table("tezos.operations").
		Where("kind = 'transaction'").
		Where("status = 'applied'").
		Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *Repository) OperationsNumber(from, to int64, period string) (data []models.ChartData, err error) {
	db := r.getDB(from, to, period)

	err = db.Select(fmt.Sprintf("date_trunc('%s', timestamp) as timestamp, count(1) operations", period)).
		Table("tezos.operations").
		Where("status = 'applied'").
		Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *Repository) FeesVolume(from, to int64, period string) (data []models.ChartData, err error) {
	db := r.getDB(from, to, period)

	err = db.Select(fmt.Sprintf("date_trunc('%s', timestamp) as timestamp, sum(fee) fees", period)).
		Table("tezos.operations").
		Where("kind = 'transaction'").
		Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *Repository) ActivationsNumber(from, to int64, period string) (data []models.ChartData, err error) {
	db := r.getDB(from, to, period)

	err = db.Select(fmt.Sprintf("date_trunc('%s', timestamp) as timestamp, count(1) activations", period)).
		Table("tezos.operations").
		Where("kind = 'activate_account'").
		Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *Repository) AvgBlockDelay(from, to int64, period string) (data []models.ChartData, err error) {
	db := r.getDB(from, to, period)

	err = db.Select(fmt.Sprintf("date_trunc('%s', timestamp) as timestamp, extract(epoch from avg(block_delay)) average_delay", period)).
		Table("tezos.blocks_delay").
		Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *Repository) DelegationVolume(from, to int64, period string) (data []models.ChartData, err error) {
	db := r.getDB(from, to, period)

	err = db.Select(fmt.Sprintf("date_trunc('%s', timestamp) as timestamp, sum(delegation_amount) delegation_volume", period)).
		Table("tezos.delegations_view").
		Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *Repository) Bakers(from, to int64, period string) (data []models.ChartData, err error) {
	db := r.getDB(from, to, period)

	err = db.Select(fmt.Sprintf("date_trunc('%s', timestamp) as timestamp, count(distinct (baker)) bakers", period)).
		Table("tezos.blocks").
		Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *Repository) InsertWhaleAccounts(dayTime int64) (err error) {
	err = r.db.Exec("Select insert_whale_stat(?)", dayTime).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) WhaleAccounts(from, to int64, period string) (data []models.ChartData, err error) {
	err = r.db.Select("day as timestamp, whale_accounts").
		Table("tezos.whale_accounts_periods").
		Where("date_part('epoch',day)::int >= ?", from).
		Where("date_part('epoch',day)::int <= ?", to).
		Order("day asc").
		Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}
