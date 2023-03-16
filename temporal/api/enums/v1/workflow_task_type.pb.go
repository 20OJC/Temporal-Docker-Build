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

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/server/api/enums/v1/workflow_task_type.proto

package enums

import (
	fmt "fmt"
	math "math"
	strconv "strconv"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type WorkflowTaskType int32

const (
	WORKFLOW_TASK_TYPE_UNSPECIFIED WorkflowTaskType = 0
	WORKFLOW_TASK_TYPE_NORMAL      WorkflowTaskType = 1
	// TODO (alex): TRANSIENT is not current used. Needs to be set when Attempt>1.
	WORKFLOW_TASK_TYPE_TRANSIENT   WorkflowTaskType = 2
	WORKFLOW_TASK_TYPE_SPECULATIVE WorkflowTaskType = 3
)

var WorkflowTaskType_name = map[int32]string{
	0: "Unspecified",
	1: "Normal",
	2: "Transient",
	3: "Speculative",
}

var WorkflowTaskType_value = map[string]int32{
	"Unspecified": 0,
	"Normal":      1,
	"Transient":   2,
	"Speculative": 3,
}

func (WorkflowTaskType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c9364840fefcb74c, []int{0}
}

func init() {
	proto.RegisterEnum("temporal.server.api.enums.v1.WorkflowTaskType", WorkflowTaskType_name, WorkflowTaskType_value)
}

func init() {
	proto.RegisterFile("temporal/server/api/enums/v1/workflow_task_type.proto", fileDescriptor_c9364840fefcb74c)
}

var fileDescriptor_c9364840fefcb74c = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x2d, 0x49, 0xcd, 0x2d,
	0xc8, 0x2f, 0x4a, 0xcc, 0xd1, 0x2f, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x4f, 0x2c, 0xc8, 0xd4,
	0x4f, 0xcd, 0x2b, 0xcd, 0x2d, 0xd6, 0x2f, 0x33, 0xd4, 0x2f, 0xcf, 0x2f, 0xca, 0x4e, 0xcb, 0xc9,
	0x2f, 0x8f, 0x2f, 0x49, 0x2c, 0xce, 0x8e, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x92, 0x81, 0x69, 0xd3, 0x83, 0x68, 0xd3, 0x4b, 0x2c, 0xc8, 0xd4, 0x03, 0x6b, 0xd3,
	0x2b, 0x33, 0xd4, 0x9a, 0xcd, 0xc8, 0x25, 0x10, 0x0e, 0xd5, 0x1a, 0x92, 0x58, 0x9c, 0x1d, 0x52,
	0x59, 0x90, 0x2a, 0xa4, 0xc4, 0x25, 0x17, 0xee, 0x1f, 0xe4, 0xed, 0xe6, 0xe3, 0x1f, 0x1e, 0x1f,
	0xe2, 0x18, 0xec, 0x1d, 0x1f, 0x12, 0x19, 0xe0, 0x1a, 0x1f, 0xea, 0x17, 0x1c, 0xe0, 0xea, 0xec,
	0xe9, 0xe6, 0xe9, 0xea, 0x22, 0xc0, 0x20, 0x24, 0xcb, 0x25, 0x89, 0x45, 0x8d, 0x9f, 0x7f, 0x90,
	0xaf, 0xa3, 0x8f, 0x00, 0xa3, 0x90, 0x02, 0x97, 0x0c, 0x16, 0xe9, 0x90, 0x20, 0x47, 0xbf, 0x60,
	0x4f, 0x57, 0xbf, 0x10, 0x01, 0x26, 0x1c, 0x96, 0x80, 0xac, 0x08, 0xf5, 0x71, 0x0c, 0xf1, 0x0c,
	0x73, 0x15, 0x60, 0x76, 0x8a, 0xbb, 0xf0, 0x50, 0x8e, 0xe1, 0xc6, 0x43, 0x39, 0x86, 0x0f, 0x0f,
	0xe5, 0x18, 0x1b, 0x1e, 0xc9, 0x31, 0xae, 0x78, 0x24, 0xc7, 0x78, 0xe2, 0x91, 0x1c, 0xe3, 0x85,
	0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0xbe, 0x78, 0x24, 0xc7, 0xf0, 0xe1, 0x91, 0x1c, 0xe3,
	0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x91, 0x9e,
	0xaf, 0x07, 0xf7, 0x74, 0x66, 0x3e, 0xb6, 0xe0, 0xb2, 0x06, 0x33, 0x92, 0xd8, 0xc0, 0x41, 0x64,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x11, 0xba, 0xfa, 0x5b, 0x01, 0x00, 0x00,
}

func (x WorkflowTaskType) String() string {
	s, ok := WorkflowTaskType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
