package models

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type FutureEndorsementRight struct {
	BlockLevel    int64         `json:"block_level"`
	BlockHash     string        `json:"block_hash"`
	Delegate      string        `json:"delegate"`
	DelegateName  string        `json:"delegate_name" gorm:"-"`
	Cycle         sql.NullInt64 `json:"cycle"`
	Slots         pq.Int64Array `json:"slots" gorm: "type:integer[]"`
	EstimatedTime time.Time     `json:"estimated_time"`

	ForkId  string `json:"fork_id"`
	Reward  int64  `json:"reward"  gorm:"-"`
	Deposit int64  `json:"deposit"  gorm:"-"`
}

type FutureBlockEndorsementRight struct {
	Level  int64                    `json:"level"`
	Rights []FutureEndorsementRight `json:"rights"`
}

type EndorsementRight struct {
	BlockLevel    int64          `json:"block_level"`
	BlockHash     sql.NullString `json:"block_hash"`
	Cycle         sql.NullInt64  `json:"cycle"`
	Delegate      string         `json:"delegate"`
	ForkId        string         `json:"fork_id"`
	Slot          int64          `json:"slot"`
	EstimatedTime time.Time      `json:"estimated_time"`
}
