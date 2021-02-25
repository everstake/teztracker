package daily_stats

import (
	"fmt"
	"github.com/everstake/teztracker/services"
	"github.com/shopspring/decimal"
)

func AccountsWithLowBalance(p services.Provider) (value decimal.Decimal, err error) {
	count, err := p.GetAccount().GetCountWhereBalance(1000000)
	if err != nil {
		return value, fmt.Errorf("GetCountWhereBalance: %s", err.Error())
	}
	return decimal.New(count, 0), nil
}
