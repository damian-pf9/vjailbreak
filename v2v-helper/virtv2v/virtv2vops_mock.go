// Code generated by MockGen. DO NOT EDIT.
// Source: ../virtv2v/virtv2vops.go

// Package virtv2v is a generated GoMock package.
package virtv2v

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockVirtV2VOperations is a mock of VirtV2VOperations interface.
type MockVirtV2VOperations struct {
	ctrl     *gomock.Controller
	recorder *MockVirtV2VOperationsMockRecorder
}

// MockVirtV2VOperationsMockRecorder is the mock recorder for MockVirtV2VOperations.
type MockVirtV2VOperationsMockRecorder struct {
	mock *MockVirtV2VOperations
}

// NewMockVirtV2VOperations creates a new mock instance.
func NewMockVirtV2VOperations(ctrl *gomock.Controller) *MockVirtV2VOperations {
	mock := &MockVirtV2VOperations{ctrl: ctrl}
	mock.recorder = &MockVirtV2VOperationsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtV2VOperations) EXPECT() *MockVirtV2VOperationsMockRecorder {
	return m.recorder
}

// AddWildcardNetplan mocks base method.
func (m *MockVirtV2VOperations) AddWildcardNetplan(path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddWildcardNetplan", path)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddWildcardNetplan indicates an expected call of AddWildcardNetplan.
func (mr *MockVirtV2VOperationsMockRecorder) AddWildcardNetplan(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWildcardNetplan", reflect.TypeOf((*MockVirtV2VOperations)(nil).AddWildcardNetplan), path)
}

// ConvertDisk mocks base method.
func (m *MockVirtV2VOperations) ConvertDisk(ctx context.Context, path, ostype, virtiowindriver string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvertDisk", ctx, path, ostype, virtiowindriver)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConvertDisk indicates an expected call of ConvertDisk.
func (mr *MockVirtV2VOperationsMockRecorder) ConvertDisk(ctx, path, ostype, virtiowindriver interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvertDisk", reflect.TypeOf((*MockVirtV2VOperations)(nil).ConvertDisk), ctx, path, ostype, virtiowindriver)
}

// GetPartitions mocks base method.
func (m *MockVirtV2VOperations) GetPartitions(disk string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPartitions", disk)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartitions indicates an expected call of GetPartitions.
func (mr *MockVirtV2VOperationsMockRecorder) GetPartitions(disk interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartitions", reflect.TypeOf((*MockVirtV2VOperations)(nil).GetPartitions), disk)
}

// NTFSFix mocks base method.
func (m *MockVirtV2VOperations) NTFSFix(path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NTFSFix", path)
	ret0, _ := ret[0].(error)
	return ret0
}

// NTFSFix indicates an expected call of NTFSFix.
func (mr *MockVirtV2VOperationsMockRecorder) NTFSFix(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NTFSFix", reflect.TypeOf((*MockVirtV2VOperations)(nil).NTFSFix), path)
}

// RetainAlphanumeric mocks base method.
func (m *MockVirtV2VOperations) RetainAlphanumeric(input string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetainAlphanumeric", input)
	ret0, _ := ret[0].(string)
	return ret0
}

// RetainAlphanumeric indicates an expected call of RetainAlphanumeric.
func (mr *MockVirtV2VOperationsMockRecorder) RetainAlphanumeric(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetainAlphanumeric", reflect.TypeOf((*MockVirtV2VOperations)(nil).RetainAlphanumeric), input)
}
