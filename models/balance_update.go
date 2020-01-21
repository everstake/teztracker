package models

type BalanceUpdate struct {
	ID                 uint    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Source             string  `json:"source"`
	SourceID           uint    `json:"source_id"`
	SourceHash         string  `json:"source_hash"`
	Kind               string  `json:"kind"`
	Contract           string  `json:"contract"`
	Change             float64 `json:"change"`
	Level              float64 `json:"level"`
	Delegate           string  `json:"delegate"`
	Category           string  `json:"category"`
	OperationGroupHash string  `json:"operation_group_hash" gorm:"column:operation_group_hash"`
}
