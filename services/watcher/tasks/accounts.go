package tasks

import (
	"encoding/json"
	"fmt"
	"github.com/everstake/teztracker/api/render"
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/services"
	"strings"
)

type Account struct {
	repos services.Provider
}

func NewAccountTask(repos services.Provider) Account {

	return Account{repos: repos}
}

func (ac Account) GetEventData(data interface{}) ([]string, interface{}, error) {
	bt, err := json.Marshal(data)
	if err != nil {
		return nil, nil, err
	}

	account := models.Account{}

	err = json.Unmarshal(bt, &account)
	if err != nil {
		return nil, nil, err
	}

	found, acc, err := ac.repos.GetAccount().Find(account)
	if err != nil {
		return nil, acc, err
	}

	if !found {
		return nil, acc, fmt.Errorf("Account not found: %s", account.AccountID.String)
	}

	apiAccount := render.Account(acc)

	accountType := "accounts"
	if strings.Contains(acc.AccountID.String, models.ContractAccountPrefix) {
		accountType = "contracts"
	}

	return []string{accountType}, apiAccount, nil
}
