// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.12.4
// source: proto/register.proto

package pb

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

type Register struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Latitude  string `protobuf:"bytes,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude string `protobuf:"bytes,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Distance  string `protobuf:"bytes,4,opt,name=distance,proto3" json:"distance,omitempty"`
	Closer    bool   `protobuf:"varint,5,opt,name=closer,proto3" json:"closer,omitempty"`
}

func (x *Register) Reset() {
	*x = Register{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_register_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Register) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Register) ProtoMessage() {}

func (x *Register) ProtoReflect() protoreflect.Message {
	mi := &file_proto_register_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Register.ProtoReflect.Descriptor instead.
func (*Register) Descriptor() ([]byte, []int) {
	return file_proto_register_proto_rawDescGZIP(), []int{0}
}

func (x *Register) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Register) GetLatitude() string {
	if x != nil {
		return x.Latitude
	}
	return ""
}

func (x *Register) GetLongitude() string {
	if x != nil {
		return x.Longitude
	}
	return ""
}

func (x *Register) GetDistance() string {
	if x != nil {
		return x.Distance
	}
	return ""
}

func (x *Register) GetCloser() bool {
	if x != nil {
		return x.Closer
	}
	return false
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  string `protobuf:"bytes,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude string `protobuf:"bytes,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_register_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_register_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_register_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRequest) GetLatitude() string {
	if x != nil {
		return x.Latitude
	}
	return ""
}

func (x *CreateRequest) GetLongitude() string {
	if x != nil {
		return x.Longitude
	}
	return ""
}

type Registers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Regiters []*Register `protobuf:"bytes,1,rep,name=regiters,proto3" json:"regiters,omitempty"`
}

func (x *Registers) Reset() {
	*x = Registers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_register_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registers) ProtoMessage() {}

func (x *Registers) ProtoReflect() protoreflect.Message {
	mi := &file_proto_register_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registers.ProtoReflect.Descriptor instead.
func (*Registers) Descriptor() ([]byte, []int) {
	return file_proto_register_proto_rawDescGZIP(), []int{2}
}

func (x *Registers) GetRegiters() []*Register {
	if x != nil {
		return x.Regiters
	}
	return nil
}

var File_proto_register_proto protoreflect.FileDescriptor

var file_proto_register_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x88, 0x01, 0x0a, 0x08, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x63,
	0x6c, 0x6f, 0x73, 0x65, 0x72, 0x22, 0x49, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x22, 0x35, 0x0a, 0x09, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x28, 0x0a,
	0x08, 0x72, 0x65, 0x67, 0x69, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x08, 0x72,
	0x65, 0x67, 0x69, 0x74, 0x65, 0x72, 0x73, 0x32, 0xcb, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x0e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x11, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x73, 0x22,
	0x00, 0x12, 0x3c, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x70,
	0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x73, 0x22, 0x00, 0x28, 0x01, 0x12,
	0x44, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x42, 0x69, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x11,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x22,
	0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0d, 0x5a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_register_proto_rawDescOnce sync.Once
	file_proto_register_proto_rawDescData = file_proto_register_proto_rawDesc
)

func file_proto_register_proto_rawDescGZIP() []byte {
	file_proto_register_proto_rawDescOnce.Do(func() {
		file_proto_register_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_register_proto_rawDescData)
	})
	return file_proto_register_proto_rawDescData
}

var file_proto_register_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_register_proto_goTypes = []interface{}{
	(*Register)(nil),      // 0: pb.Register
	(*CreateRequest)(nil), // 1: pb.CreateRequest
	(*Registers)(nil),     // 2: pb.Registers
}
var file_proto_register_proto_depIdxs = []int32{
	0, // 0: pb.Registers.regiters:type_name -> pb.Register
	1, // 1: pb.RegisterService.CreateRegister:input_type -> pb.CreateRequest
	1, // 2: pb.RegisterService.CreateRegisterStream:input_type -> pb.CreateRequest
	1, // 3: pb.RegisterService.CreateRegisterBidirectional:input_type -> pb.CreateRequest
	2, // 4: pb.RegisterService.CreateRegister:output_type -> pb.Registers
	2, // 5: pb.RegisterService.CreateRegisterStream:output_type -> pb.Registers
	0, // 6: pb.RegisterService.CreateRegisterBidirectional:output_type -> pb.Register
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_register_proto_init() }
func file_proto_register_proto_init() {
	if File_proto_register_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_register_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Register); i {
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
		file_proto_register_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_proto_register_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registers); i {
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
			RawDescriptor: file_proto_register_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_register_proto_goTypes,
		DependencyIndexes: file_proto_register_proto_depIdxs,
		MessageInfos:      file_proto_register_proto_msgTypes,
	}.Build()
	File_proto_register_proto = out.File
	file_proto_register_proto_rawDesc = nil
	file_proto_register_proto_goTypes = nil
	file_proto_register_proto_depIdxs = nil
}
