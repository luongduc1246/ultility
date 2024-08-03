// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: test.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid          string                    `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name          *wrapperspb.StringValue   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status        *wrapperspb.Int32Value    `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Description   string                    `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Test          *string                   `protobuf:"bytes,5,opt,name=test,proto3,oneof" json:"test,omitempty"`
	Mame          *wrapperspb.StringValue   `protobuf:"bytes,6,opt,name=mame,proto3" json:"mame,omitempty"`
	User          *User                     `protobuf:"bytes,7,opt,name=user,proto3,oneof" json:"user,omitempty"`
	Users         []*User                   `protobuf:"bytes,8,rep,name=users,proto3" json:"users,omitempty"`
	Address       []string                  `protobuf:"bytes,9,rep,name=address,proto3" json:"address,omitempty"`
	Meme          []*wrapperspb.StringValue `protobuf:"bytes,10,rep,name=meme,proto3" json:"meme,omitempty"`
	MyTime        *timestamppb.Timestamp    `protobuf:"bytes,11,opt,name=my_time,json=myTime,proto3" json:"my_time,omitempty"`
	MyDataJson    *structpb.Struct          `protobuf:"bytes,12,opt,name=my_data_json,json=myDataJson,proto3" json:"my_data_json,omitempty"`
	MyDataJsonMap *structpb.Struct          `protobuf:"bytes,13,opt,name=my_data_json_map,json=myDataJsonMap,proto3" json:"my_data_json_map,omitempty"`
	MyDataDate    *timestamppb.Timestamp    `protobuf:"bytes,14,opt,name=my_data_date,json=myDataDate,proto3" json:"my_data_date,omitempty"`
	MyDataTime    *durationpb.Duration      `protobuf:"bytes,15,opt,name=my_data_time,json=myDataTime,proto3" json:"my_data_time,omitempty"`
	MyByte        []byte                    `protobuf:"bytes,16,opt,name=my_byte,json=myByte,proto3" json:"my_byte,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{1}
}

func (x *Request) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Request) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *Request) GetStatus() *wrapperspb.Int32Value {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *Request) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Request) GetTest() string {
	if x != nil && x.Test != nil {
		return *x.Test
	}
	return ""
}

func (x *Request) GetMame() *wrapperspb.StringValue {
	if x != nil {
		return x.Mame
	}
	return nil
}

func (x *Request) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Request) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *Request) GetAddress() []string {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Request) GetMeme() []*wrapperspb.StringValue {
	if x != nil {
		return x.Meme
	}
	return nil
}

func (x *Request) GetMyTime() *timestamppb.Timestamp {
	if x != nil {
		return x.MyTime
	}
	return nil
}

func (x *Request) GetMyDataJson() *structpb.Struct {
	if x != nil {
		return x.MyDataJson
	}
	return nil
}

func (x *Request) GetMyDataJsonMap() *structpb.Struct {
	if x != nil {
		return x.MyDataJsonMap
	}
	return nil
}

func (x *Request) GetMyDataDate() *timestamppb.Timestamp {
	if x != nil {
		return x.MyDataDate
	}
	return nil
}

func (x *Request) GetMyDataTime() *durationpb.Duration {
	if x != nil {
		return x.MyDataTime
	}
	return nil
}

func (x *Request) GetMyByte() []byte {
	if x != nil {
		return x.MyByte
	}
	return nil
}

var File_test_proto protoreflect.FileDescriptor

var file_test_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0xde, 0x05, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x12, 0x30, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x04, 0x74,
	0x65, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x74, 0x65, 0x73,
	0x74, 0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x04, 0x6d, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x04, 0x6d, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x48, 0x01, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x05,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x30, 0x0a, 0x04, 0x6d, 0x65, 0x6d,
	0x65, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x6d,
	0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x6d, 0x79, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x39, 0x0a, 0x0c, 0x6d, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6a, 0x73, 0x6f, 0x6e,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52,
	0x0a, 0x6d, 0x79, 0x44, 0x61, 0x74, 0x61, 0x4a, 0x73, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x10, 0x6d,
	0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x70, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0d,
	0x6d, 0x79, 0x44, 0x61, 0x74, 0x61, 0x4a, 0x73, 0x6f, 0x6e, 0x4d, 0x61, 0x70, 0x12, 0x3c, 0x0a,
	0x0c, 0x6d, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x6d, 0x79, 0x44, 0x61, 0x74, 0x61, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x6d,
	0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6d, 0x79,
	0x44, 0x61, 0x74, 0x61, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x79, 0x5f, 0x62,
	0x79, 0x74, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6d, 0x79, 0x42, 0x79, 0x74,
	0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x32, 0x32, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_proto_rawDescOnce sync.Once
	file_test_proto_rawDescData = file_test_proto_rawDesc
)

func file_test_proto_rawDescGZIP() []byte {
	file_test_proto_rawDescOnce.Do(func() {
		file_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_proto_rawDescData)
	})
	return file_test_proto_rawDescData
}

var file_test_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_test_proto_goTypes = []interface{}{
	(*User)(nil),                   // 0: proto.User
	(*Request)(nil),                // 1: proto.Request
	(*wrapperspb.StringValue)(nil), // 2: google.protobuf.StringValue
	(*wrapperspb.Int32Value)(nil),  // 3: google.protobuf.Int32Value
	(*timestamppb.Timestamp)(nil),  // 4: google.protobuf.Timestamp
	(*structpb.Struct)(nil),        // 5: google.protobuf.Struct
	(*durationpb.Duration)(nil),    // 6: google.protobuf.Duration
}
var file_test_proto_depIdxs = []int32{
	2,  // 0: proto.Request.name:type_name -> google.protobuf.StringValue
	3,  // 1: proto.Request.status:type_name -> google.protobuf.Int32Value
	2,  // 2: proto.Request.mame:type_name -> google.protobuf.StringValue
	0,  // 3: proto.Request.user:type_name -> proto.User
	0,  // 4: proto.Request.users:type_name -> proto.User
	2,  // 5: proto.Request.meme:type_name -> google.protobuf.StringValue
	4,  // 6: proto.Request.my_time:type_name -> google.protobuf.Timestamp
	5,  // 7: proto.Request.my_data_json:type_name -> google.protobuf.Struct
	5,  // 8: proto.Request.my_data_json_map:type_name -> google.protobuf.Struct
	4,  // 9: proto.Request.my_data_date:type_name -> google.protobuf.Timestamp
	6,  // 10: proto.Request.my_data_time:type_name -> google.protobuf.Duration
	1,  // 11: proto.TestService.Test:input_type -> proto.Request
	0,  // 12: proto.TestService.Test:output_type -> proto.User
	12, // [12:13] is the sub-list for method output_type
	11, // [11:12] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_test_proto_init() }
func file_test_proto_init() {
	if File_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
	file_test_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test_proto_goTypes,
		DependencyIndexes: file_test_proto_depIdxs,
		MessageInfos:      file_test_proto_msgTypes,
	}.Build()
	File_test_proto = out.File
	file_test_proto_rawDesc = nil
	file_test_proto_goTypes = nil
	file_test_proto_depIdxs = nil
}