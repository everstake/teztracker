package models

import (
	"time"

	"github.com/guregu/null"
)

const (
	SortAsc  = "asc"
	SortDesc = "desc"
)

type Operation struct {
	OperationID           null.Int    `gorm:"primary_key;AUTO_INCREMENT" json:"operation_id"`
	OperationGroupHash    null.String `json:"operation_group_hash"`
	Kind                  null.String `json:"kind"`
	Level                 int64       `json:"level"`
	Delegate              string      `json:"delegate"`
	Slots                 string      `json:"slots"`
	Nonce                 string      `json:"nonce"`
	Pkh                   string      `json:"pkh"`
	Secret                string      `json:"secret"`
	Source                string      `json:"source"`
	Fee                   int64       `json:"fee"`
	Counter               int64       `json:"counter"`
	GasLimit              int64       `json:"gas_limit"`
	StorageLimit          int64       `json:"storage_limit"`
	PublicKey             string      `json:"public_key"`
	Amount                int64       `json:"amount"`
	Destination           string      `json:"destination"`
	Parameters            string      `json:"parameters"`
	ParametersEntrypoints string      `json:"parameters_entrypoints"`
	ParametersMicheline   string      `json:"parameters_micheline"`
	ManagerPubkey         string      `json:"manager_pubkey"`
	Balance               int64       `json:"balance"`
	Spendable             bool        `json:"spendable"`
	Delegatable           bool        `json:"delegatable"`
	DelegationAmount      int64       `json:"delegation_amount" gorm:"column:balance"`
	Script                string      `json:"script"`
	Storage               string      `json:"storage"`
	Status                string      `json:"status"`
	Errors                string      `json:"errors"`
	ConsumedGas           int64       `json:"consumed_gas"`
	StorageSize           int64       `json:"storage_size"`
	PaidStorageSizeDiff   int64       `json:"paid_storage_size_diff"`
	OriginatedContracts   string      `json:"originated_contracts"`
	BlockHash             null.String `json:"block_hash"`
	BlockLevel            null.Int    `json:"block_level"`
	//Temp not unmarshal timestamp from json db because time.Time not support ISO without timezone
	Timestamp          time.Time `json:"-"`
	Branch             string    `json:"branch" gorm:"column:branch"`
	NumberOfSlots      int64     `json:"number_of_slots" gorm:"column:number_of_slots"`
	Cycle              int64     `json:"cycle" gorm:"column:cycle"`
	Proposal           string    `json:"proposal" gorm:"column:proposal"`
	Ballot             string    `json:"ballot" gorm:"column:ballot"`
	Internal           bool      `json:"internal" gorm:"column:internal"`
	Period             int64     `json:"period" gorm:"column:period"`
	Reward             int64     `json:"reward" gorm:"column:change"`
	DelegateName       string    `json:"delegate_name" gorm:"column:delegate_name"`
	SourceName         string    `json:"source_name" gorm:"column:source_name"`
	DestinationName    string    `json:"destination_name" gorm:"column:destination_name"`
	Confirmations      int64     `json:"confirmations"`
	EndorsementReward  int64     `json:"endorsement_reward" gorm:"column:endorsement_reward"`
	EndorsementDeposit int64     `json:"endorsement_reward" gorm:"column:endorsement_deposit"`
	ClaimedAmount      int64     `json:"claimed_amount" gorm:"column:claimed_amount"`
	Deposit            int64     `json:"deposit"`
	*DoubleOperationEvidenceExtended
}

type OperationCount struct {
	Kind  string `json:"kind"`
	Count int64  `json:"count"`
}
