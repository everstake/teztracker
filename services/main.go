package services

import (
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/repos/account"
	"github.com/everstake/teztracker/repos/baker"
	"github.com/everstake/teztracker/repos/block"
	"github.com/everstake/teztracker/repos/double_baking"
	"github.com/everstake/teztracker/repos/future_baking_rights"
	"github.com/everstake/teztracker/repos/operation"
	"github.com/everstake/teztracker/repos/operation_groups"
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
		GetFutureBakingRight() future_baking_rights.Repo
		GetSnapshots() snapshots.Repo
		GetDoubleBaking() double_baking.Repo
		GetVotingPeriod() voting_periods.Repo
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

const BlocksInMainnetCycle = 4096

func (t *TezTracker) BlocksInCycle() int64 {
	if t.net == models.NetworkMain {
		return BlocksInMainnetCycle
	}
	return BlocksInMainnetCycle / 2
}
