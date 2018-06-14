// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/genomics/v1/position.proto

package genomics // import "google.golang.org/genproto/googleapis/genomics/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// An abstraction for referring to a genomic position, in relation to some
// already known reference. For now, represents a genomic position as a
// reference name, a base number on that reference (0-based), and a
// determination of forward or reverse strand.
type Position struct {
	// The name of the reference in whatever reference set is being used.
	ReferenceName string `protobuf:"bytes,1,opt,name=reference_name,json=referenceName,proto3" json:"reference_name,omitempty"`
	// The 0-based offset from the start of the forward strand for that reference.
	Position int64 `protobuf:"varint,2,opt,name=position,proto3" json:"position,omitempty"`
	// Whether this position is on the reverse strand, as opposed to the forward
	// strand.
	ReverseStrand        bool     `protobuf:"varint,3,opt,name=reverse_strand,json=reverseStrand,proto3" json:"reverse_strand,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Position) Reset()         { *m = Position{} }
func (m *Position) String() string { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()    {}
func (*Position) Descriptor() ([]byte, []int) {
	return fileDescriptor_position_c03dd355ebc45a7d, []int{0}
}
func (m *Position) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Position.Unmarshal(m, b)
}
func (m *Position) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Position.Marshal(b, m, deterministic)
}
func (dst *Position) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Position.Merge(dst, src)
}
func (m *Position) XXX_Size() int {
	return xxx_messageInfo_Position.Size(m)
}
func (m *Position) XXX_DiscardUnknown() {
	xxx_messageInfo_Position.DiscardUnknown(m)
}

var xxx_messageInfo_Position proto.InternalMessageInfo

func (m *Position) GetReferenceName() string {
	if m != nil {
		return m.ReferenceName
	}
	return ""
}

func (m *Position) GetPosition() int64 {
	if m != nil {
		return m.Position
	}
	return 0
}

func (m *Position) GetReverseStrand() bool {
	if m != nil {
		return m.ReverseStrand
	}
	return false
}

func init() {
	proto.RegisterType((*Position)(nil), "google.genomics.v1.Position")
}

func init() {
	proto.RegisterFile("google/genomics/v1/position.proto", fileDescriptor_position_c03dd355ebc45a7d)
}

var fileDescriptor_position_c03dd355ebc45a7d = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x14, 0x84, 0x89, 0x05, 0x59, 0x03, 0xf5, 0xb0, 0x07, 0x59, 0x8a, 0x87, 0x55, 0x10, 0xf6, 0x94,
	0x50, 0xbc, 0xe9, 0xad, 0x3f, 0x40, 0x96, 0x7a, 0xf3, 0x52, 0x9e, 0xeb, 0x33, 0x06, 0xba, 0xef,
	0x85, 0x24, 0xec, 0x6f, 0xf7, 0x28, 0x49, 0x9a, 0x22, 0xf4, 0x96, 0x4c, 0x66, 0x26, 0x1f, 0x23,
	0x1f, 0x0c, 0xb3, 0x39, 0xa2, 0x36, 0x48, 0x3c, 0xdb, 0x29, 0xe8, 0x65, 0xab, 0x1d, 0x07, 0x1b,
	0x2d, 0x93, 0x72, 0x9e, 0x23, 0xb7, 0x6d, 0xb1, 0xa8, 0x6a, 0x51, 0xcb, 0x76, 0x73, 0x7f, 0x8a,
	0x81, 0xb3, 0x1a, 0x88, 0x38, 0x42, 0x0a, 0x84, 0x92, 0x78, 0x8c, 0xb2, 0x19, 0x4f, 0x1d, 0xed,
	0x93, 0xbc, 0xf5, 0xf8, 0x8d, 0x1e, 0x69, 0xc2, 0x03, 0xc1, 0x8c, 0x9d, 0xe8, 0xc5, 0x70, 0xb3,
	0x5f, 0x9f, 0xd5, 0x37, 0x98, 0xb1, 0xdd, 0xc8, 0xa6, 0x7e, 0xdb, 0x5d, 0xf5, 0x62, 0x58, 0xed,
	0xcf, 0xf7, 0x52, 0xb1, 0xa0, 0x0f, 0x78, 0x08, 0xd1, 0x03, 0x7d, 0x75, 0xab, 0x5e, 0x0c, 0x4d,
	0xaa, 0xc8, 0xea, 0x7b, 0x16, 0x77, 0x3f, 0xf2, 0x6e, 0xe2, 0x59, 0x5d, 0xd2, 0xee, 0xd6, 0x95,
	0x66, 0x4c, 0x78, 0xa3, 0xf8, 0x78, 0xa9, 0x26, 0x3e, 0x02, 0x19, 0xc5, 0xde, 0xa4, 0x01, 0x32,
	0xbc, 0x2e, 0x4f, 0xe0, 0x6c, 0xf8, 0x3f, 0xca, 0x6b, 0x3d, 0xff, 0x0a, 0xf1, 0x79, 0x9d, 0x9d,
	0xcf, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5c, 0xc6, 0x22, 0xea, 0x3d, 0x01, 0x00, 0x00,
}
