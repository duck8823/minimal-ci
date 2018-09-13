// Code generated by MockGen. DO NOT EDIT.
// Source: application/service/runner/runner.go

// Package mock_runner is a generated GoMock package.
package mock_runner

import (
	context "github.com/duck8823/duci/application/context"
	runner "github.com/duck8823/duci/application/service/runner"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRunner is a mock of Runner interface
type MockRunner struct {
	ctrl     *gomock.Controller
	recorder *MockRunnerMockRecorder
}

// MockRunnerMockRecorder is the mock recorder for MockRunner
type MockRunnerMockRecorder struct {
	mock *MockRunner
}

// NewMockRunner creates a new mock instance
func NewMockRunner(ctrl *gomock.Controller) *MockRunner {
	mock := &MockRunner{ctrl: ctrl}
	mock.recorder = &MockRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRunner) EXPECT() *MockRunnerMockRecorder {
	return m.recorder
}

// Run mocks base method
func (m *MockRunner) Run(ctx context.Context, src runner.TargetSource, command ...string) error {
	varargs := []interface{}{ctx, src}
	for _, a := range command {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Run", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockRunnerMockRecorder) Run(ctx, src interface{}, command ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, src}, command...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockRunner)(nil).Run), varargs...)
}
