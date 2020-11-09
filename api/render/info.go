package render

import (
	genModels "github.com/everstake/teztracker/gen/models"
	"github.com/everstake/teztracker/models"
)

const annualYield = 7.12

// Info renders price info into OpenAPI model.
func Info(currency string, mi models.MarketInfo, ratio float64, blocks int64) *genModels.Info {
	p := mi.GetPrice()
	p24 := mi.GetPriceChange()
	ratioInPercent := ratio * 100
	vol := mi.GetVolume()
	mc := mi.GetMarketCap()
	return &genModels.Info{
		Currency:          currency,
		Price:             &p,
		Price24hChange:    &p24,
		StakingRatio:      &ratioInPercent,
		AnnualYield:       annualYield,
		MarketCap:         mc,
		Volume24h:         vol,
		CirculatingSupply: mi.GetSupply(),
		BlocksInCycle:     blocks,
	}
}

func ChartData(chd []models.ChartData) []*genModels.ChartsData {
	chds := make([]*genModels.ChartsData, len(chd))
	for i := range chd {
		chds[i] = ChartElement(chd[i])
	}
	return chds
}

func ChartElement(chd models.ChartData) *genModels.ChartsData {
	tm := chd.Timestamp.Unix()

	data := &genModels.ChartsData{
		Timestamp:         &tm,
		Activations:       chd.Activations,
		AverageDelay:      chd.AverageDelay,
		Blocks:            chd.Blocks,
		DelegationVolume:  chd.DelegationVolume,
		Fees:              chd.Fees,
		Operations:        chd.Operations,
		TransactionVolume: chd.TransactionVolume,
		Bakers:            chd.Bakers,
		WhaleAccounts:     chd.WhaleAccounts,
	}

	return data
}

func BakerChartData(chd []models.BakerChartData) []*genModels.BakerChartData {
	chds := make([]*genModels.BakerChartData, len(chd))
	for i := range chd {
		chds[i] = BakerChartElement(chd[i])
	}
	return chds
}

func BakerChartElement(chd models.BakerChartData) *genModels.BakerChartData {

	return &genModels.BakerChartData{
		Baker:     chd.Baker,
		BakerName: chd.BakerName,
		Rolls:     chd.Rolls,
		Percent:   chd.Percent,
	}
}

func BlocksPriorityChartData(chd []models.BlockPriority) []*genModels.BlockPriorityChartData {
	chds := make([]*genModels.BlockPriorityChartData, len(chd))
	for i := range chd {
		chds[i] = BlocksPriorityChartElement(chd[i])
	}
	return chds
}

func BlocksPriorityChartElement(chd models.BlockPriority) *genModels.BlockPriorityChartData {

	return &genModels.BlockPriorityChartData{
		Cycle:          &chd.Cycle,
		Blocks:         &chd.Blocks,
		FirstPriority:  &chd.FirstPriority,
		SecondPriority: &chd.SecondPriority,
		ZeroPriority:   &chd.ZeroPriority,
		ThirdPriority:  &chd.ThirdPriority,
	}
}
