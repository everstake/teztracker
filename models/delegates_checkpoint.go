package models

type DelegatesCheckpoint struct {
	DelegatePkh string `json:"delegate_pkh"`
	BlockID     string `json:"block_id"`
	Block       *Block `json:"block"` // This line is infered from column name "block_id".
	BlockLevel  uint   `json:"block_level" sql:"DEFAULT:'-1'::integer"`
}
