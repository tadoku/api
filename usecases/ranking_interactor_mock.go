// Code generated by MockGen. DO NOT EDIT.
// Source: ranking_interactor.go

// Package usecases is a generated GoMock package.
package usecases

import (
	gomock "github.com/golang/mock/gomock"
	domain "github.com/tadoku/api/domain"
	reflect "reflect"
)

// MockRankingInteractor is a mock of RankingInteractor interface
type MockRankingInteractor struct {
	ctrl     *gomock.Controller
	recorder *MockRankingInteractorMockRecorder
}

// MockRankingInteractorMockRecorder is the mock recorder for MockRankingInteractor
type MockRankingInteractorMockRecorder struct {
	mock *MockRankingInteractor
}

// NewMockRankingInteractor creates a new mock instance
func NewMockRankingInteractor(ctrl *gomock.Controller) *MockRankingInteractor {
	mock := &MockRankingInteractor{ctrl: ctrl}
	mock.recorder = &MockRankingInteractorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRankingInteractor) EXPECT() *MockRankingInteractorMockRecorder {
	return m.recorder
}

// CreateRanking mocks base method
func (m *MockRankingInteractor) CreateRanking(userID, contestID uint64, languages domain.LanguageCodes) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRanking", userID, contestID, languages)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRanking indicates an expected call of CreateRanking
func (mr *MockRankingInteractorMockRecorder) CreateRanking(userID, contestID, languages interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRanking", reflect.TypeOf((*MockRankingInteractor)(nil).CreateRanking), userID, contestID, languages)
}

// CreateLog mocks base method
func (m *MockRankingInteractor) CreateLog(log domain.ContestLog) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLog", log)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateLog indicates an expected call of CreateLog
func (mr *MockRankingInteractorMockRecorder) CreateLog(log interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLog", reflect.TypeOf((*MockRankingInteractor)(nil).CreateLog), log)
}

// UpdateRanking mocks base method
func (m *MockRankingInteractor) UpdateRanking(userID, contestID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRanking", userID, contestID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRanking indicates an expected call of UpdateRanking
func (mr *MockRankingInteractorMockRecorder) UpdateRanking(userID, contestID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRanking", reflect.TypeOf((*MockRankingInteractor)(nil).UpdateRanking), userID, contestID)
}
