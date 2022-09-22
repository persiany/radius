// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/project-radius/radius/pkg/cli/prompt (interfaces: Interface)

// Package prompt is a generated GoMock package.
package prompt

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	promptui "github.com/manifoldco/promptui"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// RunPrompt mocks base method.
func (m *MockInterface) RunPrompt(arg0 promptui.Prompt) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunPrompt", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunPrompt indicates an expected call of RunPrompt.
func (mr *MockInterfaceMockRecorder) RunPrompt(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunPrompt", reflect.TypeOf((*MockInterface)(nil).RunPrompt), arg0)
}

// RunSelect mocks base method.
func (m *MockInterface) RunSelect(arg0 promptui.Select) (int, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunSelect", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RunSelect indicates an expected call of RunSelect.
func (mr *MockInterfaceMockRecorder) RunSelect(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunSelect", reflect.TypeOf((*MockInterface)(nil).RunSelect), arg0)
}
