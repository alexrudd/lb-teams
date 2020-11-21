// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: infra/lb/message.proto

package lb

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset      int64                  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	PubTimstamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=pub_timstamp,json=pubTimstamp,proto3" json:"pub_timstamp,omitempty"`
	StreamName  string                 `protobuf:"bytes,3,opt,name=stream_name,json=streamName,proto3" json:"stream_name,omitempty"`
	RawData     *anypb.Any             `protobuf:"bytes,4,opt,name=raw_data,json=rawData,proto3" json:"raw_data,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_lb_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_infra_lb_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_infra_lb_message_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *Event) GetPubTimstamp() *timestamppb.Timestamp {
	if x != nil {
		return x.PubTimstamp
	}
	return nil
}

func (x *Event) GetStreamName() string {
	if x != nil {
		return x.StreamName
	}
	return ""
}

func (x *Event) GetRawData() *anypb.Any {
	if x != nil {
		return x.RawData
	}
	return nil
}

var File_infra_lb_message_proto protoreflect.FileDescriptor

var file_infra_lb_message_proto_rawDesc = []byte{
	0x0a, 0x16, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x6c, 0x62, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6c, 0x69, 0x66, 0x74, 0x62, 0x72,
	0x69, 0x64, 0x67, 0x65, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb0, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x5f, 0x74, 0x69, 0x6d, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x54, 0x69, 0x6d, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x72, 0x61, 0x77, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x72, 0x61, 0x77, 0x44,
	0x61, 0x74, 0x61, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x6c, 0x65, 0x78, 0x72, 0x75, 0x64, 0x64, 0x2f, 0x6c, 0x62, 0x2d, 0x74, 0x65,
	0x61, 0x6d, 0x73, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x6c, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_lb_message_proto_rawDescOnce sync.Once
	file_infra_lb_message_proto_rawDescData = file_infra_lb_message_proto_rawDesc
)

func file_infra_lb_message_proto_rawDescGZIP() []byte {
	file_infra_lb_message_proto_rawDescOnce.Do(func() {
		file_infra_lb_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_lb_message_proto_rawDescData)
	})
	return file_infra_lb_message_proto_rawDescData
}

var file_infra_lb_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_infra_lb_message_proto_goTypes = []interface{}{
	(*Event)(nil),                 // 0: liftbridge.Event
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*anypb.Any)(nil),             // 2: google.protobuf.Any
}
var file_infra_lb_message_proto_depIdxs = []int32{
	1, // 0: liftbridge.Event.pub_timstamp:type_name -> google.protobuf.Timestamp
	2, // 1: liftbridge.Event.raw_data:type_name -> google.protobuf.Any
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_infra_lb_message_proto_init() }
func file_infra_lb_message_proto_init() {
	if File_infra_lb_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_lb_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
			RawDescriptor: file_infra_lb_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_lb_message_proto_goTypes,
		DependencyIndexes: file_infra_lb_message_proto_depIdxs,
		MessageInfos:      file_infra_lb_message_proto_msgTypes,
	}.Build()
	File_infra_lb_message_proto = out.File
	file_infra_lb_message_proto_rawDesc = nil
	file_infra_lb_message_proto_goTypes = nil
	file_infra_lb_message_proto_depIdxs = nil
}
