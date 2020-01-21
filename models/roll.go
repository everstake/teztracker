package models

type Roll struct {
	Pkh        string `json:"pkh"`
	Rolls      int64  `json:"rolls"`
	BlockID    string `json:"block_id"`
	Block      *Block `json:"block"` // This line is infered from column name "block_id".
	BlockLevel int64  `json:"block_level"`
}
