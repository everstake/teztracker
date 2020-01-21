package models

import (
	"time"
)

type BakingRight struct {
	BlockHash     string    `json:"block_hash"`
	Level         int64     `json:"level"`
	Delegate      string    `json:"delegate"`
	Priority      int       `json:"priority"`
	EstimatedTime time.Time `json:"estimated_time"`
}

type BakingRightFilter struct {
	BlockFilter
	Delegates    []string
	PriorityFrom int
	PriorityTo   int
}

type FutureBakingRight struct {
	Level         int64     `json:"level"`
	Delegate      string    `json:"delegate"`
	Priority      int       `json:"priority"`
	EstimatedTime time.Time `json:"estimated_time"`
}

type FutureBlockBakingRight struct {
	Level  int64               `json:"level"`
	Rights []FutureBakingRight `json:"rights"`
}

type EndorsingRight struct {
	BlockHash     string    `json:"block_hash"`
	Level         int64     `json:"level"`
	Delegate      string    `json:"delegate"`
	Slot          int       `json:"slot"`
	EstimatedTime time.Time `json:"estimated_time"`
}
