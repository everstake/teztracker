package models

import (
	"github.com/shopspring/decimal"
	"time"
)

const (
	LowBalanceAccountsStatKey = "low_balance_accounts"
)

type DailyStat struct {
	Key   string          `gorm:"column:key"`
	Date  time.Time       `gorm:"column:date"`
	Value decimal.Decimal `gorm:"column:value"`
}
