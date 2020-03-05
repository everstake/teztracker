package models

import (
	"encoding/hex"
	"encoding/json"
	"github.com/lib/pq"
)

type Baker struct {
	AccountID string `json:"pkh"`
	Name      string `json:"name"`
	BakerStats
}

type BakerStats struct {
	Balance           int64 `json:"balance"`
	StakingBalance    int64 `json:"staking_balance"`
	Blocks            int64 `json:"blocks"`
	Endorsements      int64 `json:"endorsements"`
	Fees              int64 `json:"fees"`
	FirstBlock        int64 `json:"first_block"`
	ActiveDelegations int64 `json:"active_delegations"`
	StakingCapacity   int64 `json:"staking_capacity"`
}

type PublicBaker struct {
	Delegate                               string         `gorm:"primary_key" json:"delegate"`
	BakerName                              string         `json:"bakerName"`
	BakerOffchainRegistryUrl               string         `json:"bakerOffchainRegistryUrl"`
	BakerPaysFromAccounts                  pq.StringArray `gorm:"type:varchar[]" json:"bakerPaysFromAccounts"`
	BakerChargesTransactionFee             bool           `json:"bakerChargesTransactionFee"`
	MinDelegation                          int64          `json:"minDelegation,string"`
	MinPayout                              int64          `json:"minPayout,string"`
	OverDelegationThreshold                int64          `json:"overDelegationThreshold,string"`
	PayoutDelay                            int64          `json:"payoutDelay,string"`
	PaymentConfigMask                      string         `json:"paymentConfigMask"`
	PayoutFrequency                        int64          `json:"payoutFrequency,string"`
	ReporterAccount                        pq.StringArray `gorm:"type:varchar[]" json:"reporterAccount"`
	Split                                  int64          `json:"split,string""`
	OpenForDelegation                      bool           `json:"openForDelegation"`
	SubtractPayoutsLessThanMin             bool           `json:"subtractPayoutsLessThanMin"`
	SubtractRewardsFromUninvitedDelegation bool           `json:"subtractRewardsFromUninvitedDelegation"`
	LastUpdateId                           int64          `json:"-"`
}

func (pb *PublicBaker) Unmarshal(data []byte) (err error) {

	err = json.Unmarshal(data, &pb)
	if err != nil {
		return err
	}

	bytes, err := hex.DecodeString(pb.BakerName)
	if err != nil {
		return err
	}
	pb.BakerName = string(bytes)

	bytes, err = hex.DecodeString(pb.BakerOffchainRegistryUrl)
	if err != nil {
		return err
	}

	pb.BakerOffchainRegistryUrl = string(bytes)

	return nil
}

type BakerInfo struct {
	Delegate
	BakingDeposits      int64
	EndorsementDeposits int64
	BakingRewards       int64
	EndorsementRewards  int64
}
