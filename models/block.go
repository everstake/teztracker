package models

import (
	"time"

	"github.com/guregu/null"
)

type Block struct {
	Level                    null.Int               `json:"level"`
	Proto                    null.Int               `json:"proto"`
	Predecessor              null.String            `json:"predecessor"`
	Timestamp                time.Time              `json:"timestamp"`
	BlockTime                int64                  `json:"block_time"`
	ValidationPass           null.Int               `json:"validation_pass"`
	Fitness                  null.String            `json:"fitness"`
	Context                  string                 `json:"context"`
	Signature                string                 `json:"signature"`
	Protocol                 null.String            `json:"protocol"`
	ChainID                  string                 `json:"chain_id"`
	Hash                     null.String            `json:"hash"`
	OperationsHash           string                 `json:"operations_hash"`
	PeriodKind               string                 `json:"period_kind"`
	CurrentExpectedQuorum    int64                  `json:"current_expected_quorum"`
	ActiveProposal           string                 `json:"active_proposal"`
	Baker                    string                 `json:"baker"`
	BakerName                string                 `json:"baker_name"`
	Reward                   int64                  `json:"reward"`
	NonceHash                string                 `json:"nonce_hash"`
	ConsumedGas              int64                  `json:"consumed_gas"`
	MetaLevel                int64                  `json:"meta_level"`
	MetaLevelPosition        int64                  `json:"meta_level_position"`
	MetaCycle                int64                  `json:"meta_cycle"`
	MetaCyclePosition        int64                  `json:"meta_cycle_position"`
	MetaVotingPeriod         int64                  `json:"meta_voting_period"`
	MetaVotingPeriodPosition int64                  `json:"meta_voting_period_position"`
	ExpectedCommitment       bool                   `json:"expected_commitment"`
	Priority                 null.Int               `json:"priority" gorm:"column:priority"`
	BlockAggregation         *BlockAggregationView  `json:"-"`
	Delegates                []*Delegate            `json:"delegates"`            // This line is infered from other tables.
	Proposals                []*Proposal            `json:"proposals"`            // This line is infered from other tables.
	Rolls                    []*Roll                `json:"rolls"`                // This line is infered from other tables.
	Ballots                  []*Ballot              `json:"ballots"`              // This line is infered from other tables.
	AccountsCheckpoint       []*AccountsCheckpoint  `json:"accounts_checkpoint"`  // This line is infered from other tables.
	OperationGroups          []*OperationGroup      `json:"operation_groups"`     // This line is infered from other tables.
	DelegatesCheckpoint      []*DelegatesCheckpoint `json:"delegates_checkpoint"` // This line is infered from other tables.
	Accounts                 []*Account             `json:"accounts"`             // This line is infered from other tables.
	BakingRights             []FutureBakingRight    `json:"baking_rights,omitempty"`
}

type BlockFilter struct {
	FromID      null.Int
	ToID        null.Int
	BlockLevels []int64
	BlockHashes []string
}
