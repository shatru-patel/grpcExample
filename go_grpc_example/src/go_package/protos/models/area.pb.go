//
//Nudm_SDM
//
//Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
//
//The version of the OpenAPI document: 2.1.7
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0-devel
// 	protoc        v3.6.1
// source: area.proto

package models

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Area struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tacs     []string `protobuf:"bytes,3552157,rep,name=tacs,proto3" json:"tacs,omitempty"`
	AreaCode string   `protobuf:"bytes,56040425,opt,name=areaCode,proto3" json:"areaCode,omitempty"`
}

func (x *Area) Reset() {
	*x = Area{}
	if protoimpl.UnsafeEnabled {
		mi := &file_area_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Area) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Area) ProtoMessage() {}

func (x *Area) ProtoReflect() protoreflect.Message {
	mi := &file_area_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Area.ProtoReflect.Descriptor instead.
func (*Area) Descriptor() ([]byte, []int) {
	return file_area_proto_rawDescGZIP(), []int{0}
}

func (x *Area) GetTacs() []string {
	if x != nil {
		return x.Tacs
	}
	return nil
}

func (x *Area) GetAreaCode() string {
	if x != nil {
		return x.AreaCode
	}
	return ""
}

var File_area_proto protoreflect.FileDescriptor

var file_area_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x72, 0x65, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x67, 0x6f,
	0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22,
	0x3c, 0x0a, 0x04, 0x41, 0x72, 0x65, 0x61, 0x12, 0x15, 0x0a, 0x04, 0x74, 0x61, 0x63, 0x73, 0x18,
	0x9d, 0xe7, 0xd8, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x63, 0x73, 0x12, 0x1d,
	0x0a, 0x08, 0x61, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x18, 0xe9, 0xb7, 0xdc, 0x1a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x1a, 0x5a,
	0x18, 0x67, 0x6f, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_area_proto_rawDescOnce sync.Once
	file_area_proto_rawDescData = file_area_proto_rawDesc
)

func file_area_proto_rawDescGZIP() []byte {
	file_area_proto_rawDescOnce.Do(func() {
		file_area_proto_rawDescData = protoimpl.X.CompressGZIP(file_area_proto_rawDescData)
	})
	return file_area_proto_rawDescData
}

var file_area_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_area_proto_goTypes = []interface{}{
	(*Area)(nil), // 0: go_package.protos.Area
}
var file_area_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_area_proto_init() }
func file_area_proto_init() {
	if File_area_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_area_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Area); i {
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
			RawDescriptor: file_area_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_area_proto_goTypes,
		DependencyIndexes: file_area_proto_depIdxs,
		MessageInfos:      file_area_proto_msgTypes,
	}.Build()
	File_area_proto = out.File
	file_area_proto_rawDesc = nil
	file_area_proto_goTypes = nil
	file_area_proto_depIdxs = nil
}
