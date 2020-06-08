// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mwsconfig/mwsconfig.proto

package mwsconfig

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MWSConfig struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserID               string   `protobuf:"bytes,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	SellerID             string   `protobuf:"bytes,3,opt,name=SellerID,proto3" json:"SellerID,omitempty"`
	MarketplaceID        string   `protobuf:"bytes,4,opt,name=MarketplaceID,proto3" json:"MarketplaceID,omitempty"`
	BaseURL              string   `protobuf:"bytes,5,opt,name=BaseURL,proto3" json:"BaseURL,omitempty"`
	SignatureVersion     string   `protobuf:"bytes,6,opt,name=SignatureVersion,proto3" json:"SignatureVersion,omitempty"`
	SignatureMethod      string   `protobuf:"bytes,7,opt,name=SignatureMethod,proto3" json:"SignatureMethod,omitempty"`
	AuthToken            string   `protobuf:"bytes,8,opt,name=AuthToken,proto3" json:"AuthToken,omitempty"`
	AccessKey            string   `protobuf:"bytes,9,opt,name=AccessKey,proto3" json:"AccessKey,omitempty"`
	AccessSecret         string   `protobuf:"bytes,10,opt,name=AccessSecret,proto3" json:"AccessSecret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MWSConfig) Reset()         { *m = MWSConfig{} }
func (m *MWSConfig) String() string { return proto.CompactTextString(m) }
func (*MWSConfig) ProtoMessage()    {}
func (*MWSConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_07b380766a1cf088, []int{0}
}

func (m *MWSConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MWSConfig.Unmarshal(m, b)
}
func (m *MWSConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MWSConfig.Marshal(b, m, deterministic)
}
func (m *MWSConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MWSConfig.Merge(m, src)
}
func (m *MWSConfig) XXX_Size() int {
	return xxx_messageInfo_MWSConfig.Size(m)
}
func (m *MWSConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MWSConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MWSConfig proto.InternalMessageInfo

func (m *MWSConfig) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *MWSConfig) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *MWSConfig) GetSellerID() string {
	if m != nil {
		return m.SellerID
	}
	return ""
}

func (m *MWSConfig) GetMarketplaceID() string {
	if m != nil {
		return m.MarketplaceID
	}
	return ""
}

func (m *MWSConfig) GetBaseURL() string {
	if m != nil {
		return m.BaseURL
	}
	return ""
}

func (m *MWSConfig) GetSignatureVersion() string {
	if m != nil {
		return m.SignatureVersion
	}
	return ""
}

func (m *MWSConfig) GetSignatureMethod() string {
	if m != nil {
		return m.SignatureMethod
	}
	return ""
}

func (m *MWSConfig) GetAuthToken() string {
	if m != nil {
		return m.AuthToken
	}
	return ""
}

func (m *MWSConfig) GetAccessKey() string {
	if m != nil {
		return m.AccessKey
	}
	return ""
}

func (m *MWSConfig) GetAccessSecret() string {
	if m != nil {
		return m.AccessSecret
	}
	return ""
}

func init() {
	proto.RegisterType((*MWSConfig)(nil), "report.MWSConfig")
}

func init() { proto.RegisterFile("mwsconfig/mwsconfig.proto", fileDescriptor_07b380766a1cf088) }

var fileDescriptor_07b380766a1cf088 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0x87, 0x69, 0xd4, 0xb4, 0x19, 0xfc, 0xc7, 0x1e, 0x64, 0x14, 0x0f, 0x52, 0x3c, 0x94, 0x22,
	0xcd, 0xc1, 0x27, 0xb0, 0xe6, 0x12, 0x34, 0x97, 0xc6, 0x2a, 0x78, 0x4b, 0xd7, 0x69, 0x12, 0x9a,
	0xee, 0x86, 0xdd, 0x0d, 0xd2, 0x07, 0xf4, 0xbd, 0xa4, 0x63, 0x9b, 0x52, 0x7b, 0xcb, 0xef, 0xfb,
	0x3e, 0x08, 0xcc, 0xc2, 0xf5, 0xf2, 0xdb, 0x4a, 0xad, 0xe6, 0x65, 0x1e, 0xb6, 0x5f, 0xa3, 0xda,
	0x68, 0xa7, 0x85, 0x6f, 0xa8, 0xd6, 0xc6, 0xf5, 0x7f, 0x3c, 0x08, 0x92, 0x8f, 0xf4, 0x99, 0x9d,
	0x38, 0x07, 0x2f, 0x8e, 0xb0, 0x73, 0xd7, 0x19, 0x04, 0x13, 0x2f, 0x8e, 0xc4, 0x15, 0xf8, 0x53,
	0x4b, 0x26, 0x8e, 0xd0, 0x63, 0xb6, 0x59, 0xe2, 0x06, 0x7a, 0x29, 0x55, 0x15, 0x9b, 0x23, 0x36,
	0xed, 0x16, 0xf7, 0x70, 0x96, 0x64, 0x66, 0x41, 0xae, 0xae, 0x32, 0x49, 0x71, 0x84, 0xc7, 0x1c,
	0xec, 0x43, 0x81, 0xd0, 0x1d, 0x67, 0x96, 0xa6, 0x93, 0x57, 0x3c, 0x61, 0xbf, 0x9d, 0x62, 0x08,
	0x97, 0x69, 0x99, 0xab, 0xcc, 0x35, 0x86, 0xde, 0xc9, 0xd8, 0x52, 0x2b, 0xf4, 0x39, 0x39, 0xe0,
	0x62, 0x00, 0x17, 0x2d, 0x4b, 0xc8, 0x15, 0xfa, 0x0b, 0xbb, 0x9c, 0xfe, 0xc7, 0xe2, 0x16, 0x82,
	0xa7, 0xc6, 0x15, 0x6f, 0x7a, 0x41, 0x0a, 0x7b, 0xdc, 0xec, 0x00, 0x5b, 0x29, 0xc9, 0xda, 0x17,
	0x5a, 0x61, 0xb0, 0xb1, 0x5b, 0x20, 0xfa, 0x70, 0xfa, 0x37, 0x52, 0x92, 0x86, 0x1c, 0x02, 0x07,
	0x7b, 0x6c, 0xfc, 0xf0, 0x39, 0xcc, 0x4b, 0x57, 0x34, 0xb3, 0x91, 0xd4, 0xcb, 0xd0, 0xae, 0x94,
	0x9c, 0x37, 0xeb, 0xdf, 0xaf, 0x0f, 0x1f, 0xf2, 0xc9, 0xe5, 0xee, 0x0d, 0x66, 0x3e, 0x93, 0xc7,
	0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x92, 0xc5, 0x2f, 0x47, 0xa1, 0x01, 0x00, 0x00,
}