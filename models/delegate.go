package models

type Delegate struct {
	Pkh              string `gorm:"primary_key;AUTO_INCREMENT" json:"pkh"`
	BlockID          string `json:"block_id"`
	Block            *Block `json:"block"` // This line is infered from column name "block_id".
	Balance          int64  `json:"balance"`
	FrozenBalance    int64  `json:"frozen_balance"`
	StakingBalance   int64  `json:"staking_balance"`
	DelegatedBalance int64  `json:"delegated_balance"`
	Deactivated      bool   `json:"deactivated"`
	GracePeriod      uint   `json:"grace_period"`
	BlockLevel       uint   `json:"block_level" sql:"DEFAULT:'-1'::integer"`
}
