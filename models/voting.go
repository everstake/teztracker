package models

import "time"

type PeriodType string

type PeriodInfo struct {
	Rolls       int64
	Bakers      int64
	BlockLevel  int64
	Period      int64
	Kind        string
	StartBlock  int64
	EndBlock    int64
	Cycle       int8
	StartTime   time.Time
	EndTime     time.Time
	TotalBakers int64
	TotalRolls  int64
}
