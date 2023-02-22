// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: messages.proto

package laptiming

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

type PitStatus int32

const (
	PitStatus_UKNOWN    PitStatus = 0
	PitStatus_PIT_ENTRY PitStatus = 1
	PitStatus_PIT_EXIT  PitStatus = 2
)

// Enum value maps for PitStatus.
var (
	PitStatus_name = map[int32]string{
		0: "UKNOWN",
		1: "PIT_ENTRY",
		2: "PIT_EXIT",
	}
	PitStatus_value = map[string]int32{
		"UKNOWN":    0,
		"PIT_ENTRY": 1,
		"PIT_EXIT":  2,
	}
)

func (x PitStatus) Enum() *PitStatus {
	p := new(PitStatus)
	*p = x
	return p
}

func (x PitStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PitStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_messages_proto_enumTypes[0].Descriptor()
}

func (PitStatus) Type() protoreflect.EnumType {
	return &file_messages_proto_enumTypes[0]
}

func (x PitStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PitStatus.Descriptor instead.
func (PitStatus) EnumDescriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{0}
}

type PitStatusMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp *int64     `protobuf:"varint,2,opt,name=timestamp,proto3,oneof" json:"timestamp,omitempty"`
	PitStatus *PitStatus `protobuf:"varint,3,opt,name=pitStatus,proto3,enum=iot.laptiming.PitStatus,oneof" json:"pitStatus,omitempty"`
}

func (x *PitStatusMessage) Reset() {
	*x = PitStatusMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PitStatusMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PitStatusMessage) ProtoMessage() {}

func (x *PitStatusMessage) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PitStatusMessage.ProtoReflect.Descriptor instead.
func (*PitStatusMessage) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{0}
}

func (x *PitStatusMessage) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PitStatusMessage) GetTimestamp() int64 {
	if x != nil && x.Timestamp != nil {
		return *x.Timestamp
	}
	return 0
}

func (x *PitStatusMessage) GetPitStatus() PitStatus {
	if x != nil && x.PitStatus != nil {
		return *x.PitStatus
	}
	return PitStatus_UKNOWN
}

type PitStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PitStatusResponse) Reset() {
	*x = PitStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PitStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PitStatusResponse) ProtoMessage() {}

func (x *PitStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PitStatusResponse.ProtoReflect.Descriptor instead.
func (*PitStatusResponse) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{1}
}

type LapTimeMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	LapTime   int64 `protobuf:"varint,2,opt,name=lapTime,proto3" json:"lapTime,omitempty"`
	Timestamp int64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *LapTimeMessage) Reset() {
	*x = LapTimeMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LapTimeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LapTimeMessage) ProtoMessage() {}

func (x *LapTimeMessage) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LapTimeMessage.ProtoReflect.Descriptor instead.
func (*LapTimeMessage) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{2}
}

func (x *LapTimeMessage) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LapTimeMessage) GetLapTime() int64 {
	if x != nil {
		return x.LapTime
	}
	return 0
}

func (x *LapTimeMessage) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type LapTimeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LapTimeResponse) Reset() {
	*x = LapTimeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LapTimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LapTimeResponse) ProtoMessage() {}

func (x *LapTimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LapTimeResponse.ProtoReflect.Descriptor instead.
func (*LapTimeResponse) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{3}
}

type TestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled   bool  `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Timestamp int64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *TestMessage) Reset() {
	*x = TestMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestMessage) ProtoMessage() {}

func (x *TestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestMessage.ProtoReflect.Descriptor instead.
func (*TestMessage) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{4}
}

func (x *TestMessage) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *TestMessage) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_messages_proto protoreflect.FileDescriptor

var file_messages_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x69, 0x6f, 0x74, 0x2e, 0x6c, 0x61, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x6e, 0x67, 0x22,
	0x9e, 0x01, 0x0a, 0x10, 0x50, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x09, 0x70, 0x69, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x69, 0x6f, 0x74,
	0x2e, 0x6c, 0x61, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x69, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x48, 0x01, 0x52, 0x09, 0x70, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x13, 0x0a, 0x11, 0x50, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x58, 0x0a, 0x0e, 0x4c, 0x61, 0x70, 0x54, 0x69, 0x6d, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x61, 0x70, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6c, 0x61, 0x70, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22,
	0x11, 0x0a, 0x0f, 0x4c, 0x61, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x45, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2a, 0x34, 0x0a, 0x09, 0x50, 0x69, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x49, 0x54, 0x5f, 0x45, 0x4e, 0x54, 0x52, 0x59, 0x10,
	0x01, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x49, 0x54, 0x5f, 0x45, 0x58, 0x49, 0x54, 0x10, 0x02, 0x42,
	0x0f, 0x5a, 0x0d, 0x69, 0x6f, 0x74, 0x2e, 0x6c, 0x61, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x6e, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messages_proto_rawDescOnce sync.Once
	file_messages_proto_rawDescData = file_messages_proto_rawDesc
)

func file_messages_proto_rawDescGZIP() []byte {
	file_messages_proto_rawDescOnce.Do(func() {
		file_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_proto_rawDescData)
	})
	return file_messages_proto_rawDescData
}

var file_messages_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_messages_proto_goTypes = []interface{}{
	(PitStatus)(0),            // 0: iot.laptiming.PitStatus
	(*PitStatusMessage)(nil),  // 1: iot.laptiming.PitStatusMessage
	(*PitStatusResponse)(nil), // 2: iot.laptiming.PitStatusResponse
	(*LapTimeMessage)(nil),    // 3: iot.laptiming.LapTimeMessage
	(*LapTimeResponse)(nil),   // 4: iot.laptiming.LapTimeResponse
	(*TestMessage)(nil),       // 5: iot.laptiming.TestMessage
}
var file_messages_proto_depIdxs = []int32{
	0, // 0: iot.laptiming.PitStatusMessage.pitStatus:type_name -> iot.laptiming.PitStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_messages_proto_init() }
func file_messages_proto_init() {
	if File_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PitStatusMessage); i {
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
		file_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PitStatusResponse); i {
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
		file_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LapTimeMessage); i {
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
		file_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LapTimeResponse); i {
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
		file_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestMessage); i {
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
	file_messages_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_messages_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_proto_goTypes,
		DependencyIndexes: file_messages_proto_depIdxs,
		EnumInfos:         file_messages_proto_enumTypes,
		MessageInfos:      file_messages_proto_msgTypes,
	}.Build()
	File_messages_proto = out.File
	file_messages_proto_rawDesc = nil
	file_messages_proto_goTypes = nil
	file_messages_proto_depIdxs = nil
}
