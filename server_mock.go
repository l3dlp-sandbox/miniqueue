// Code generated by MockGen. DO NOT EDIT.
// Source: server.go

// Package main is a generated GoMock package.
package main

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// Mockbrokerer is a mock of brokerer interface.
type Mockbrokerer struct {
	ctrl     *gomock.Controller
	recorder *MockbrokererMockRecorder
}

// MockbrokererMockRecorder is the mock recorder for Mockbrokerer.
type MockbrokererMockRecorder struct {
	mock *Mockbrokerer
}

// NewMockbrokerer creates a new mock instance.
func NewMockbrokerer(ctrl *gomock.Controller) *Mockbrokerer {
	mock := &Mockbrokerer{ctrl: ctrl}
	mock.recorder = &MockbrokererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockbrokerer) EXPECT() *MockbrokererMockRecorder {
	return m.recorder
}

// Publish mocks base method.
func (m *Mockbrokerer) Publish(topic string, value value) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", topic, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockbrokererMockRecorder) Publish(topic, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*Mockbrokerer)(nil).Publish), topic, value)
}

// Subscribe mocks base method.
func (m *Mockbrokerer) Subscribe(topic string) *consumer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", topic)
	ret0, _ := ret[0].(*consumer)
	return ret0
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockbrokererMockRecorder) Subscribe(topic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*Mockbrokerer)(nil).Subscribe), topic)
}
