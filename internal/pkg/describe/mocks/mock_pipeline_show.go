// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/describe/pipeline_show.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	codepipeline "github.com/aws/copilot-cli/internal/pkg/aws/codepipeline"
	gomock "github.com/golang/mock/gomock"
)

// MockpipelineGetter is a mock of pipelineGetter interface.
type MockpipelineGetter struct {
	ctrl     *gomock.Controller
	recorder *MockpipelineGetterMockRecorder
}

// MockpipelineGetterMockRecorder is the mock recorder for MockpipelineGetter.
type MockpipelineGetterMockRecorder struct {
	mock *MockpipelineGetter
}

// NewMockpipelineGetter creates a new mock instance.
func NewMockpipelineGetter(ctrl *gomock.Controller) *MockpipelineGetter {
	mock := &MockpipelineGetter{ctrl: ctrl}
	mock.recorder = &MockpipelineGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockpipelineGetter) EXPECT() *MockpipelineGetterMockRecorder {
	return m.recorder
}

// GetPipeline mocks base method.
func (m *MockpipelineGetter) GetPipeline(pipelineName string) (*codepipeline.Pipeline, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPipeline", pipelineName)
	ret0, _ := ret[0].(*codepipeline.Pipeline)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPipeline indicates an expected call of GetPipeline.
func (mr *MockpipelineGetterMockRecorder) GetPipeline(pipelineName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPipeline", reflect.TypeOf((*MockpipelineGetter)(nil).GetPipeline), pipelineName)
}
