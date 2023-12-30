// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/cli/progress.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// Mockprogress is a mock of progress interface.
type Mockprogress struct {
	ctrl     *gomock.Controller
	recorder *MockprogressMockRecorder
}

// MockprogressMockRecorder is the mock recorder for Mockprogress.
type MockprogressMockRecorder struct {
	mock *Mockprogress
}

// NewMockprogress creates a new mock instance.
func NewMockprogress(ctrl *gomock.Controller) *Mockprogress {
	mock := &Mockprogress{ctrl: ctrl}
	mock.recorder = &MockprogressMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockprogress) EXPECT() *MockprogressMockRecorder {
	return m.recorder
}

// Start mocks base method.
func (m *Mockprogress) Start(label string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start", label)
}

// Start indicates an expected call of Start.
func (mr *MockprogressMockRecorder) Start(label interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*Mockprogress)(nil).Start), label)
}

// Stop mocks base method.
func (m *Mockprogress) Stop(label string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop", label)
}

// Stop indicates an expected call of Stop.
func (mr *MockprogressMockRecorder) Stop(label interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*Mockprogress)(nil).Stop), label)
}
