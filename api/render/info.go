package render

import (
	genModels "github.com/bullblock-io/tezTracker/gen/models"
	"github.com/bullblock-io/tezTracker/models"
)

const annualYield = 7.12

// Info renders price info into OpenAPI model.
func Info(mi models.MarketInfo, ratio float64) *genModels.Info {
	p := mi.GetPrice()
	p24 := mi.GetPriceChange()
	ratioInPercent := ratio * 100
	return &genModels.Info{Price: &p, Price24hChange: &p24, StakingRatio: &ratioInPercent, AnnualYield: annualYield}
}
