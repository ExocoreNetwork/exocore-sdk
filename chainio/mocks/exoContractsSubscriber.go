// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ExocoreNetwork/exocore-sdk/chainio/exocontracts (interfaces: ELSubscriber)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/exoContractsSubscriber.go -package=mocks github.com/ExocoreNetwork/exocore-sdk/chainio/exocontracts ELSubscriber
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	contractBLSPublicKeyCompendium "github.com/ExocoreNetwork/exocore-sdk/contracts/bindings/BLSPublicKeyCompendium"
	event "github.com/ethereum/go-ethereum/event"
	gomock "go.uber.org/mock/gomock"
)

// MockELSubscriber is a mock of ELSubscriber interface.
type MockELSubscriber struct {
	ctrl     *gomock.Controller
	recorder *MockELSubscriberMockRecorder
}

// MockELSubscriberMockRecorder is the mock recorder for MockELSubscriber.
type MockELSubscriberMockRecorder struct {
	mock *MockELSubscriber
}

// NewMockELSubscriber creates a new mock instance.
func NewMockELSubscriber(ctrl *gomock.Controller) *MockELSubscriber {
	mock := &MockELSubscriber{ctrl: ctrl}
	mock.recorder = &MockELSubscriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockELSubscriber) EXPECT() *MockELSubscriberMockRecorder {
	return m.recorder
}

// SubscribeToNewPubkeyRegistrations mocks base method.
func (m *MockELSubscriber) SubscribeToNewPubkeyRegistrations() (chan *contractBLSPublicKeyCompendium.ContractBLSPublicKeyCompendiumNewPubkeyRegistration, event.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeToNewPubkeyRegistrations")
	ret0, _ := ret[0].(chan *contractBLSPublicKeyCompendium.ContractBLSPublicKeyCompendiumNewPubkeyRegistration)
	ret1, _ := ret[1].(event.Subscription)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SubscribeToNewPubkeyRegistrations indicates an expected call of SubscribeToNewPubkeyRegistrations.
func (mr *MockELSubscriberMockRecorder) SubscribeToNewPubkeyRegistrations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToNewPubkeyRegistrations", reflect.TypeOf((*MockELSubscriber)(nil).SubscribeToNewPubkeyRegistrations))
}
