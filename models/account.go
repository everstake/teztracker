package models

import (
	"github.com/guregu/null"
	"time"
)

type Account struct {
	AccountID          null.String           `gorm:"primary_key;AUTO_INCREMENT" json:"account_id"`
	AccountName        string                `json:"baker_name"`
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
	CreatedAt          time.Time             `json:"created_at"`
	LastActive         time.Time             `json:"last_active"`
	IsRevealed         bool                  `json:"is_revealed"`
	Transactions       int64                 `json:"transactions"`
	Operations         int64                 `json:"operations"`
}

type AccountType int

const (
	AccountTypeBoth AccountType = iota
	AccountTypeAccount
	AccountTypeContract
)

type RewardStatus string

const (
	StatusPending  RewardStatus = "pending"
	StatusActive   RewardStatus = "active"
	StatusFrozen   RewardStatus = "frozen"
	StatusUnfrozen RewardStatus = "unfrozen"
)

type AccountFilter struct {
	Type     AccountType
	Delegate string
	After    string
}

type AccountBalance struct {
	Time    time.Time
	Balance int64
}

type AccountBaking struct {
	Cycle        int64
	Status       RewardStatus
	Count        int64
	Missed       int64
	Reward       int64
	AvgPriority  float32
	Stolen       int64
	TotalDeposit int64
}

type AccountEndorsing struct {
	Cycle        int64
	Status       RewardStatus
	Count        int64
	Missed       int64
	Reward       int64
	TotalDeposit int64
}

type AccountReward struct {
	Status                 RewardStatus
	Cycle                  int64
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

type AccountDelegator struct {
	AccountId string
	Cycle     int64
	Balance   int64
	Share     float64
}
