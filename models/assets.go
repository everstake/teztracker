package models

import "time"

type AssetInfo struct {
	ID           uint64
	Name         string
	Balance      int64
	Source       string
	ContractType string
	AccountId    string
	Timestamp    time.Time
	Scale        int64
}

type AssetHolder struct {
	AccountID string
	Balance   int64
}

type AssetOperation struct {
	TokenId            uint64    `json:"token_id"`
	OperationId        int64     `json:"operation_id"`
	OperationGroupHash string    `json:"operation_group_hash"`
	Sender             string    `json:"sender"`
	Receiver           string    `json:"receiver"`
	Amount             int64     `json:"amount"`
	Type               string    `json:"type"`
	Timestamp          time.Time `json:"timestamp"`
}

type AssetOperationReport struct {
	AssetOperation
	Fee          int64
	GasLimit     int64
	StorageLimit int64
}

type RegisteredToken struct {
	ID           uint64
	Name         string
	ContractType string
	AccountId    string
	Scale        uint64
	Ticker       string
}
