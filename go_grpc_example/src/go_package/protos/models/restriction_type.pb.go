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
// source: restriction_type.proto

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

type RestrictionType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RestrictionType string `protobuf:"bytes,1,opt,name=restrictionType,proto3" json:"restrictionType,omitempty"`
}

func (x *RestrictionType) Reset() {
	*x = RestrictionType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restriction_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestrictionType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestrictionType) ProtoMessage() {}

func (x *RestrictionType) ProtoReflect() protoreflect.Message {
	mi := &file_restriction_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestrictionType.ProtoReflect.Descriptor instead.
func (*RestrictionType) Descriptor() ([]byte, []int) {
	return file_restriction_type_proto_rawDescGZIP(), []int{0}
}

func (x *RestrictionType) GetRestrictionType() string {
	if x != nil {
		return x.RestrictionType
	}
	return ""
}

var File_restriction_type_proto protoreflect.FileDescriptor

var file_restriction_type_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x67, 0x6f, 0x5f, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x3b, 0x0a, 0x0f, 0x52,
	0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28,
	0x0a, 0x0f, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x6f, 0x5f, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_restriction_type_proto_rawDescOnce sync.Once
	file_restriction_type_proto_rawDescData = file_restriction_type_proto_rawDesc
)

func file_restriction_type_proto_rawDescGZIP() []byte {
	file_restriction_type_proto_rawDescOnce.Do(func() {
		file_restriction_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_restriction_type_proto_rawDescData)
	})
	return file_restriction_type_proto_rawDescData
}

var file_restriction_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_restriction_type_proto_goTypes = []interface{}{
	(*RestrictionType)(nil), // 0: go_package.protos.RestrictionType
}
var file_restriction_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_restriction_type_proto_init() }
func file_restriction_type_proto_init() {
	if File_restriction_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_restriction_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestrictionType); i {
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
			RawDescriptor: file_restriction_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_restriction_type_proto_goTypes,
		DependencyIndexes: file_restriction_type_proto_depIdxs,
		MessageInfos:      file_restriction_type_proto_msgTypes,
	}.Build()
	File_restriction_type_proto = out.File
	file_restriction_type_proto_rawDesc = nil
	file_restriction_type_proto_goTypes = nil
	file_restriction_type_proto_depIdxs = nil
}
