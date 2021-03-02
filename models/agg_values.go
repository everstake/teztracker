package models

import (
	"fmt"
	"time"
)

const (
	DayPeriod   = "day"
	WeekPeriod  = "week"
	MonthPeriod = "month"
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
	return ValidatePeriod(agg.Period)
}

func GetChartPeriods() map[string]time.Duration {
	return map[string]time.Duration{
		DayPeriod:   time.Hour * 24 * 30,  // 30 days, 30 values
		WeekPeriod:  time.Hour * 24 * 180, // 180 days, 25 values
		MonthPeriod: time.Hour * 24 * 365, // 365 days, 12 values
	}
}

func ValidatePeriod(period string) error {
	switch {
	case period == DayPeriod:
	case period == MonthPeriod:
	case period == WeekPeriod:
	default:
		return fmt.Errorf("unknown period: %s", period)
	}
	return nil
}
