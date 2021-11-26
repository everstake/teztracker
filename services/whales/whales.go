package whales

import (
	"fmt"
	"sync"

	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/repos"
	"github.com/jinzhu/gorm"
)

const (
	numberWhaleAccounts = 500
)

var Service *Whales

type (
	Whales struct {
		mu       *sync.RWMutex
		networks map[models.Network]network
	}
	network struct {
		db   *gorm.DB
		data Data
	}
	Data struct {
		Accounts  []models.Account
		Transfers []models.Operation
	}
)

func init() {
	Service = newInstance()
}

func newInstance() *Whales {
	return &Whales{
		mu:       &sync.RWMutex{},
		networks: make(map[models.Network]network),
	}
}

func (w *Whales) AddNetwork(net models.Network, db *gorm.DB) {
	w.mu.Lock()
	w.networks[net] = network{db: db}
	w.mu.Unlock()
}

func (w *Whales) Update() error {
	err := w.updateWhaleAccounts()
	if err != nil {
		return fmt.Errorf("updateWhaleAccounts: %s", err.Error())
	}

	return nil
}

func (w *Whales) GetData(net models.Network) Data {
	w.mu.RLock()
	defer w.mu.RUnlock()
	n, ok := w.networks[net]
	if !ok {
		return Data{}
	}
	return n.data
}

func (w *Whales) updateWhaleAccounts() error {
	for net, item := range w.networks {
		reposProvider := repos.New(item.db)
		accountsRepo := reposProvider.GetAccount()
		richAccounts, err := accountsRepo.RichAccounts(numberWhaleAccounts)
		if err != nil {
			return fmt.Errorf("accountsRepo.RichAccounts: %s", err.Error())
		}
		accountsIDs := make([]string, len(richAccounts))
		for i := range richAccounts {
			accountsIDs[i] = richAccounts[i].AccountID.String
		}
		operationsRepo := reposProvider.GetOperation()
		kind := []string{"transaction"}
		transfers, err := operationsRepo.ListAsc(kind, accountsIDs, 10000, 0, 0)
		if err != nil {
			return fmt.Errorf("operationsRepo.List: %s", err.Error())
		}
		w.mu.Lock()
		n := w.networks[net]
		n.data.Transfers = transfers
		n.data.Accounts = richAccounts
		w.networks[net] = n
		w.mu.Unlock()
	}
	return nil
}
