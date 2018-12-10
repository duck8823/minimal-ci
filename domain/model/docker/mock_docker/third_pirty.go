// Code generated by MockGen. DO NOT EDIT.
// Source: domain/model/docker/third_pirty.go

// Package mock_docker is a generated GoMock package.
package mock_docker

import (
	context "context"
	types "github.com/docker/docker/api/types"
	container "github.com/docker/docker/api/types/container"
	network "github.com/docker/docker/api/types/network"
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockMoby is a mock of Moby interface
type MockMoby struct {
	ctrl     *gomock.Controller
	recorder *MockMobyMockRecorder
}

// MockMobyMockRecorder is the mock recorder for MockMoby
type MockMobyMockRecorder struct {
	mock *MockMoby
}

// NewMockMoby creates a new mock instance
func NewMockMoby(ctrl *gomock.Controller) *MockMoby {
	mock := &MockMoby{ctrl: ctrl}
	mock.recorder = &MockMobyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMoby) EXPECT() *MockMobyMockRecorder {
	return m.recorder
}

// ImageBuild mocks base method
func (m *MockMoby) ImageBuild(ctx context.Context, buildContext io.Reader, options types.ImageBuildOptions) (types.ImageBuildResponse, error) {
	ret := m.ctrl.Call(m, "ImageBuild", ctx, buildContext, options)
	ret0, _ := ret[0].(types.ImageBuildResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImageBuild indicates an expected call of ImageBuild
func (mr *MockMobyMockRecorder) ImageBuild(ctx, buildContext, options interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageBuild", reflect.TypeOf((*MockMoby)(nil).ImageBuild), ctx, buildContext, options)
}

// ContainerCreate mocks base method
func (m *MockMoby) ContainerCreate(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string) (container.ContainerCreateCreatedBody, error) {
	ret := m.ctrl.Call(m, "ContainerCreate", ctx, config, hostConfig, networkingConfig, containerName)
	ret0, _ := ret[0].(container.ContainerCreateCreatedBody)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerCreate indicates an expected call of ContainerCreate
func (mr *MockMobyMockRecorder) ContainerCreate(ctx, config, hostConfig, networkingConfig, containerName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerCreate", reflect.TypeOf((*MockMoby)(nil).ContainerCreate), ctx, config, hostConfig, networkingConfig, containerName)
}

// ContainerStart mocks base method
func (m *MockMoby) ContainerStart(ctx context.Context, containerID string, options types.ContainerStartOptions) error {
	ret := m.ctrl.Call(m, "ContainerStart", ctx, containerID, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// ContainerStart indicates an expected call of ContainerStart
func (mr *MockMobyMockRecorder) ContainerStart(ctx, containerID, options interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerStart", reflect.TypeOf((*MockMoby)(nil).ContainerStart), ctx, containerID, options)
}

// ContainerLogs mocks base method
func (m *MockMoby) ContainerLogs(ctx context.Context, container string, options types.ContainerLogsOptions) (io.ReadCloser, error) {
	ret := m.ctrl.Call(m, "ContainerLogs", ctx, container, options)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerLogs indicates an expected call of ContainerLogs
func (mr *MockMobyMockRecorder) ContainerLogs(ctx, container, options interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerLogs", reflect.TypeOf((*MockMoby)(nil).ContainerLogs), ctx, container, options)
}

// ContainerRemove mocks base method
func (m *MockMoby) ContainerRemove(ctx context.Context, containerID string, options types.ContainerRemoveOptions) error {
	ret := m.ctrl.Call(m, "ContainerRemove", ctx, containerID, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// ContainerRemove indicates an expected call of ContainerRemove
func (mr *MockMobyMockRecorder) ContainerRemove(ctx, containerID, options interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerRemove", reflect.TypeOf((*MockMoby)(nil).ContainerRemove), ctx, containerID, options)
}

// ImageRemove mocks base method
func (m *MockMoby) ImageRemove(ctx context.Context, imageID string, options types.ImageRemoveOptions) ([]types.ImageDeleteResponseItem, error) {
	ret := m.ctrl.Call(m, "ImageRemove", ctx, imageID, options)
	ret0, _ := ret[0].([]types.ImageDeleteResponseItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImageRemove indicates an expected call of ImageRemove
func (mr *MockMobyMockRecorder) ImageRemove(ctx, imageID, options interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageRemove", reflect.TypeOf((*MockMoby)(nil).ImageRemove), ctx, imageID, options)
}

// ContainerWait mocks base method
func (m *MockMoby) ContainerWait(ctx context.Context, containerID string, condition container.WaitCondition) (<-chan container.ContainerWaitOKBody, <-chan error) {
	ret := m.ctrl.Call(m, "ContainerWait", ctx, containerID, condition)
	ret0, _ := ret[0].(<-chan container.ContainerWaitOKBody)
	ret1, _ := ret[1].(<-chan error)
	return ret0, ret1
}

// ContainerWait indicates an expected call of ContainerWait
func (mr *MockMobyMockRecorder) ContainerWait(ctx, containerID, condition interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerWait", reflect.TypeOf((*MockMoby)(nil).ContainerWait), ctx, containerID, condition)
}

// Info mocks base method
func (m *MockMoby) Info(ctx context.Context) (types.Info, error) {
	ret := m.ctrl.Call(m, "Info", ctx)
	ret0, _ := ret[0].(types.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info indicates an expected call of Info
func (mr *MockMobyMockRecorder) Info(ctx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockMoby)(nil).Info), ctx)
}
