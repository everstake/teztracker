package repos

import (
	"github.com/bullblock-io/tezTracker/repos/account"
	"github.com/bullblock-io/tezTracker/repos/baker"
	"github.com/bullblock-io/tezTracker/repos/block"
	"github.com/bullblock-io/tezTracker/repos/operation"
	"github.com/bullblock-io/tezTracker/repos/operation_groups"
	"github.com/jinzhu/gorm"
)

// Provider is the repository provider.
type Provider struct {
	db *gorm.DB
}

// New creates a new instance of Provider with the underlying DB instance.
func New(db *gorm.DB) *Provider {
	return &Provider{
		db: db,
	}
}

// GetBlock returns a new block repository.
func (u *Provider) GetBlock() block.Repo {
	return block.New(u.db)
}

// GetOperationGroup returns a new operation group repository.
func (u *Provider) GetOperationGroup() operation_groups.Repo {
	return operation_groups.New(u.db)
}

// GetOperation returns a new operation repository.
func (u *Provider) GetOperation() operation.Repo {
	return operation.New(u.db)
}

// GetAccount returns a new account repository.
func (u *Provider) GetAccount() account.Repo {
	return account.New(u.db)
}

// GetBaker returns a new baker repository.
func (u *Provider) GetBaker() baker.Repo {
	return baker.New(u.db)
}
