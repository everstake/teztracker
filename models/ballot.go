package models

type Ballot struct {
	Pkh        string `json:"pkh"`
	Ballot     string `json:"ballot"`
	BlockID    string `json:"block_id"`
	Block      *Block `json:"block"` // This line is infered from column name "block_id".
	BlockLevel uint   `json:"block_level"`
}
