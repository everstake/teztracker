package storage

import (
	"encoding/json"
	"fmt"

	"github.com/everstake/teztracker/models"
	"github.com/jinzhu/gorm"
)

const StorageTable = "tezos.storage"

type (
	// Repository is the storage repo implementation.
	Repository struct {
		db *gorm.DB
	}

	Repo interface {
		Set(key string, value interface{}) error
		Get(key string, dst interface{}) (bool, error)
	}
)

// New creates an instance of repository using the provided db.
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) getDb() *gorm.DB {
	db := r.db.Table(StorageTable)
	return db
}

func (r *Repository) Set(key string, value interface{}) error {
	if value == nil || key == "" {
		return fmt.Errorf("invalid key or data")
	}
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("json.Marshal: %s", err.Error())
	}
	var storage models.Storage
	res := r.getDb().Where("key = ?", key).First(&storage)
	storage = models.Storage{
		Key:   key,
		Value: string(data),
	}
	if res.Error != nil && res.RecordNotFound() {
		return r.getDb().Create(&storage).Error
	}
	if res.Error != nil {
		return res.Error
	}
	return r.getDb().Where("key = ?", key).Updates(&storage).Error
}

func (r *Repository) Get(key string, dst interface{}) (bool, error) {
	if dst == nil || key == "" {
		return false, fmt.Errorf("invalid key or data")
	}
	var model models.Storage
	res := r.getDb().Select("value").Where("key = ?", key).First(&model)

	if res.RecordNotFound() {
		return false, nil
	}

	if res.Error != nil {
		return false, res.Error
	}

	err := json.Unmarshal([]byte(model.Value), dst)
	if err != nil {
		return false, fmt.Errorf("json.Unmarshal: %s", err.Error())
	}

	return true, nil
}
