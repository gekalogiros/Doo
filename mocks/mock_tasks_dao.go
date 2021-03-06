// Code generated by MockGen. DO NOT EDIT.
// Source: dao/dao.go

// Package mock_dao is a generated GoMock package.
package mock_dao

import (
	model "github.com/gekalogiros/Doo/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockTaskDao is a mock of TaskDao interface
type MockTaskDao struct {
	ctrl     *gomock.Controller
	recorder *MockTaskDaoMockRecorder
}

// MockTaskDaoMockRecorder is the mock recorder for MockTaskDao
type MockTaskDaoMockRecorder struct {
	mock *MockTaskDao
}

// NewMockTaskDao creates a new mock instance
func NewMockTaskDao(ctrl *gomock.Controller) *MockTaskDao {
	mock := &MockTaskDao{ctrl: ctrl}
	mock.recorder = &MockTaskDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTaskDao) EXPECT() *MockTaskDaoMockRecorder {
	return m.recorder
}

// Save mocks base method
func (m *MockTaskDao) Save(n *model.Task) {
	m.ctrl.Call(m, "Save", n)
}

// Save indicates an expected call of Save
func (mr *MockTaskDaoMockRecorder) Save(n interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockTaskDao)(nil).Save), n)
}

// RemoveByDate mocks base method
func (m *MockTaskDao) RemoveByDate(date time.Time) {
	m.ctrl.Call(m, "RemoveByDate", date)
}

// RemoveByDate indicates an expected call of RemoveByDate
func (mr *MockTaskDaoMockRecorder) RemoveByDate(date interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveByDate", reflect.TypeOf((*MockTaskDao)(nil).RemoveByDate), date)
}

// RetrieveByDate mocks base method
func (m *MockTaskDao) RetrieveByDate(date time.Time) []model.Task {
	ret := m.ctrl.Call(m, "RetrieveByDate", date)
	ret0, _ := ret[0].([]model.Task)
	return ret0
}

// RetrieveByDate indicates an expected call of RetrieveByDate
func (mr *MockTaskDaoMockRecorder) RetrieveByDate(date interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveByDate", reflect.TypeOf((*MockTaskDao)(nil).RetrieveByDate), date)
}

// RemovePast mocks base method
func (m *MockTaskDao) RemovePast() {
	m.ctrl.Call(m, "RemovePast")
}

// RemovePast indicates an expected call of RemovePast
func (mr *MockTaskDaoMockRecorder) RemovePast() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePast", reflect.TypeOf((*MockTaskDao)(nil).RemovePast))
}

// Move mocks base method
func (m *MockTaskDao) Move(id string, source, target time.Time) {
	m.ctrl.Call(m, "Move", id, source, target)
}

// Move indicates an expected call of Move
func (mr *MockTaskDaoMockRecorder) Move(id, source, target interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Move", reflect.TypeOf((*MockTaskDao)(nil).Move), id, source, target)
}
