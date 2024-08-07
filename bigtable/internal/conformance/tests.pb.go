// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: tests.proto

package conformance

import (
	reflect "reflect"
	sync "sync"

	v2 "cloud.google.com/go/bigtable/apiv2/bigtablepb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TestFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReadRowsTests []*ReadRowsTest `protobuf:"bytes,1,rep,name=read_rows_tests,json=readRowsTests,proto3" json:"read_rows_tests,omitempty"`
}

func (x *TestFile) Reset() {
	*x = TestFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestFile) ProtoMessage() {}

func (x *TestFile) ProtoReflect() protoreflect.Message {
	mi := &file_tests_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestFile.ProtoReflect.Descriptor instead.
func (*TestFile) Descriptor() ([]byte, []int) {
	return file_tests_proto_rawDescGZIP(), []int{0}
}

func (x *TestFile) GetReadRowsTests() []*ReadRowsTest {
	if x != nil {
		return x.ReadRowsTests
	}
	return nil
}

type ReadRowsTest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string                           `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Chunks      []*v2.ReadRowsResponse_CellChunk `protobuf:"bytes,2,rep,name=chunks,proto3" json:"chunks,omitempty"`
	Results     []*ReadRowsTest_Result           `protobuf:"bytes,3,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *ReadRowsTest) Reset() {
	*x = ReadRowsTest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRowsTest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRowsTest) ProtoMessage() {}

func (x *ReadRowsTest) ProtoReflect() protoreflect.Message {
	mi := &file_tests_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRowsTest.ProtoReflect.Descriptor instead.
func (*ReadRowsTest) Descriptor() ([]byte, []int) {
	return file_tests_proto_rawDescGZIP(), []int{1}
}

func (x *ReadRowsTest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ReadRowsTest) GetChunks() []*v2.ReadRowsResponse_CellChunk {
	if x != nil {
		return x.Chunks
	}
	return nil
}

func (x *ReadRowsTest) GetResults() []*ReadRowsTest_Result {
	if x != nil {
		return x.Results
	}
	return nil
}

// Expected results of reading the row.
// Only the last result can be an error.
type ReadRowsTest_Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RowKey          string `protobuf:"bytes,1,opt,name=row_key,json=rowKey,proto3" json:"row_key,omitempty"`
	FamilyName      string `protobuf:"bytes,2,opt,name=family_name,json=familyName,proto3" json:"family_name,omitempty"`
	Qualifier       string `protobuf:"bytes,3,opt,name=qualifier,proto3" json:"qualifier,omitempty"`
	TimestampMicros int64  `protobuf:"varint,4,opt,name=timestamp_micros,json=timestampMicros,proto3" json:"timestamp_micros,omitempty"`
	Value           string `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	Label           string `protobuf:"bytes,6,opt,name=label,proto3" json:"label,omitempty"`
	Error           bool   `protobuf:"varint,7,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ReadRowsTest_Result) Reset() {
	*x = ReadRowsTest_Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRowsTest_Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRowsTest_Result) ProtoMessage() {}

func (x *ReadRowsTest_Result) ProtoReflect() protoreflect.Message {
	mi := &file_tests_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRowsTest_Result.ProtoReflect.Descriptor instead.
func (*ReadRowsTest_Result) Descriptor() ([]byte, []int) {
	return file_tests_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ReadRowsTest_Result) GetRowKey() string {
	if x != nil {
		return x.RowKey
	}
	return ""
}

func (x *ReadRowsTest_Result) GetFamilyName() string {
	if x != nil {
		return x.FamilyName
	}
	return ""
}

func (x *ReadRowsTest_Result) GetQualifier() string {
	if x != nil {
		return x.Qualifier
	}
	return ""
}

func (x *ReadRowsTest_Result) GetTimestampMicros() int64 {
	if x != nil {
		return x.TimestampMicros
	}
	return 0
}

func (x *ReadRowsTest_Result) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *ReadRowsTest_Result) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *ReadRowsTest_Result) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

var File_tests_proto protoreflect.FileDescriptor

var file_tests_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x2e, 0x76, 0x32, 0x1a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x62, 0x69, 0x67, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x08, 0x54, 0x65, 0x73, 0x74, 0x46, 0x69,
	0x6c, 0x65, 0x12, 0x5a, 0x0a, 0x0f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x72, 0x6f, 0x77, 0x73, 0x5f,
	0x74, 0x65, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e,
	0x76, 0x32, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x6f, 0x77, 0x73, 0x54, 0x65, 0x73, 0x74, 0x52,
	0x0d, 0x72, 0x65, 0x61, 0x64, 0x52, 0x6f, 0x77, 0x73, 0x54, 0x65, 0x73, 0x74, 0x73, 0x22, 0x9d,
	0x03, 0x0a, 0x0c, 0x52, 0x65, 0x61, 0x64, 0x52, 0x6f, 0x77, 0x73, 0x54, 0x65, 0x73, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x46, 0x0a, 0x06, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x6f, 0x77, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x65, 0x6c, 0x6c, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x52, 0x06, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x53, 0x0a, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76,
	0x32, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x6f, 0x77, 0x73, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x1a, 0xcd,
	0x01, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x77,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x77, 0x4b,
	0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x9a,
	0x01, 0x0a, 0x28, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e,
	0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x32, 0x42, 0x0e, 0x54, 0x65, 0x73,
	0x74, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5a, 0x31, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2f, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0xaa, 0x02,
	0x2a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x42, 0x69,
	0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x56, 0x32, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x73, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_tests_proto_rawDescOnce sync.Once
	file_tests_proto_rawDescData = file_tests_proto_rawDesc
)

func file_tests_proto_rawDescGZIP() []byte {
	file_tests_proto_rawDescOnce.Do(func() {
		file_tests_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_proto_rawDescData)
	})
	return file_tests_proto_rawDescData
}

var file_tests_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tests_proto_goTypes = []interface{}{
	(*TestFile)(nil),                      // 0: google.cloud.conformance.bigtable.v2.TestFile
	(*ReadRowsTest)(nil),                  // 1: google.cloud.conformance.bigtable.v2.ReadRowsTest
	(*ReadRowsTest_Result)(nil),           // 2: google.cloud.conformance.bigtable.v2.ReadRowsTest.Result
	(*v2.ReadRowsResponse_CellChunk)(nil), // 3: google.bigtable.v2.ReadRowsResponse.CellChunk
}
var file_tests_proto_depIdxs = []int32{
	1, // 0: google.cloud.conformance.bigtable.v2.TestFile.read_rows_tests:type_name -> google.cloud.conformance.bigtable.v2.ReadRowsTest
	3, // 1: google.cloud.conformance.bigtable.v2.ReadRowsTest.chunks:type_name -> google.bigtable.v2.ReadRowsResponse.CellChunk
	2, // 2: google.cloud.conformance.bigtable.v2.ReadRowsTest.results:type_name -> google.cloud.conformance.bigtable.v2.ReadRowsTest.Result
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tests_proto_init() }
func file_tests_proto_init() {
	if File_tests_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tests_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestFile); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tests_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRowsTest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tests_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRowsTest_Result); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tests_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tests_proto_goTypes,
		DependencyIndexes: file_tests_proto_depIdxs,
		MessageInfos:      file_tests_proto_msgTypes,
	}.Build()
	File_tests_proto = out.File
	file_tests_proto_rawDesc = nil
	file_tests_proto_goTypes = nil
	file_tests_proto_depIdxs = nil
}
