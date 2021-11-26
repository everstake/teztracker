package daily_stats

import (
	"fmt"
	"github.com/everstake/teztracker/services"
	"github.com/shopspring/decimal"
	"time"
)

const (
	inactiveAccountsPeriod = time.Hour * 24 * 30 * 6 // ~6 month
)

func AccountsWithLowBalance(p services.Provider) (value decimal.Decimal, err error) {
	count, err := p.GetAccount().GetCountWhereBalance(1000000)
	if err != nil {
		return value, fmt.Errorf("GetCountWhereBalance: %s", err.Error())
	}
	return decimal.New(count, 0), nil
}

func InActiveAccounts(p services.Provider) (value decimal.Decimal, err error) {
	totalAccounts, err := p.GetAccount().GetCount(time.Time{}, time.Time{})
	if err != nil {
		return value, fmt.Errorf("GetCount: %s", err.Error())
	}
	activeAccounts, err := p.GetAccount().GetCountActive(time.Now().Add(-inactiveAccountsPeriod))
	if err != nil {
		return value, fmt.Errorf("GetCount: %s", err.Error())
	}
	return decimal.New(totalAccounts-activeAccounts, 0), nil
}
