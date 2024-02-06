// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ExocoreNetwork/exocore-sdk/chainio/avsregistry (interfaces: AvsRegistryWriter)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/avsRegistryContractsWriter.go -package=mocks github.com/ExocoreNetwork/exocore-sdk/chainio/avsregistry AvsRegistryWriter
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	contractBLSPubkeyRegistry "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/BLSPubkeyRegistry"
	contractBLSRegistryCoordinatorWithIndices "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/BLSRegistryCoordinatorWithIndices"
	common "github.com/ethereum/go-ethereum/common"
	types "github.com/ethereum/go-ethereum/core/types"
	gomock "go.uber.org/mock/gomock"
)

// MockAvsRegistryWriter is a mock of AvsRegistryWriter interface.
type MockAvsRegistryWriter struct {
	ctrl     *gomock.Controller
	recorder *MockAvsRegistryWriterMockRecorder
}

// MockAvsRegistryWriterMockRecorder is the mock recorder for MockAvsRegistryWriter.
type MockAvsRegistryWriterMockRecorder struct {
	mock *MockAvsRegistryWriter
}

// NewMockAvsRegistryWriter creates a new mock instance.
func NewMockAvsRegistryWriter(ctrl *gomock.Controller) *MockAvsRegistryWriter {
	mock := &MockAvsRegistryWriter{ctrl: ctrl}
	mock.recorder = &MockAvsRegistryWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAvsRegistryWriter) EXPECT() *MockAvsRegistryWriterMockRecorder {
	return m.recorder
}

// DeregisterOperator mocks base method.
func (m *MockAvsRegistryWriter) DeregisterOperator(arg0 context.Context, arg1 common.Address, arg2 []byte, arg3 contractBLSPubkeyRegistry.BN254G1Point) (*types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeregisterOperator", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*types.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeregisterOperator indicates an expected call of DeregisterOperator.
func (mr *MockAvsRegistryWriterMockRecorder) DeregisterOperator(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeregisterOperator", reflect.TypeOf((*MockAvsRegistryWriter)(nil).DeregisterOperator), arg0, arg1, arg2, arg3)
}

// RegisterOperatorWithAVSRegistryCoordinator mocks base method.
func (m *MockAvsRegistryWriter) RegisterOperatorWithAVSRegistryCoordinator(arg0 context.Context, arg1 []byte, arg2 contractBLSRegistryCoordinatorWithIndices.BN254G1Point, arg3 string) (*types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterOperatorWithAVSRegistryCoordinator", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*types.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterOperatorWithAVSRegistryCoordinator indicates an expected call of RegisterOperatorWithAVSRegistryCoordinator.
func (mr *MockAvsRegistryWriterMockRecorder) RegisterOperatorWithAVSRegistryCoordinator(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterOperatorWithAVSRegistryCoordinator", reflect.TypeOf((*MockAvsRegistryWriter)(nil).RegisterOperatorWithAVSRegistryCoordinator), arg0, arg1, arg2, arg3)
}

// UpdateStakes mocks base method.
func (m *MockAvsRegistryWriter) UpdateStakes(arg0 context.Context, arg1 []common.Address) (*types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStakes", arg0, arg1)
	ret0, _ := ret[0].(*types.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStakes indicates an expected call of UpdateStakes.
func (mr *MockAvsRegistryWriterMockRecorder) UpdateStakes(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStakes", reflect.TypeOf((*MockAvsRegistryWriter)(nil).UpdateStakes), arg0, arg1)
}
