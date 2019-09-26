package models

type Roll struct {
	Pkh        string `json:"pkh"`
	Rolls      uint   `json:"rolls"`
	BlockID    string `json:"block_id"`
	Block      *Block `json:"block"` // This line is infered from column name "block_id".
	BlockLevel uint   `json:"block_level"`
}
