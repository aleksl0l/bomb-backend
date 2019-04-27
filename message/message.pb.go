// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package message

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MessageType int32

const (
	MessageType_JOIN_GAME               MessageType = 0
	MessageType_LEAVE_GAME              MessageType = 1
	MessageType_ACTION                  MessageType = 2
	MessageType_MAP_RESPONSE            MessageType = 3
	MessageType_PLAYER_POSITION_REQUEST MessageType = 4
	MessageType_CREATE_GAME_REQUEST     MessageType = 5
)

var MessageType_name = map[int32]string{
	0: "JOIN_GAME",
	1: "LEAVE_GAME",
	2: "ACTION",
	3: "MAP_RESPONSE",
	4: "PLAYER_POSITION_REQUEST",
	5: "CREATE_GAME_REQUEST",
}

var MessageType_value = map[string]int32{
	"JOIN_GAME":               0,
	"LEAVE_GAME":              1,
	"ACTION":                  2,
	"MAP_RESPONSE":            3,
	"PLAYER_POSITION_REQUEST": 4,
	"CREATE_GAME_REQUEST":     5,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}

func (MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

type TypeObject int32

const (
	TypeObject_BREAKABLE_BLOCK  TypeObject = 0
	TypeObject_STONE_BLOCK      TypeObject = 1
	TypeObject_SPEED_BUFF       TypeObject = 2
	TypeObject_FIRE_BUFF        TypeObject = 3
	TypeObject_BOMB_NUMBER_BUFF TypeObject = 4
)

var TypeObject_name = map[int32]string{
	0: "BREAKABLE_BLOCK",
	1: "STONE_BLOCK",
	2: "SPEED_BUFF",
	3: "FIRE_BUFF",
	4: "BOMB_NUMBER_BUFF",
}

var TypeObject_value = map[string]int32{
	"BREAKABLE_BLOCK":  0,
	"STONE_BLOCK":      1,
	"SPEED_BUFF":       2,
	"FIRE_BUFF":        3,
	"BOMB_NUMBER_BUFF": 4,
}

func (x TypeObject) String() string {
	return proto.EnumName(TypeObject_name, int32(x))
}

func (TypeObject) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

type Message struct {
	Type                 MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=message.MessageType" json:"type,omitempty"`
	Content              []byte      `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_JOIN_GAME
}

func (m *Message) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type JoinGame struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinGame) Reset()         { *m = JoinGame{} }
func (m *JoinGame) String() string { return proto.CompactTextString(m) }
func (*JoinGame) ProtoMessage()    {}
func (*JoinGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *JoinGame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinGame.Unmarshal(m, b)
}
func (m *JoinGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinGame.Marshal(b, m, deterministic)
}
func (m *JoinGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinGame.Merge(m, src)
}
func (m *JoinGame) XXX_Size() int {
	return xxx_messageInfo_JoinGame.Size(m)
}
func (m *JoinGame) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinGame.DiscardUnknown(m)
}

var xxx_messageInfo_JoinGame proto.InternalMessageInfo

func (m *JoinGame) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *JoinGame) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type LeaveGame struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaveGame) Reset()         { *m = LeaveGame{} }
func (m *LeaveGame) String() string { return proto.CompactTextString(m) }
func (*LeaveGame) ProtoMessage()    {}
func (*LeaveGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{2}
}

func (m *LeaveGame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaveGame.Unmarshal(m, b)
}
func (m *LeaveGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaveGame.Marshal(b, m, deterministic)
}
func (m *LeaveGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveGame.Merge(m, src)
}
func (m *LeaveGame) XXX_Size() int {
	return xxx_messageInfo_LeaveGame.Size(m)
}
func (m *LeaveGame) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveGame.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveGame proto.InternalMessageInfo

func (m *LeaveGame) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type Point struct {
	X                    int32    `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int32    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}
func (*Point) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{3}
}

func (m *Point) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Point.Unmarshal(m, b)
}
func (m *Point) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Point.Marshal(b, m, deterministic)
}
func (m *Point) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Point.Merge(m, src)
}
func (m *Point) XXX_Size() int {
	return xxx_messageInfo_Point.Size(m)
}
func (m *Point) XXX_DiscardUnknown() {
	xxx_messageInfo_Point.DiscardUnknown(m)
}

var xxx_messageInfo_Point proto.InternalMessageInfo

func (m *Point) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Point) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

type GameObject struct {
	Type                 TypeObject `protobuf:"varint,1,opt,name=type,proto3,enum=message.TypeObject" json:"type,omitempty"`
	HealthPoints         uint32     `protobuf:"varint,2,opt,name=health_points,json=healthPoints,proto3" json:"health_points,omitempty"`
	ID                   int32      `protobuf:"varint,3,opt,name=ID,proto3" json:"ID,omitempty"`
	Position             *Point     `protobuf:"bytes,4,opt,name=position,proto3" json:"position,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GameObject) Reset()         { *m = GameObject{} }
func (m *GameObject) String() string { return proto.CompactTextString(m) }
func (*GameObject) ProtoMessage()    {}
func (*GameObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{4}
}

func (m *GameObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameObject.Unmarshal(m, b)
}
func (m *GameObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameObject.Marshal(b, m, deterministic)
}
func (m *GameObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameObject.Merge(m, src)
}
func (m *GameObject) XXX_Size() int {
	return xxx_messageInfo_GameObject.Size(m)
}
func (m *GameObject) XXX_DiscardUnknown() {
	xxx_messageInfo_GameObject.DiscardUnknown(m)
}

var xxx_messageInfo_GameObject proto.InternalMessageInfo

func (m *GameObject) GetType() TypeObject {
	if m != nil {
		return m.Type
	}
	return TypeObject_BREAKABLE_BLOCK
}

func (m *GameObject) GetHealthPoints() uint32 {
	if m != nil {
		return m.HealthPoints
	}
	return 0
}

func (m *GameObject) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *GameObject) GetPosition() *Point {
	if m != nil {
		return m.Position
	}
	return nil
}

type PlayerPosition struct {
	X                    float32  `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    float32  `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayerPosition) Reset()         { *m = PlayerPosition{} }
func (m *PlayerPosition) String() string { return proto.CompactTextString(m) }
func (*PlayerPosition) ProtoMessage()    {}
func (*PlayerPosition) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{5}
}

func (m *PlayerPosition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerPosition.Unmarshal(m, b)
}
func (m *PlayerPosition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerPosition.Marshal(b, m, deterministic)
}
func (m *PlayerPosition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerPosition.Merge(m, src)
}
func (m *PlayerPosition) XXX_Size() int {
	return xxx_messageInfo_PlayerPosition.Size(m)
}
func (m *PlayerPosition) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerPosition.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerPosition proto.InternalMessageInfo

func (m *PlayerPosition) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *PlayerPosition) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

type Player struct {
	Username             string          `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Score                uint32          `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
	Position             *PlayerPosition `protobuf:"bytes,3,opt,name=position,proto3" json:"position,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}
func (*Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{6}
}

func (m *Player) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Player.Unmarshal(m, b)
}
func (m *Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Player.Marshal(b, m, deterministic)
}
func (m *Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Player.Merge(m, src)
}
func (m *Player) XXX_Size() int {
	return xxx_messageInfo_Player.Size(m)
}
func (m *Player) XXX_DiscardUnknown() {
	xxx_messageInfo_Player.DiscardUnknown(m)
}

var xxx_messageInfo_Player proto.InternalMessageInfo

func (m *Player) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Player) GetScore() uint32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *Player) GetPosition() *PlayerPosition {
	if m != nil {
		return m.Position
	}
	return nil
}

type GameMap struct {
	GameObjects          []*GameObject `protobuf:"bytes,1,rep,name=gameObjects,proto3" json:"gameObjects,omitempty"`
	Players              []*Player     `protobuf:"bytes,2,rep,name=players,proto3" json:"players,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GameMap) Reset()         { *m = GameMap{} }
func (m *GameMap) String() string { return proto.CompactTextString(m) }
func (*GameMap) ProtoMessage()    {}
func (*GameMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{7}
}

func (m *GameMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameMap.Unmarshal(m, b)
}
func (m *GameMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameMap.Marshal(b, m, deterministic)
}
func (m *GameMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameMap.Merge(m, src)
}
func (m *GameMap) XXX_Size() int {
	return xxx_messageInfo_GameMap.Size(m)
}
func (m *GameMap) XXX_DiscardUnknown() {
	xxx_messageInfo_GameMap.DiscardUnknown(m)
}

var xxx_messageInfo_GameMap proto.InternalMessageInfo

func (m *GameMap) GetGameObjects() []*GameObject {
	if m != nil {
		return m.GameObjects
	}
	return nil
}

func (m *GameMap) GetPlayers() []*Player {
	if m != nil {
		return m.Players
	}
	return nil
}

type PlayerPositionRequest struct {
	Position             *PlayerPosition `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	Token                string          `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *PlayerPositionRequest) Reset()         { *m = PlayerPositionRequest{} }
func (m *PlayerPositionRequest) String() string { return proto.CompactTextString(m) }
func (*PlayerPositionRequest) ProtoMessage()    {}
func (*PlayerPositionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{8}
}

func (m *PlayerPositionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerPositionRequest.Unmarshal(m, b)
}
func (m *PlayerPositionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerPositionRequest.Marshal(b, m, deterministic)
}
func (m *PlayerPositionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerPositionRequest.Merge(m, src)
}
func (m *PlayerPositionRequest) XXX_Size() int {
	return xxx_messageInfo_PlayerPositionRequest.Size(m)
}
func (m *PlayerPositionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerPositionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerPositionRequest proto.InternalMessageInfo

func (m *PlayerPositionRequest) GetPosition() *PlayerPosition {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *PlayerPositionRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterEnum("message.MessageType", MessageType_name, MessageType_value)
	proto.RegisterEnum("message.TypeObject", TypeObject_name, TypeObject_value)
	proto.RegisterType((*Message)(nil), "message.Message")
	proto.RegisterType((*JoinGame)(nil), "message.JoinGame")
	proto.RegisterType((*LeaveGame)(nil), "message.LeaveGame")
	proto.RegisterType((*Point)(nil), "message.Point")
	proto.RegisterType((*GameObject)(nil), "message.GameObject")
	proto.RegisterType((*PlayerPosition)(nil), "message.PlayerPosition")
	proto.RegisterType((*Player)(nil), "message.Player")
	proto.RegisterType((*GameMap)(nil), "message.GameMap")
	proto.RegisterType((*PlayerPositionRequest)(nil), "message.PlayerPositionRequest")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 539 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4f, 0x6f, 0xda, 0x4e,
	0x10, 0xcd, 0x1a, 0x08, 0x30, 0xfc, 0xb3, 0x16, 0x7e, 0xc2, 0xfa, 0xf5, 0x42, 0x9d, 0x43, 0x29,
	0xaa, 0x72, 0x48, 0xd4, 0x5b, 0x2f, 0x36, 0x2c, 0xc8, 0x09, 0xc6, 0xee, 0x1a, 0x2a, 0xf5, 0x64,
	0x19, 0xb4, 0x22, 0x34, 0xc1, 0x76, 0xb1, 0x53, 0x85, 0x5b, 0x3f, 0x45, 0x3f, 0x6f, 0xb5, 0xbb,
	0x60, 0x4c, 0x54, 0xa9, 0xbd, 0xf9, 0xcd, 0xbc, 0x99, 0x79, 0xef, 0x69, 0x0d, 0x8d, 0x2d, 0x4b,
	0x92, 0x60, 0xcd, 0xae, 0xe3, 0x5d, 0x94, 0x46, 0xb8, 0x7c, 0x80, 0xba, 0x0d, 0x65, 0x5b, 0x7e,
	0xe2, 0x3e, 0x14, 0xd3, 0x7d, 0xcc, 0x34, 0xd4, 0x43, 0xfd, 0xe6, 0x4d, 0xe7, 0xfa, 0x38, 0x71,
	0xe8, 0xcf, 0xf7, 0x31, 0xa3, 0x82, 0x81, 0x35, 0x28, 0xaf, 0xa2, 0x30, 0x65, 0x61, 0xaa, 0x29,
	0x3d, 0xd4, 0xaf, 0xd3, 0x23, 0xd4, 0x3f, 0x41, 0xe5, 0x2e, 0xda, 0x84, 0x93, 0x60, 0xcb, 0xf0,
	0xff, 0x50, 0x79, 0x4e, 0xd8, 0x2e, 0x0c, 0xb6, 0x72, 0x67, 0x95, 0x66, 0x18, 0x77, 0xa0, 0x94,
	0x46, 0x8f, 0x2c, 0x14, 0xf3, 0x55, 0x2a, 0x81, 0xfe, 0x16, 0xaa, 0x53, 0x16, 0xfc, 0x60, 0x93,
	0x33, 0x0a, 0xca, 0x53, 0xae, 0xa0, 0xe4, 0x46, 0x9b, 0x30, 0xc5, 0x75, 0x40, 0x2f, 0xa2, 0x55,
	0xa2, 0xe8, 0x85, 0xa3, 0xbd, 0xd8, 0x55, 0xa2, 0x68, 0xaf, 0xff, 0x42, 0x00, 0x7c, 0x87, 0xb3,
	0xfc, 0xc6, 0x56, 0x29, 0x7e, 0x77, 0x66, 0xac, 0x9d, 0x19, 0xe3, 0x8e, 0x24, 0xe5, 0xe0, 0xeb,
	0x0a, 0x1a, 0x0f, 0x2c, 0x78, 0x4a, 0x1f, 0xfc, 0x98, 0xdf, 0x48, 0xc4, 0xc6, 0x06, 0xad, 0xcb,
	0xa2, 0xb8, 0x9b, 0xe0, 0x26, 0x28, 0xd6, 0x48, 0x2b, 0x88, 0x5b, 0x8a, 0x35, 0xc2, 0x03, 0xa8,
	0xc4, 0x51, 0xb2, 0x49, 0x37, 0x51, 0xa8, 0x15, 0x7b, 0xa8, 0x5f, 0xbb, 0x69, 0x66, 0x17, 0xc4,
	0x08, 0xcd, 0xfa, 0xfa, 0x07, 0x68, 0xba, 0x4f, 0xc1, 0x9e, 0xed, 0xdc, 0x43, 0xe5, 0x64, 0x43,
	0x39, 0xb3, 0xa1, 0x70, 0x1b, 0x11, 0x5c, 0x4a, 0xf6, 0xdf, 0xa2, 0x4c, 0x56, 0xd1, 0x8e, 0x1d,
	0xc4, 0x4a, 0x80, 0x6f, 0x73, 0xaa, 0x0a, 0x42, 0x55, 0xf7, 0xa4, 0xea, 0x4c, 0x42, 0x4e, 0xde,
	0x23, 0x94, 0x79, 0x6c, 0x76, 0x10, 0xe3, 0x8f, 0x50, 0x5b, 0x67, 0x09, 0x26, 0x1a, 0xea, 0x15,
	0xfa, 0xb5, 0x5c, 0x74, 0xa7, 0x74, 0x69, 0x9e, 0x87, 0xdf, 0x43, 0x39, 0x16, 0xdb, 0x79, 0x76,
	0x7c, 0xa4, 0xf5, 0xea, 0x2a, 0x3d, 0xf6, 0xf5, 0x25, 0xfc, 0xf7, 0x4a, 0x08, 0xfb, 0xfe, 0xcc,
	0x92, 0xf4, 0x4c, 0x3a, 0xfa, 0x47, 0xe9, 0x7f, 0x7e, 0x50, 0x83, 0x9f, 0x08, 0x6a, 0xb9, 0xe7,
	0x8b, 0x1b, 0x50, 0xbd, 0x73, 0xac, 0x99, 0x3f, 0x31, 0x6c, 0xa2, 0x5e, 0xe0, 0x26, 0xc0, 0x94,
	0x18, 0x5f, 0x88, 0xc4, 0x08, 0x03, 0x5c, 0x1a, 0xc3, 0xb9, 0xe5, 0xcc, 0x54, 0x05, 0xab, 0x50,
	0xb7, 0x0d, 0xd7, 0xa7, 0xc4, 0x73, 0x9d, 0x99, 0x47, 0xd4, 0x02, 0x7e, 0x03, 0x5d, 0x77, 0x6a,
	0x7c, 0x25, 0xd4, 0x77, 0x1d, 0xcf, 0xe2, 0x34, 0x9f, 0x92, 0xcf, 0x0b, 0xe2, 0xcd, 0xd5, 0x22,
	0xee, 0x42, 0x7b, 0x48, 0x89, 0x31, 0x97, 0xbb, 0xb2, 0x46, 0x69, 0xb0, 0x06, 0x38, 0xbd, 0x33,
	0xdc, 0x86, 0x96, 0x49, 0x89, 0x71, 0x6f, 0x98, 0x53, 0xe2, 0x9b, 0x53, 0x67, 0x78, 0xaf, 0x5e,
	0xe0, 0x16, 0xd4, 0xbc, 0xb9, 0x33, 0x3b, 0x16, 0x10, 0xd7, 0xe5, 0xb9, 0x84, 0x8c, 0x7c, 0x73,
	0x31, 0x1e, 0xab, 0x0a, 0x97, 0x3d, 0xb6, 0x28, 0x91, 0xb0, 0x80, 0x3b, 0xa0, 0x9a, 0x8e, 0x6d,
	0xfa, 0xb3, 0x85, 0x6d, 0x12, 0x2a, 0xab, 0xc5, 0xe5, 0xa5, 0xf8, 0xb3, 0x6f, 0x7f, 0x07, 0x00,
	0x00, 0xff, 0xff, 0xa0, 0xa1, 0xe4, 0x0e, 0xea, 0x03, 0x00, 0x00,
}
