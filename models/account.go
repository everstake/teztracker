package models

import (
	"github.com/guregu/null"
	"time"
)

type Account struct {
	AccountID          null.String           `gorm:"primary_key;AUTO_INCREMENT" json:"account_id"`
	BlockID            null.String           `json:"block_id"`
	Block              *Block                `json:"block"` // This line is infered from column name "block_id".
	Manager            null.String           `json:"manager"`
	Spendable          null.Bool             `json:"spendable"`
	DelegateSetable    null.Bool             `json:"delegate_setable"`
	DelegateValue      string                `json:"delegate_value"`
	Counter            null.Int              `json:"counter"`
	Script             string                `json:"script"`
	Storage            string                `json:"storage"`
	Balance            null.Int              `json:"balance"`
	BlockLevel         null.Int              `json:"block_level" sql:"DEFAULT:'-1'::integer"`
	AccountsCheckpoint []*AccountsCheckpoint `json:"accounts_checkpoint"` // This line is infered from other tables.
	DelegatedContracts []*DelegatedContract  `json:"delegated_contracts"` // This line is infered from other tables.
	BakerInfo          *Baker                `json:"baker_info"`
	IsBaker            bool                  `json:"is_baker"`
	Asof               time.Time             `json:"asof"`
}

type AccountType int

const (
	AccountTypeBoth AccountType = iota
	AccountTypeAccount
	AccountTypeContract
)

type AccountFilter struct {
	Type     AccountType
	Delegate string
	After    string
}
