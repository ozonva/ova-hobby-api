// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozonva/ova-hobby-api/internal/repo (interfaces: Repo)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	models "github.com/ozonva/ova-hobby-api/pkg/models"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// AddHobbies mocks base method.
func (m *MockRepo) AddHobbies(arg0 []models.Hobby) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddHobbies", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddHobbies indicates an expected call of AddHobbies.
func (mr *MockRepoMockRecorder) AddHobbies(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddHobbies", reflect.TypeOf((*MockRepo)(nil).AddHobbies), arg0)
}

// AddHobby mocks base method.
func (m *MockRepo) AddHobby(arg0 models.Hobby) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddHobby", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddHobby indicates an expected call of AddHobby.
func (mr *MockRepoMockRecorder) AddHobby(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddHobby", reflect.TypeOf((*MockRepo)(nil).AddHobby), arg0)
}

// DescribeHobby mocks base method.
func (m *MockRepo) DescribeHobby(arg0 uuid.UUID) (*models.Hobby, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeHobby", arg0)
	ret0, _ := ret[0].(*models.Hobby)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeHobby indicates an expected call of DescribeHobby.
func (mr *MockRepoMockRecorder) DescribeHobby(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeHobby", reflect.TypeOf((*MockRepo)(nil).DescribeHobby), arg0)
}

// ListHobbies mocks base method.
func (m *MockRepo) ListHobbies(arg0, arg1 uint64) ([]models.Hobby, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListHobbies", arg0, arg1)
	ret0, _ := ret[0].([]models.Hobby)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListHobbies indicates an expected call of ListHobbies.
func (mr *MockRepoMockRecorder) ListHobbies(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListHobbies", reflect.TypeOf((*MockRepo)(nil).ListHobbies), arg0, arg1)
}

// RemoveHobby mocks base method.
func (m *MockRepo) RemoveHobby(arg0 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveHobby", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveHobby indicates an expected call of RemoveHobby.
func (mr *MockRepoMockRecorder) RemoveHobby(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveHobby", reflect.TypeOf((*MockRepo)(nil).RemoveHobby), arg0)
}
