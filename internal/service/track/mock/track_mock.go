// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/RaceSimHub/race-hub-backend/internal/service/track (interfaces: TrackContract)
//
// Generated by this command:
//
//	mockgen -destination=./mock/track_mock.go -package=mock . TrackContract
//

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	sqlc "github.com/RaceSimHub/race-hub-backend/internal/database/sqlc"
	gomock "go.uber.org/mock/gomock"
)

// MockTrackContract is a mock of TrackContract interface.
type MockTrackContract struct {
	ctrl     *gomock.Controller
	recorder *MockTrackContractMockRecorder
	isgomock struct{}
}

// MockTrackContractMockRecorder is the mock recorder for MockTrackContract.
type MockTrackContractMockRecorder struct {
	mock *MockTrackContract
}

// NewMockTrackContract creates a new mock instance.
func NewMockTrackContract(ctrl *gomock.Controller) *MockTrackContract {
	mock := &MockTrackContract{ctrl: ctrl}
	mock.recorder = &MockTrackContractMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrackContract) EXPECT() *MockTrackContractMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTrackContract) Create(name, country string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", name, country)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTrackContractMockRecorder) Create(name, country any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTrackContract)(nil).Create), name, country)
}

// Delete mocks base method.
func (m *MockTrackContract) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTrackContractMockRecorder) Delete(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTrackContract)(nil).Delete), id)
}

// GetByID mocks base method.
func (m *MockTrackContract) GetByID(id int) (sqlc.SelectTrackByIdRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(sqlc.SelectTrackByIdRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockTrackContractMockRecorder) GetByID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockTrackContract)(nil).GetByID), id)
}

// GetList mocks base method.
func (m *MockTrackContract) GetList(offset, limit int) ([]sqlc.SelectListTracksRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetList", offset, limit)
	ret0, _ := ret[0].([]sqlc.SelectListTracksRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetList indicates an expected call of GetList.
func (mr *MockTrackContractMockRecorder) GetList(offset, limit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetList", reflect.TypeOf((*MockTrackContract)(nil).GetList), offset, limit)
}

// Update mocks base method.
func (m *MockTrackContract) Update(id int, name, country string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, name, country)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockTrackContractMockRecorder) Update(id, name, country any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTrackContract)(nil).Update), id, name, country)
}
