package models

import (
	"github.com/guregu/null"
)

type OperationGroup struct {
	Protocol  null.String `json:"protocol"`
	ChainID   string      `json:"chain_id"`
	Hash      null.String `gorm:"primary_key;AUTO_INCREMENT" json:"hash"`
	Branch    null.String `json:"branch"`
	Signature string      `json:"signature"`
	BlockID   null.String `json:"block_id"`
	Block     *Block      `json:"block"` // This line is infered from column name "block_id".

}
