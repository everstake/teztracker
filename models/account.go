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
	IsRevealed         bool                  `json:"is_revealed"`
	Transactions       int64                 `json:"transactions"`
	Operations         int64                 `json:"operations"`
	Index              int64                 `json:"index"`
}

type AccountListView struct {
	Account
	AccountName  string    `json:"account_name"`
	DelegateName string    `json:"delegate_name"`
	CreatedAt    time.Time `json:"created_at"`
	LastActive   time.Time `json:"last_active"`
}

type AccountType int

const (
	AccountTypeBoth AccountType = iota
	AccountTypeAccount
	AccountTypeContract
)

type AccountOrderField int

const (
	AccountOrderFieldBalance AccountOrderField = iota
	AccountOrderFieldCreatedAt
)

type RewardStatus string

const (
	StatusPending  RewardStatus = "pending"
	StatusActive   RewardStatus = "active"
	StatusFrozen   RewardStatus = "frozen"
	StatusUnfrozen RewardStatus = "unfrozen"
)

type AccountPrefix string

const (
	ImplicitAccountPrefix = "tz"
	ContractAccountPrefix = "KT1"
)

type AccountFilter struct {
	Type     AccountType
	OrderBy  AccountOrderField
	Delegate string
	After    string
}

type AccountBalance struct {
	Time    time.Time
	Balance int64
}

type AccountBaking struct {
	BakingCycle
	Status       RewardStatus
	Count        int64
	Missed       int64
	Reward       int64
	AvgPriority  float32
	Stolen       int64
	TotalDeposit int64
}

type AccountEndorsing struct {
	BakingCycle
	Status       RewardStatus
	Count        int64
	Missed       int64
	Reward       int64
	TotalDeposit int64
}

//TODO refactor Account rewards models
type AccountReward struct {
	BakingCycle
	Status                 RewardStatus
	Delegators             int64
	StakingBalance         int64
	BakingRewards          int64
	FutureBakingCount      int64
	EndorsementRewards     int64
	FutureEndorsementCount int64
	Fees                   int64
	MissedBaking           int64
	MissedEndorsements     int64
	Losses                 int64
}

type AccountRewardsCount struct {
	BakingCycle
	Status                 RewardStatus
	StakingBalance         int64
	BakingCount            int64
	BakingReward           int64
	StolenBaking           int64
	FutureBakingCount      int64
	FutureEndorsementCount int64
	EndorsementsCount      int64
	EndorsementsReward     int64
	//Deposit
	ActualBakingSecurityDeposit        int64
	ExpectedBakingSecurityDeposit      int64
	ActualEndorsementSecurityDeposit   int64
	ExpectedEndorsementSecurityDeposit int64
	ActualTotalSecirityDeposit         int64
	ExpectedTotalSecurityDeposit       int64
	AvailableBond                      int64
}

type AccountDelegator struct {
	AccountId string
	Cycle     int64
	Balance   int64
	Share     float64
}

type AccountReportFilter struct {
	From         int64
	To           int64
	Limit        int64
	Operations   []string
	EndorsingReq bool
}

type AccountReport struct {
	BlockLevel         uint64      `csv:"block level"`
	OperationGroupHash null.String `csv:"-"`
	Timestamp          time.Time   `csv:"timestamp"`
	Kind               string      `csv:"operation type"`
	Coin               string      `csv:"coin"`
	//DB field
	Amount      uint64 `csv:"-"`
	Fee         uint64 `csv:"fee"`
	Source      string `csv:"-"`
	Destination string `csv:"-"`
	Status      string `csv:"status"`

	//CSV field
	In   uint64 `csv:"in"`
	Out  uint64 `csv:"out"`
	Link string `csv:"link"`
}

type BakerReport struct {
	AccountReport
	//Baker operations
	Reward uint64 `csv:"reward"`
	Loss   uint64 `csv:"loss"`
}
