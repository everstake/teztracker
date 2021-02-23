package models

const (
	BakersChangesStorageKey = "bakers_changes"
)

type Storage struct {
	Key   string `gorm:"column:key" `
	Value string `gorm:"column:value" `
}
