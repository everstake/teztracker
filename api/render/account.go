package render

import (
	genModels "github.com/everstake/teztracker/gen/models"
	"github.com/everstake/teztracker/models"
)

// Account renders an app level model to a gennerated OpenAPI model.
func Account(a models.Account) *genModels.AccountsRow {
	return &genModels.AccountsRow{
		AccountID:       a.AccountID.Ptr(),
		AccountName:     a.AccountName,
		BlockID:         a.BlockID.Ptr(),
		Manager:         a.Manager.Ptr(),
		Spendable:       a.Spendable.Ptr(),
		DelegateSetable: a.DelegateSetable.Ptr(),
		DelegateValue:   a.DelegateValue,
		DelegateName:    a.DelegateName,
		Counter:         a.Counter.Ptr(),
		Script:          a.Script,
		Storage:         a.Storage,
		Balance:         a.Balance.Ptr(),
		BlockLevel:      a.BlockLevel.Ptr(),
		IsBaker:         &a.IsBaker,
		BakerInfo:       BakerInfo(a.BakerInfo),
		CreatedAt:       a.CreatedAt.Unix(),
		LastActive:      a.LastActive.Unix(),
		Transactions:    a.Transactions,
		Operations:      a.Operations,
		Revealed:        &a.IsRevealed,
	}
}

// Accounts renders a slice of app level Accounts into a slice of OpenAPI models.
func Accounts(ams []models.Account) []*genModels.AccountsRow {
	accs := make([]*genModels.AccountsRow, len(ams))
	for i := range ams {
		accs[i] = Account(ams[i])
	}
	return accs
}

func AccountBalances(acb []models.AccountBalance) []*genModels.AccountBalance {
	accs := make([]*genModels.AccountBalance, len(acb))
	for i := range acb {
		accs[i] = AccountBalance(acb[i])
	}
	return accs
}

func AccountBalance(acb models.AccountBalance) *genModels.AccountBalance {
	tm := acb.Time.Unix()
	return &genModels.AccountBalance{
		Balance:   &acb.Balance,
		Timestamp: &tm,
	}
}

func AccountRewardsList(accrl []models.AccountReward) []*genModels.AccountRewardsRow {
	accrr := make([]*genModels.AccountRewardsRow, len(accrl))
	for i := range accrl {
		accrr[i] = AccountReward(accrl[i])
	}
	return accrr
}

func AccountReward(acb models.AccountReward) *genModels.AccountRewardsRow {
	return &genModels.AccountRewardsRow{
		Cycle:          &acb.Cycle,
		Status:         string(acb.Status),
		Delegators:     &acb.Delegators,
		Baking:         &acb.BakingRewards,
		Fees:           &acb.Fees,
		StakingBalance: &acb.StakingBalance,
		Endorsements:   &acb.EndorsementRewards,
		Losses:         &acb.Losses,
	}
}

func AccountDelegators(acd []models.AccountDelegator) []*genModels.BakerDelegator {
	accd := make([]*genModels.BakerDelegator, len(acd))
	for i := range acd {
		accd[i] = AccountDelegator(acd[i])
	}
	return accd
}

func AccountDelegator(acb models.AccountDelegator) *genModels.BakerDelegator {
	return &genModels.BakerDelegator{
		Balance:   &acb.Balance,
		Cycle:     &acb.Cycle,
		Delegator: &acb.AccountId,
		Share:     &acb.Share,
	}
}
