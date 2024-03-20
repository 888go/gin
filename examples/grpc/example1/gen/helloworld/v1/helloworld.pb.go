// 该代码由 protoc-gen-go 自动生成。请勿直接编辑。
// 版本信息：
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// 源文件：pb/helloworld/v1/helloworld.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// 验证此生成的代码是否足够新。
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// 验证runtime/protoimpl是否足够新。
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 请求消息，其中包含用户的名字。
type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}


// ff:
func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_helloworld_v1_helloworld_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}


// ff:
func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}


// ff:
func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_helloworld_v1_helloworld_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// 已弃用：请改用 HelloRequest.ProtoReflect.Descriptor。
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_pb_helloworld_v1_helloworld_proto_rawDescGZIP(), []int{0}
}


// ff:
func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 响应消息，其中包含问候语
type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}


// ff:
func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_helloworld_v1_helloworld_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}


// ff:
func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}


// ff:
func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_pb_helloworld_v1_helloworld_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// 已弃用：请改用 HelloReply.ProtoReflect.Descriptor。
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_pb_helloworld_v1_helloworld_proto_rawDescGZIP(), []int{1}
}


// ff:
func (x *HelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_pb_helloworld_v1_helloworld_proto protoreflect.FileDescriptor

var file_pb_helloworld_v1_helloworld_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x62, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e,
	0x76, 0x31, 0x22, 0x22, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x4f,
	0x0a, 0x07, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x08, 0x53, 0x61, 0x79,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x1b, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42,
	0x43, 0x0a, 0x1b, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x73, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x42, 0x0f,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x11, 0x67, 0x65, 0x6e, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_helloworld_v1_helloworld_proto_rawDescOnce sync.Once
	file_pb_helloworld_v1_helloworld_proto_rawDescData = file_pb_helloworld_v1_helloworld_proto_rawDesc
)

func file_pb_helloworld_v1_helloworld_proto_rawDescGZIP() []byte {
	file_pb_helloworld_v1_helloworld_proto_rawDescOnce.Do(func() {
		file_pb_helloworld_v1_helloworld_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_helloworld_v1_helloworld_proto_rawDescData)
	})
	return file_pb_helloworld_v1_helloworld_proto_rawDescData
}

var file_pb_helloworld_v1_helloworld_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_helloworld_v1_helloworld_proto_goTypes = []interface{}{
	(*HelloRequest)(nil), // 0: helloworld.v1.HelloRequest 的中文翻译是：
// 
// 0: helloworld.v1.HelloRequest（HelloRequest 是 helloworld.v1 包中的结构体或接口类型）
	(*HelloReply)(nil),   // 1: helloworld.v1.HelloReply （这是一条关于Go语言程序的注释，翻译为：）
// 
// 1: helloworld.v1.HelloReply 表示这是一个版本为v1的helloworld包中的HelloReply结构体或接口。
}
var file_pb_helloworld_v1_helloworld_proto_depIdxs = []int32{
	0, // 0: helloworld.v1.Greeter.SayHello 方法的输入类型 -> helloworld.v1.HelloRequest
	1, // 1: helloworld.v1.Greeter.SayHello 方法的输出类型为 helloworld.v1.HelloReply
	1, // [1:2] 是方法输出类型对应的子列表
	0, // [0:1] 是用于方法输入类型的子列表
	0, // [0:0] 是用于扩展类型名的子列表
	0, // [0:0] 是扩展实体的子列表
	0, // [0:0] 是字段 type_name 的子列表
}

func init() { file_pb_helloworld_v1_helloworld_proto_init() }
func file_pb_helloworld_v1_helloworld_proto_init() {
	if File_pb_helloworld_v1_helloworld_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_helloworld_v1_helloworld_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_pb_helloworld_v1_helloworld_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
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
			RawDescriptor: file_pb_helloworld_v1_helloworld_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_helloworld_v1_helloworld_proto_goTypes,
		DependencyIndexes: file_pb_helloworld_v1_helloworld_proto_depIdxs,
		MessageInfos:      file_pb_helloworld_v1_helloworld_proto_msgTypes,
	}.Build()
	File_pb_helloworld_v1_helloworld_proto = out.File
	file_pb_helloworld_v1_helloworld_proto_rawDesc = nil
	file_pb_helloworld_v1_helloworld_proto_goTypes = nil
	file_pb_helloworld_v1_helloworld_proto_depIdxs = nil
}
