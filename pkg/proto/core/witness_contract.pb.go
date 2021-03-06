// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/contract/witness_contract.proto

package core

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

type WitnessCreateContract struct {
	OwnerAddress         []byte   `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	Url                  []byte   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WitnessCreateContract) Reset()         { *m = WitnessCreateContract{} }
func (m *WitnessCreateContract) String() string { return proto.CompactTextString(m) }
func (*WitnessCreateContract) ProtoMessage()    {}
func (*WitnessCreateContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_083020148d0431a0, []int{0}
}

func (m *WitnessCreateContract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WitnessCreateContract.Unmarshal(m, b)
}
func (m *WitnessCreateContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WitnessCreateContract.Marshal(b, m, deterministic)
}
func (m *WitnessCreateContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WitnessCreateContract.Merge(m, src)
}
func (m *WitnessCreateContract) XXX_Size() int {
	return xxx_messageInfo_WitnessCreateContract.Size(m)
}
func (m *WitnessCreateContract) XXX_DiscardUnknown() {
	xxx_messageInfo_WitnessCreateContract.DiscardUnknown(m)
}

var xxx_messageInfo_WitnessCreateContract proto.InternalMessageInfo

func (m *WitnessCreateContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *WitnessCreateContract) GetUrl() []byte {
	if m != nil {
		return m.Url
	}
	return nil
}

type WitnessUpdateContract struct {
	OwnerAddress         []byte   `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	UpdateUrl            []byte   `protobuf:"bytes,12,opt,name=update_url,json=updateUrl,proto3" json:"update_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WitnessUpdateContract) Reset()         { *m = WitnessUpdateContract{} }
func (m *WitnessUpdateContract) String() string { return proto.CompactTextString(m) }
func (*WitnessUpdateContract) ProtoMessage()    {}
func (*WitnessUpdateContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_083020148d0431a0, []int{1}
}

func (m *WitnessUpdateContract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WitnessUpdateContract.Unmarshal(m, b)
}
func (m *WitnessUpdateContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WitnessUpdateContract.Marshal(b, m, deterministic)
}
func (m *WitnessUpdateContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WitnessUpdateContract.Merge(m, src)
}
func (m *WitnessUpdateContract) XXX_Size() int {
	return xxx_messageInfo_WitnessUpdateContract.Size(m)
}
func (m *WitnessUpdateContract) XXX_DiscardUnknown() {
	xxx_messageInfo_WitnessUpdateContract.DiscardUnknown(m)
}

var xxx_messageInfo_WitnessUpdateContract proto.InternalMessageInfo

func (m *WitnessUpdateContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *WitnessUpdateContract) GetUpdateUrl() []byte {
	if m != nil {
		return m.UpdateUrl
	}
	return nil
}

type VoteWitnessContract struct {
	OwnerAddress         []byte                      `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	Votes                []*VoteWitnessContract_Vote `protobuf:"bytes,2,rep,name=votes,proto3" json:"votes,omitempty"`
	Support              bool                        `protobuf:"varint,3,opt,name=support,proto3" json:"support,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *VoteWitnessContract) Reset()         { *m = VoteWitnessContract{} }
func (m *VoteWitnessContract) String() string { return proto.CompactTextString(m) }
func (*VoteWitnessContract) ProtoMessage()    {}
func (*VoteWitnessContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_083020148d0431a0, []int{2}
}

func (m *VoteWitnessContract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoteWitnessContract.Unmarshal(m, b)
}
func (m *VoteWitnessContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoteWitnessContract.Marshal(b, m, deterministic)
}
func (m *VoteWitnessContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteWitnessContract.Merge(m, src)
}
func (m *VoteWitnessContract) XXX_Size() int {
	return xxx_messageInfo_VoteWitnessContract.Size(m)
}
func (m *VoteWitnessContract) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteWitnessContract.DiscardUnknown(m)
}

var xxx_messageInfo_VoteWitnessContract proto.InternalMessageInfo

func (m *VoteWitnessContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *VoteWitnessContract) GetVotes() []*VoteWitnessContract_Vote {
	if m != nil {
		return m.Votes
	}
	return nil
}

func (m *VoteWitnessContract) GetSupport() bool {
	if m != nil {
		return m.Support
	}
	return false
}

type VoteWitnessContract_Vote struct {
	VoteAddress          []byte   `protobuf:"bytes,1,opt,name=vote_address,json=voteAddress,proto3" json:"vote_address,omitempty"`
	VoteCount            int64    `protobuf:"varint,2,opt,name=vote_count,json=voteCount,proto3" json:"vote_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VoteWitnessContract_Vote) Reset()         { *m = VoteWitnessContract_Vote{} }
func (m *VoteWitnessContract_Vote) String() string { return proto.CompactTextString(m) }
func (*VoteWitnessContract_Vote) ProtoMessage()    {}
func (*VoteWitnessContract_Vote) Descriptor() ([]byte, []int) {
	return fileDescriptor_083020148d0431a0, []int{2, 0}
}

func (m *VoteWitnessContract_Vote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoteWitnessContract_Vote.Unmarshal(m, b)
}
func (m *VoteWitnessContract_Vote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoteWitnessContract_Vote.Marshal(b, m, deterministic)
}
func (m *VoteWitnessContract_Vote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteWitnessContract_Vote.Merge(m, src)
}
func (m *VoteWitnessContract_Vote) XXX_Size() int {
	return xxx_messageInfo_VoteWitnessContract_Vote.Size(m)
}
func (m *VoteWitnessContract_Vote) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteWitnessContract_Vote.DiscardUnknown(m)
}

var xxx_messageInfo_VoteWitnessContract_Vote proto.InternalMessageInfo

func (m *VoteWitnessContract_Vote) GetVoteAddress() []byte {
	if m != nil {
		return m.VoteAddress
	}
	return nil
}

func (m *VoteWitnessContract_Vote) GetVoteCount() int64 {
	if m != nil {
		return m.VoteCount
	}
	return 0
}

func init() {
	proto.RegisterType((*WitnessCreateContract)(nil), "protocol.WitnessCreateContract")
	proto.RegisterType((*WitnessUpdateContract)(nil), "protocol.WitnessUpdateContract")
	proto.RegisterType((*VoteWitnessContract)(nil), "protocol.VoteWitnessContract")
	proto.RegisterType((*VoteWitnessContract_Vote)(nil), "protocol.VoteWitnessContract.Vote")
}

func init() {
	proto.RegisterFile("core/contract/witness_contract.proto", fileDescriptor_083020148d0431a0)
}

var fileDescriptor_083020148d0431a0 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x31, 0x4b, 0xf4, 0x40,
	0x10, 0x86, 0xc9, 0xe5, 0xfb, 0xf4, 0x6e, 0x2e, 0x82, 0x44, 0x84, 0x45, 0x10, 0xee, 0xa2, 0x45,
	0xaa, 0x0d, 0x68, 0x63, 0x63, 0xa1, 0x69, 0xac, 0x2c, 0x02, 0xa7, 0xa0, 0x45, 0x48, 0x36, 0x6b,
	0x0c, 0xc4, 0x4c, 0x98, 0x9d, 0x78, 0x7f, 0xd5, 0x9f, 0x23, 0xbb, 0x31, 0x22, 0x62, 0x73, 0x55,
	0x92, 0x27, 0xcf, 0xbc, 0x3b, 0x3b, 0x03, 0xe7, 0x0a, 0x49, 0x27, 0x0a, 0x3b, 0xa6, 0x42, 0x71,
	0xb2, 0x6d, 0xb8, 0xd3, 0xc6, 0xe4, 0x13, 0x90, 0x3d, 0x21, 0x63, 0x38, 0x77, 0x0f, 0x85, 0x6d,
	0x74, 0x0f, 0xc7, 0x8f, 0xa3, 0x93, 0x92, 0x2e, 0x58, 0xa7, 0x5f, 0x62, 0x78, 0x06, 0x07, 0xb8,
	0xed, 0x34, 0xe5, 0x45, 0x55, 0x91, 0x36, 0x46, 0x78, 0x2b, 0x2f, 0x0e, 0xb2, 0xc0, 0xc1, 0x9b,
	0x91, 0x85, 0x87, 0xe0, 0x0f, 0xd4, 0x8a, 0x99, 0xfb, 0x65, 0x5f, 0xa3, 0xe7, 0xef, 0xbc, 0x4d,
	0x5f, 0xed, 0x9c, 0x77, 0x0a, 0x30, 0xb8, 0xb2, 0xdc, 0xc6, 0x06, 0xce, 0x58, 0x8c, 0x64, 0x43,
	0x6d, 0xf4, 0xe1, 0xc1, 0xd1, 0x03, 0xb2, 0x9e, 0x3a, 0xde, 0x29, 0xfb, 0x0a, 0xfe, 0xbf, 0x23,
	0x6b, 0x23, 0x66, 0x2b, 0x3f, 0x5e, 0x5e, 0x44, 0x72, 0x9a, 0x81, 0xfc, 0x23, 0xd2, 0xb1, 0x6c,
	0x2c, 0x08, 0x05, 0xec, 0x9b, 0xa1, 0xef, 0x91, 0x58, 0xf8, 0x2b, 0x2f, 0x9e, 0x67, 0xd3, 0xe7,
	0xc9, 0x1d, 0xfc, 0xb3, 0x62, 0xb8, 0x86, 0xc0, 0xaa, 0xbf, 0xce, 0x5f, 0x5a, 0xf6, 0xe3, 0x6a,
	0x4e, 0x51, 0x38, 0x74, 0xec, 0x26, 0xe6, 0x67, 0x0b, 0x4b, 0x52, 0x0b, 0x6e, 0xaf, 0x41, 0x20,
	0xd5, 0x92, 0x09, 0xbb, 0xb1, 0x31, 0x23, 0xa7, 0x9d, 0x3d, 0xad, 0xeb, 0x86, 0x5f, 0x87, 0x52,
	0x2a, 0x7c, 0x4b, 0x5e, 0x4a, 0x83, 0x25, 0xe9, 0x86, 0x8a, 0xa4, 0x46, 0x6b, 0x27, 0x76, 0xe1,
	0xe5, 0x9e, 0xab, 0xb9, 0xfc, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xea, 0xca, 0xef, 0x4c, 0xff, 0x01,
	0x00, 0x00,
}
