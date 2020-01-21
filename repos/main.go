package repos

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/bullblock-io/tezTracker/repos/account"
	"github.com/bullblock-io/tezTracker/repos/baker"
	"github.com/bullblock-io/tezTracker/repos/block"
	"github.com/bullblock-io/tezTracker/repos/double_baking"
	"github.com/bullblock-io/tezTracker/repos/future_baking_rights"
	"github.com/bullblock-io/tezTracker/repos/operation"
	"github.com/bullblock-io/tezTracker/repos/operation_counter"
	"github.com/bullblock-io/tezTracker/repos/operation_groups"
	"github.com/bullblock-io/tezTracker/repos/snapshots"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// Provider is the repository provider.
type Provider struct {
	db *gorm.DB
	tx *gorm.DB
}

// New creates a new instance of Provider with the underlying DB instance.
func New(db *gorm.DB) *Provider {
	return &Provider{
		db: db,
	}
}
func (u *Provider) getDB() *gorm.DB {
	if u.tx != nil {
		return u.tx
	}
	return u.db
}

// GetBlock returns a new block repository.
func (u *Provider) GetBlock() block.Repo {
	return block.New(u.getDB())
}

// GetOperationGroup returns a new operation group repository.
func (u *Provider) GetOperationGroup() operation_groups.Repo {
	return operation_groups.New(u.getDB())
}

// GetOperation returns a new operation repository.
func (u *Provider) GetOperation() operation.Repo {
	return operation.New(u.getDB())
}

// GetAccount returns a new account repository.
func (u *Provider) GetAccount() account.Repo {
	return account.New(u.getDB())
}

// GetBaker returns a new baker repository.
func (u *Provider) GetBaker() baker.Repo {
	return baker.New(u.getDB())
}

func (u *Provider) GetOperationCounter() operation_counter.Repo {
	return operation_counter.New(u.getDB())
}

func (u *Provider) GetFutureBakingRight() future_baking_rights.Repo {
	return future_baking_rights.New(u.getDB())
}

func (u *Provider) GetSnapshots() snapshots.Repo {
	return snapshots.New(u.getDB())
}

func (u *Provider) GetDoubleBaking() double_baking.Repo {
	return double_baking.New(u.getDB())
}

func (u *Provider) Start(ctx context.Context) {
	u.tx = u.db.BeginTx(ctx, &sql.TxOptions{})
}

func (u *Provider) RollbackUnlessCommitted() {
	if u.tx != nil {
		if err := u.tx.RollbackUnlessCommitted().Error; err != nil {
			logrus.Printf("error on rollback: %s", err.Error())
		}
		u.tx = nil
	}
}

func (u *Provider) Commit() error {
	if u.tx == nil {
		return fmt.Errorf("tx is empty")
	}
	if err := u.tx.Commit().Error; err != nil {
		u.RollbackUnlessCommitted()
		return err
	}
	u.tx = nil
	return nil
}
