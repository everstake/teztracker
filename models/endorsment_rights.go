package models

import (
	"github.com/lib/pq"
	"time"
)

type FutureEndorsementRight struct {
	Level         int64         `json:"level"`
	Delegate      string        `json:"delegate"`
	DelegateName  string        `json:"delegate_name" gorm:"-"`
	Cycle         int64         `json:"cycle"`
	Slots         pq.Int64Array `json:"slots" gorm: "type:integer[]"`
	EstimatedTime time.Time     `json:"estimated_time"`
	Reward        int64         `json:"reward"  gorm:"-"`
	Deposit       int64         `json:"deposit"  gorm:"-"`
}

type FutureBlockEndorsementRight struct {
	Level  int64                    `json:"level"`
	Rights []FutureEndorsementRight `json:"rights"`
}

type EndorsingRight struct {
	BlockHash     string    `json:"block_hash"`
	Level         int64     `json:"level"`
	Delegate      string    `json:"delegate"`
	Slot          int       `json:"slot"`
	EstimatedTime time.Time `json:"estimated_time"`
}
