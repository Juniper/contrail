// Code generated by MockGen. DO NOT EDIT.
// Source: input.go

// Package empty_interface is a generated GoMock package.
package empty_interface

import (
	gomock "github.com/golang/mock/gomock"
)

// MockEmpty is a mock of Empty interface
type MockEmpty struct {
	ctrl     *gomock.Controller
	recorder *MockEmptyMockRecorder
}

// MockEmptyMockRecorder is the mock recorder for MockEmpty
type MockEmptyMockRecorder struct {
	mock *MockEmpty
}

// NewMockEmpty creates a new mock instance
func NewMockEmpty(ctrl *gomock.Controller) *MockEmpty {
	mock := &MockEmpty{ctrl: ctrl}
	mock.recorder = &MockEmptyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEmpty) EXPECT() *MockEmptyMockRecorder {
	return m.recorder
}
