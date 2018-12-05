// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Account struct {
	Address              []byte   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{0}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

type AccountList struct {
	Accounts             []*Account `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *AccountList) Reset()         { *m = AccountList{} }
func (m *AccountList) String() string { return proto.CompactTextString(m) }
func (*AccountList) ProtoMessage()    {}
func (*AccountList) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{1}
}
func (m *AccountList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountList.Unmarshal(m, b)
}
func (m *AccountList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountList.Marshal(b, m, deterministic)
}
func (m *AccountList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountList.Merge(m, src)
}
func (m *AccountList) XXX_Size() int {
	return xxx_messageInfo_AccountList.Size(m)
}
func (m *AccountList) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountList.DiscardUnknown(m)
}

var xxx_messageInfo_AccountList proto.InternalMessageInfo

func (m *AccountList) GetAccounts() []*Account {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func init() {
	proto.RegisterType((*Account)(nil), "types.Account")
	proto.RegisterType((*AccountList)(nil), "types.AccountList")
}

func init() { proto.RegisterFile("account.proto", fileDescriptor_8e28828dcb8d24f0) }

var fileDescriptor_8e28828dcb8d24f0 = []byte{
	// 137 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4c, 0x4e, 0xce,
	0x2f, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d, 0xa9, 0x2c, 0x48, 0x2d,
	0x56, 0x52, 0xe6, 0x62, 0x77, 0x84, 0x88, 0x0b, 0x49, 0x70, 0xb1, 0x27, 0xa6, 0xa4, 0x14, 0xa5,
	0x16, 0x17, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x04, 0xc1, 0xb8, 0x4a, 0x96, 0x5c, 0xdc, 0x50,
	0x45, 0x3e, 0x99, 0xc5, 0x25, 0x42, 0x5a, 0x5c, 0x1c, 0x50, 0xb3, 0x40, 0x2a, 0x99, 0x35, 0xb8,
	0x8d, 0xf8, 0xf4, 0xc0, 0xa6, 0xe9, 0x41, 0x55, 0x05, 0xc1, 0xe5, 0x9d, 0x14, 0xa2, 0xe4, 0xd2,
	0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x13, 0x53, 0x8b, 0xd2, 0xf3, 0x33,
	0xf3, 0x21, 0xb4, 0x3e, 0x58, 0x4f, 0x12, 0x1b, 0xd8, 0x3d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xdc, 0x1a, 0xad, 0xeb, 0xa0, 0x00, 0x00, 0x00,
}
