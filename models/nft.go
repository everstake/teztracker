package models

import "time"

type NFTContract struct {
	ID               int64  `gorm:"column:id"`
	Name             string `gorm:"column:name"`
	ContractType     string `gorm:"column:contract_type"`
	AccountId        string `gorm:"column:account_id"`
	SwapContract     string `gorm:"column:swap_contract"`
	Description      string `gorm:"column:description"`
	LedgerBigMap     int64  `gorm:"column:ledger_big_map"`
	TokensBigMap     int64  `gorm:"column:tokens_big_map"`
	OperationsNum    int64  `gorm:"column:operations_num"`
	LastHeight       int64  `gorm:"column:last_height"`
	LastUpdateHeight int64  `gorm:"column:last_update_height"`
	NFTsNumber       int64  `gorm:"column:nfts_number"`
}

type NFTToken struct {
	ContractID   int64     `gorm:"column:contract_id"`
	ID           uint64    `gorm:"column:token_id"`
	Name         string    `gorm:"column:name"`
	Category     string    `gorm:"column:category"`
	Description  string    `gorm:"column:description"`
	Decimals     int64     `gorm:"column:decimals"`
	Amount       int64     `gorm:"column:amount"`
	LastPrice    int64     `gorm:"column:last_price"`
	IssuedBy     string    `gorm:"column:issued_by"`
	IsForSale    bool      `gorm:"column:is_for_sale"`
	IpfsSource   string    `gorm:"column:ipfs_source"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	LastActiveAt time.Time `gorm:"column:last_active_at"`
}

type NFTDistribution struct {
	Holders          []AssetHolder
	TokenNum         int64
	UniqueHoldersNum int64
}

type NFTOperations struct {
	BlockLevel         int64
	TokenID            int64
	OperationId        int64
	OperationGroupHash string
	Sender             string
	Receiver           string
	Amount             int64
}

type NFTContractOwnership struct {
	UniqueHoldersNum   int64
	SingleTokenHolders int64
	MultiTokenHolders  int64
	WhaleTokenHolders  int64
}
