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
		Get(key string, dst interface{}) error
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
	fmt.Println(key, value == nil)
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

func (r *Repository) Get(key string, dst interface{}) error {
	if dst == nil || key == "" {
		return fmt.Errorf("invalid key or data")
	}
	var data string
	res := r.getDb().Where("key = ?", key).First(&data)
	if res.Error != nil {
		return res.Error
	}
	if res.RecordNotFound() {
		return fmt.Errorf("not found by key: %s", key)
	}
	err := json.Unmarshal([]byte(data), dst)
	if err != nil {
		return fmt.Errorf("json.Unmarshal: %s", err.Error())
	}
	return nil
}
