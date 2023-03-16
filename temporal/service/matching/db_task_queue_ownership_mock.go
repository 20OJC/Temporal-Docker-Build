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
// Source: db_task_queue_ownership.go

// Package matching is a generated GoMock package.
package matching

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	persistence "go.temporal.io/server/api/persistence/v1"
)

// MockdbTaskQueueOwnership is a mock of dbTaskQueueOwnership interface.
type MockdbTaskQueueOwnership struct {
	ctrl     *gomock.Controller
	recorder *MockdbTaskQueueOwnershipMockRecorder
}

// MockdbTaskQueueOwnershipMockRecorder is the mock recorder for MockdbTaskQueueOwnership.
type MockdbTaskQueueOwnershipMockRecorder struct {
	mock *MockdbTaskQueueOwnership
}

// NewMockdbTaskQueueOwnership creates a new mock instance.
func NewMockdbTaskQueueOwnership(ctrl *gomock.Controller) *MockdbTaskQueueOwnership {
	mock := &MockdbTaskQueueOwnership{ctrl: ctrl}
	mock.recorder = &MockdbTaskQueueOwnershipMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockdbTaskQueueOwnership) EXPECT() *MockdbTaskQueueOwnershipMockRecorder {
	return m.recorder
}

// flushTasks mocks base method.
func (m *MockdbTaskQueueOwnership) flushTasks(ctx context.Context, taskInfos ...*persistence.TaskInfo) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range taskInfos {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "flushTasks", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// flushTasks indicates an expected call of flushTasks.
func (mr *MockdbTaskQueueOwnershipMockRecorder) flushTasks(ctx interface{}, taskInfos ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, taskInfos...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "flushTasks", reflect.TypeOf((*MockdbTaskQueueOwnership)(nil).flushTasks), varargs...)
}

// getAckedTaskID mocks base method.
func (m *MockdbTaskQueueOwnership) getAckedTaskID() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getAckedTaskID")
	ret0, _ := ret[0].(int64)
	return ret0
}

// getAckedTaskID indicates an expected call of getAckedTaskID.
func (mr *MockdbTaskQueueOwnershipMockRecorder) getAckedTaskID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getAckedTaskID", reflect.TypeOf((*MockdbTaskQueueOwnership)(nil).getAckedTaskID))
}

// getLastAllocatedTaskID mocks base method.
func (m *MockdbTaskQueueOwnership) getLastAllocatedTaskID() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getLastAllocatedTaskID")
	ret0, _ := ret[0].(int64)
	return ret0
}

// getLastAllocatedTaskID indicates an expected call of getLastAllocatedTaskID.
func (mr *MockdbTaskQueueOwnershipMockRecorder) getLastAllocatedTaskID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getLastAllocatedTaskID", reflect.TypeOf((*MockdbTaskQueueOwnership)(nil).getLastAllocatedTaskID))
}

// getShutdownChan mocks base method.
func (m *MockdbTaskQueueOwnership) getShutdownChan() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getShutdownChan")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// getShutdownChan indicates an expected call of getShutdownChan.
func (mr *MockdbTaskQueueOwnershipMockRecorder) getShutdownChan() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getShutdownChan", reflect.TypeOf((*MockdbTaskQueueOwnership)(nil).getShutdownChan))
}

// persistTaskQueue mocks base method.
func (m *MockdbTaskQueueOwnership) persistTaskQueue(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "persistTaskQueue", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// persistTaskQueue indicates an expected call of persistTaskQueue.
func (mr *MockdbTaskQueueOwnershipMockRecorder) persistTaskQueue(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "persistTaskQueue", reflect.TypeOf((*MockdbTaskQueueOwnership)(nil).persistTaskQueue), ctx)
}

// takeTaskQueueOwnership mocks base method.
func (m *MockdbTaskQueueOwnership) takeTaskQueueOwnership(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "takeTaskQueueOwnership", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// takeTaskQueueOwnership indicates an expected call of takeTaskQueueOwnership.
func (mr *MockdbTaskQueueOwnershipMockRecorder) takeTaskQueueOwnership(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "takeTaskQueueOwnership", reflect.TypeOf((*MockdbTaskQueueOwnership)(nil).takeTaskQueueOwnership), ctx)
}

// updateAckedTaskID mocks base method.
func (m *MockdbTaskQueueOwnership) updateAckedTaskID(taskID int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "updateAckedTaskID", taskID)
}

// updateAckedTaskID indicates an expected call of updateAckedTaskID.
func (mr *MockdbTaskQueueOwnershipMockRecorder) updateAckedTaskID(taskID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "updateAckedTaskID", reflect.TypeOf((*MockdbTaskQueueOwnership)(nil).updateAckedTaskID), taskID)
}
