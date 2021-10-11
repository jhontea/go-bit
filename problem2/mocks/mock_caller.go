// Code generated by MockGen. DO NOT EDIT.
// Source: callers/contract.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	responses "go-bit/entities/responses"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIMDBCaller is a mock of IMDBCaller interface.
type MockIMDBCaller struct {
	ctrl     *gomock.Controller
	recorder *MockIMDBCallerMockRecorder
}

// MockIMDBCallerMockRecorder is the mock recorder for MockIMDBCaller.
type MockIMDBCallerMockRecorder struct {
	mock *MockIMDBCaller
}

// NewMockIMDBCaller creates a new mock instance.
func NewMockIMDBCaller(ctrl *gomock.Controller) *MockIMDBCaller {
	mock := &MockIMDBCaller{ctrl: ctrl}
	mock.recorder = &MockIMDBCallerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIMDBCaller) EXPECT() *MockIMDBCallerMockRecorder {
	return m.recorder
}

// GetDetail mocks base method.
func (m *MockIMDBCaller) GetDetail(ctx context.Context, id string) (responses.IMDBGetDetailResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDetail", ctx, id)
	ret0, _ := ret[0].(responses.IMDBGetDetailResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDetail indicates an expected call of GetDetail.
func (mr *MockIMDBCallerMockRecorder) GetDetail(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetail", reflect.TypeOf((*MockIMDBCaller)(nil).GetDetail), ctx, id)
}

// Search mocks base method.
func (m *MockIMDBCaller) Search(ctx context.Context, search string, page int32) (responses.IMDBSearchResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, search, page)
	ret0, _ := ret[0].(responses.IMDBSearchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockIMDBCallerMockRecorder) Search(ctx, search, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockIMDBCaller)(nil).Search), ctx, search, page)
}