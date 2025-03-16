// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.3
// source: packets.proto

package packets

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

type LoginRequestMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginRequestMessage) Reset() {
	*x = LoginRequestMessage{}
	mi := &file_packets_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequestMessage) ProtoMessage() {}

func (x *LoginRequestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_packets_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequestMessage.ProtoReflect.Descriptor instead.
func (*LoginRequestMessage) Descriptor() ([]byte, []int) {
	return file_packets_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequestMessage) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginRequestMessage) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterRequestMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterRequestMessage) Reset() {
	*x = RegisterRequestMessage{}
	mi := &file_packets_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterRequestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequestMessage) ProtoMessage() {}

func (x *RegisterRequestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_packets_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequestMessage.ProtoReflect.Descriptor instead.
func (*RegisterRequestMessage) Descriptor() ([]byte, []int) {
	return file_packets_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterRequestMessage) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterRequestMessage) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type OkResponseMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OkResponseMessage) Reset() {
	*x = OkResponseMessage{}
	mi := &file_packets_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OkResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OkResponseMessage) ProtoMessage() {}

func (x *OkResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_packets_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OkResponseMessage.ProtoReflect.Descriptor instead.
func (*OkResponseMessage) Descriptor() ([]byte, []int) {
	return file_packets_proto_rawDescGZIP(), []int{2}
}

type DenyResponseMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Reason        string                 `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DenyResponseMessage) Reset() {
	*x = DenyResponseMessage{}
	mi := &file_packets_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DenyResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DenyResponseMessage) ProtoMessage() {}

func (x *DenyResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_packets_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DenyResponseMessage.ProtoReflect.Descriptor instead.
func (*DenyResponseMessage) Descriptor() ([]byte, []int) {
	return file_packets_proto_rawDescGZIP(), []int{3}
}

func (x *DenyResponseMessage) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type PlayerMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	X             float64                `protobuf:"fixed64,3,opt,name=x,proto3" json:"x,omitempty"`
	Y             float64                `protobuf:"fixed64,4,opt,name=y,proto3" json:"y,omitempty"`
	Radius        float64                `protobuf:"fixed64,5,opt,name=radius,proto3" json:"radius,omitempty"`
	Direction     float64                `protobuf:"fixed64,6,opt,name=direction,proto3" json:"direction,omitempty"`
	Speed         float64                `protobuf:"fixed64,7,opt,name=speed,proto3" json:"speed,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlayerMessage) Reset() {
	*x = PlayerMessage{}
	mi := &file_packets_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlayerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerMessage) ProtoMessage() {}

func (x *PlayerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_packets_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerMessage.ProtoReflect.Descriptor instead.
func (*PlayerMessage) Descriptor() ([]byte, []int) {
	return file_packets_proto_rawDescGZIP(), []int{4}
}

func (x *PlayerMessage) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PlayerMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PlayerMessage) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *PlayerMessage) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *PlayerMessage) GetRadius() float64 {
	if x != nil {
		return x.Radius
	}
	return 0
}

func (x *PlayerMessage) GetDirection() float64 {
	if x != nil {
		return x.Direction
	}
	return 0
}

func (x *PlayerMessage) GetSpeed() float64 {
	if x != nil {
		return x.Speed
	}
	return 0
}

type PlayerDirectionMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Direction     float64                `protobuf:"fixed64,1,opt,name=direction,proto3" json:"direction,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlayerDirectionMessage) Reset() {
	*x = PlayerDirectionMessage{}
	mi := &file_packets_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlayerDirectionMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerDirectionMessage) ProtoMessage() {}

func (x *PlayerDirectionMessage) ProtoReflect() protoreflect.Message {
	mi := &file_packets_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerDirectionMessage.ProtoReflect.Descriptor instead.
func (*PlayerDirectionMessage) Descriptor() ([]byte, []int) {
	return file_packets_proto_rawDescGZIP(), []int{5}
}

func (x *PlayerDirectionMessage) GetDirection() float64 {
	if x != nil {
		return x.Direction
	}
	return 0
}

type SporeMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	X             float64                `protobuf:"fixed64,2,opt,name=x,proto3" json:"x,omitempty"`
	Y             float64                `protobuf:"fixed64,3,opt,name=y,proto3" json:"y,omitempty"`
	Radius        float64                `protobuf:"fixed64,4,opt,name=radius,proto3" json:"radius,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SporeMessage) Reset() {
	*x = SporeMessage{}
	mi := &file_packets_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SporeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SporeMessage) ProtoMessage() {}

func (x *SporeMessage) ProtoReflect() protoreflect.Message {
	mi := &file_packets_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SporeMessage.ProtoReflect.Descriptor instead.
func (*SporeMessage) Descriptor() ([]byte, []int) {
	return file_packets_proto_rawDescGZIP(), []int{6}
}

func (x *SporeMessage) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SporeMessage) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *SporeMessage) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *SporeMessage) GetRadius() float64 {
	if x != nil {
		return x.Radius
	}
	return 0
}

// Define your messages
type ChatMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Msg           string                 `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChatMessage) Reset() {
	*x = ChatMessage{}
	mi := &file_packets_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessage) ProtoMessage() {}

func (x *ChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_packets_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessage.ProtoReflect.Descriptor instead.
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return file_packets_proto_rawDescGZIP(), []int{7}
}

func (x *ChatMessage) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type IdMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IdMessage) Reset() {
	*x = IdMessage{}
	mi := &file_packets_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IdMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdMessage) ProtoMessage() {}

func (x *IdMessage) ProtoReflect() protoreflect.Message {
	mi := &file_packets_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdMessage.ProtoReflect.Descriptor instead.
func (*IdMessage) Descriptor() ([]byte, []int) {
	return file_packets_proto_rawDescGZIP(), []int{8}
}

func (x *IdMessage) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Define the main Packet message
type Packet struct {
	state    protoimpl.MessageState `protogen:"open.v1"`
	SenderId uint64                 `protobuf:"varint,1,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	// Types that are valid to be assigned to Msg:
	//
	//	*Packet_Chat
	//	*Packet_Id
	//	*Packet_LoginRequest
	//	*Packet_RegisterRequest
	//	*Packet_OkResponse
	//	*Packet_DenyResponse
	//	*Packet_Player
	//	*Packet_PlayerDirection
	//	*Packet_Spore
	Msg           isPacket_Msg `protobuf_oneof:"msg"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Packet) Reset() {
	*x = Packet{}
	mi := &file_packets_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Packet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Packet) ProtoMessage() {}

func (x *Packet) ProtoReflect() protoreflect.Message {
	mi := &file_packets_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Packet.ProtoReflect.Descriptor instead.
func (*Packet) Descriptor() ([]byte, []int) {
	return file_packets_proto_rawDescGZIP(), []int{9}
}

func (x *Packet) GetSenderId() uint64 {
	if x != nil {
		return x.SenderId
	}
	return 0
}

func (x *Packet) GetMsg() isPacket_Msg {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *Packet) GetChat() *ChatMessage {
	if x != nil {
		if x, ok := x.Msg.(*Packet_Chat); ok {
			return x.Chat
		}
	}
	return nil
}

func (x *Packet) GetId() *IdMessage {
	if x != nil {
		if x, ok := x.Msg.(*Packet_Id); ok {
			return x.Id
		}
	}
	return nil
}

func (x *Packet) GetLoginRequest() *LoginRequestMessage {
	if x != nil {
		if x, ok := x.Msg.(*Packet_LoginRequest); ok {
			return x.LoginRequest
		}
	}
	return nil
}

func (x *Packet) GetRegisterRequest() *RegisterRequestMessage {
	if x != nil {
		if x, ok := x.Msg.(*Packet_RegisterRequest); ok {
			return x.RegisterRequest
		}
	}
	return nil
}

func (x *Packet) GetOkResponse() *OkResponseMessage {
	if x != nil {
		if x, ok := x.Msg.(*Packet_OkResponse); ok {
			return x.OkResponse
		}
	}
	return nil
}

func (x *Packet) GetDenyResponse() *DenyResponseMessage {
	if x != nil {
		if x, ok := x.Msg.(*Packet_DenyResponse); ok {
			return x.DenyResponse
		}
	}
	return nil
}

func (x *Packet) GetPlayer() *PlayerMessage {
	if x != nil {
		if x, ok := x.Msg.(*Packet_Player); ok {
			return x.Player
		}
	}
	return nil
}

func (x *Packet) GetPlayerDirection() *PlayerDirectionMessage {
	if x != nil {
		if x, ok := x.Msg.(*Packet_PlayerDirection); ok {
			return x.PlayerDirection
		}
	}
	return nil
}

func (x *Packet) GetSpore() *SporeMessage {
	if x != nil {
		if x, ok := x.Msg.(*Packet_Spore); ok {
			return x.Spore
		}
	}
	return nil
}

type isPacket_Msg interface {
	isPacket_Msg()
}

type Packet_Chat struct {
	Chat *ChatMessage `protobuf:"bytes,2,opt,name=chat,proto3,oneof"`
}

type Packet_Id struct {
	Id *IdMessage `protobuf:"bytes,3,opt,name=id,proto3,oneof"`
}

type Packet_LoginRequest struct {
	LoginRequest *LoginRequestMessage `protobuf:"bytes,4,opt,name=login_request,json=loginRequest,proto3,oneof"`
}

type Packet_RegisterRequest struct {
	RegisterRequest *RegisterRequestMessage `protobuf:"bytes,5,opt,name=register_request,json=registerRequest,proto3,oneof"`
}

type Packet_OkResponse struct {
	OkResponse *OkResponseMessage `protobuf:"bytes,6,opt,name=ok_response,json=okResponse,proto3,oneof"`
}

type Packet_DenyResponse struct {
	DenyResponse *DenyResponseMessage `protobuf:"bytes,7,opt,name=deny_response,json=denyResponse,proto3,oneof"`
}

type Packet_Player struct {
	Player *PlayerMessage `protobuf:"bytes,8,opt,name=player,proto3,oneof"`
}

type Packet_PlayerDirection struct {
	PlayerDirection *PlayerDirectionMessage `protobuf:"bytes,9,opt,name=player_direction,json=playerDirection,proto3,oneof"`
}

type Packet_Spore struct {
	Spore *SporeMessage `protobuf:"bytes,10,opt,name=spore,proto3,oneof"`
}

func (*Packet_Chat) isPacket_Msg() {}

func (*Packet_Id) isPacket_Msg() {}

func (*Packet_LoginRequest) isPacket_Msg() {}

func (*Packet_RegisterRequest) isPacket_Msg() {}

func (*Packet_OkResponse) isPacket_Msg() {}

func (*Packet_DenyResponse) isPacket_Msg() {}

func (*Packet_Player) isPacket_Msg() {}

func (*Packet_PlayerDirection) isPacket_Msg() {}

func (*Packet_Spore) isPacket_Msg() {}

var File_packets_proto protoreflect.FileDescriptor

var file_packets_proto_rawDesc = string([]byte{
	0x0a, 0x0d, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x22, 0x4d, 0x0a, 0x13, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x50, 0x0a, 0x16, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x4f, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2d,
	0x0a, 0x13, 0x44, 0x65, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x9b, 0x01,
	0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01,
	0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x22, 0x36, 0x0a, 0x16, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x52, 0x0a, 0x0c, 0x53, 0x70, 0x6f, 0x72, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01,
	0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x22, 0x1f, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x1b, 0x0a, 0x09, 0x49, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0xc4, 0x04, 0x0a, 0x06, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a,
	0x04, 0x63, 0x68, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x48, 0x00, 0x52, 0x04, 0x63, 0x68, 0x61, 0x74, 0x12, 0x24, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e,
	0x49, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x43, 0x0a, 0x0d, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x10, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48,
	0x00, 0x52, 0x0f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x3d, 0x0a, 0x0b, 0x6f, 0x6b, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x73, 0x2e, 0x4f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x43, 0x0a, 0x0d, 0x64, 0x65, 0x6e, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x73, 0x2e, 0x44, 0x65, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x64, 0x65, 0x6e, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73,
	0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00,
	0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x4c, 0x0a, 0x10, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x05, 0x73, 0x70, 0x6f, 0x72, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e,
	0x53, 0x70, 0x6f, 0x72, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x05,
	0x73, 0x70, 0x6f, 0x72, 0x65, 0x42, 0x05, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x0d, 0x5a, 0x0b,
	0x70, 0x6b, 0x67, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_packets_proto_rawDescOnce sync.Once
	file_packets_proto_rawDescData []byte
)

func file_packets_proto_rawDescGZIP() []byte {
	file_packets_proto_rawDescOnce.Do(func() {
		file_packets_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_packets_proto_rawDesc), len(file_packets_proto_rawDesc)))
	})
	return file_packets_proto_rawDescData
}

var file_packets_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_packets_proto_goTypes = []any{
	(*LoginRequestMessage)(nil),    // 0: packets.LoginRequestMessage
	(*RegisterRequestMessage)(nil), // 1: packets.RegisterRequestMessage
	(*OkResponseMessage)(nil),      // 2: packets.OkResponseMessage
	(*DenyResponseMessage)(nil),    // 3: packets.DenyResponseMessage
	(*PlayerMessage)(nil),          // 4: packets.PlayerMessage
	(*PlayerDirectionMessage)(nil), // 5: packets.PlayerDirectionMessage
	(*SporeMessage)(nil),           // 6: packets.SporeMessage
	(*ChatMessage)(nil),            // 7: packets.ChatMessage
	(*IdMessage)(nil),              // 8: packets.IdMessage
	(*Packet)(nil),                 // 9: packets.Packet
}
var file_packets_proto_depIdxs = []int32{
	7, // 0: packets.Packet.chat:type_name -> packets.ChatMessage
	8, // 1: packets.Packet.id:type_name -> packets.IdMessage
	0, // 2: packets.Packet.login_request:type_name -> packets.LoginRequestMessage
	1, // 3: packets.Packet.register_request:type_name -> packets.RegisterRequestMessage
	2, // 4: packets.Packet.ok_response:type_name -> packets.OkResponseMessage
	3, // 5: packets.Packet.deny_response:type_name -> packets.DenyResponseMessage
	4, // 6: packets.Packet.player:type_name -> packets.PlayerMessage
	5, // 7: packets.Packet.player_direction:type_name -> packets.PlayerDirectionMessage
	6, // 8: packets.Packet.spore:type_name -> packets.SporeMessage
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_packets_proto_init() }
func file_packets_proto_init() {
	if File_packets_proto != nil {
		return
	}
	file_packets_proto_msgTypes[9].OneofWrappers = []any{
		(*Packet_Chat)(nil),
		(*Packet_Id)(nil),
		(*Packet_LoginRequest)(nil),
		(*Packet_RegisterRequest)(nil),
		(*Packet_OkResponse)(nil),
		(*Packet_DenyResponse)(nil),
		(*Packet_Player)(nil),
		(*Packet_PlayerDirection)(nil),
		(*Packet_Spore)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_packets_proto_rawDesc), len(file_packets_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_packets_proto_goTypes,
		DependencyIndexes: file_packets_proto_depIdxs,
		MessageInfos:      file_packets_proto_msgTypes,
	}.Build()
	File_packets_proto = out.File
	file_packets_proto_goTypes = nil
	file_packets_proto_depIdxs = nil
}
