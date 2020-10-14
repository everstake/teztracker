package services

import (
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/repos/account"
	"github.com/everstake/teztracker/repos/assets"
	"github.com/everstake/teztracker/repos/baker"
	"github.com/everstake/teztracker/repos/baking"
	"github.com/everstake/teztracker/repos/block"
	"github.com/everstake/teztracker/repos/chart"
	"github.com/everstake/teztracker/repos/double_baking"
	"github.com/everstake/teztracker/repos/double_endorsement"
	"github.com/everstake/teztracker/repos/endorsing"
	"github.com/everstake/teztracker/repos/future_baking_rights"
	"github.com/everstake/teztracker/repos/future_endorsement_rights"
	"github.com/everstake/teztracker/repos/operation"
	"github.com/everstake/teztracker/repos/operation_groups"
	"github.com/everstake/teztracker/repos/rolls"
	"github.com/everstake/teztracker/repos/snapshots"
	"github.com/everstake/teztracker/repos/voting_periods"
)

//go:generate mockgen -source ./main.go -destination ./mock_service/main.go Provider
type (
	// TezTracker is the main service for tezos tracker. It has methods to process all the user's requests.
	TezTracker struct {
		repoProvider Provider
		net          models.Network
	}

	// Provider is the abstract interface to get any repository.
	Provider interface {
		GetBlock() block.Repo
		GetOperationGroup() operation_groups.Repo
		GetOperation() operation.Repo
		GetAccount() account.Repo
		GetBaker() baker.Repo
		GetBaking() baking.Repo
		GetEndorsing() endorsing.Repo
		GetFutureBakingRight() future_baking_rights.Repo
		GetFutureEndorsementRight() future_endorsement_rights.Repo
		GetSnapshots() snapshots.Repo
		GetRolls() rolls.Repo
		GetDoubleBaking() double_baking.Repo
		GetDoubleEndorsement() double_endorsement.Repo
		GetVotingPeriod() voting_periods.Repo
		GetChart() chart.Repo
		GetAssets() assets.Repo
	}

	Limiter interface {
		Limit() uint
		Offset() uint
	}
)

// New creates a new TexTracker service using the repository provider.
func New(rp Provider, net models.Network) *TezTracker {
	return &TezTracker{repoProvider: rp, net: net}
}

const (
	BlocksInMainnetCycle = 4096
)

func (t *TezTracker) BlocksInCycle() int64 {
	if t.net == models.NetworkMain {
		return BlocksInMainnetCycle
	}
	return BlocksInMainnetCycle / 2
}
