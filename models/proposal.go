package models

type Proposal struct {
	ProtocolHash string `json:"protocol_hash"`
	BlockID      string `json:"block_id"`
	Block        *Block `json:"block"` // This line is infered from column name "block_id".
	BlockLevel   uint   `json:"block_level"`
	Supporters   uint   `json:"supporters"`
}
