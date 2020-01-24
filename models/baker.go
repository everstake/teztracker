package models

type Baker struct {
	AccountID      string `json:"pkh"`
	StakingBalance int64  `json:"staking_balance"`
	Blocks         int64  `json:"blocks"`
	Endorsements   int64  `json:"endorsements"`
	Fees           int64  `json:"fees"`
}

type BakerInfo struct {
	Delegate
	BakingDeposits      int64
	EndorsementDeposits int64
	BakingRewards       int64
	EndorsementRewards  int64
}
