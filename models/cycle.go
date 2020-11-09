package models

import (
	"github.com/guregu/null"
)

type BakingCycle struct {
	Cycle      int64
	CycleStart null.Time
	CycleEnd   null.Time
}
