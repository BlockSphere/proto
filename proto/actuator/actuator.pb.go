// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.3
// source: proto/actuator.proto

package actuatorpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TriggerRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	GreenhouseId  string                 `protobuf:"bytes,1,opt,name=greenhouse_id,json=greenhouseId,proto3" json:"greenhouse_id,omitempty"`
	Component     string                 `protobuf:"bytes,2,opt,name=component,proto3" json:"component,omitempty"`
	Command       string                 `protobuf:"bytes,3,opt,name=command,proto3" json:"command,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TriggerRequest) Reset() {
	*x = TriggerRequest{}
	mi := &file_proto_actuator_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TriggerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TriggerRequest) ProtoMessage() {}

func (x *TriggerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actuator_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriggerRequest.ProtoReflect.Descriptor instead.
func (*TriggerRequest) Descriptor() ([]byte, []int) {
	return file_proto_actuator_proto_rawDescGZIP(), []int{0}
}

func (x *TriggerRequest) GetGreenhouseId() string {
	if x != nil {
		return x.GreenhouseId
	}
	return ""
}

func (x *TriggerRequest) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *TriggerRequest) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

type TriggerResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TriggerResponse) Reset() {
	*x = TriggerResponse{}
	mi := &file_proto_actuator_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TriggerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TriggerResponse) ProtoMessage() {}

func (x *TriggerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actuator_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriggerResponse.ProtoReflect.Descriptor instead.
func (*TriggerResponse) Descriptor() ([]byte, []int) {
	return file_proto_actuator_proto_rawDescGZIP(), []int{1}
}

func (x *TriggerResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_actuator_proto protoreflect.FileDescriptor

const file_proto_actuator_proto_rawDesc = "" +
	"\n" +
	"\x14proto/actuator.proto\x12\bactuator\"m\n" +
	"\x0eTriggerRequest\x12#\n" +
	"\rgreenhouse_id\x18\x01 \x01(\tR\fgreenhouseId\x12\x1c\n" +
	"\tcomponent\x18\x02 \x01(\tR\tcomponent\x12\x18\n" +
	"\acommand\x18\x03 \x01(\tR\acommand\")\n" +
	"\x0fTriggerResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06status2W\n" +
	"\x0fActuatorService\x12D\n" +
	"\rTriggerAction\x12\x18.actuator.TriggerRequest\x1a\x19.actuator.TriggerResponseB\x1bZ\x19proto/actuator;actuatorpbb\x06proto3"

var (
	file_proto_actuator_proto_rawDescOnce sync.Once
	file_proto_actuator_proto_rawDescData []byte
)

func file_proto_actuator_proto_rawDescGZIP() []byte {
	file_proto_actuator_proto_rawDescOnce.Do(func() {
		file_proto_actuator_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_actuator_proto_rawDesc), len(file_proto_actuator_proto_rawDesc)))
	})
	return file_proto_actuator_proto_rawDescData
}

var file_proto_actuator_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_actuator_proto_goTypes = []any{
	(*TriggerRequest)(nil),  // 0: actuator.TriggerRequest
	(*TriggerResponse)(nil), // 1: actuator.TriggerResponse
}
var file_proto_actuator_proto_depIdxs = []int32{
	0, // 0: actuator.ActuatorService.TriggerAction:input_type -> actuator.TriggerRequest
	1, // 1: actuator.ActuatorService.TriggerAction:output_type -> actuator.TriggerResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_actuator_proto_init() }
func file_proto_actuator_proto_init() {
	if File_proto_actuator_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_actuator_proto_rawDesc), len(file_proto_actuator_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_actuator_proto_goTypes,
		DependencyIndexes: file_proto_actuator_proto_depIdxs,
		MessageInfos:      file_proto_actuator_proto_msgTypes,
	}.Build()
	File_proto_actuator_proto = out.File
	file_proto_actuator_proto_goTypes = nil
	file_proto_actuator_proto_depIdxs = nil
}
