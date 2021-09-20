package models

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
)

type BakingRight struct {
	BlockHash     string    `json:"block_hash"`
	Level         int64     `json:"level"`
	Delegate      string    `json:"delegate"`
	Priority      int       `json:"priority"`
	EstimatedTime time.Time `json:"estimated_time"`
}

type RightFilter struct {
	BlockFilter
	Delegates    []string
	PriorityFrom int
	PriorityTo   int
	Limit        null.Int
	Offset       null.Int
}

type FutureBakingRight struct {
	BlockLevel    int64         `json:"block_level"`
	Delegate      string        `json:"delegate"`
	DelegateName  string        `json:"delegate_name" gorm:"-"`
	Cycle         sql.NullInt64 `json:"cycle"`
	Priority      int           `json:"priority"`
	EstimatedTime time.Time     `json:"estimated_time"`
	Deposit       int64         `json:"deposit" gorm:"-"`
	Reward        int64         `json:"reward"  gorm:"-"`
	ForkId        string        `json:"fork_id"`
}

type FutureBlockBakingRight struct {
	Level  int64               `json:"level"`
	Rights []FutureBakingRight `json:"rights"`
}
