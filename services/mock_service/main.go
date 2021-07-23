// Code generated by MockGen. DO NOT EDIT.
// Source: ./main.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	account "github.com/everstake/teztracker/repos/account"
	assets "github.com/everstake/teztracker/repos/assets"
	baker "github.com/everstake/teztracker/repos/baker"
	baking "github.com/everstake/teztracker/repos/baking"
	block "github.com/everstake/teztracker/repos/block"
	chart "github.com/everstake/teztracker/repos/chart"
	double_baking "github.com/everstake/teztracker/repos/double_baking"
	double_endorsement "github.com/everstake/teztracker/repos/double_endorsement"
	endorsing "github.com/everstake/teztracker/repos/endorsing"
	future_baking_rights "github.com/everstake/teztracker/repos/future_baking_rights"
	future_endorsement_rights "github.com/everstake/teztracker/repos/future_endorsement_rights"
	operation "github.com/everstake/teztracker/repos/operation"
	operation_groups "github.com/everstake/teztracker/repos/operation_groups"
	rolls "github.com/everstake/teztracker/repos/rolls"
	snapshots "github.com/everstake/teztracker/repos/snapshots"
	thirdparty_bakers "github.com/everstake/teztracker/repos/thirdparty_bakers"
	user_profile "github.com/everstake/teztracker/repos/user_profile"
	voting_periods "github.com/everstake/teztracker/repos/voting_periods"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockProvider is a mock of Provider interface
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// Health mocks base method
func (m *MockProvider) Health() error {
	ret := m.ctrl.Call(m, "Health")
	ret0, _ := ret[0].(error)
	return ret0
}

// Health indicates an expected call of Health
func (mr *MockProviderMockRecorder) Health() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Health", reflect.TypeOf((*MockProvider)(nil).Health))
}

// GetBlock mocks base method
func (m *MockProvider) GetBlock() block.Repo {
	ret := m.ctrl.Call(m, "GetBlock")
	ret0, _ := ret[0].(block.Repo)
	return ret0
}

// GetBlock indicates an expected call of GetBlock
func (mr *MockProviderMockRecorder) GetBlock() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlock", reflect.TypeOf((*MockProvider)(nil).GetBlock))
}

// GetOperationGroup mocks base method
func (m *MockProvider) GetOperationGroup() operation_groups.Repo {
	ret := m.ctrl.Call(m, "GetOperationGroup")
	ret0, _ := ret[0].(operation_groups.Repo)
	return ret0
}

// GetOperationGroup indicates an expected call of GetOperationGroup
func (mr *MockProviderMockRecorder) GetOperationGroup() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperationGroup", reflect.TypeOf((*MockProvider)(nil).GetOperationGroup))
}

// GetOperation mocks base method
func (m *MockProvider) GetOperation() operation.Repo {
	ret := m.ctrl.Call(m, "GetOperation")
	ret0, _ := ret[0].(operation.Repo)
	return ret0
}

// GetOperation indicates an expected call of GetOperation
func (mr *MockProviderMockRecorder) GetOperation() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperation", reflect.TypeOf((*MockProvider)(nil).GetOperation))
}

// GetAccount mocks base method
func (m *MockProvider) GetAccount() account.Repo {
	ret := m.ctrl.Call(m, "GetAccount")
	ret0, _ := ret[0].(account.Repo)
	return ret0
}

// GetAccount indicates an expected call of GetAccount
func (mr *MockProviderMockRecorder) GetAccount() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockProvider)(nil).GetAccount))
}

// GetBaker mocks base method
func (m *MockProvider) GetBaker() baker.Repo {
	ret := m.ctrl.Call(m, "GetBaker")
	ret0, _ := ret[0].(baker.Repo)
	return ret0
}

// GetBaker indicates an expected call of GetBaker
func (mr *MockProviderMockRecorder) GetBaker() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBaker", reflect.TypeOf((*MockProvider)(nil).GetBaker))
}

// GetBaking mocks base method
func (m *MockProvider) GetBaking() baking.Repo {
	ret := m.ctrl.Call(m, "GetBaking")
	ret0, _ := ret[0].(baking.Repo)
	return ret0
}

// GetBaking indicates an expected call of GetBaking
func (mr *MockProviderMockRecorder) GetBaking() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBaking", reflect.TypeOf((*MockProvider)(nil).GetBaking))
}

// GetEndorsing mocks base method
func (m *MockProvider) GetEndorsing() endorsing.Repo {
	ret := m.ctrl.Call(m, "GetEndorsing")
	ret0, _ := ret[0].(endorsing.Repo)
	return ret0
}

// GetEndorsing indicates an expected call of GetEndorsing
func (mr *MockProviderMockRecorder) GetEndorsing() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEndorsing", reflect.TypeOf((*MockProvider)(nil).GetEndorsing))
}

// GetFutureBakingRight mocks base method
func (m *MockProvider) GetFutureBakingRight() future_baking_rights.Repo {
	ret := m.ctrl.Call(m, "GetFutureBakingRight")
	ret0, _ := ret[0].(future_baking_rights.Repo)
	return ret0
}

// GetFutureBakingRight indicates an expected call of GetFutureBakingRight
func (mr *MockProviderMockRecorder) GetFutureBakingRight() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFutureBakingRight", reflect.TypeOf((*MockProvider)(nil).GetFutureBakingRight))
}

// GetFutureEndorsementRight mocks base method
func (m *MockProvider) GetFutureEndorsementRight() future_endorsement_rights.Repo {
	ret := m.ctrl.Call(m, "GetFutureEndorsementRight")
	ret0, _ := ret[0].(future_endorsement_rights.Repo)
	return ret0
}

// GetFutureEndorsementRight indicates an expected call of GetFutureEndorsementRight
func (mr *MockProviderMockRecorder) GetFutureEndorsementRight() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFutureEndorsementRight", reflect.TypeOf((*MockProvider)(nil).GetFutureEndorsementRight))
}

// GetSnapshots mocks base method
func (m *MockProvider) GetSnapshots() snapshots.Repo {
	ret := m.ctrl.Call(m, "GetSnapshots")
	ret0, _ := ret[0].(snapshots.Repo)
	return ret0
}

// GetSnapshots indicates an expected call of GetSnapshots
func (mr *MockProviderMockRecorder) GetSnapshots() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSnapshots", reflect.TypeOf((*MockProvider)(nil).GetSnapshots))
}

// GetRolls mocks base method
func (m *MockProvider) GetRolls() rolls.Repo {
	ret := m.ctrl.Call(m, "GetRolls")
	ret0, _ := ret[0].(rolls.Repo)
	return ret0
}

// GetRolls indicates an expected call of GetRolls
func (mr *MockProviderMockRecorder) GetRolls() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRolls", reflect.TypeOf((*MockProvider)(nil).GetRolls))
}

// GetDoubleBaking mocks base method
func (m *MockProvider) GetDoubleBaking() double_baking.Repo {
	ret := m.ctrl.Call(m, "GetDoubleBaking")
	ret0, _ := ret[0].(double_baking.Repo)
	return ret0
}

// GetDoubleBaking indicates an expected call of GetDoubleBaking
func (mr *MockProviderMockRecorder) GetDoubleBaking() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDoubleBaking", reflect.TypeOf((*MockProvider)(nil).GetDoubleBaking))
}

// GetDoubleEndorsement mocks base method
func (m *MockProvider) GetDoubleEndorsement() double_endorsement.Repo {
	ret := m.ctrl.Call(m, "GetDoubleEndorsement")
	ret0, _ := ret[0].(double_endorsement.Repo)
	return ret0
}

// GetDoubleEndorsement indicates an expected call of GetDoubleEndorsement
func (mr *MockProviderMockRecorder) GetDoubleEndorsement() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDoubleEndorsement", reflect.TypeOf((*MockProvider)(nil).GetDoubleEndorsement))
}

// GetVotingPeriod mocks base method
func (m *MockProvider) GetVotingPeriod() voting_periods.Repo {
	ret := m.ctrl.Call(m, "GetVotingPeriod")
	ret0, _ := ret[0].(voting_periods.Repo)
	return ret0
}

// GetVotingPeriod indicates an expected call of GetVotingPeriod
func (mr *MockProviderMockRecorder) GetVotingPeriod() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVotingPeriod", reflect.TypeOf((*MockProvider)(nil).GetVotingPeriod))
}

// GetChart mocks base method
func (m *MockProvider) GetChart() chart.Repo {
	ret := m.ctrl.Call(m, "GetChart")
	ret0, _ := ret[0].(chart.Repo)
	return ret0
}

// GetChart indicates an expected call of GetChart
func (mr *MockProviderMockRecorder) GetChart() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChart", reflect.TypeOf((*MockProvider)(nil).GetChart))
}

// GetAssets mocks base method
func (m *MockProvider) GetAssets() assets.Repo {
	ret := m.ctrl.Call(m, "GetAssets")
	ret0, _ := ret[0].(assets.Repo)
	return ret0
}

// GetAssets indicates an expected call of GetAssets
func (mr *MockProviderMockRecorder) GetAssets() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAssets", reflect.TypeOf((*MockProvider)(nil).GetAssets))
}

// GetThirdPartyBakers mocks base method
func (m *MockProvider) GetThirdPartyBakers() thirdparty_bakers.Repo {
	ret := m.ctrl.Call(m, "GetThirdPartyBakers")
	ret0, _ := ret[0].(thirdparty_bakers.Repo)
	return ret0
}

// GetThirdPartyBakers indicates an expected call of GetThirdPartyBakers
func (mr *MockProviderMockRecorder) GetThirdPartyBakers() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThirdPartyBakers", reflect.TypeOf((*MockProvider)(nil).GetThirdPartyBakers))
}

// GetUserProfile mocks base method
func (m *MockProvider) GetUserProfile() user_profile.Repo {
	ret := m.ctrl.Call(m, "GetUserProfile")
	ret0, _ := ret[0].(user_profile.Repo)
	return ret0
}

// GetUserProfile indicates an expected call of GetUserProfile
func (mr *MockProviderMockRecorder) GetUserProfile() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserProfile", reflect.TypeOf((*MockProvider)(nil).GetUserProfile))
}

// MockLimiter is a mock of Limiter interface
type MockLimiter struct {
	ctrl     *gomock.Controller
	recorder *MockLimiterMockRecorder
}

// MockLimiterMockRecorder is the mock recorder for MockLimiter
type MockLimiterMockRecorder struct {
	mock *MockLimiter
}

// NewMockLimiter creates a new mock instance
func NewMockLimiter(ctrl *gomock.Controller) *MockLimiter {
	mock := &MockLimiter{ctrl: ctrl}
	mock.recorder = &MockLimiterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLimiter) EXPECT() *MockLimiterMockRecorder {
	return m.recorder
}

// Limit mocks base method
func (m *MockLimiter) Limit() uint {
	ret := m.ctrl.Call(m, "Limit")
	ret0, _ := ret[0].(uint)
	return ret0
}

// Limit indicates an expected call of Limit
func (mr *MockLimiterMockRecorder) Limit() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Limit", reflect.TypeOf((*MockLimiter)(nil).Limit))
}

// Offset mocks base method
func (m *MockLimiter) Offset() uint {
	ret := m.ctrl.Call(m, "Offset")
	ret0, _ := ret[0].(uint)
	return ret0
}

// Offset indicates an expected call of Offset
func (mr *MockLimiterMockRecorder) Offset() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Offset", reflect.TypeOf((*MockLimiter)(nil).Offset))
}
