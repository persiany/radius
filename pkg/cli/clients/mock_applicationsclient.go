// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/project-radius/radius/pkg/cli/clients (interfaces: ApplicationsManagementClient)

// Package clients is a generated GoMock package.
package clients

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	generated "github.com/project-radius/radius/pkg/cli/clients_new/generated"
	v20220315privatepreview "github.com/project-radius/radius/pkg/corerp/api/v20220315privatepreview"
	v20220315privatepreview0 "github.com/project-radius/radius/pkg/ucp/api/v20220315privatepreview"
)

// MockApplicationsManagementClient is a mock of ApplicationsManagementClient interface.
type MockApplicationsManagementClient struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationsManagementClientMockRecorder
}

// MockApplicationsManagementClientMockRecorder is the mock recorder for MockApplicationsManagementClient.
type MockApplicationsManagementClientMockRecorder struct {
	mock *MockApplicationsManagementClient
}

// NewMockApplicationsManagementClient creates a new mock instance.
func NewMockApplicationsManagementClient(ctrl *gomock.Controller) *MockApplicationsManagementClient {
	mock := &MockApplicationsManagementClient{ctrl: ctrl}
	mock.recorder = &MockApplicationsManagementClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplicationsManagementClient) EXPECT() *MockApplicationsManagementClientMockRecorder {
	return m.recorder
}

// CreateEnvironment mocks base method.
func (m *MockApplicationsManagementClient) CreateEnvironment(arg0 context.Context, arg1, arg2, arg3, arg4, arg5 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEnvironment", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEnvironment indicates an expected call of CreateEnvironment.
func (mr *MockApplicationsManagementClientMockRecorder) CreateEnvironment(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEnvironment", reflect.TypeOf((*MockApplicationsManagementClient)(nil).CreateEnvironment), arg0, arg1, arg2, arg3, arg4, arg5)
}

// CreateUCPGroup mocks base method.
func (m *MockApplicationsManagementClient) CreateUCPGroup(arg0 context.Context, arg1, arg2, arg3 string, arg4 v20220315privatepreview0.ResourceGroupResource) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUCPGroup", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUCPGroup indicates an expected call of CreateUCPGroup.
func (mr *MockApplicationsManagementClientMockRecorder) CreateUCPGroup(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUCPGroup", reflect.TypeOf((*MockApplicationsManagementClient)(nil).CreateUCPGroup), arg0, arg1, arg2, arg3, arg4)
}

// DeleteApplication mocks base method.
func (m *MockApplicationsManagementClient) DeleteApplication(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteApplication", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteApplication indicates an expected call of DeleteApplication.
func (mr *MockApplicationsManagementClientMockRecorder) DeleteApplication(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteApplication", reflect.TypeOf((*MockApplicationsManagementClient)(nil).DeleteApplication), arg0, arg1)
}

// DeleteEnv mocks base method.
func (m *MockApplicationsManagementClient) DeleteEnv(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEnv", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteEnv indicates an expected call of DeleteEnv.
func (mr *MockApplicationsManagementClientMockRecorder) DeleteEnv(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEnv", reflect.TypeOf((*MockApplicationsManagementClient)(nil).DeleteEnv), arg0, arg1)
}

// DeleteResource mocks base method.
func (m *MockApplicationsManagementClient) DeleteResource(arg0 context.Context, arg1, arg2 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteResource", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteResource indicates an expected call of DeleteResource.
func (mr *MockApplicationsManagementClientMockRecorder) DeleteResource(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteResource", reflect.TypeOf((*MockApplicationsManagementClient)(nil).DeleteResource), arg0, arg1, arg2)
}

// DeleteUCPGroup mocks base method.
func (m *MockApplicationsManagementClient) DeleteUCPGroup(arg0 context.Context, arg1, arg2, arg3 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUCPGroup", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUCPGroup indicates an expected call of DeleteUCPGroup.
func (mr *MockApplicationsManagementClientMockRecorder) DeleteUCPGroup(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUCPGroup", reflect.TypeOf((*MockApplicationsManagementClient)(nil).DeleteUCPGroup), arg0, arg1, arg2, arg3)
}

// GetEnvDetails mocks base method.
func (m *MockApplicationsManagementClient) GetEnvDetails(arg0 context.Context, arg1 string) (v20220315privatepreview.EnvironmentResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnvDetails", arg0, arg1)
	ret0, _ := ret[0].(v20220315privatepreview.EnvironmentResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEnvDetails indicates an expected call of GetEnvDetails.
func (mr *MockApplicationsManagementClientMockRecorder) GetEnvDetails(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnvDetails", reflect.TypeOf((*MockApplicationsManagementClient)(nil).GetEnvDetails), arg0, arg1)
}

// ListAllResourcesByApplication mocks base method.
func (m *MockApplicationsManagementClient) ListAllResourcesByApplication(arg0 context.Context, arg1 string) ([]generated.GenericResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllResourcesByApplication", arg0, arg1)
	ret0, _ := ret[0].([]generated.GenericResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllResourcesByApplication indicates an expected call of ListAllResourcesByApplication.
func (mr *MockApplicationsManagementClientMockRecorder) ListAllResourcesByApplication(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllResourcesByApplication", reflect.TypeOf((*MockApplicationsManagementClient)(nil).ListAllResourcesByApplication), arg0, arg1)
}

// ListAllResourcesByEnvironment mocks base method.
func (m *MockApplicationsManagementClient) ListAllResourcesByEnvironment(arg0 context.Context, arg1 string) ([]generated.GenericResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllResourcesByEnvironment", arg0, arg1)
	ret0, _ := ret[0].([]generated.GenericResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllResourcesByEnvironment indicates an expected call of ListAllResourcesByEnvironment.
func (mr *MockApplicationsManagementClientMockRecorder) ListAllResourcesByEnvironment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllResourcesByEnvironment", reflect.TypeOf((*MockApplicationsManagementClient)(nil).ListAllResourcesByEnvironment), arg0, arg1)
}

// ListAllResourcesByType mocks base method.
func (m *MockApplicationsManagementClient) ListAllResourcesByType(arg0 context.Context, arg1 string) ([]generated.GenericResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllResourcesByType", arg0, arg1)
	ret0, _ := ret[0].([]generated.GenericResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllResourcesByType indicates an expected call of ListAllResourcesByType.
func (mr *MockApplicationsManagementClientMockRecorder) ListAllResourcesByType(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllResourcesByType", reflect.TypeOf((*MockApplicationsManagementClient)(nil).ListAllResourcesByType), arg0, arg1)
}

// ListAllResourcesOfTypeInApplication mocks base method.
func (m *MockApplicationsManagementClient) ListAllResourcesOfTypeInApplication(arg0 context.Context, arg1, arg2 string) ([]generated.GenericResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllResourcesOfTypeInApplication", arg0, arg1, arg2)
	ret0, _ := ret[0].([]generated.GenericResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllResourcesOfTypeInApplication indicates an expected call of ListAllResourcesOfTypeInApplication.
func (mr *MockApplicationsManagementClientMockRecorder) ListAllResourcesOfTypeInApplication(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllResourcesOfTypeInApplication", reflect.TypeOf((*MockApplicationsManagementClient)(nil).ListAllResourcesOfTypeInApplication), arg0, arg1, arg2)
}

// ListAllResourcesOfTypeInEnvironment mocks base method.
func (m *MockApplicationsManagementClient) ListAllResourcesOfTypeInEnvironment(arg0 context.Context, arg1, arg2 string) ([]generated.GenericResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllResourcesOfTypeInEnvironment", arg0, arg1, arg2)
	ret0, _ := ret[0].([]generated.GenericResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllResourcesOfTypeInEnvironment indicates an expected call of ListAllResourcesOfTypeInEnvironment.
func (mr *MockApplicationsManagementClientMockRecorder) ListAllResourcesOfTypeInEnvironment(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllResourcesOfTypeInEnvironment", reflect.TypeOf((*MockApplicationsManagementClient)(nil).ListAllResourcesOfTypeInEnvironment), arg0, arg1, arg2)
}

// ListApplications mocks base method.
func (m *MockApplicationsManagementClient) ListApplications(arg0 context.Context) ([]v20220315privatepreview.ApplicationResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListApplications", arg0)
	ret0, _ := ret[0].([]v20220315privatepreview.ApplicationResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListApplications indicates an expected call of ListApplications.
func (mr *MockApplicationsManagementClientMockRecorder) ListApplications(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListApplications", reflect.TypeOf((*MockApplicationsManagementClient)(nil).ListApplications), arg0)
}

// ListEnv mocks base method.
func (m *MockApplicationsManagementClient) ListEnv(arg0 context.Context) ([]v20220315privatepreview.EnvironmentResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEnv", arg0)
	ret0, _ := ret[0].([]v20220315privatepreview.EnvironmentResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEnv indicates an expected call of ListEnv.
func (mr *MockApplicationsManagementClientMockRecorder) ListEnv(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEnv", reflect.TypeOf((*MockApplicationsManagementClient)(nil).ListEnv), arg0)
}

// ListUCPGroup mocks base method.
func (m *MockApplicationsManagementClient) ListUCPGroup(arg0 context.Context, arg1, arg2 string) ([]v20220315privatepreview0.ResourceGroupResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUCPGroup", arg0, arg1, arg2)
	ret0, _ := ret[0].([]v20220315privatepreview0.ResourceGroupResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUCPGroup indicates an expected call of ListUCPGroup.
func (mr *MockApplicationsManagementClientMockRecorder) ListUCPGroup(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUCPGroup", reflect.TypeOf((*MockApplicationsManagementClient)(nil).ListUCPGroup), arg0, arg1, arg2)
}

// ShowApplication mocks base method.
func (m *MockApplicationsManagementClient) ShowApplication(arg0 context.Context, arg1 string) (v20220315privatepreview.ApplicationResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowApplication", arg0, arg1)
	ret0, _ := ret[0].(v20220315privatepreview.ApplicationResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShowApplication indicates an expected call of ShowApplication.
func (mr *MockApplicationsManagementClientMockRecorder) ShowApplication(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowApplication", reflect.TypeOf((*MockApplicationsManagementClient)(nil).ShowApplication), arg0, arg1)
}

// ShowResource mocks base method.
func (m *MockApplicationsManagementClient) ShowResource(arg0 context.Context, arg1, arg2 string) (generated.GenericResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowResource", arg0, arg1, arg2)
	ret0, _ := ret[0].(generated.GenericResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShowResource indicates an expected call of ShowResource.
func (mr *MockApplicationsManagementClientMockRecorder) ShowResource(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowResource", reflect.TypeOf((*MockApplicationsManagementClient)(nil).ShowResource), arg0, arg1, arg2)
}

// ShowUCPGroup mocks base method.
func (m *MockApplicationsManagementClient) ShowUCPGroup(arg0 context.Context, arg1, arg2, arg3 string) (v20220315privatepreview0.ResourceGroupResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowUCPGroup", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(v20220315privatepreview0.ResourceGroupResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShowUCPGroup indicates an expected call of ShowUCPGroup.
func (mr *MockApplicationsManagementClientMockRecorder) ShowUCPGroup(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowUCPGroup", reflect.TypeOf((*MockApplicationsManagementClient)(nil).ShowUCPGroup), arg0, arg1, arg2, arg3)
}
