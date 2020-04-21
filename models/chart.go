package models

import "time"

type ChartData struct {
	Timestamp         time.Time
	Activations       int64
	AverageDelay      int64
	Blocks            int64
	DelegationVolume  int64
	Fees              int64
	Operations        int64
	TransactionVolume int64
}
