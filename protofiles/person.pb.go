// Code generated by protoc-gen-go. DO NOT EDIT.
// source: person.proto

package protofiles

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

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}

var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}

func (Person_PhoneType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{0, 0}
}

type Person struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   int32                 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Email                string                `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phones               []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones,proto3" json:"phones,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{0}
}
func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

type Person_PhoneNumber struct {
	Number               string           `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Type                 Person_PhoneType `protobuf:"varint,2,opt,name=type,proto3,enum=protofiles.Person_PhoneType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Person_PhoneNumber) Reset()         { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()    {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{0, 0}
}
func (m *Person_PhoneNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person_PhoneNumber.Unmarshal(m, b)
}
func (m *Person_PhoneNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person_PhoneNumber.Marshal(b, m, deterministic)
}
func (m *Person_PhoneNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person_PhoneNumber.Merge(m, src)
}
func (m *Person_PhoneNumber) XXX_Size() int {
	return xxx_messageInfo_Person_PhoneNumber.Size(m)
}
func (m *Person_PhoneNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_Person_PhoneNumber.DiscardUnknown(m)
}

var xxx_messageInfo_Person_PhoneNumber proto.InternalMessageInfo

func (m *Person_PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Person_PhoneNumber) GetType() Person_PhoneType {
	if m != nil {
		return m.Type
	}
	return Person_MOBILE
}

// Our address book file is just one of these.
type AddressBook struct {
	People               []*Person `protobuf:"bytes,1,rep,name=people,proto3" json:"people,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AddressBook) Reset()         { *m = AddressBook{} }
func (m *AddressBook) String() string { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()    {}
func (*AddressBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{1}
}
func (m *AddressBook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressBook.Unmarshal(m, b)
}
func (m *AddressBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressBook.Marshal(b, m, deterministic)
}
func (m *AddressBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressBook.Merge(m, src)
}
func (m *AddressBook) XXX_Size() int {
	return xxx_messageInfo_AddressBook.Size(m)
}
func (m *AddressBook) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressBook.DiscardUnknown(m)
}

var xxx_messageInfo_AddressBook proto.InternalMessageInfo

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterType((*Person)(nil), "protofiles.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "protofiles.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "protofiles.AddressBook")
	proto.RegisterEnum("protofiles.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
}

func init() { proto.RegisterFile("person.proto", fileDescriptor_4c9e10cf24b1156d) }

var fileDescriptor_4c9e10cf24b1156d = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xcd, 0x36, 0x5d, 0xec, 0x8b, 0x94, 0x30, 0x88, 0x04, 0x11, 0x09, 0x39, 0x05, 0x85,
	0x20, 0x15, 0x04, 0x8f, 0x16, 0x0a, 0x8a, 0xd6, 0x94, 0x45, 0xe8, 0xb9, 0x25, 0x23, 0x06, 0x93,
	0xec, 0x92, 0xad, 0x87, 0xde, 0xfc, 0xe9, 0x92, 0x4d, 0x50, 0x41, 0x3c, 0xed, 0x9b, 0x99, 0x8f,
	0x79, 0x6f, 0x07, 0x47, 0x86, 0x5b, 0xab, 0x9b, 0xcc, 0xb4, 0x7a, 0xa7, 0x09, 0xee, 0x79, 0x2d,
	0x2b, 0xb6, 0xc9, 0xa7, 0x80, 0x5c, 0xb9, 0x21, 0x11, 0xfc, 0x66, 0x53, 0x73, 0xe4, 0xc5, 0x5e,
	0x3a, 0x51, 0x4e, 0xd3, 0x14, 0xa2, 0x2c, 0x22, 0x11, 0x7b, 0xe9, 0x58, 0x89, 0xb2, 0xa0, 0x63,
	0x8c, 0xb9, 0xde, 0x94, 0x55, 0x34, 0x72, 0x50, 0x5f, 0xd0, 0x0d, 0xa4, 0x79, 0xd3, 0x0d, 0xdb,
	0xc8, 0x8f, 0x47, 0x69, 0x30, 0x3b, 0xcf, 0x7e, 0x1c, 0xb2, 0x7e, 0x7b, 0xb6, 0xea, 0x80, 0xe7,
	0x8f, 0x7a, 0xcb, 0xad, 0x1a, 0xe8, 0xd3, 0x35, 0x82, 0x5f, 0x6d, 0x3a, 0x81, 0x6c, 0x9c, 0x1a,
	0x22, 0x0c, 0x15, 0x5d, 0xc1, 0xdf, 0xed, 0x0d, 0xbb, 0x18, 0xd3, 0xd9, 0xd9, 0x7f, 0xcb, 0x5f,
	0xf6, 0x86, 0x95, 0x23, 0x93, 0x4b, 0x4c, 0xbe, 0x5b, 0x04, 0xc8, 0x65, 0x3e, 0x7f, 0x78, 0x5a,
	0x84, 0x07, 0x74, 0x08, 0xff, 0x3e, 0x5f, 0x2e, 0x42, 0xaf, 0x53, 0xeb, 0x5c, 0x3d, 0x86, 0x22,
	0xb9, 0x45, 0x70, 0x57, 0x14, 0x2d, 0x5b, 0x3b, 0xd7, 0xfa, 0x9d, 0x2e, 0x20, 0x0d, 0x6b, 0x53,
	0x75, 0x87, 0xe8, 0x3e, 0x43, 0x7f, 0xfd, 0xd4, 0x40, 0x6c, 0xa5, 0x1b, 0x5d, 0x7f, 0x05, 0x00,
	0x00, 0xff, 0xff, 0x7e, 0x71, 0x4d, 0x40, 0x60, 0x01, 0x00, 0x00,
}
