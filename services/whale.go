package services

import (
	"fmt"
	"time"

	"github.com/everstake/teztracker/models"
)

const minAmountForLargeTransfer = 3e11

func (t *TezTracker) GetWhaleTransfers(limits Limiter, period string) (transfers []models.Operation, err error) {

	timeSince := time.Time{}
	tt := time.Now()

	var cycle int64
	switch period {
	case "C":
		block, err := t.repoProvider.GetBlock().Last()
		if err != nil {
			return transfers, err
		}

		cycle = block.MetaCycle
	case "D":
		timeSince = tt.AddDate(0, 0, -1)
	case "W":
		timeSince = tt.AddDate(0, 0, -7)
	case "M":
		timeSince = tt.AddDate(0, -1, 0)
	default:
		return transfers, fmt.Errorf("wrong period")
	}

	transfers, err = t.repoProvider.GetOperation().LargeTransfers(minAmountForLargeTransfer, cycle, limits.Limit(), limits.Offset(), timeSince)
	if err != nil {
		return transfers, err
	}

	return transfers, nil
}
