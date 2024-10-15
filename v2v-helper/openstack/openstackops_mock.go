// Code generated by MockGen. DO NOT EDIT.
// Source: ../openstack/openstackops.go

// Package openstack is a generated GoMock package.
package openstack

import (
	reflect "reflect"
	vm "vjailbreak/vm"

	gomock "github.com/golang/mock/gomock"
	volumes "github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
	flavors "github.com/gophercloud/gophercloud/openstack/compute/v2/flavors"
	servers "github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	networks "github.com/gophercloud/gophercloud/openstack/networking/v2/networks"
	ports "github.com/gophercloud/gophercloud/openstack/networking/v2/ports"
)

// MockOpenstackOperations is a mock of OpenstackOperations interface.
type MockOpenstackOperations struct {
	ctrl     *gomock.Controller
	recorder *MockOpenstackOperationsMockRecorder
}

// MockOpenstackOperationsMockRecorder is the mock recorder for MockOpenstackOperations.
type MockOpenstackOperationsMockRecorder struct {
	mock *MockOpenstackOperations
}

// NewMockOpenstackOperations creates a new mock instance.
func NewMockOpenstackOperations(ctrl *gomock.Controller) *MockOpenstackOperations {
	mock := &MockOpenstackOperations{ctrl: ctrl}
	mock.recorder = &MockOpenstackOperationsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOpenstackOperations) EXPECT() *MockOpenstackOperationsMockRecorder {
	return m.recorder
}

// AttachVolumeToVM mocks base method.
func (m *MockOpenstackOperations) AttachVolumeToVM(volumeID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachVolumeToVM", volumeID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AttachVolumeToVM indicates an expected call of AttachVolumeToVM.
func (mr *MockOpenstackOperationsMockRecorder) AttachVolumeToVM(volumeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachVolumeToVM", reflect.TypeOf((*MockOpenstackOperations)(nil).AttachVolumeToVM), volumeID)
}

// CreatePort mocks base method.
func (m *MockOpenstackOperations) CreatePort(networkid *networks.Network, mac, ip, vmname string) (*ports.Port, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePort", networkid, mac, ip, vmname)
	ret0, _ := ret[0].(*ports.Port)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePort indicates an expected call of CreatePort.
func (mr *MockOpenstackOperationsMockRecorder) CreatePort(networkid, mac, ip, vmname interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePort", reflect.TypeOf((*MockOpenstackOperations)(nil).CreatePort), networkid, mac, ip, vmname)
}

// CreateVM mocks base method.
func (m *MockOpenstackOperations) CreateVM(flavor *flavors.Flavor, networkIDs, portIDs []string, vminfo vm.VMInfo) (*servers.Server, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVM", flavor, networkIDs, portIDs, vminfo)
	ret0, _ := ret[0].(*servers.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVM indicates an expected call of CreateVM.
func (mr *MockOpenstackOperationsMockRecorder) CreateVM(flavor, networkIDs, portIDs, vminfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVM", reflect.TypeOf((*MockOpenstackOperations)(nil).CreateVM), flavor, networkIDs, portIDs, vminfo)
}

// CreateVolume mocks base method.
func (m *MockOpenstackOperations) CreateVolume(name string, size int64, ostype string, uefi bool, volumetype string) (*volumes.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVolume", name, size, ostype, uefi, volumetype)
	ret0, _ := ret[0].(*volumes.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVolume indicates an expected call of CreateVolume.
func (mr *MockOpenstackOperationsMockRecorder) CreateVolume(name, size, ostype, uefi, volumetype interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVolume", reflect.TypeOf((*MockOpenstackOperations)(nil).CreateVolume), name, size, ostype, uefi, volumetype)
}

// DeleteVolume mocks base method.
func (m *MockOpenstackOperations) DeleteVolume(volumeID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVolume", volumeID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVolume indicates an expected call of DeleteVolume.
func (mr *MockOpenstackOperationsMockRecorder) DeleteVolume(volumeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVolume", reflect.TypeOf((*MockOpenstackOperations)(nil).DeleteVolume), volumeID)
}

// DetachVolumeFromVM mocks base method.
func (m *MockOpenstackOperations) DetachVolumeFromVM(volumeID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetachVolumeFromVM", volumeID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DetachVolumeFromVM indicates an expected call of DetachVolumeFromVM.
func (mr *MockOpenstackOperationsMockRecorder) DetachVolumeFromVM(volumeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetachVolumeFromVM", reflect.TypeOf((*MockOpenstackOperations)(nil).DetachVolumeFromVM), volumeID)
}

// EnableQGA mocks base method.
func (m *MockOpenstackOperations) EnableQGA(volume *volumes.Volume) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableQGA", volume)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableQGA indicates an expected call of EnableQGA.
func (mr *MockOpenstackOperationsMockRecorder) EnableQGA(volume interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableQGA", reflect.TypeOf((*MockOpenstackOperations)(nil).EnableQGA), volume)
}

// FindDevice mocks base method.
func (m *MockOpenstackOperations) FindDevice(volumeID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindDevice", volumeID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindDevice indicates an expected call of FindDevice.
func (mr *MockOpenstackOperationsMockRecorder) FindDevice(volumeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindDevice", reflect.TypeOf((*MockOpenstackOperations)(nil).FindDevice), volumeID)
}

// GetClosestFlavour mocks base method.
func (m *MockOpenstackOperations) GetClosestFlavour(cpu, memory int32) (*flavors.Flavor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClosestFlavour", cpu, memory)
	ret0, _ := ret[0].(*flavors.Flavor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClosestFlavour indicates an expected call of GetClosestFlavour.
func (mr *MockOpenstackOperationsMockRecorder) GetClosestFlavour(cpu, memory interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClosestFlavour", reflect.TypeOf((*MockOpenstackOperations)(nil).GetClosestFlavour), cpu, memory)
}

// GetNetwork mocks base method.
func (m *MockOpenstackOperations) GetNetwork(networkname string) (*networks.Network, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetwork", networkname)
	ret0, _ := ret[0].(*networks.Network)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNetwork indicates an expected call of GetNetwork.
func (mr *MockOpenstackOperationsMockRecorder) GetNetwork(networkname interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetwork", reflect.TypeOf((*MockOpenstackOperations)(nil).GetNetwork), networkname)
}

// GetPort mocks base method.
func (m *MockOpenstackOperations) GetPort(portID string) (*ports.Port, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPort", portID)
	ret0, _ := ret[0].(*ports.Port)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPort indicates an expected call of GetPort.
func (mr *MockOpenstackOperationsMockRecorder) GetPort(portID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPort", reflect.TypeOf((*MockOpenstackOperations)(nil).GetPort), portID)
}

// SetVolumeBootable mocks base method.
func (m *MockOpenstackOperations) SetVolumeBootable(volume *volumes.Volume) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetVolumeBootable", volume)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetVolumeBootable indicates an expected call of SetVolumeBootable.
func (mr *MockOpenstackOperationsMockRecorder) SetVolumeBootable(volume interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetVolumeBootable", reflect.TypeOf((*MockOpenstackOperations)(nil).SetVolumeBootable), volume)
}

// SetVolumeImageMetadata mocks base method.
func (m *MockOpenstackOperations) SetVolumeImageMetadata(volume *volumes.Volume) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetVolumeImageMetadata", volume)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetVolumeImageMetadata indicates an expected call of SetVolumeImageMetadata.
func (mr *MockOpenstackOperationsMockRecorder) SetVolumeImageMetadata(volume interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetVolumeImageMetadata", reflect.TypeOf((*MockOpenstackOperations)(nil).SetVolumeImageMetadata), volume)
}

// SetVolumeUEFI mocks base method.
func (m *MockOpenstackOperations) SetVolumeUEFI(volume *volumes.Volume) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetVolumeUEFI", volume)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetVolumeUEFI indicates an expected call of SetVolumeUEFI.
func (mr *MockOpenstackOperationsMockRecorder) SetVolumeUEFI(volume interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetVolumeUEFI", reflect.TypeOf((*MockOpenstackOperations)(nil).SetVolumeUEFI), volume)
}

// WaitForVolume mocks base method.
func (m *MockOpenstackOperations) WaitForVolume(volumeID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForVolume", volumeID)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForVolume indicates an expected call of WaitForVolume.
func (mr *MockOpenstackOperationsMockRecorder) WaitForVolume(volumeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForVolume", reflect.TypeOf((*MockOpenstackOperations)(nil).WaitForVolume), volumeID)
}

// WaitForVolumeAttachment mocks base method.
func (m *MockOpenstackOperations) WaitForVolumeAttachment(volumeID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForVolumeAttachment", volumeID)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForVolumeAttachment indicates an expected call of WaitForVolumeAttachment.
func (mr *MockOpenstackOperationsMockRecorder) WaitForVolumeAttachment(volumeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForVolumeAttachment", reflect.TypeOf((*MockOpenstackOperations)(nil).WaitForVolumeAttachment), volumeID)
}
