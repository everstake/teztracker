package services

import (
	"github.com/everstake/teztracker/models"
	"github.com/guregu/null"
)

func (t *TezTracker) GetAccountFutureEndorsementRights(accountID string, cycle int64, limits Limiter) (count int64, futureRights []models.FutureEndorsementRight, err error) {
	lastBlock, err := t.repoProvider.GetBlock().Last()
	if err != nil {
		return 0, nil, err
	}

	cycleFirstBlock := cycle*t.BlocksInCycle() + 1
	//Return future part of active cycle
	if lastBlock.MetaCycle == cycle {
		cycleFirstBlock = lastBlock.MetaLevel + 1
	}

	repo := t.repoProvider.GetFutureEndorsementRight()
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
		futureRights[i].Reward = getEndorsementRewardByCycle(futureRights[i].Cycle.Int64) * int64(len(futureRights[i].Slots))
		futureRights[i].Deposit = getEndorsementSecurityDepositByCycle(futureRights[i].Cycle.Int64)
	}

	return count, futureRights, nil
}
