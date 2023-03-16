// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package common is a generated GoMock package.
package common

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	worker "go.temporal.io/sdk/worker"
	namespace "go.temporal.io/server/common/namespace"
)

// MockWorkerComponent is a mock of WorkerComponent interface.
type MockWorkerComponent struct {
	ctrl     *gomock.Controller
	recorder *MockWorkerComponentMockRecorder
}

// MockWorkerComponentMockRecorder is the mock recorder for MockWorkerComponent.
type MockWorkerComponentMockRecorder struct {
	mock *MockWorkerComponent
}

// NewMockWorkerComponent creates a new mock instance.
func NewMockWorkerComponent(ctrl *gomock.Controller) *MockWorkerComponent {
	mock := &MockWorkerComponent{ctrl: ctrl}
	mock.recorder = &MockWorkerComponentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkerComponent) EXPECT() *MockWorkerComponentMockRecorder {
	return m.recorder
}

// DedicatedWorkerOptions mocks base method.
func (m *MockWorkerComponent) DedicatedWorkerOptions() *DedicatedWorkerOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DedicatedWorkerOptions")
	ret0, _ := ret[0].(*DedicatedWorkerOptions)
	return ret0
}

// DedicatedWorkerOptions indicates an expected call of DedicatedWorkerOptions.
func (mr *MockWorkerComponentMockRecorder) DedicatedWorkerOptions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DedicatedWorkerOptions", reflect.TypeOf((*MockWorkerComponent)(nil).DedicatedWorkerOptions))
}

// Register mocks base method.
func (m *MockWorkerComponent) Register(arg0 worker.Worker) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Register", arg0)
}

// Register indicates an expected call of Register.
func (mr *MockWorkerComponentMockRecorder) Register(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockWorkerComponent)(nil).Register), arg0)
}

// MockPerNSWorkerComponent is a mock of PerNSWorkerComponent interface.
type MockPerNSWorkerComponent struct {
	ctrl     *gomock.Controller
	recorder *MockPerNSWorkerComponentMockRecorder
}

// MockPerNSWorkerComponentMockRecorder is the mock recorder for MockPerNSWorkerComponent.
type MockPerNSWorkerComponentMockRecorder struct {
	mock *MockPerNSWorkerComponent
}

// NewMockPerNSWorkerComponent creates a new mock instance.
func NewMockPerNSWorkerComponent(ctrl *gomock.Controller) *MockPerNSWorkerComponent {
	mock := &MockPerNSWorkerComponent{ctrl: ctrl}
	mock.recorder = &MockPerNSWorkerComponentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPerNSWorkerComponent) EXPECT() *MockPerNSWorkerComponentMockRecorder {
	return m.recorder
}

// DedicatedWorkerOptions mocks base method.
func (m *MockPerNSWorkerComponent) DedicatedWorkerOptions(arg0 *namespace.Namespace) *PerNSDedicatedWorkerOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DedicatedWorkerOptions", arg0)
	ret0, _ := ret[0].(*PerNSDedicatedWorkerOptions)
	return ret0
}

// DedicatedWorkerOptions indicates an expected call of DedicatedWorkerOptions.
func (mr *MockPerNSWorkerComponentMockRecorder) DedicatedWorkerOptions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DedicatedWorkerOptions", reflect.TypeOf((*MockPerNSWorkerComponent)(nil).DedicatedWorkerOptions), arg0)
}

// Register mocks base method.
func (m *MockPerNSWorkerComponent) Register(arg0 worker.Worker, arg1 *namespace.Namespace, arg2 RegistrationDetails) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Register", arg0, arg1, arg2)
}

// Register indicates an expected call of Register.
func (mr *MockPerNSWorkerComponentMockRecorder) Register(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockPerNSWorkerComponent)(nil).Register), arg0, arg1, arg2)
}