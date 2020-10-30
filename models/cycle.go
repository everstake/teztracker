package models

import "time"

type BakingCycle struct {
	Cycle      int64
	CycleStart time.Time
	CycleEnd   time.Time
}
