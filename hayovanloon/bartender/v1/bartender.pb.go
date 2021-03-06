//
// Copyright 2020 Hayo van Loon
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: hayovanloon/bartender/v1/bartender.proto

package bartender

import (
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// The main resource
type Beer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the beer (type)
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are assignable to Packaging:
	//	*Beer_BottleCl
	//	*Beer_GlassCl
	//	*Beer_KegL
	Packaging isBeer_Packaging `protobuf_oneof:"packaging"`
}

func (x *Beer) Reset() {
	*x = Beer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hayovanloon_bartender_v1_bartender_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Beer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Beer) ProtoMessage() {}

func (x *Beer) ProtoReflect() protoreflect.Message {
	mi := &file_hayovanloon_bartender_v1_bartender_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Beer.ProtoReflect.Descriptor instead.
func (*Beer) Descriptor() ([]byte, []int) {
	return file_hayovanloon_bartender_v1_bartender_proto_rawDescGZIP(), []int{0}
}

func (x *Beer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *Beer) GetPackaging() isBeer_Packaging {
	if m != nil {
		return m.Packaging
	}
	return nil
}

func (x *Beer) GetBottleCl() int32 {
	if x, ok := x.GetPackaging().(*Beer_BottleCl); ok {
		return x.BottleCl
	}
	return 0
}

func (x *Beer) GetGlassCl() int32 {
	if x, ok := x.GetPackaging().(*Beer_GlassCl); ok {
		return x.GlassCl
	}
	return 0
}

func (x *Beer) GetKegL() int32 {
	if x, ok := x.GetPackaging().(*Beer_KegL); ok {
		return x.KegL
	}
	return 0
}

type isBeer_Packaging interface {
	isBeer_Packaging()
}

type Beer_BottleCl struct {
	// Bottle size in centiliters
	BottleCl int32 `protobuf:"varint,10,opt,name=bottle_cl,json=bottleCl,proto3,oneof"`
}

type Beer_GlassCl struct {
	// Glass size in centiliters
	GlassCl int32 `protobuf:"varint,11,opt,name=glass_cl,json=glassCl,proto3,oneof"`
}

type Beer_KegL struct {
	// Keg size in liters
	KegL int32 `protobuf:"varint,12,opt,name=keg_l,json=kegL,proto3,oneof"`
}

func (*Beer_BottleCl) isBeer_Packaging() {}

func (*Beer_GlassCl) isBeer_Packaging() {}

func (*Beer_KegL) isBeer_Packaging() {}

type GetBeerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Brand string `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
}

func (x *GetBeerRequest) Reset() {
	*x = GetBeerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hayovanloon_bartender_v1_bartender_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBeerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBeerRequest) ProtoMessage() {}

func (x *GetBeerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hayovanloon_bartender_v1_bartender_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBeerRequest.ProtoReflect.Descriptor instead.
func (*GetBeerRequest) Descriptor() ([]byte, []int) {
	return file_hayovanloon_bartender_v1_bartender_proto_rawDescGZIP(), []int{1}
}

func (x *GetBeerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetBeerRequest) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

type CreateBeerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Brand string `protobuf:"bytes,1,opt,name=brand,proto3" json:"brand,omitempty"`
	Beer  *Beer  `protobuf:"bytes,2,opt,name=beer,proto3" json:"beer,omitempty"`
}

func (x *CreateBeerRequest) Reset() {
	*x = CreateBeerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hayovanloon_bartender_v1_bartender_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBeerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBeerRequest) ProtoMessage() {}

func (x *CreateBeerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hayovanloon_bartender_v1_bartender_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBeerRequest.ProtoReflect.Descriptor instead.
func (*CreateBeerRequest) Descriptor() ([]byte, []int) {
	return file_hayovanloon_bartender_v1_bartender_proto_rawDescGZIP(), []int{2}
}

func (x *CreateBeerRequest) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *CreateBeerRequest) GetBeer() *Beer {
	if x != nil {
		return x.Beer
	}
	return nil
}

var File_hayovanloon_bartender_v1_bartender_proto protoreflect.FileDescriptor

var file_hayovanloon_bartender_v1_bartender_proto_rawDesc = []byte{
	0x0a, 0x28, 0x68, 0x61, 0x79, 0x6f, 0x76, 0x61, 0x6e, 0x6c, 0x6f, 0x6f, 0x6e, 0x2f, 0x62, 0x61,
	0x72, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x72, 0x74, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x68, 0x61, 0x79, 0x6f,
	0x76, 0x61, 0x6e, 0x6c, 0x6f, 0x6f, 0x6e, 0x2e, 0x62, 0x61, 0x72, 0x74, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x04, 0x42, 0x65, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x09, 0x62, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x63, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x00, 0x52, 0x08, 0x62, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x43, 0x6c, 0x12, 0x1b, 0x0a,
	0x08, 0x67, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x63, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x00, 0x52, 0x07, 0x67, 0x6c, 0x61, 0x73, 0x73, 0x43, 0x6c, 0x12, 0x15, 0x0a, 0x05, 0x6b, 0x65,
	0x67, 0x5f, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x04, 0x6b, 0x65, 0x67,
	0x4c, 0x42, 0x0b, 0x0a, 0x09, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x22, 0x3a,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x22, 0x5d, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x62, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x32, 0x0a, 0x04, 0x62, 0x65, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x68, 0x61, 0x79, 0x6f, 0x76, 0x61, 0x6e, 0x6c, 0x6f, 0x6f,
	0x6e, 0x2e, 0x62, 0x61, 0x72, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x65, 0x65, 0x72, 0x52, 0x04, 0x62, 0x65, 0x65, 0x72, 0x32, 0x88, 0x02, 0x0a, 0x09, 0x42, 0x61,
	0x72, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x7f, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x65, 0x65, 0x72, 0x12, 0x2b, 0x2e, 0x68, 0x61, 0x79, 0x6f, 0x76, 0x61, 0x6e, 0x6c,
	0x6f, 0x6f, 0x6e, 0x2e, 0x62, 0x61, 0x72, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x68, 0x61, 0x79, 0x6f, 0x76, 0x61, 0x6e, 0x6c, 0x6f, 0x6f, 0x6e,
	0x2e, 0x62, 0x61, 0x72, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x65,
	0x65, 0x72, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x16, 0x2f, 0x76, 0x31, 0x2f,
	0x62, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x7d, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x3a, 0x04, 0x62, 0x65, 0x65, 0x72, 0x12, 0x7a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42,
	0x65, 0x65, 0x72, 0x12, 0x28, 0x2e, 0x68, 0x61, 0x79, 0x6f, 0x76, 0x61, 0x6e, 0x6c, 0x6f, 0x6f,
	0x6e, 0x2e, 0x62, 0x61, 0x72, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x68, 0x61, 0x79, 0x6f, 0x76, 0x61, 0x6e, 0x6c, 0x6f, 0x6f, 0x6e, 0x2e, 0x62, 0x61, 0x72, 0x74,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x65, 0x65, 0x72, 0x22, 0x25, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x65, 0x65, 0x72, 0x73,
	0x2f, 0x7b, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x7d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x7b, 0x6e,
	0x61, 0x6d, 0x65, 0x7d, 0x42, 0x52, 0x0a, 0x0a, 0x67, 0x6c, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x65, 0x72, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x48, 0x61, 0x79, 0x6f, 0x56, 0x61, 0x6e, 0x4c, 0x6f, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x61, 0x79, 0x6f, 0x76, 0x61, 0x6e, 0x6c, 0x6f, 0x6f,
	0x6e, 0x2f, 0x62, 0x61, 0x72, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x62,
	0x61, 0x72, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hayovanloon_bartender_v1_bartender_proto_rawDescOnce sync.Once
	file_hayovanloon_bartender_v1_bartender_proto_rawDescData = file_hayovanloon_bartender_v1_bartender_proto_rawDesc
)

func file_hayovanloon_bartender_v1_bartender_proto_rawDescGZIP() []byte {
	file_hayovanloon_bartender_v1_bartender_proto_rawDescOnce.Do(func() {
		file_hayovanloon_bartender_v1_bartender_proto_rawDescData = protoimpl.X.CompressGZIP(file_hayovanloon_bartender_v1_bartender_proto_rawDescData)
	})
	return file_hayovanloon_bartender_v1_bartender_proto_rawDescData
}

var file_hayovanloon_bartender_v1_bartender_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_hayovanloon_bartender_v1_bartender_proto_goTypes = []interface{}{
	(*Beer)(nil),              // 0: hayovanloon.bartender.v1.Beer
	(*GetBeerRequest)(nil),    // 1: hayovanloon.bartender.v1.GetBeerRequest
	(*CreateBeerRequest)(nil), // 2: hayovanloon.bartender.v1.CreateBeerRequest
}
var file_hayovanloon_bartender_v1_bartender_proto_depIdxs = []int32{
	0, // 0: hayovanloon.bartender.v1.CreateBeerRequest.beer:type_name -> hayovanloon.bartender.v1.Beer
	2, // 1: hayovanloon.bartender.v1.Bartender.CreateBeer:input_type -> hayovanloon.bartender.v1.CreateBeerRequest
	1, // 2: hayovanloon.bartender.v1.Bartender.GetBeer:input_type -> hayovanloon.bartender.v1.GetBeerRequest
	0, // 3: hayovanloon.bartender.v1.Bartender.CreateBeer:output_type -> hayovanloon.bartender.v1.Beer
	0, // 4: hayovanloon.bartender.v1.Bartender.GetBeer:output_type -> hayovanloon.bartender.v1.Beer
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_hayovanloon_bartender_v1_bartender_proto_init() }
func file_hayovanloon_bartender_v1_bartender_proto_init() {
	if File_hayovanloon_bartender_v1_bartender_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hayovanloon_bartender_v1_bartender_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Beer); i {
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
		file_hayovanloon_bartender_v1_bartender_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBeerRequest); i {
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
		file_hayovanloon_bartender_v1_bartender_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBeerRequest); i {
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
	file_hayovanloon_bartender_v1_bartender_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Beer_BottleCl)(nil),
		(*Beer_GlassCl)(nil),
		(*Beer_KegL)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hayovanloon_bartender_v1_bartender_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hayovanloon_bartender_v1_bartender_proto_goTypes,
		DependencyIndexes: file_hayovanloon_bartender_v1_bartender_proto_depIdxs,
		MessageInfos:      file_hayovanloon_bartender_v1_bartender_proto_msgTypes,
	}.Build()
	File_hayovanloon_bartender_v1_bartender_proto = out.File
	file_hayovanloon_bartender_v1_bartender_proto_rawDesc = nil
	file_hayovanloon_bartender_v1_bartender_proto_goTypes = nil
	file_hayovanloon_bartender_v1_bartender_proto_depIdxs = nil
}
