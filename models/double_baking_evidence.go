package models

type DoubleOperationType string

const (
	DoubleOperationTypeBaking      DoubleOperationType = "baking"
	DoubleOperationTypeEndorsement DoubleOperationType = "endorsement"
)

type DoubleOperationEvidence struct {
	OperationID    int64               `json:"operation_id" gorm:"column:operation_id"`
	Operation      Operation           `gorm:"save_associations:false;association_foreignkey:OperationID;foreignkey:OperationID"`
	Type           DoubleOperationType `json:"type" gorm:"column:doe_type"`
	BlockHash      string              `json:"block_hash" gorm:"column:doe_block_hash"`
	BlockLevel     int64               `json:"block_level" gorm:"column:doe_block_level"`
	DenouncedLevel int64               `json:"denounced_level" gorm:"column:doe_denounced_level"`
	Offender       string              `json:"offender" gorm:"column:doe_offender"`
	Priority       int                 `json:"priority" gorm:"column:doe_priority"`
	EvidenceBaker  string              `json:"evidence_baker" gorm:"column:doe_evidence_baker"`
	BakerReward    int64               `json:"baker_reward" gorm:"column:doe_baker_reward"`
	LostDeposits   int64               `json:"lost_deposits" gorm:"column:doe_lost_deposits"`
	LostRewards    int64               `json:"lost_rewards" gorm:"column:doe_lost_rewards"`
	LostFees       int64               `json:"lost_fees" gorm:"column:doe_lost_fees"`
}

type DoubleOperationEvidenceQueryOptions struct {
	BlockIDs        []string
	OperationHashes []string
	Type            DoubleOperationType
	LoadOperation   bool
	Limit           uint
	Offset          uint
}
