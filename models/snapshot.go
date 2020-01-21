package models

type Snapshot struct {
	Cycle      int64 `gorm:"column:snp_cycle" json:"cycle"`
	BlockLevel int64 `gorm:"column:snp_block_level" json:"block_level"`
	Block      Block `gorm:"column:snp_block_level save_associations:false" json:"block"`
	Rolls      int64 `gorm:"column:snp_rolls" json:"rolls"`
}
