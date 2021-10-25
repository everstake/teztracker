package services

import (
	"github.com/everstake/teztracker/models"
)

func (t *TezTracker) GetChartsInfo(from, to int64, period string, columns []string) (data []models.ChartData, err error) {
	repo := t.repoProvider.GetChart()

	switch period {
	case "D":
		period = "day"
	}

	//TODO Refactor
	for i := range columns {
		switch columns[i] {
		case "blocks":
			data, err = repo.BlocksNumber(from, to, period)
		case "volume":
			data, err = repo.TransactionsVolume(from, to, period)
		case "operations":
			data, err = repo.OperationsNumber(from, to, period, "", nil, nil)
		case "avg_block_delay":
			data, err = repo.AvgBlockDelay(from, to, period)
		case "fees":
			data, err = repo.FeesVolume(from, to, period)
		case "activations":
			data, err = repo.ActivationsNumber(from, to, period)
		case "delegation_volume":
			data, err = repo.DelegationVolume(from, to, period)
		case "bakers":
			data, err = repo.Bakers(from, to, period)
		case "whale_accounts":
			data, err = repo.WhaleAccounts(from, to, period)
		}
		if err != nil {
			return data, err
		}
	}

	return data, nil
}

func (t *TezTracker) GetBakerChartInfo(limits Limiter) (data []models.BakerChartData, err error) {
	br := t.repoProvider.GetBaker()
	stakedBalance, err := br.TotalStakingBalance()
	if err != nil {
		return nil, err
	}

	totalRolls := stakedBalance / TokensPerRoll / XTZ

	bakers, err := br.List(limits.Limit(), 0, nil)
	if err != nil {
		return nil, err
	}

	for i := range bakers {
		percent := float64(bakers[i].Rolls) / float64(totalRolls)
		data = append(data, models.BakerChartData{
			Baker:     bakers[i].AccountID,
			BakerName: bakers[i].Name,
			Rolls:     bakers[i].Rolls,
			Percent:   percent,
		})
	}

	return data, nil
}

func (t *TezTracker) GetBlocksPriorityByCycle(limits Limiter) (data []models.BlockPriority, err error) {
	repo := t.repoProvider.GetBlock()

	data, err = repo.BlocksPriority(limits.Limit())
	if err != nil {
		return data, err
	}

	return data, err
}
