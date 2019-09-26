package services

import (
	"github.com/bullblock-io/tezTracker/repos/account"
	"github.com/bullblock-io/tezTracker/repos/baker"
	"github.com/bullblock-io/tezTracker/repos/block"
	"github.com/bullblock-io/tezTracker/repos/operation"
	"github.com/bullblock-io/tezTracker/repos/operation_groups"
)

//go:generate mockgen -source ./main.go -destination ./mock_service/main.go Provider
type (
	// TezTracker is the main service for tezos tracker. It has methods to process all the user's requests.
	TezTracker struct {
		repoProvider Provider
	}

	// Provider is the abstract interface to get any repository.
	Provider interface {
		GetBlock() block.Repo
		GetOperationGroup() operation_groups.Repo
		GetOperation() operation.Repo
		GetAccount() account.Repo
		GetBaker() baker.Repo
	}

	Limiter interface {
		Limit() uint
		Offset() uint
	}
)

// New creates a new TexTracker service using the repository provider.
func New(rp Provider) *TezTracker {
	return &TezTracker{repoProvider: rp}
}
