package services

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bullblock-io/tezTracker/models"
	"github.com/bullblock-io/tezTracker/repos/account/mock_account"
	mock_services "github.com/bullblock-io/tezTracker/services/mock_service"
	"github.com/golang/mock/gomock"
	"github.com/guregu/null"
)

func TestTezTracker_GetAccount(t *testing.T) {
	type mockRet struct {
		found bool
		acc   models.Account
		err   error
	}
	validAcc := models.Account{
		AccountID:       null.StringFrom("1"),
		BlockID:         null.StringFrom("2"),
		Manager:         null.StringFrom("3"),
		Spendable:       null.BoolFrom(true),
		DelegateSetable: null.BoolFrom(true),
		DelegateValue:   "a.DelegateValue",
		Counter:         null.IntFrom(10),
		Script:          "a.Script",
		Storage:         "a.Storage",
		Balance:         null.IntFrom(13),
		BlockLevel:      null.IntFrom(16),
	}
	tests := []struct {
		name    string
		id      string
		ret     mockRet
		wantAcc models.Account
		wantErr bool
	}{
		{
			id:      "1",
			ret:     mockRet{found: false, err: nil},
			wantErr: true,
		},
		{
			id:      "2",
			ret:     mockRet{err: fmt.Errorf("")},
			wantErr: true,
		},
		{
			id:      "3",
			ret:     mockRet{found: true, acc: validAcc},
			wantAcc: validAcc,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			prov := mock_services.NewMockProvider(ctrl)
			repo := mock_account.NewMockRepo(ctrl)

			tracker := &TezTracker{
				repoProvider: prov,
			}
			prov.EXPECT().GetAccount().Return(repo)
			repo.EXPECT().Find(models.Account{AccountID: null.StringFrom(tt.id)}).Return(tt.ret.found, tt.ret.acc, tt.ret.err)

			gotAcc, err := tracker.GetAccount(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("TezTracker.GetAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotAcc, tt.wantAcc) {
				t.Errorf("TezTracker.GetAccount() = %v, want %v", gotAcc, tt.wantAcc)
			}
		})
	}
}
