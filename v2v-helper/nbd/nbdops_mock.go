// Code generated by MockGen. DO NOT EDIT.
// Source: ../nbd/nbdops.go

// Package nbd is a generated GoMock package.
package nbd

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	object "github.com/vmware/govmomi/object"
	types "github.com/vmware/govmomi/vim25/types"
)

// MockNBDOperations is a mock of NBDOperations interface.
type MockNBDOperations struct {
	ctrl     *gomock.Controller
	recorder *MockNBDOperationsMockRecorder
}

// MockNBDOperationsMockRecorder is the mock recorder for MockNBDOperations.
type MockNBDOperationsMockRecorder struct {
	mock *MockNBDOperations
}

// NewMockNBDOperations creates a new mock instance.
func NewMockNBDOperations(ctrl *gomock.Controller) *MockNBDOperations {
	mock := &MockNBDOperations{ctrl: ctrl}
	mock.recorder = &MockNBDOperationsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNBDOperations) EXPECT() *MockNBDOperationsMockRecorder {
	return m.recorder
}

// CopyChangedBlocks mocks base method.
func (m *MockNBDOperations) CopyChangedBlocks(changedAreas types.DiskChangeInfo, path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CopyChangedBlocks", changedAreas, path)
	ret0, _ := ret[0].(error)
	return ret0
}

// CopyChangedBlocks indicates an expected call of CopyChangedBlocks.
func (mr *MockNBDOperationsMockRecorder) CopyChangedBlocks(changedAreas, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CopyChangedBlocks", reflect.TypeOf((*MockNBDOperations)(nil).CopyChangedBlocks), changedAreas, path)
}

// CopyDisk mocks base method.
func (m *MockNBDOperations) CopyDisk(ctx context.Context, dest string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CopyDisk", ctx, dest)
	ret0, _ := ret[0].(error)
	return ret0
}

// CopyDisk indicates an expected call of CopyDisk.
func (mr *MockNBDOperationsMockRecorder) CopyDisk(ctx, dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CopyDisk", reflect.TypeOf((*MockNBDOperations)(nil).CopyDisk), ctx, dest)
}

// StartNBDServer mocks base method.
func (m *MockNBDOperations) StartNBDServer(vm *object.VirtualMachine, server, username, password, thumbprint, snapref, file string, progchan chan string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartNBDServer", vm, server, username, password, thumbprint, snapref, file, progchan)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartNBDServer indicates an expected call of StartNBDServer.
func (mr *MockNBDOperationsMockRecorder) StartNBDServer(vm, server, username, password, thumbprint, snapref, file, progchan interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartNBDServer", reflect.TypeOf((*MockNBDOperations)(nil).StartNBDServer), vm, server, username, password, thumbprint, snapref, file, progchan)
}

// StopNBDServer mocks base method.
func (m *MockNBDOperations) StopNBDServer() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopNBDServer")
	ret0, _ := ret[0].(error)
	return ret0
}

// StopNBDServer indicates an expected call of StopNBDServer.
func (mr *MockNBDOperationsMockRecorder) StopNBDServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopNBDServer", reflect.TypeOf((*MockNBDOperations)(nil).StopNBDServer))
}
