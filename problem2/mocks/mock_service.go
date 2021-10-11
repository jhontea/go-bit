// Code generated by MockGen. DO NOT EDIT.
// Source: services/contract.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	responses "go-bit/entities/responses"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIMDBService is a mock of IMDBService interface.
type MockIMDBService struct {
	ctrl     *gomock.Controller
	recorder *MockIMDBServiceMockRecorder
}

// MockIMDBServiceMockRecorder is the mock recorder for MockIMDBService.
type MockIMDBServiceMockRecorder struct {
	mock *MockIMDBService
}

// NewMockIMDBService creates a new mock instance.
func NewMockIMDBService(ctrl *gomock.Controller) *MockIMDBService {
	mock := &MockIMDBService{ctrl: ctrl}
	mock.recorder = &MockIMDBServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIMDBService) EXPECT() *MockIMDBServiceMockRecorder {
	return m.recorder
}

// GetDetail mocks base method.
func (m *MockIMDBService) GetDetail(ctx context.Context, id string) (responses.IMDBGetDetailResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDetail", ctx, id)
	ret0, _ := ret[0].(responses.IMDBGetDetailResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDetail indicates an expected call of GetDetail.
func (mr *MockIMDBServiceMockRecorder) GetDetail(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetail", reflect.TypeOf((*MockIMDBService)(nil).GetDetail), ctx, id)
}

// Search mocks base method.
func (m *MockIMDBService) Search(ctx context.Context, search string, page int32) (responses.IMDBSearchResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, search, page)
	ret0, _ := ret[0].(responses.IMDBSearchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockIMDBServiceMockRecorder) Search(ctx, search, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockIMDBService)(nil).Search), ctx, search, page)
}
