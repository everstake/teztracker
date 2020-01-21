package models

import "time"

type AccountsCheckpoint struct {
	AccountID  string    `json:"account_id"`
	Account    *Account  `json:"account"` // This line is infered from column name "account_id".
	BlockID    string    `json:"block_id"`
	Block      *Block    `json:"block"` // This line is infered from column name "block_id".
	BlockLevel uint      `json:"block_level" sql:"DEFAULT:'-1'::integer"`
	AsOf       time.Time `json:"asof" gorm:"column:asof"`
	IsBaker    bool      `json:"is_baker"`
}
