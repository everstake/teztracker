package models

import "time"

type ChartData struct {
	Timestamp         time.Time
	Activations       int64
	AverageDelay      float64
	Blocks            int64
	DelegationVolume  int64
	Fees              int64
	Operations        int64
	TransactionVolume int64
	Bakers            int64
	BlockPriority     *BlockPriority
}

type BlockPriority struct {
	Timestamp      time.Time
	Blocks         int64
	ZeroPriority   int64
	FirstPriority  int64
	SecondPriority int64
}

type BakerChartData struct {
	Baker     string
	BakerName string
	Rolls     int64
	Percent   float64
}
