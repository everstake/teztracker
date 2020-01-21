package models

import (
	"time"
)

type OperationCounter struct {
	ID              int64     `gorm:"column:cnt_id;primary_key;AUTO_INCREMENT"`
	LastOperationID int64     `gorm:"column:cnt_last_op_id"`
	OperationType   string    `gorm:"column:cnt_operation_type"`
	Count           int64     `gorm:"column:cnt_count"`
	CreatedAt       time.Time `gorm:"column:cnt_created_at"`
}
