package services

import (
	"encoding/hex"
	"time"

	"fmt"

	chain "blockwatch.cc/tzgo/tezos"
	"github.com/everstake/teztracker/models"
	"github.com/guregu/null"
)

const activeBalanceCacheKey = "active_balance"

// AccountList retrives up to limit of account before the specified id.
func (t *TezTracker) AccountList(before string, limits Limiter, favorites []string) (accs []models.AccountListView, count int64, err error) {
	r := t.repoProvider.GetAccount()
	filter := models.AccountFilter{
		Type:      models.AccountTypeAccount,
		OrderBy:   models.AccountOrderFieldCreatedAt,
		After:     before,
		Favorites: favorites,
	}
	count, accs, err = r.List(limits.Limit(), limits.Offset(), filter)
	return accs, count, err
}

func (t *TezTracker) AccountTopBalanceList(before string, limits Limiter, favorites []string) (accs []models.AccountListView, count int64, err error) {
	r := t.repoProvider.GetAccount()
	filter := models.AccountFilter{
		Type:      models.AccountTypeBoth,
		OrderBy:   models.AccountOrderFieldBalance,
		After:     before,
		Favorites: favorites,
	}
	count, accs, err = r.List(limits.Limit(), limits.Offset(), filter)
	if err != nil {
		return accs, count, err
	}

	for i := range accs {
		accs[i].Index = int64(int(limits.Offset()) + i + 1)
	}
	return accs, count, err
}

// ContractList retrives up to limit of contract before the specified id.
func (t *TezTracker) ContractList(before string, limits Limiter, favorites []string) (accs []models.AccountListView, count int64, err error) {
	r := t.repoProvider.GetAccount()
	filter := models.AccountFilter{
		Type:      models.AccountTypeContract,
		OrderBy:   models.AccountOrderFieldCreatedAt,
		After:     before,
		Favorites: favorites,
	}
	count, accs, err = r.List(limits.Limit(), limits.Offset(), filter)
	return accs, count, err
}

// AccountDelegatorsList retrives up to limit of delegators accounts for the specified accountID.
func (t *TezTracker) AccountDelegatorsList(accountID string, limits Limiter) ([]models.AccountListView, int64, error) {
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
func (t *TezTracker) GetAccount(id string) (acc models.AccountListView, err error) {
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

	acc.Operations = total

	bi, err := t.GetBakerInfo(id)
	if err != nil {
		return acc, err
	}

	acc.BakerInfo = bi

	//Account identified as baker
	if bi != nil {
		//Set real value for front
		acc.IsBaker = true
	}

	return acc, nil
}

func (t *TezTracker) GetAccountBalanceHistory(id string, from, to time.Time) (balances []models.AccountBalance, err error) {
	repo := t.repoProvider.GetAccount()
	balances, err = repo.Balances(id, from, to)
	if err != nil {
		return balances, err
	}

	found, bal, err := repo.PrevBalance(id, from)
	if err != nil {
		return balances, err
	}

	if found {
		balances = append([]models.AccountBalance{bal}, balances...)
	}

	return balances, nil
}

func (t *TezTracker) GetAccountAssetsBalance(address string) (balances []models.AccountAssetBalance, err error) {

	adr, err := base58ToHexAddress(address)
	if err != nil {
		return balances, err
	}

	balances, err = t.repoProvider.GetAssets().GetAccountAssetsBalances(adr)
	if err != nil {
		return balances, nil
	}

	return balances, nil
}

func base58ToHexAddress(address string) (string, error) {
	adr, err := chain.ParseAddress(address)
	if err != nil {
		return "", nil
	}

	bt, err := adr.MarshalBinary()
	if err != nil {
		return "", nil
	}

	return hex.EncodeToString(bt), nil
}

func (t *TezTracker) GetAccountCountByPeriod(filter models.AggTimeFilter) (items []models.AggTimeInt, err error) {
	repo := t.repoProvider.GetAccount()
	return repo.GetCountByPeriod(filter)
}

func (t *TezTracker) GetTotalAccountCountByPeriod(filter models.AggTimeFilter) (items []models.AggTimeInt, err error) {
	repo := t.repoProvider.GetAccount()
	items, err = repo.GetCountByPeriod(filter)
	if err != nil {
		return nil, fmt.Errorf("GetCountByPeriod: %s", err.Error())
	}
	total, err := repo.GetCount(time.Time{}, filter.To)
	if err != nil {
		return nil, fmt.Errorf("GetCount: %s", err.Error())
	}
	for i := 0; i < len(items); i++ {
		v := items[i]
		v.Value += total
		items[i] = v
	}
	return items, nil
}

func (t *TezTracker) GetContractCountByPeriod(filter models.AggTimeFilter) (items []models.AggTimeInt, err error) {
	repo := t.repoProvider.GetAccount()
	return repo.GetContractsCountByPeriod(filter)
}

func (t *TezTracker) GetTotalContractCountByPeriod(filter models.AggTimeFilter) (items []models.AggTimeInt, err error) {
	repo := t.repoProvider.GetAccount()
	items, err = repo.GetContractsCountByPeriod(filter)
	if err != nil {
		return nil, fmt.Errorf("GetContractsCountByPeriod: %s", err.Error())
	}
	total, err := repo.GetContractsCount(time.Time{}, filter.To)
	if err != nil {
		return nil, fmt.Errorf("GetContractsCount: %s", err.Error())
	}
	for i := 0; i < len(items); i++ {
		v := items[i]
		v.Value += total
		items[i] = v
	}
	return items, nil
}

func (t *TezTracker) SaveActiveAccountsInCache() error {
	repo := t.repoProvider.GetAccount()
	for period, duration := range models.GetChartPeriods() {
		items, err := repo.GetCountActiveByPeriod(models.AggTimeFilter{
			From:   time.Now().Add(-duration),
			Period: period,
		})
		if err != nil {
			return fmt.Errorf("repo.GetCountActiveByPeriod: %s", err.Error())
		}
		storageKey := fmt.Sprintf("%s_%s", activeBalanceCacheKey, period)
		err = t.repoProvider.GetStorage().Set(storageKey, items)
		if err != nil {
			return fmt.Errorf("GetStorage: Set: %s", err.Error())
		}
	}
	return nil
}

func (t *TezTracker) GetActiveAccounts(period string) (items []models.AggTimeInt, err error) {
	err = models.ValidatePeriod(period)
	if err != nil {
		return items, fmt.Errorf("ValidatePeriod: %s", err.Error())
	}
	storageKey := fmt.Sprintf("%s_%s", activeBalanceCacheKey, period)
	//Return only if storage error
	_, err = t.repoProvider.GetStorage().Get(storageKey, &items)
	if err != nil {
		return items, fmt.Errorf("GetStorage: Set: %s", err.Error())
	}
	return items, nil
}

func (t *TezTracker) GetAccountsWithLowBalance(filter models.AggTimeFilter) (items []models.AggTimeInt, err error) {
	return t.repoProvider.GetDailyStats().GetDailyStats(models.LowBalanceAccountsStatKey, "avg", filter)
}

func (t *TezTracker) GetInactiveAccounts(filter models.AggTimeFilter) (items []models.AggTimeInt, err error) {
	return t.repoProvider.GetDailyStats().GetDailyStats(models.InactiveAccountsStatKey, "avg", filter)
}
