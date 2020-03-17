package services

import (
	"github.com/everstake/teztracker/models"
	"github.com/guregu/null"
	"time"
)

// AccountList retrives up to limit of account before the specified id.
func (t *TezTracker) AccountList(before string, limits Limiter) (accs []models.Account, count int64, err error) {
	r := t.repoProvider.GetAccount()
	filter := models.AccountFilter{
		Type:  models.AccountTypeAccount,
		After: before,
	}
	count, accs, err = r.List(limits.Limit(), limits.Offset(), filter)
	return accs, count, err
}

// ContractList retrives up to limit of contract before the specified id.
func (t *TezTracker) ContractList(before string, limits Limiter) (accs []models.Account, count int64, err error) {
	r := t.repoProvider.GetAccount()
	filter := models.AccountFilter{
		Type:  models.AccountTypeContract,
		After: before,
	}
	count, accs, err = r.List(limits.Limit(), limits.Offset(), filter)
	return accs, count, err
}

// AccountDelegatorsList retrives up to limit of delegators accounts for the specified accountID.
func (t *TezTracker) AccountDelegatorsList(accountID string, limits Limiter) ([]models.Account, int64, error) {
	r := t.repoProvider.GetAccount()
	filter := models.Account{DelegateValue: accountID}
	count, err := r.Count(filter)
	if err != nil {
		return nil, 0, err
	}
	accs, err := r.Filter(filter, limits.Limit(), limits.Offset())
	return accs, count, err
}

// GetAccount retrieves an account by its ID.
func (t *TezTracker) GetAccount(id string) (acc models.Account, err error) {
	r := t.repoProvider.GetAccount()

	filter := models.Account{AccountID: null.StringFrom(id)}

	found, acc, err := r.Find(filter)
	if err != nil {
		return acc, err
	}
	if !found {
		return acc, ErrNotFound
	}

	counts, err := t.repoProvider.GetOperation().AccountOperationCount(acc.AccountID.String)
	if err != nil {
		return acc, err
	}

	var total int64
	for i := range counts {
		if counts[i].Kind == "transaction" {
			acc.Transactions = counts[i].Count
		}
		if counts[i].Kind == "reveal" {
			acc.IsRevealed = true
		}

		total += counts[i].Count
	}

	bi, err := t.GetBakerInfo(id)
	if err != nil {
		return acc, err
	}

	acc.BakerInfo = bi
	return acc, nil
}

func (t *TezTracker) GetAccountBalanceHistory(id string, from, to time.Time) (balances []models.AccountBalance, err error) {
	balances, err = t.repoProvider.GetAccount().Balances(id, from, to)
	if err != nil {
		return balances, err
	}

	return balances, nil
}
