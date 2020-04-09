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
		Counter:         a.Counter.Ptr(),
		Script:          a.Script,
		Storage:         a.Storage,
		Balance:         a.Balance.Ptr(),
		BlockLevel:      a.BlockLevel.Ptr(),
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
	return &genModels.AccountBalance{
		Balance:   acb.Balance,
		Timestamp: acb.Time.Unix(),
	}
}

func AccountBaking(acb models.AccountBaking) *genModels.AccountBakingRow {
	return &genModels.AccountBakingRow{
		AvgPriority: &acb.AvgPriority,
		Blocks:      &acb.Count,
		Cycle:       &acb.Cycle,
		Missed:      &acb.Missed,
		Rewards:     &acb.Reward,
		Stolen:      &acb.Stolen,
	}
}

func AccountBakingList(accb []models.AccountBaking) []*genModels.AccountBakingRow {
	accbs := make([]*genModels.AccountBakingRow, len(accb))
	for i := range accb {
		accbs[i] = AccountBaking(accb[i])
	}
	return accbs
}

func AccountEndorsing(acb models.AccountEndorsing) *genModels.AccountEndorsingRow {
	return &genModels.AccountEndorsingRow{
		Slots:   &acb.Count,
		Cycle:   &acb.Cycle,
		Missed:  &acb.Missed,
		Rewards: &acb.Reward,
	}
}

func AccountEndorsingList(acce []models.AccountEndorsing) []*genModels.AccountEndorsingRow {
	accbs := make([]*genModels.AccountEndorsingRow, len(acce))
	for i := range acce {
		accbs[i] = AccountEndorsing(acce[i])
	}
	return accbs
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
		Delegators:     &acb.Delegators,
		Baking:         &acb.BakingRewards,
		StakingBalance: &acb.StakingBalance,
		Endorsements:   &acb.EndorsementRewards,
		Losses:         &acb.Losses,
	}
}
