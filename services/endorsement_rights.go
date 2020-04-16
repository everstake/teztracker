package services

import (
	"github.com/everstake/teztracker/models"
	"github.com/guregu/null"
)

func (t *TezTracker) GetAccountFutureEndorsementRights(accountID string, cycle int64, limits Limiter) (count int64, futureRights []models.FutureEndorsementRight, err error) {
	repo := t.repoProvider.GetFutureEndorsementRight()
	cycleFirstBlock := cycle*t.BlocksInCycle() + 1
	filter := models.RightFilter{
		BlockFilter: models.BlockFilter{
			FromID: null.IntFrom(cycleFirstBlock),
			ToID:   null.IntFrom(cycleFirstBlock + t.BlocksInCycle()),
		},
		Delegates: []string{accountID},
	}

	count, err = repo.Count(filter)
	if err != nil {
		return 0, nil, err
	}

	futureRights, err = repo.List(filter, limits.Limit(), limits.Offset())
	if err != nil {
		return 0, nil, err
	}

	for i := range futureRights {
		futureRights[i].Reward = int64(EndorsementReward * len(futureRights[i].Slots))
		futureRights[i].Deposit = BlockSecurityDeposit
	}

	return count, futureRights, nil
}
