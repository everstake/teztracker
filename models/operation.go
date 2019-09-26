package models

import (
	"time"

	"github.com/guregu/null"
)

type Operation struct {
	OperationID         null.Int    `gorm:"primary_key;AUTO_INCREMENT" json:"operation_id"`
	OperationGroupHash  null.String `json:"operation_group_hash"`
	Kind                null.String `json:"kind"`
	Level               int64       `json:"level"`
	Delegate            string      `json:"delegate"`
	Slots               string      `json:"slots"`
	Nonce               string      `json:"nonce"`
	Pkh                 string      `json:"pkh"`
	Secret              string      `json:"secret"`
	Source              string      `json:"source"`
	Fee                 int64     `json:"fee"`
	Counter             int64     `json:"counter"`
	GasLimit            int64     `json:"gas_limit"`
	StorageLimit        int64     `json:"storage_limit"`
	PublicKey           string      `json:"public_key"`
	Amount              int64     `json:"amount"`
	Destination         string      `json:"destination"`
	Parameters          string      `json:"parameters"`
	ManagerPubkey       string      `json:"manager_pubkey"`
	Balance             int64       `json:"balance"`
	Spendable           bool        `json:"spendable"`
	Delegatable         bool        `json:"delegatable"`
	Script              string      `json:"script"`
	Storage             string      `json:"storage"`
	Status              string      `json:"status"`
	ConsumedGas         int64     `json:"consumed_gas"`
	StorageSize         int64     `json:"storage_size"`
	PaidStorageSizeDiff int64     `json:"paid_storage_size_diff"`
	OriginatedContracts string      `json:"originated_contracts"`
	BlockHash           null.String `json:"block_hash"`
	BlockLevel          null.Int    `json:"block_level"`
	Timestamp           time.Time   `json:"timestamp"`
}
