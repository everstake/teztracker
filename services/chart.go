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
			if err != nil {
				return data, err
			}
		case "volume":
			data, err = repo.TransactionsVolume(from, to, period)
			if err != nil {
				return data, err
			}
		case "operations":
			data, err = repo.OperationsNumber(from, to, period)
			if err != nil {
				return data, err
			}
		case "avg_block_delay":
			data, err = repo.AvgBlockDelay(from, to, period)
			if err != nil {
				return data, err
			}

		case "fees":
			data, err = repo.FeesVolume(from, to, period)
			if err != nil {
				return data, err
			}
		case "activations":
			data, err = repo.ActivationsNumber(from, to, period)
			if err != nil {
				return data, err
			}
		case "delegation_volume":
			data, err = repo.DelegationVolume(from, to, period)
			if err != nil {
				return data, err
			}
		case "bakers":
			data, err = repo.Bakers(from, to, period)
			if err != nil {
				return data, err
			}
		}
	}

	return data, nil
}

func (t *TezTracker) GetBakerChartInfo(limits Limiter) (data []models.BakerChartData, err error) {
	br := t.repoProvider.GetBaker()

	bakers, err := br.List(limits.Limit(), 0)
	if err != nil {
		return nil, err
	}

	for i := range bakers {
		data = append(data, models.BakerChartData{
			Baker:     bakers[i].AccountID,
			BakerName: bakers[i].Name,
			Rolls:     bakers[i].Rolls,
		})
	}

	return data, nil
}
