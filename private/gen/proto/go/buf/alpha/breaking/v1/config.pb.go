// Copyright 2020-2024 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: buf/alpha/breaking/v1/config.proto

package breakingv1

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

// Config represents the breaking change configuration for a module. The rule and category IDs are defined
// by the version and apply across the config. The version is independent of the version of
// the package. The package version refers to the config shape, the version encoded in the Config message
// indicates which rule and category IDs should be used.
//
// The rule and category IDs are not encoded as enums in this package because we may want to support custom rule
// and category IDs in the future. Callers will need to resolve the rule and category ID strings.
type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// version represents the version of the breaking change rule and category IDs that should be used with this config.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// use_ids lists the rule and/or category IDs that are included in the breaking change check.
	UseIds []string `protobuf:"bytes,2,rep,name=use_ids,json=useIds,proto3" json:"use_ids,omitempty"`
	// except_ids lists the rule and/or category IDs that are excluded from the breaking change check.
	ExceptIds []string `protobuf:"bytes,3,rep,name=except_ids,json=exceptIds,proto3" json:"except_ids,omitempty"`
	// ignore_paths lists the paths of directories and/or files that should be ignored by the breaking change check.
	// All paths are relative to the root of the module.
	IgnorePaths []string `protobuf:"bytes,4,rep,name=ignore_paths,json=ignorePaths,proto3" json:"ignore_paths,omitempty"`
	// ignore_id_paths is a map of rule and/or category IDs to directory and/or file paths to exclude from the
	// breaking change check. This corresponds with the ignore_only configuration key.
	IgnoreIdPaths []*IDPaths `protobuf:"bytes,5,rep,name=ignore_id_paths,json=ignoreIdPaths,proto3" json:"ignore_id_paths,omitempty"`
	// ignore_unstable_packages ignores packages with a last component that is one of the unstable forms recognised
	// by the PACKAGE_VERSION_SUFFIX:
	//
	//	v\d+test.*
	//	v\d+(alpha|beta)\d+
	//	v\d+p\d+(alpha|beta)\d+
	IgnoreUnstablePackages bool `protobuf:"varint,6,opt,name=ignore_unstable_packages,json=ignoreUnstablePackages,proto3" json:"ignore_unstable_packages,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_breaking_v1_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_breaking_v1_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_buf_alpha_breaking_v1_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Config) GetUseIds() []string {
	if x != nil {
		return x.UseIds
	}
	return nil
}

func (x *Config) GetExceptIds() []string {
	if x != nil {
		return x.ExceptIds
	}
	return nil
}

func (x *Config) GetIgnorePaths() []string {
	if x != nil {
		return x.IgnorePaths
	}
	return nil
}

func (x *Config) GetIgnoreIdPaths() []*IDPaths {
	if x != nil {
		return x.IgnoreIdPaths
	}
	return nil
}

func (x *Config) GetIgnoreUnstablePackages() bool {
	if x != nil {
		return x.IgnoreUnstablePackages
	}
	return false
}

// IDPaths represents a rule or category ID and the file and/or directory paths that are ignored for the rule.
type IDPaths struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Paths []string `protobuf:"bytes,2,rep,name=paths,proto3" json:"paths,omitempty"`
}

func (x *IDPaths) Reset() {
	*x = IDPaths{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_breaking_v1_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDPaths) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDPaths) ProtoMessage() {}

func (x *IDPaths) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_breaking_v1_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDPaths.ProtoReflect.Descriptor instead.
func (*IDPaths) Descriptor() ([]byte, []int) {
	return file_buf_alpha_breaking_v1_config_proto_rawDescGZIP(), []int{1}
}

func (x *IDPaths) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *IDPaths) GetPaths() []string {
	if x != nil {
		return x.Paths
	}
	return nil
}

var File_buf_alpha_breaking_v1_config_proto protoreflect.FileDescriptor

var file_buf_alpha_breaking_v1_config_proto_rawDesc = []byte{
	0x0a, 0x22, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x62, 0x72, 0x65, 0x61,
	0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x62, 0x72, 0x65, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x22, 0xff, 0x01, 0x0a, 0x06,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x49, 0x64, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x63,
	0x65, 0x70, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x65,
	0x78, 0x63, 0x65, 0x70, 0x74, 0x49, 0x64, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x67, 0x6e, 0x6f,
	0x72, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b,
	0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x50, 0x61, 0x74, 0x68, 0x73, 0x12, 0x46, 0x0a, 0x0f, 0x69,
	0x67, 0x6e, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x44, 0x50,
	0x61, 0x74, 0x68, 0x73, 0x52, 0x0d, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x50, 0x61,
	0x74, 0x68, 0x73, 0x12, 0x38, 0x0a, 0x18, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x5f, 0x75, 0x6e,
	0x73, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x55, 0x6e, 0x73,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x22, 0x2f, 0x0a,
	0x07, 0x49, 0x44, 0x50, 0x61, 0x74, 0x68, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x68,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x42, 0xee,
	0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x2f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x3b,
	0x62, 0x72, 0x65, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x41, 0x42,
	0xaa, 0x02, 0x15, 0x42, 0x75, 0x66, 0x2e, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x42, 0x72, 0x65,
	0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x15, 0x42, 0x75, 0x66, 0x5c, 0x41,
	0x6c, 0x70, 0x68, 0x61, 0x5c, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x21, 0x42, 0x75, 0x66, 0x5c, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x5c, 0x42, 0x72, 0x65,
	0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x18, 0x42, 0x75, 0x66, 0x3a, 0x3a, 0x41, 0x6c, 0x70, 0x68,
	0x61, 0x3a, 0x3a, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buf_alpha_breaking_v1_config_proto_rawDescOnce sync.Once
	file_buf_alpha_breaking_v1_config_proto_rawDescData = file_buf_alpha_breaking_v1_config_proto_rawDesc
)

func file_buf_alpha_breaking_v1_config_proto_rawDescGZIP() []byte {
	file_buf_alpha_breaking_v1_config_proto_rawDescOnce.Do(func() {
		file_buf_alpha_breaking_v1_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_alpha_breaking_v1_config_proto_rawDescData)
	})
	return file_buf_alpha_breaking_v1_config_proto_rawDescData
}

var file_buf_alpha_breaking_v1_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_buf_alpha_breaking_v1_config_proto_goTypes = []any{
	(*Config)(nil),  // 0: buf.alpha.breaking.v1.Config
	(*IDPaths)(nil), // 1: buf.alpha.breaking.v1.IDPaths
}
var file_buf_alpha_breaking_v1_config_proto_depIdxs = []int32{
	1, // 0: buf.alpha.breaking.v1.Config.ignore_id_paths:type_name -> buf.alpha.breaking.v1.IDPaths
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_buf_alpha_breaking_v1_config_proto_init() }
func file_buf_alpha_breaking_v1_config_proto_init() {
	if File_buf_alpha_breaking_v1_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_buf_alpha_breaking_v1_config_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Config); i {
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
		file_buf_alpha_breaking_v1_config_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*IDPaths); i {
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
			RawDescriptor: file_buf_alpha_breaking_v1_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_alpha_breaking_v1_config_proto_goTypes,
		DependencyIndexes: file_buf_alpha_breaking_v1_config_proto_depIdxs,
		MessageInfos:      file_buf_alpha_breaking_v1_config_proto_msgTypes,
	}.Build()
	File_buf_alpha_breaking_v1_config_proto = out.File
	file_buf_alpha_breaking_v1_config_proto_rawDesc = nil
	file_buf_alpha_breaking_v1_config_proto_goTypes = nil
	file_buf_alpha_breaking_v1_config_proto_depIdxs = nil
}
