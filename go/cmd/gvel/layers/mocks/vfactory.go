// Code generated by MockGen. DO NOT EDIT.
// Source: ./layers/repositories/vfactory/interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	entity "github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	ivclient "github.com/velo-protocol/DRSv2/go/libs/vclient/ivclient"
	reflect "reflect"
)

// MockVFactoryRepo is a mock of Repository interface
type MockVFactoryRepo struct {
	ctrl     *gomock.Controller
	recorder *MockVFactoryRepoMockRecorder
}

// MockVFactoryRepoMockRecorder is the mock recorder for MockVFactoryRepo
type MockVFactoryRepoMockRecorder struct {
	mock *MockVFactoryRepo
}

// NewMockVFactoryRepo creates a new mock instance
func NewMockVFactoryRepo(ctrl *gomock.Controller) *MockVFactoryRepo {
	mock := &MockVFactoryRepo{ctrl: ctrl}
	mock.recorder = &MockVFactoryRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVFactoryRepo) EXPECT() *MockVFactoryRepoMockRecorder {
	return m.recorder
}

// NewClient mocks base method
func (m *MockVFactoryRepo) NewClient(input *entity.NewClientInput) (ivclient.VClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewClient", input)
	ret0, _ := ret[0].(ivclient.VClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewClient indicates an expected call of NewClient
func (mr *MockVFactoryRepoMockRecorder) NewClient(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewClient", reflect.TypeOf((*MockVFactoryRepo)(nil).NewClient), input)
}

// NewClientFromConfig mocks base method
func (m *MockVFactoryRepo) NewClientFromConfig(privateKey string) (ivclient.VClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewClientFromConfig", privateKey)
	ret0, _ := ret[0].(ivclient.VClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewClientFromConfig indicates an expected call of NewClientFromConfig
func (mr *MockVFactoryRepoMockRecorder) NewClientFromConfig(privateKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewClientFromConfig", reflect.TypeOf((*MockVFactoryRepo)(nil).NewClientFromConfig), privateKey)
}
