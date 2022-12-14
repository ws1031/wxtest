// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.19.0
// source: weixin-service.proto

package proto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	_ "team.wphr.vip/technology-group/infrastructure/protoc-sdk"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//服务对外错误码, 会同步至idl/error/error.go文件中，E_ErrOk 和 E_DEFAULT_ERROR 不要删，会影响框架正常功能，错误码的值可以改
type ErrCode int32

const (
	ErrCode_E_ErrOk                    ErrCode = 0     //成功
	ErrCode_E_DEFAULT_ERROR            ErrCode = 99999 //系统错误
	ErrCode_E_FORBIDDEN                ErrCode = 403   //没有权限
	ErrCode_E_PARAM_ERROR              ErrCode = 10400 //参数错误
	ErrCode_E_TOKEN_ERROR              ErrCode = 10401 //token 错误
	ErrCode_E_DOWN_STREAM_RETURN_ERROR ErrCode = 10001 //调用下游错误
	ErrCode_E_GET_DATA_ERROR           ErrCode = 10002 //获取数据异常
	ErrCode_E_INSERT_ERROR             ErrCode = 10003 //数据入库失败
	// 文件错误
	ErrCode_E_UPLOAD_FILE_ERROR ErrCode = 20001 //文件上传错误
)

// Enum value maps for ErrCode.
var (
	ErrCode_name = map[int32]string{
		0:     "E_ErrOk",
		99999: "E_DEFAULT_ERROR",
		403:   "E_FORBIDDEN",
		10400: "E_PARAM_ERROR",
		10401: "E_TOKEN_ERROR",
		10001: "E_DOWN_STREAM_RETURN_ERROR",
		10002: "E_GET_DATA_ERROR",
		10003: "E_INSERT_ERROR",
		20001: "E_UPLOAD_FILE_ERROR",
	}
	ErrCode_value = map[string]int32{
		"E_ErrOk":                    0,
		"E_DEFAULT_ERROR":            99999,
		"E_FORBIDDEN":                403,
		"E_PARAM_ERROR":              10400,
		"E_TOKEN_ERROR":              10401,
		"E_DOWN_STREAM_RETURN_ERROR": 10001,
		"E_GET_DATA_ERROR":           10002,
		"E_INSERT_ERROR":             10003,
		"E_UPLOAD_FILE_ERROR":        20001,
	}
	ErrCode_desc = map[int32]string{
		0:     "成功",
		99999: "系统错误",
		403:   "没有权限",
		10400: "参数错误",
		10401: "token 错误",
		10001: "调用下游错误",
		10002: "获取数据异常",
		10003: "数据入库失败",
		20001: "文件上传错误",
	}
	ErrCode_desc_without_default = map[int32]string{
		99999: "系统错误",
		403:   "没有权限",
		10400: "参数错误",
		10401: "token 错误",
		10001: "调用下游错误",
		10002: "获取数据异常",
		10003: "数据入库失败",
		20001: "文件上传错误",
	}
)

func (x ErrCode) Enum() *ErrCode {
	p := new(ErrCode)
	*p = x
	return p
}

func (x ErrCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrCode) Descriptor() protoreflect.EnumDescriptor {
	return file_weixin_service_proto_enumTypes[0].Descriptor()
}

func (ErrCode) Type() protoreflect.EnumType {
	return &file_weixin_service_proto_enumTypes[0]
}

func (x ErrCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrCode.Descriptor instead.
func (ErrCode) EnumDescriptor() ([]byte, []int) {
	return file_weixin_service_proto_rawDescGZIP(), []int{0}
}

type PingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"` //姓名
}

func (x *PingReq) Reset() {
	*x = PingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weixin_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingReq) ProtoMessage() {}

func (x *PingReq) ProtoReflect() protoreflect.Message {
	mi := &file_weixin_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingReq.ProtoReflect.Descriptor instead.
func (*PingReq) Descriptor() ([]byte, []int) {
	return file_weixin_service_proto_rawDescGZIP(), []int{0}
}

func (x *PingReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

//Ping返回
type PingRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errno  int64  `protobuf:"varint,1,opt,name=errno,proto3" json:"errno"`  //错误码
	Errmsg string `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg"` //错误信息
	Data   string `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`     //返回值
}

func (x *PingRsp) Reset() {
	*x = PingRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weixin_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRsp) ProtoMessage() {}

func (x *PingRsp) ProtoReflect() protoreflect.Message {
	mi := &file_weixin_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRsp.ProtoReflect.Descriptor instead.
func (*PingRsp) Descriptor() ([]byte, []int) {
	return file_weixin_service_proto_rawDescGZIP(), []int{1}
}

func (x *PingRsp) GetErrno() int64 {
	if x != nil {
		return x.Errno
	}
	return 0
}

func (x *PingRsp) GetErrmsg() string {
	if x != nil {
		return x.Errmsg
	}
	return ""
}

func (x *PingRsp) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_weixin_service_proto protoreflect.FileDescriptor

var file_weixin_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x77, 0x65, 0x69, 0x78, 0x69, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x57, 0x65, 0x69, 0x78, 0x69, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x78, 0x64, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x64, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x1d, 0x0a, 0x07, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x4b, 0x0a, 0x07, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6e,
	0x6f, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0xcf, 0x01,
	0x0a, 0x07, 0x45, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x5f, 0x45,
	0x72, 0x72, 0x4f, 0x6b, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x0f, 0x45, 0x5f, 0x44, 0x45, 0x46, 0x41,
	0x55, 0x4c, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x9f, 0x8d, 0x06, 0x12, 0x10, 0x0a,
	0x0b, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x42, 0x49, 0x44, 0x44, 0x45, 0x4e, 0x10, 0x93, 0x03, 0x12,
	0x12, 0x0a, 0x0d, 0x45, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0xa0, 0x51, 0x12, 0x12, 0x0a, 0x0d, 0x45, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0xa1, 0x51, 0x12, 0x1f, 0x0a, 0x1a, 0x45, 0x5f, 0x44, 0x4f, 0x57,
	0x4e, 0x5f, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x5f, 0x52, 0x45, 0x54, 0x55, 0x52, 0x4e, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x91, 0x4e, 0x12, 0x15, 0x0a, 0x10, 0x45, 0x5f, 0x47, 0x45,
	0x54, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x92, 0x4e, 0x12,
	0x13, 0x0a, 0x0e, 0x45, 0x5f, 0x49, 0x4e, 0x53, 0x45, 0x52, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x93, 0x4e, 0x12, 0x19, 0x0a, 0x13, 0x45, 0x5f, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44,
	0x5f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xa1, 0x9c, 0x01, 0x32,
	0xe1, 0x01, 0x0a, 0x0d, 0x57, 0x65, 0x69, 0x78, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x45, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x57, 0x65, 0x69, 0x78,
	0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x1a, 0x16, 0x2e, 0x57, 0x65, 0x69, 0x78, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x73, 0x70, 0x22, 0x0d, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x07, 0x12, 0x05, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x57, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64,
	0x4e, 0x65, 0x77, 0x73, 0x12, 0x1a, 0x2e, 0x57, 0x65, 0x69, 0x78, 0x69, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71,
	0x1a, 0x1b, 0x2e, 0x57, 0x65, 0x69, 0x78, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x12, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x22, 0x0a, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2f, 0x73, 0x65, 0x6e,
	0x64, 0x1a, 0x30, 0xb2, 0xb8, 0x87, 0x4d, 0x2b, 0x0a, 0x05, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x12,
	0x0e, 0x77, 0x65, 0x69, 0x78, 0x69, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18,
	0xe8, 0x07, 0x20, 0x02, 0x2a, 0x0d, 0x78, 0x78, 0x78, 0x2e, 0x76, 0x61, 0x6e, 0x68, 0x72, 0x2e,
	0x6e, 0x65, 0x74, 0x42, 0x07, 0x5a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_weixin_service_proto_rawDescOnce sync.Once
	file_weixin_service_proto_rawDescData = file_weixin_service_proto_rawDesc
)

func file_weixin_service_proto_rawDescGZIP() []byte {
	file_weixin_service_proto_rawDescOnce.Do(func() {
		file_weixin_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_weixin_service_proto_rawDescData)
	})
	return file_weixin_service_proto_rawDescData
}

var file_weixin_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_weixin_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_weixin_service_proto_goTypes = []interface{}{
	(ErrCode)(0),         // 0: WeixinService.ErrCode
	(*PingReq)(nil),      // 1: WeixinService.PingReq
	(*PingRsp)(nil),      // 2: WeixinService.PingRsp
	(*SendNewsReq)(nil),  // 3: WeixinService.SendNewsReq
	(*SendNewsResp)(nil), // 4: WeixinService.SendNewsResp
}
var file_weixin_service_proto_depIdxs = []int32{
	1, // 0: WeixinService.WeixinService.Ping:input_type -> WeixinService.PingReq
	3, // 1: WeixinService.WeixinService.SendNews:input_type -> WeixinService.SendNewsReq
	2, // 2: WeixinService.WeixinService.Ping:output_type -> WeixinService.PingRsp
	4, // 3: WeixinService.WeixinService.SendNews:output_type -> WeixinService.SendNewsResp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_weixin_service_proto_init() }
func file_weixin_service_proto_init() {
	if File_weixin_service_proto != nil {
		return
	}
	file_basic_api_proto_init()
	file_news_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_weixin_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingReq); i {
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
		file_weixin_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRsp); i {
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
			RawDescriptor: file_weixin_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_weixin_service_proto_goTypes,
		DependencyIndexes: file_weixin_service_proto_depIdxs,
		EnumInfos:         file_weixin_service_proto_enumTypes,
		MessageInfos:      file_weixin_service_proto_msgTypes,
	}.Build()
	File_weixin_service_proto = out.File
	file_weixin_service_proto_rawDesc = nil
	file_weixin_service_proto_goTypes = nil
	file_weixin_service_proto_depIdxs = nil
}
