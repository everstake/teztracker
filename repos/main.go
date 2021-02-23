package repos

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/everstake/teztracker/repos/assets"
	"github.com/everstake/teztracker/repos/baking"
	"github.com/everstake/teztracker/repos/chart"
	"github.com/everstake/teztracker/repos/double_endorsement"
	"github.com/everstake/teztracker/repos/endorsing"
	"github.com/everstake/teztracker/repos/future_endorsement_rights"
	"github.com/everstake/teztracker/repos/rolls"
	"github.com/everstake/teztracker/repos/storage"
	"github.com/everstake/teztracker/repos/thirdparty_bakers"
	"github.com/everstake/teztracker/repos/user_profile"
	"github.com/everstake/teztracker/repos/voting_periods"

	"github.com/everstake/teztracker/repos/account"
	"github.com/everstake/teztracker/repos/baker"
	"github.com/everstake/teztracker/repos/block"
	"github.com/everstake/teztracker/repos/double_baking"
	"github.com/everstake/teztracker/repos/future_baking_rights"
	"github.com/everstake/teztracker/repos/operation"
	"github.com/everstake/teztracker/repos/operation_counter"
	"github.com/everstake/teztracker/repos/operation_groups"
	"github.com/everstake/teztracker/repos/snapshots"
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

//Heath returns a new health check of repository provider.
func (u *Provider) Health() (err error) {
	err = u.db.DB().Ping()
	if err != nil {
		return err
	}

	return nil
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

func (u *Provider) GetBaking() baking.Repo {
	return baking.New(u.getDB())
}

func (u *Provider) GetEndorsing() endorsing.Repo {
	return endorsing.New(u.getDB())
}

func (u *Provider) GetOperationCounter() operation_counter.Repo {
	return operation_counter.New(u.getDB())
}

func (u *Provider) GetFutureBakingRight() future_baking_rights.Repo {
	return future_baking_rights.New(u.getDB())
}

func (u *Provider) GetFutureEndorsementRight() future_endorsement_rights.Repo {
	return future_endorsement_rights.New(u.getDB())
}

func (u *Provider) GetSnapshots() snapshots.Repo {
	return snapshots.New(u.getDB())
}

func (u *Provider) GetRolls() rolls.Repo {
	return rolls.New(u.getDB())
}

func (u *Provider) GetDoubleBaking() double_baking.Repo {
	return double_baking.New(u.getDB())
}

func (u *Provider) GetDoubleEndorsement() double_endorsement.Repo {
	return double_endorsement.New(u.getDB())
}

func (u *Provider) GetVotingPeriod() voting_periods.Repo {
	return voting_periods.New(u.getDB())
}

func (u *Provider) GetChart() chart.Repo {
	return chart.New(u.getDB())
}

func (u *Provider) GetAssets() assets.Repo {
	return assets.New(u.getDB())
}

func (u *Provider) GetThirdPartyBakers() thirdparty_bakers.Repo {
	return thirdparty_bakers.New(u.getDB())
}

func (u *Provider) GetUserProfile() user_profile.Repo {
	return user_profile.New(u.getDB())
}

func (u *Provider) GetStorage() storage.Repo {
	return storage.New(u.getDB())
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
