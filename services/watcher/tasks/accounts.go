package tasks

import (
	"encoding/json"
	"fmt"
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/services"
)

type Account struct {
	repos services.Provider
}

func NewAccountTask(repos services.Provider) Account {

	return Account{repos: repos}
}

func (ac Account) GetEventData(data interface{}) (interface{}, error) {
	bt, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	account := models.Account{}

	err = json.Unmarshal(bt, &account)
	if err != nil {
		return nil, err
	}

	found, acc, err := ac.repos.GetAccount().Find(account)
	if err != nil {
		return acc, err
	}

	if !found {
		return acc, fmt.Errorf("Account not found: %s", account.AccountID.String)
	}

	return acc, nil
}
