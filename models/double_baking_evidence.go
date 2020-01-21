package models

type DoubleBakingEvidence struct {
	OperationID    int64     `json:"operation_id" gorm:"column:operation_id"`
	Operation      Operation `gorm:"save_associations:false;association_foreignkey:OperationID;foreignkey:OperationID"`
	BlockHash      string    `json:"block_hash" gorm:"column:dbe_block_hash"`
	BlockLevel     int64     `json:"block_level" gorm:"column:dbe_block_level"`
	DenouncedLevel int64     `json:"denounced_level" gorm:"column:dbe_denounced_level"`
	Offender       string    `json:"offender" gorm:"column:dbe_offender"`
	Priority       int       `json:"priority" gorm:"column:dbe_priority"`
	EvidenceBaker  string    `json:"evidence_baker" gorm:"column:dbe_evidence_baker"`
	BakerReward    int64     `json:"baker_reward" gorm:"column:dbe_baker_reward"`
	LostDeposits   int64     `json:"lost_deposits" gorm:"column:dbe_lost_deposits"`
	LostRewards    int64     `json:"lost_rewards" gorm:"column:dbe_lost_rewards"`
	LostFees       int64     `json:"lost_fees" gorm:"column:dbe_lost_fees"`
}

type DoubleBakingEvidenceQueryOptions struct {
	BlockIDs        []string
	OperationHashes []string
	LoadOperation   bool
	Limit           uint
	Offset          uint
}
