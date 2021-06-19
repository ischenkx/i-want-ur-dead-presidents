// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: grabber_service.proto

package grabbing

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

// empty message
type EmptyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyReply) Reset() {
	*x = EmptyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grabber_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyReply) ProtoMessage() {}

func (x *EmptyReply) ProtoReflect() protoreflect.Message {
	mi := &file_grabber_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyReply.ProtoReflect.Descriptor instead.
func (*EmptyReply) Descriptor() ([]byte, []int) {
	return file_grabber_service_proto_rawDescGZIP(), []int{0}
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Inn string `protobuf:"bytes,2,opt,name=inn,proto3" json:"inn,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grabber_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_grabber_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_grabber_service_proto_rawDescGZIP(), []int{1}
}

func (x *Product) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Product) GetInn() string {
	if x != nil {
		return x.Inn
	}
	return ""
}

// model
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Score    *Score `protobuf:"bytes,1,opt,name=score,proto3" json:"score,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	FullName string `protobuf:"bytes,3,opt,name=fullName,proto3" json:"fullName,omitempty"`
	Inn      string `protobuf:"bytes,4,opt,name=inn,proto3" json:"inn,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grabber_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_grabber_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_grabber_service_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetScore() *Score {
	if x != nil {
		return x.Score
	}
	return nil
}

func (x *Response) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Response) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *Response) GetInn() string {
	if x != nil {
		return x.Inn
	}
	return ""
}

type Score struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Score        int32 `protobuf:"varint,1,opt,name=score,proto3" json:"score,omitempty"`
	CourtScore   int32 `protobuf:"varint,2,opt,name=courtScore,proto3" json:"courtScore,omitempty"`
	FinKoefScore int32 `protobuf:"varint,3,opt,name=finKoefScore,proto3" json:"finKoefScore,omitempty"`
	SmartScore   int32 `protobuf:"varint,4,opt,name=smartScore,proto3" json:"smartScore,omitempty"`
}

func (x *Score) Reset() {
	*x = Score{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grabber_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Score) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Score) ProtoMessage() {}

func (x *Score) ProtoReflect() protoreflect.Message {
	mi := &file_grabber_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Score.ProtoReflect.Descriptor instead.
func (*Score) Descriptor() ([]byte, []int) {
	return file_grabber_service_proto_rawDescGZIP(), []int{3}
}

func (x *Score) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Score) GetCourtScore() int32 {
	if x != nil {
		return x.CourtScore
	}
	return 0
}

func (x *Score) GetFinKoefScore() int32 {
	if x != nil {
		return x.FinKoefScore
	}
	return 0
}

func (x *Score) GetSmartScore() int32 {
	if x != nil {
		return x.SmartScore
	}
	return 0
}

var File_grabber_service_proto protoreflect.FileDescriptor

var file_grabber_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x67, 0x72, 0x61, 0x62, 0x62, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0c, 0x0a, 0x0a, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2b, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x69, 0x6e, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69,
	0x6e, 0x6e, 0x22, 0x6a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c,
	0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x69, 0x6e, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x6e, 0x6e, 0x22, 0x81,
	0x01, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x6f, 0x75, 0x72, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x63, 0x6f, 0x75, 0x72, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x66, 0x69, 0x6e, 0x4b, 0x6f, 0x65, 0x66, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x66, 0x69, 0x6e, 0x4b, 0x6f, 0x65, 0x66, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x32, 0x28, 0x0a, 0x08, 0x47, 0x72, 0x61, 0x62, 0x62, 0x69, 0x6e, 0x67, 0x12, 0x1c,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x08, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a,
	0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x38, 0x5a, 0x36,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6e, 0x69, 0x67, 0x68,
	0x74, 0x73, 0x5f, 0x6f, 0x66, 0x5f, 0x6e, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2f, 0x69, 0x6e, 0x6e,
	0x6f, 0x74, 0x65, 0x63, 0x68, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x72,
	0x61, 0x62, 0x62, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grabber_service_proto_rawDescOnce sync.Once
	file_grabber_service_proto_rawDescData = file_grabber_service_proto_rawDesc
)

func file_grabber_service_proto_rawDescGZIP() []byte {
	file_grabber_service_proto_rawDescOnce.Do(func() {
		file_grabber_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_grabber_service_proto_rawDescData)
	})
	return file_grabber_service_proto_rawDescData
}

var file_grabber_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_grabber_service_proto_goTypes = []interface{}{
	(*EmptyReply)(nil), // 0: EmptyReply
	(*Product)(nil),    // 1: Product
	(*Response)(nil),   // 2: Response
	(*Score)(nil),      // 3: Score
}
var file_grabber_service_proto_depIdxs = []int32{
	3, // 0: Response.score:type_name -> Score
	1, // 1: Grabbing.Get:input_type -> Product
	2, // 2: Grabbing.Get:output_type -> Response
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_grabber_service_proto_init() }
func file_grabber_service_proto_init() {
	if File_grabber_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grabber_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyReply); i {
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
		file_grabber_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_grabber_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_grabber_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Score); i {
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
			RawDescriptor: file_grabber_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grabber_service_proto_goTypes,
		DependencyIndexes: file_grabber_service_proto_depIdxs,
		MessageInfos:      file_grabber_service_proto_msgTypes,
	}.Build()
	File_grabber_service_proto = out.File
	file_grabber_service_proto_rawDesc = nil
	file_grabber_service_proto_goTypes = nil
	file_grabber_service_proto_depIdxs = nil
}