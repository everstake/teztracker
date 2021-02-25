package models

import (
	"fmt"
	"time"
)

const (
	DayPeriod   = "day"
	MonthPeriod = "month"
	WeekPeriod  = "week"
)

type (
	AggTimeInt struct {
		Date  time.Time `json:"date"`
		Value int64     `json:"value"`
	}
	AggTimeFilter struct {
		From   time.Time
		To     time.Time
		Period string
	}
)

func (agg AggTimeFilter) Validate() error {
	switch {
	case agg.Period == DayPeriod:
	case agg.Period == MonthPeriod:
	case agg.Period == WeekPeriod:
	default:
		return fmt.Errorf("unknown period: %s", agg.Period)
	}
	return nil
}
