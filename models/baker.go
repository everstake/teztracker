package models

import (
	"encoding/hex"
	"encoding/json"
	"github.com/lib/pq"
	"time"
)

type Baker struct {
	AccountID string `json:"pkh"`
	BakerStats
}

type PublicBakerSearch struct {
	Delegate  string
	BakerName string
}

type BakerBalance struct {
	Pkh              string
	Balance          int64
	FrozenBalance    uint64
	StakingBalance   uint64
	DelegatedBalance uint64
}

type BakerStats struct {
	Name                     string    `json:"name"`
	Fee                      int64     `json:"fee"`
	BakingSince              time.Time `json:"baking_since"` //first baking or endorsement
	Balance                  int64     `json:"balance"`
	StakingBalance           int64     `json:"staking_balance"`
	FrozenBalance            int64     `json:"frozen_balance"`
	Rolls                    int64     `json:"rolls"`
	Blocks                   int64     `json:"blocks"`
	Endorsements             int64     `json:"endorsements"`
	TotalPaidFees            int64     `json:"fees"`
	ActiveDelegations        int64     `json:"active_delegations"`
	StakingCapacity          int64     `json:"staking_capacity"`
	FrozenEndorsementRewards int64     `json:"frozen_endorsement_rewards"`
	FrozenBakingRewards      int64     `json:"frozen_baking_rewards"`
	EndorsementCount         int64     `json:"endorsement_count"`
	BakingCount              int64     `json:"baking_count"`

	//	From old resp
	BakingDeposits      int64
	EndorsementDeposits int64
	BakingRewards       int64
	EndorsementRewards  int64
}

type BakerRegistry struct {
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
	IsHidden                               bool           `json:"is_hidden"`
}

func (pb *BakerRegistry) Unmarshal(data []byte) (err error) {

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
