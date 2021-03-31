// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: tb_online_list.proto

package tcaplusservice

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

type TbOnlineList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Openid    int32                `protobuf:"varint,1,opt,name=openid,proto3" json:"openid,omitempty"` //QQ Uin
	Tconndid  int32                `protobuf:"varint,2,opt,name=tconndid,proto3" json:"tconndid,omitempty"`
	Timekey   string               `protobuf:"bytes,3,opt,name=timekey,proto3" json:"timekey,omitempty"`
	Gamesvrid string               `protobuf:"bytes,4,opt,name=gamesvrid,proto3" json:"gamesvrid,omitempty"`
	Logintime int32                `protobuf:"varint,5,opt,name=logintime,proto3" json:"logintime,omitempty"`
	Lockid    []int64              `protobuf:"varint,6,rep,packed,name=lockid,proto3" json:"lockid,omitempty"`
	Pay       *TbOnlineListPayInfo `protobuf:"bytes,7,opt,name=pay,proto3" json:"pay,omitempty"`
}

func (x *TbOnlineList) Reset() {
	*x = TbOnlineList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tb_online_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TbOnlineList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TbOnlineList) ProtoMessage() {}

func (x *TbOnlineList) ProtoReflect() protoreflect.Message {
	mi := &file_tb_online_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TbOnlineList.ProtoReflect.Descriptor instead.
func (*TbOnlineList) Descriptor() ([]byte, []int) {
	return file_tb_online_list_proto_rawDescGZIP(), []int{0}
}

func (x *TbOnlineList) GetOpenid() int32 {
	if x != nil {
		return x.Openid
	}
	return 0
}

func (x *TbOnlineList) GetTconndid() int32 {
	if x != nil {
		return x.Tconndid
	}
	return 0
}

func (x *TbOnlineList) GetTimekey() string {
	if x != nil {
		return x.Timekey
	}
	return ""
}

func (x *TbOnlineList) GetGamesvrid() string {
	if x != nil {
		return x.Gamesvrid
	}
	return ""
}

func (x *TbOnlineList) GetLogintime() int32 {
	if x != nil {
		return x.Logintime
	}
	return 0
}

func (x *TbOnlineList) GetLockid() []int64 {
	if x != nil {
		return x.Lockid
	}
	return nil
}

func (x *TbOnlineList) GetPay() *TbOnlineListPayInfo {
	if x != nil {
		return x.Pay
	}
	return nil
}

type TbOnlineListPayInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalMoney uint64 `protobuf:"varint,1,opt,name=total_money,json=totalMoney,proto3" json:"total_money,omitempty"`
	PayTimes   uint64 `protobuf:"varint,2,opt,name=pay_times,json=payTimes,proto3" json:"pay_times,omitempty"`
}

func (x *TbOnlineListPayInfo) Reset() {
	*x = TbOnlineListPayInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tb_online_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TbOnlineListPayInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TbOnlineListPayInfo) ProtoMessage() {}

func (x *TbOnlineListPayInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tb_online_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TbOnlineListPayInfo.ProtoReflect.Descriptor instead.
func (*TbOnlineListPayInfo) Descriptor() ([]byte, []int) {
	return file_tb_online_list_proto_rawDescGZIP(), []int{0, 0}
}

func (x *TbOnlineListPayInfo) GetTotalMoney() uint64 {
	if x != nil {
		return x.TotalMoney
	}
	return 0
}

func (x *TbOnlineListPayInfo) GetPayTimes() uint64 {
	if x != nil {
		return x.PayTimes
	}
	return 0
}

var File_tb_online_list_proto protoreflect.FileDescriptor

var file_tb_online_list_proto_rawDesc = []byte{
	0x0a, 0x14, 0x74, 0x62, 0x5f, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x63, 0x61, 0x70, 0x6c, 0x75, 0x73, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1d, 0x74, 0x63, 0x61, 0x70, 0x6c, 0x75, 0x73, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x02, 0x0a, 0x0e, 0x74, 0x62, 0x5f, 0x6f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x65, 0x6e,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x74, 0x63, 0x6f, 0x6e, 0x6e, 0x64, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x74, 0x63, 0x6f, 0x6e, 0x6e, 0x64, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x74, 0x69, 0x6d, 0x65, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74,
	0x69, 0x6d, 0x65, 0x6b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x76,
	0x72, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x73,
	0x76, 0x72, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x69, 0x64, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x03, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x03, 0x70, 0x61,
	0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x74, 0x63, 0x61, 0x70, 0x6c, 0x75,
	0x73, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x74, 0x62, 0x5f, 0x6f, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x61, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x52, 0x03, 0x70, 0x61, 0x79, 0x1a, 0x48, 0x0a, 0x08, 0x70, 0x61, 0x79, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4d, 0x6f, 0x6e,
	0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x61, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x3a,
	0x3a, 0x82, 0xa6, 0x1d, 0x17, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x2c, 0x74, 0x63, 0x6f, 0x6e,
	0x6e, 0x64, 0x69, 0x64, 0x2c, 0x74, 0x69, 0x6d, 0x65, 0x6b, 0x65, 0x79, 0xb2, 0xa6, 0x1d, 0x1b,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x3d, 0x4c, 0x49, 0x53, 0x54, 0x3b, 0x4c,
	0x69, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x3d, 0x31, 0x30, 0x32, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_tb_online_list_proto_rawDescOnce sync.Once
	file_tb_online_list_proto_rawDescData = file_tb_online_list_proto_rawDesc
)

func file_tb_online_list_proto_rawDescGZIP() []byte {
	file_tb_online_list_proto_rawDescOnce.Do(func() {
		file_tb_online_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_tb_online_list_proto_rawDescData)
	})
	return file_tb_online_list_proto_rawDescData
}

var file_tb_online_list_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tb_online_list_proto_goTypes = []interface{}{
	(*TbOnlineList)(nil),        // 0: tcaplusservice.tb_online_list
	(*TbOnlineListPayInfo)(nil), // 1: tcaplusservice.tb_online_list.pay_info
}
var file_tb_online_list_proto_depIdxs = []int32{
	1, // 0: tcaplusservice.tb_online_list.pay:type_name -> tcaplusservice.tb_online_list.pay_info
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tb_online_list_proto_init() }
func file_tb_online_list_proto_init() {
	if File_tb_online_list_proto != nil {
		return
	}
	file_tcaplusservice_optionv1_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tb_online_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TbOnlineList); i {
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
		file_tb_online_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TbOnlineListPayInfo); i {
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
			RawDescriptor: file_tb_online_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tb_online_list_proto_goTypes,
		DependencyIndexes: file_tb_online_list_proto_depIdxs,
		MessageInfos:      file_tb_online_list_proto_msgTypes,
	}.Build()
	File_tb_online_list_proto = out.File
	file_tb_online_list_proto_rawDesc = nil
	file_tb_online_list_proto_goTypes = nil
	file_tb_online_list_proto_depIdxs = nil
}
