// Code generated by MockGen. DO NOT EDIT.
// Source: greeter.go

// Package greeter is a generated GoMock package.
package greeter

import (
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/golang/mock/mockgen/tests/custom_package_name/client/v1"
	reflect "reflect"
)

// MockInputMaker is a mock of InputMaker interface
type MockInputMaker struct {
	ctrl     *gomock.Controller
	recorder *MockInputMakerMockRecorder
}

// MockInputMakerMockRecorder is the mock recorder for MockInputMaker
type MockInputMakerMockRecorder struct {
	mock *MockInputMaker
}

// NewMockInputMaker creates a new mock instance
func NewMockInputMaker(ctrl *gomock.Controller) *MockInputMaker {
	mock := &MockInputMaker{ctrl: ctrl}
	mock.recorder = &MockInputMakerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInputMaker) EXPECT() *MockInputMakerMockRecorder {
	return m.recorder
}

// MakeInput mocks base method
func (m *MockInputMaker) MakeInput() v1.GreetInput {
	ret := m.ctrl.Call(m, "MakeInput")
	ret0, _ := ret[0].(v1.GreetInput)
	return ret0
}

// MakeInput indicates an expected call of MakeInput
func (mr *MockInputMakerMockRecorder) MakeInput() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeInput", reflect.TypeOf((*MockInputMaker)(nil).MakeInput))
}
