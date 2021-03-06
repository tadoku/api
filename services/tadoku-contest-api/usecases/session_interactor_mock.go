// Code generated by MockGen. DO NOT EDIT.
// Source: session_interactor.go

// Package usecases is a generated GoMock package.
package usecases

import (
	gomock "github.com/golang/mock/gomock"
	domain "github.com/tadoku/tadoku/services/tadoku-contest-api/domain"
	reflect "reflect"
)

// MockSessionInteractor is a mock of SessionInteractor interface
type MockSessionInteractor struct {
	ctrl     *gomock.Controller
	recorder *MockSessionInteractorMockRecorder
}

// MockSessionInteractorMockRecorder is the mock recorder for MockSessionInteractor
type MockSessionInteractorMockRecorder struct {
	mock *MockSessionInteractor
}

// NewMockSessionInteractor creates a new mock instance
func NewMockSessionInteractor(ctrl *gomock.Controller) *MockSessionInteractor {
	mock := &MockSessionInteractor{ctrl: ctrl}
	mock.recorder = &MockSessionInteractorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSessionInteractor) EXPECT() *MockSessionInteractorMockRecorder {
	return m.recorder
}

// CreateSession mocks base method
func (m *MockSessionInteractor) CreateSession(email, password string) (domain.User, string, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", email, password)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(int64)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// CreateSession indicates an expected call of CreateSession
func (mr *MockSessionInteractorMockRecorder) CreateSession(email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockSessionInteractor)(nil).CreateSession), email, password)
}

// RefreshSession mocks base method
func (m *MockSessionInteractor) RefreshSession(user domain.User) (domain.User, string, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshSession", user)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(int64)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// RefreshSession indicates an expected call of RefreshSession
func (mr *MockSessionInteractorMockRecorder) RefreshSession(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshSession", reflect.TypeOf((*MockSessionInteractor)(nil).RefreshSession), user)
}
