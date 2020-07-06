// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/pbMessage.proto

/*
Package protobuf_pbMessage is a generated protocol buffer package.

It is generated from these files:
	protobuf/pbMessage.proto

It has these top-level messages:
	Person
*/
package protobuf_pbMessage

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

type Person struct {
	Year        int32  `protobuf:"varint,1,opt,name=year" json:"year,omitempty"`
	Month       string `protobuf:"bytes,2,opt,name=month" json:"month,omitempty"`
	Day         int32  `protobuf:"varint,3,opt,name=day" json:"day,omitempty"`
	FullName    string `protobuf:"bytes,4,opt,name=fullName" json:"fullName,omitempty"`
	FirstName   string `protobuf:"bytes,5,opt,name=firstName" json:"firstName,omitempty"`
	LastName    string `protobuf:"bytes,6,opt,name=lastName" json:"lastName,omitempty"`
	Dob         string `protobuf:"bytes,7,opt,name=dob" json:"dob,omitempty"`
	Country     string `protobuf:"bytes,8,opt,name=country" json:"country,omitempty"`
	Description string `protobuf:"bytes,9,opt,name=description" json:"description,omitempty"`
	Weekday     string `protobuf:"bytes,10,opt,name=weekday" json:"weekday,omitempty"`
}

func (m *Person) Reset()                    { *m = Person{} }
func (m *Person) String() string            { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()               {}
func (*Person) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Person) GetYear() int32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func (m *Person) GetMonth() string {
	if m != nil {
		return m.Month
	}
	return ""
}

func (m *Person) GetDay() int32 {
	if m != nil {
		return m.Day
	}
	return 0
}

func (m *Person) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *Person) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Person) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Person) GetDob() string {
	if m != nil {
		return m.Dob
	}
	return ""
}

func (m *Person) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Person) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Person) GetWeekday() string {
	if m != nil {
		return m.Weekday
	}
	return ""
}

func init() {
	proto.RegisterType((*Person)(nil), "protobuf.pbMessage.Person")
}

func init() { proto.RegisterFile("protobuf/pbMessage.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xcf, 0x4e, 0x84, 0x30,
	0x10, 0xc6, 0xc3, 0xee, 0xc2, 0x2e, 0xe3, 0xc5, 0x4c, 0x3c, 0x4c, 0x8c, 0x07, 0xe2, 0x69, 0x4f,
	0x78, 0xf0, 0x39, 0x34, 0x86, 0x37, 0x68, 0x61, 0x50, 0x22, 0xb4, 0xa4, 0x2d, 0x31, 0x5c, 0x7d,
	0x72, 0xd3, 0xe1, 0x8f, 0x7b, 0xfb, 0xbe, 0xdf, 0x6f, 0xbe, 0x40, 0x0a, 0x34, 0x3a, 0x1b, 0xac,
	0x9e, 0xda, 0x97, 0x51, 0xbf, 0xb1, 0xf7, 0xea, 0x93, 0x4b, 0x41, 0x88, 0x9b, 0x29, 0x77, 0xf3,
	0xfc, 0x7b, 0x80, 0xec, 0x83, 0x9d, 0xb7, 0x06, 0x11, 0x4e, 0x33, 0x2b, 0x47, 0x49, 0x91, 0x5c,
	0xd3, 0x4a, 0x32, 0x3e, 0x40, 0x3a, 0x58, 0x13, 0xbe, 0xe8, 0x50, 0x24, 0xd7, 0xbc, 0x5a, 0x0a,
	0xde, 0xc3, 0xb1, 0x51, 0x33, 0x1d, 0xe5, 0x30, 0x46, 0x7c, 0x84, 0x4b, 0x3b, 0xf5, 0xfd, 0xbb,
	0x1a, 0x98, 0x4e, 0x72, 0xba, 0x77, 0x7c, 0x82, 0xbc, 0xed, 0x9c, 0x0f, 0x22, 0x53, 0x91, 0xff,
	0x20, 0x2e, 0x7b, 0xb5, 0xca, 0x6c, 0x59, 0x6e, 0x5d, 0xbe, 0x63, 0x35, 0x9d, 0x05, 0xc7, 0x88,
	0x04, 0xe7, 0xda, 0x4e, 0x26, 0xb8, 0x99, 0x2e, 0x42, 0xb7, 0x8a, 0x05, 0xdc, 0x35, 0xec, 0x6b,
	0xd7, 0x8d, 0xa1, 0xb3, 0x86, 0x72, 0xb1, 0xb7, 0x28, 0x6e, 0x7f, 0x98, 0xbf, 0xe3, 0x9f, 0xc3,
	0xb2, 0x5d, 0xab, 0xce, 0xe4, 0x61, 0x5e, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x32, 0x92,
	0x4a, 0x3b, 0x01, 0x00, 0x00,
}
