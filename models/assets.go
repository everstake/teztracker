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
	TokenId            uint64
	OperationId        int64
	OperationGroupHash string
	Sender             string
	Receiver           string
	Amount             int64
	Type               string
	Timestamp          time.Time
}

type AssetOperationReport struct {
	AssetOperation
	Fee          int64
	GasLimit     int64
	StorageLimit int64
}
