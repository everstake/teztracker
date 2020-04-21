package services

import "github.com/everstake/teztracker/models"

func (t *TezTracker) GetAccountDelegatorsByCycle(accountID string, cycle int64, limits Limiter) (count int64, accountDelegators []models.AccountDelegator, err error) {
	accRepo := t.repoProvider.GetAccount()
	cycleTotal, err := accRepo.CycleDelegatorsTotal(accountID, cycle)
	if err != nil {
		return 0, nil, err
	}

	delegators, err := accRepo.CycleDelegators(accountID, cycle, limits.Limit(), limits.Offset())
	if err != nil {
		return 0, nil, err
	}

	for i := range delegators {
		delegators[i].Share = float64(delegators[i].Balance) / float64(cycleTotal.StakingBalance)
	}

	return cycleTotal.Delegators, delegators, nil
}
