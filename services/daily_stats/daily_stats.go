package daily_stats

import (
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/services"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"reflect"
	"runtime"
	"time"
)

const (
	actionTimeInHours = 12
)

type (
	statsFunc  func(p services.Provider) (value decimal.Decimal, err error)
	DailyStats struct {
		providers []services.Provider
		tasks     map[string]statsFunc
	}
)

func NewDailyStats(providers []services.Provider) *DailyStats {
	return &DailyStats{
		providers: providers,
		tasks: map[string]statsFunc{
			models.LowBalanceAccountsStatKey: AccountsWithLowBalance,
			models.InactiveAccountsStatKey:   InActiveAccounts,
		},
	}
}

func (stats DailyStats) Run() {
	for {
		tn := time.Now()
		year, month, day := tn.Date()
		actionDay := time.Date(year, month, day, actionTimeInHours, 0, 0, 0, time.Local)
		waitDuration := actionDay.Sub(tn)
		if waitDuration < 0 {
			waitDuration = actionDay.Add(time.Hour * 24).Sub(tn)
		}
		<-time.After(waitDuration)
		now := time.Now()
		for _, provider := range stats.providers {
			for key, task := range stats.tasks {
				value, err := task(provider)
				if err != nil {
					funcName := runtime.FuncForPC(reflect.ValueOf(task).Pointer()).Name()
					log.Error("DailyStats: %s: %s", funcName, err.Error())
					continue
				}
				err = provider.GetDailyStats().Create(models.DailyStat{
					Key:   key,
					Date:  now,
					Value: value,
				})
				if err != nil {
					log.Error("DailyStats: GetDailyStats: Create: %s", err.Error())
					continue
				}
			}
		}
	}
}
