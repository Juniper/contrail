// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: xxxfields.proto

package test

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type NativeWithSizeCache struct {
	Field1        *float64 `protobuf:"fixed64,1,opt,name=Field1" json:"Field1,omitempty"`
	Field2        *float32 `protobuf:"fixed32,2,opt,name=Field2" json:"Field2,omitempty"`
	Field3        *int32   `protobuf:"varint,3,opt,name=Field3" json:"Field3,omitempty"`
	Field4        *int64   `protobuf:"varint,4,opt,name=Field4" json:"Field4,omitempty"`
	Field11       uint64   `protobuf:"fixed64,11,opt,name=Field11" json:"Field11"`
	Field13       *bool    `protobuf:"varint,13,opt,name=Field13" json:"Field13,omitempty"`
	Field14       *string  `protobuf:"bytes,14,opt,name=Field14" json:"Field14,omitempty"`
	Field15       []byte   `protobuf:"bytes,15,opt,name=Field15" json:"Field15,omitempty"`
	XXX_sizecache int32    `json:"-"`
}

func (m *NativeWithSizeCache) Reset()         { *m = NativeWithSizeCache{} }
func (m *NativeWithSizeCache) String() string { return proto.CompactTextString(m) }
func (*NativeWithSizeCache) ProtoMessage()    {}
func (*NativeWithSizeCache) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba18cdebbd18dc5d, []int{0}
}
func (m *NativeWithSizeCache) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NativeWithSizeCache.Unmarshal(m, b)
}
func (m *NativeWithSizeCache) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NativeWithSizeCache.Marshal(b, m, deterministic)
}
func (m *NativeWithSizeCache) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NativeWithSizeCache.Merge(m, src)
}
func (m *NativeWithSizeCache) XXX_Size() int {
	return xxx_messageInfo_NativeWithSizeCache.Size(m)
}
func (m *NativeWithSizeCache) XXX_DiscardUnknown() {
	xxx_messageInfo_NativeWithSizeCache.DiscardUnknown(m)
}

var xxx_messageInfo_NativeWithSizeCache proto.InternalMessageInfo

func (m *NativeWithSizeCache) GetField1() float64 {
	if m != nil && m.Field1 != nil {
		return *m.Field1
	}
	return 0
}

func (m *NativeWithSizeCache) GetField2() float32 {
	if m != nil && m.Field2 != nil {
		return *m.Field2
	}
	return 0
}

func (m *NativeWithSizeCache) GetField3() int32 {
	if m != nil && m.Field3 != nil {
		return *m.Field3
	}
	return 0
}

func (m *NativeWithSizeCache) GetField4() int64 {
	if m != nil && m.Field4 != nil {
		return *m.Field4
	}
	return 0
}

func (m *NativeWithSizeCache) GetField11() uint64 {
	if m != nil {
		return m.Field11
	}
	return 0
}

func (m *NativeWithSizeCache) GetField13() bool {
	if m != nil && m.Field13 != nil {
		return *m.Field13
	}
	return false
}

func (m *NativeWithSizeCache) GetField14() string {
	if m != nil && m.Field14 != nil {
		return *m.Field14
	}
	return ""
}

func (m *NativeWithSizeCache) GetField15() []byte {
	if m != nil {
		return m.Field15
	}
	return nil
}

type StructWithSizeCache struct {
	Field1        float64                `protobuf:"fixed64,1,opt,name=Field1" json:"Field1"`
	Field2        float32                `protobuf:"fixed32,2,opt,name=Field2" json:"Field2"`
	Field3        NativeWithSizeCache    `protobuf:"bytes,3,opt,name=Field3" json:"Field3"`
	Field6        *uint64                `protobuf:"varint,6,opt,name=Field6" json:"Field6,omitempty"`
	Field7        int32                  `protobuf:"zigzag32,7,opt,name=Field7" json:"Field7"`
	Field8        NativeWithoutSizeCache `protobuf:"bytes,8,opt,name=Field8" json:"Field8"`
	Field13       []bool                 `protobuf:"varint,13,rep,name=Field13" json:"Field13,omitempty"`
	Field14       *string                `protobuf:"bytes,14,opt,name=Field14" json:"Field14,omitempty"`
	Field15       []byte                 `protobuf:"bytes,15,opt,name=Field15" json:"Field15"`
	XXX_sizecache int32                  `json:"-"`
}

func (m *StructWithSizeCache) Reset()         { *m = StructWithSizeCache{} }
func (m *StructWithSizeCache) String() string { return proto.CompactTextString(m) }
func (*StructWithSizeCache) ProtoMessage()    {}
func (*StructWithSizeCache) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba18cdebbd18dc5d, []int{1}
}
func (m *StructWithSizeCache) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StructWithSizeCache.Unmarshal(m, b)
}
func (m *StructWithSizeCache) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StructWithSizeCache.Marshal(b, m, deterministic)
}
func (m *StructWithSizeCache) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StructWithSizeCache.Merge(m, src)
}
func (m *StructWithSizeCache) XXX_Size() int {
	return xxx_messageInfo_StructWithSizeCache.Size(m)
}
func (m *StructWithSizeCache) XXX_DiscardUnknown() {
	xxx_messageInfo_StructWithSizeCache.DiscardUnknown(m)
}

var xxx_messageInfo_StructWithSizeCache proto.InternalMessageInfo

func (m *StructWithSizeCache) GetField1() float64 {
	if m != nil {
		return m.Field1
	}
	return 0
}

func (m *StructWithSizeCache) GetField2() float32 {
	if m != nil {
		return m.Field2
	}
	return 0
}

func (m *StructWithSizeCache) GetField3() NativeWithSizeCache {
	if m != nil {
		return m.Field3
	}
	return NativeWithSizeCache{}
}

func (m *StructWithSizeCache) GetField6() uint64 {
	if m != nil && m.Field6 != nil {
		return *m.Field6
	}
	return 0
}

func (m *StructWithSizeCache) GetField7() int32 {
	if m != nil {
		return m.Field7
	}
	return 0
}

func (m *StructWithSizeCache) GetField8() NativeWithoutSizeCache {
	if m != nil {
		return m.Field8
	}
	return NativeWithoutSizeCache{}
}

func (m *StructWithSizeCache) GetField13() []bool {
	if m != nil {
		return m.Field13
	}
	return nil
}

func (m *StructWithSizeCache) GetField14() string {
	if m != nil && m.Field14 != nil {
		return *m.Field14
	}
	return ""
}

func (m *StructWithSizeCache) GetField15() []byte {
	if m != nil {
		return m.Field15
	}
	return nil
}

type NativeWithoutSizeCache struct {
	Field1  *float64 `protobuf:"fixed64,1,opt,name=Field1" json:"Field1,omitempty"`
	Field2  *float32 `protobuf:"fixed32,2,opt,name=Field2" json:"Field2,omitempty"`
	Field3  *int32   `protobuf:"varint,3,opt,name=Field3" json:"Field3,omitempty"`
	Field4  *int64   `protobuf:"varint,4,opt,name=Field4" json:"Field4,omitempty"`
	Field11 uint64   `protobuf:"fixed64,11,opt,name=Field11" json:"Field11"`
	Field13 *bool    `protobuf:"varint,13,opt,name=Field13" json:"Field13,omitempty"`
	Field14 *string  `protobuf:"bytes,14,opt,name=Field14" json:"Field14,omitempty"`
	Field15 []byte   `protobuf:"bytes,15,opt,name=Field15" json:"Field15,omitempty"`
}

func (m *NativeWithoutSizeCache) Reset()         { *m = NativeWithoutSizeCache{} }
func (m *NativeWithoutSizeCache) String() string { return proto.CompactTextString(m) }
func (*NativeWithoutSizeCache) ProtoMessage()    {}
func (*NativeWithoutSizeCache) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba18cdebbd18dc5d, []int{2}
}
func (m *NativeWithoutSizeCache) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NativeWithoutSizeCache.Unmarshal(m, b)
}
func (m *NativeWithoutSizeCache) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NativeWithoutSizeCache.Marshal(b, m, deterministic)
}
func (m *NativeWithoutSizeCache) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NativeWithoutSizeCache.Merge(m, src)
}
func (m *NativeWithoutSizeCache) XXX_Size() int {
	return xxx_messageInfo_NativeWithoutSizeCache.Size(m)
}
func (m *NativeWithoutSizeCache) XXX_DiscardUnknown() {
	xxx_messageInfo_NativeWithoutSizeCache.DiscardUnknown(m)
}

var xxx_messageInfo_NativeWithoutSizeCache proto.InternalMessageInfo

func (m *NativeWithoutSizeCache) GetField1() float64 {
	if m != nil && m.Field1 != nil {
		return *m.Field1
	}
	return 0
}

func (m *NativeWithoutSizeCache) GetField2() float32 {
	if m != nil && m.Field2 != nil {
		return *m.Field2
	}
	return 0
}

func (m *NativeWithoutSizeCache) GetField3() int32 {
	if m != nil && m.Field3 != nil {
		return *m.Field3
	}
	return 0
}

func (m *NativeWithoutSizeCache) GetField4() int64 {
	if m != nil && m.Field4 != nil {
		return *m.Field4
	}
	return 0
}

func (m *NativeWithoutSizeCache) GetField11() uint64 {
	if m != nil {
		return m.Field11
	}
	return 0
}

func (m *NativeWithoutSizeCache) GetField13() bool {
	if m != nil && m.Field13 != nil {
		return *m.Field13
	}
	return false
}

func (m *NativeWithoutSizeCache) GetField14() string {
	if m != nil && m.Field14 != nil {
		return *m.Field14
	}
	return ""
}

func (m *NativeWithoutSizeCache) GetField15() []byte {
	if m != nil {
		return m.Field15
	}
	return nil
}

type StructWithoutSizeCache struct {
	Field1  float64                `protobuf:"fixed64,1,opt,name=Field1" json:"Field1"`
	Field2  float32                `protobuf:"fixed32,2,opt,name=Field2" json:"Field2"`
	Field3  NativeWithSizeCache    `protobuf:"bytes,3,opt,name=Field3" json:"Field3"`
	Field6  *uint64                `protobuf:"varint,6,opt,name=Field6" json:"Field6,omitempty"`
	Field7  int32                  `protobuf:"zigzag32,7,opt,name=Field7" json:"Field7"`
	Field8  NativeWithoutSizeCache `protobuf:"bytes,8,opt,name=Field8" json:"Field8"`
	Field13 []bool                 `protobuf:"varint,13,rep,name=Field13" json:"Field13,omitempty"`
	Field14 *string                `protobuf:"bytes,14,opt,name=Field14" json:"Field14,omitempty"`
	Field15 []byte                 `protobuf:"bytes,15,opt,name=Field15" json:"Field15"`
}

func (m *StructWithoutSizeCache) Reset()         { *m = StructWithoutSizeCache{} }
func (m *StructWithoutSizeCache) String() string { return proto.CompactTextString(m) }
func (*StructWithoutSizeCache) ProtoMessage()    {}
func (*StructWithoutSizeCache) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba18cdebbd18dc5d, []int{3}
}
func (m *StructWithoutSizeCache) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StructWithoutSizeCache.Unmarshal(m, b)
}
func (m *StructWithoutSizeCache) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StructWithoutSizeCache.Marshal(b, m, deterministic)
}
func (m *StructWithoutSizeCache) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StructWithoutSizeCache.Merge(m, src)
}
func (m *StructWithoutSizeCache) XXX_Size() int {
	return xxx_messageInfo_StructWithoutSizeCache.Size(m)
}
func (m *StructWithoutSizeCache) XXX_DiscardUnknown() {
	xxx_messageInfo_StructWithoutSizeCache.DiscardUnknown(m)
}

var xxx_messageInfo_StructWithoutSizeCache proto.InternalMessageInfo

func (m *StructWithoutSizeCache) GetField1() float64 {
	if m != nil {
		return m.Field1
	}
	return 0
}

func (m *StructWithoutSizeCache) GetField2() float32 {
	if m != nil {
		return m.Field2
	}
	return 0
}

func (m *StructWithoutSizeCache) GetField3() NativeWithSizeCache {
	if m != nil {
		return m.Field3
	}
	return NativeWithSizeCache{}
}

func (m *StructWithoutSizeCache) GetField6() uint64 {
	if m != nil && m.Field6 != nil {
		return *m.Field6
	}
	return 0
}

func (m *StructWithoutSizeCache) GetField7() int32 {
	if m != nil {
		return m.Field7
	}
	return 0
}

func (m *StructWithoutSizeCache) GetField8() NativeWithoutSizeCache {
	if m != nil {
		return m.Field8
	}
	return NativeWithoutSizeCache{}
}

func (m *StructWithoutSizeCache) GetField13() []bool {
	if m != nil {
		return m.Field13
	}
	return nil
}

func (m *StructWithoutSizeCache) GetField14() string {
	if m != nil && m.Field14 != nil {
		return *m.Field14
	}
	return ""
}

func (m *StructWithoutSizeCache) GetField15() []byte {
	if m != nil {
		return m.Field15
	}
	return nil
}

func init() {
	proto.RegisterType((*NativeWithSizeCache)(nil), "test.NativeWithSizeCache")
	proto.RegisterType((*StructWithSizeCache)(nil), "test.StructWithSizeCache")
	proto.RegisterType((*NativeWithoutSizeCache)(nil), "test.NativeWithoutSizeCache")
	proto.RegisterType((*StructWithoutSizeCache)(nil), "test.StructWithoutSizeCache")
}

func init() { proto.RegisterFile("xxxfields.proto", fileDescriptor_ba18cdebbd18dc5d) }

var fileDescriptor_ba18cdebbd18dc5d = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xaf, 0xa8, 0xa8, 0x48,
	0xcb, 0x4c, 0xcd, 0x49, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0x49, 0x2d,
	0x2e, 0x91, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf,
	0x4f, 0xcf, 0xd7, 0x07, 0x4b, 0x26, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0xa4,
	0xf4, 0x96, 0x91, 0x4b, 0xd8, 0x2f, 0xb1, 0x24, 0xb3, 0x2c, 0x35, 0x3c, 0xb3, 0x24, 0x23, 0x38,
	0xb3, 0x2a, 0xd5, 0x39, 0x31, 0x39, 0x23, 0x55, 0x48, 0x8c, 0x8b, 0xcd, 0x0d, 0x64, 0xb8, 0xa1,
	0x04, 0xa3, 0x02, 0xa3, 0x06, 0x63, 0x10, 0x94, 0x07, 0x17, 0x37, 0x92, 0x60, 0x52, 0x60, 0xd4,
	0x60, 0x82, 0x8a, 0x1b, 0xc1, 0xc5, 0x8d, 0x25, 0x98, 0x15, 0x18, 0x35, 0x58, 0xa1, 0xe2, 0xc6,
	0x70, 0x71, 0x13, 0x09, 0x16, 0x05, 0x46, 0x0d, 0x66, 0xa8, 0xb8, 0x89, 0x90, 0x1c, 0x17, 0x3b,
	0xc4, 0x44, 0x43, 0x09, 0x6e, 0x05, 0x46, 0x0d, 0x36, 0x27, 0x96, 0x13, 0xf7, 0xe4, 0x19, 0x82,
	0x60, 0x82, 0x42, 0x12, 0x30, 0x79, 0x63, 0x09, 0x5e, 0x05, 0x46, 0x0d, 0x0e, 0x98, 0x8c, 0x31,
	0x42, 0xc6, 0x44, 0x82, 0x4f, 0x81, 0x51, 0x83, 0x13, 0x26, 0x63, 0x82, 0x90, 0x31, 0x95, 0xe0,
	0x57, 0x60, 0xd4, 0xe0, 0x81, 0xc9, 0x98, 0x5a, 0x71, 0x5c, 0x58, 0x28, 0xcf, 0x30, 0x63, 0x91,
	0x3c, 0x83, 0xd2, 0x2d, 0x26, 0x2e, 0xe1, 0xe0, 0x92, 0xa2, 0xd2, 0xe4, 0x12, 0x54, 0xff, 0xca,
	0xa0, 0xfa, 0x17, 0xea, 0x1c, 0x98, 0xaf, 0x65, 0x50, 0x7d, 0x8d, 0x22, 0x6b, 0x24, 0x64, 0x8e,
	0xe2, 0x77, 0x6e, 0x23, 0x49, 0x3d, 0x50, 0x4c, 0xe8, 0x61, 0x09, 0x56, 0x14, 0x8d, 0x88, 0xc0,
	0x31, 0x93, 0x60, 0x53, 0x60, 0xd4, 0x60, 0x81, 0x8a, 0x9b, 0xc1, 0xad, 0x33, 0x97, 0x60, 0x57,
	0x60, 0xd4, 0x10, 0x44, 0xd1, 0x65, 0x2e, 0x64, 0x05, 0x95, 0xb5, 0x90, 0xe0, 0x00, 0x5b, 0x27,
	0x83, 0x6e, 0x5d, 0x7e, 0x69, 0x09, 0x76, 0x1b, 0x2d, 0x50, 0x83, 0x95, 0x99, 0xb8, 0x60, 0x95,
	0x43, 0x0b, 0x56, 0xd4, 0xa8, 0x42, 0x0e, 0xdc, 0x2f, 0x8c, 0x5c, 0x62, 0xd8, 0x9d, 0x31, 0xcc,
	0xd2, 0x13, 0x0f, 0xc8, 0xcb, 0x13, 0x16, 0x41, 0xbd, 0xfd, 0x90, 0x89, 0x4b, 0x0c, 0x91, 0xa6,
	0x50, 0xbc, 0x3d, 0x9a, 0xac, 0xc8, 0x4e, 0x56, 0x28, 0x61, 0xec, 0xc4, 0xf3, 0xe3, 0xa1, 0x1c,
	0xe3, 0x8a, 0x47, 0x72, 0x8c, 0x3b, 0x1e, 0xc9, 0x31, 0x02, 0x02, 0x00, 0x00, 0xff, 0xff, 0x72,
	0x3b, 0x20, 0xc7, 0xfc, 0x04, 0x00, 0x00,
}

func (this *NativeWithSizeCache) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*NativeWithSizeCache)
	if !ok {
		that2, ok := that.(NativeWithSizeCache)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Field1 != nil && that1.Field1 != nil {
		if *this.Field1 != *that1.Field1 {
			return false
		}
	} else if this.Field1 != nil {
		return false
	} else if that1.Field1 != nil {
		return false
	}
	if this.Field2 != nil && that1.Field2 != nil {
		if *this.Field2 != *that1.Field2 {
			return false
		}
	} else if this.Field2 != nil {
		return false
	} else if that1.Field2 != nil {
		return false
	}
	if this.Field3 != nil && that1.Field3 != nil {
		if *this.Field3 != *that1.Field3 {
			return false
		}
	} else if this.Field3 != nil {
		return false
	} else if that1.Field3 != nil {
		return false
	}
	if this.Field4 != nil && that1.Field4 != nil {
		if *this.Field4 != *that1.Field4 {
			return false
		}
	} else if this.Field4 != nil {
		return false
	} else if that1.Field4 != nil {
		return false
	}
	if this.Field11 != that1.Field11 {
		return false
	}
	if this.Field13 != nil && that1.Field13 != nil {
		if *this.Field13 != *that1.Field13 {
			return false
		}
	} else if this.Field13 != nil {
		return false
	} else if that1.Field13 != nil {
		return false
	}
	if this.Field14 != nil && that1.Field14 != nil {
		if *this.Field14 != *that1.Field14 {
			return false
		}
	} else if this.Field14 != nil {
		return false
	} else if that1.Field14 != nil {
		return false
	}
	if !bytes.Equal(this.Field15, that1.Field15) {
		return false
	}
	return true
}
func (this *StructWithSizeCache) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*StructWithSizeCache)
	if !ok {
		that2, ok := that.(StructWithSizeCache)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Field1 != that1.Field1 {
		return false
	}
	if this.Field2 != that1.Field2 {
		return false
	}
	if !this.Field3.Equal(&that1.Field3) {
		return false
	}
	if this.Field6 != nil && that1.Field6 != nil {
		if *this.Field6 != *that1.Field6 {
			return false
		}
	} else if this.Field6 != nil {
		return false
	} else if that1.Field6 != nil {
		return false
	}
	if this.Field7 != that1.Field7 {
		return false
	}
	if !this.Field8.Equal(&that1.Field8) {
		return false
	}
	if len(this.Field13) != len(that1.Field13) {
		return false
	}
	for i := range this.Field13 {
		if this.Field13[i] != that1.Field13[i] {
			return false
		}
	}
	if this.Field14 != nil && that1.Field14 != nil {
		if *this.Field14 != *that1.Field14 {
			return false
		}
	} else if this.Field14 != nil {
		return false
	} else if that1.Field14 != nil {
		return false
	}
	if !bytes.Equal(this.Field15, that1.Field15) {
		return false
	}
	return true
}
func (this *NativeWithoutSizeCache) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*NativeWithoutSizeCache)
	if !ok {
		that2, ok := that.(NativeWithoutSizeCache)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Field1 != nil && that1.Field1 != nil {
		if *this.Field1 != *that1.Field1 {
			return false
		}
	} else if this.Field1 != nil {
		return false
	} else if that1.Field1 != nil {
		return false
	}
	if this.Field2 != nil && that1.Field2 != nil {
		if *this.Field2 != *that1.Field2 {
			return false
		}
	} else if this.Field2 != nil {
		return false
	} else if that1.Field2 != nil {
		return false
	}
	if this.Field3 != nil && that1.Field3 != nil {
		if *this.Field3 != *that1.Field3 {
			return false
		}
	} else if this.Field3 != nil {
		return false
	} else if that1.Field3 != nil {
		return false
	}
	if this.Field4 != nil && that1.Field4 != nil {
		if *this.Field4 != *that1.Field4 {
			return false
		}
	} else if this.Field4 != nil {
		return false
	} else if that1.Field4 != nil {
		return false
	}
	if this.Field11 != that1.Field11 {
		return false
	}
	if this.Field13 != nil && that1.Field13 != nil {
		if *this.Field13 != *that1.Field13 {
			return false
		}
	} else if this.Field13 != nil {
		return false
	} else if that1.Field13 != nil {
		return false
	}
	if this.Field14 != nil && that1.Field14 != nil {
		if *this.Field14 != *that1.Field14 {
			return false
		}
	} else if this.Field14 != nil {
		return false
	} else if that1.Field14 != nil {
		return false
	}
	if !bytes.Equal(this.Field15, that1.Field15) {
		return false
	}
	return true
}
func (this *StructWithoutSizeCache) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*StructWithoutSizeCache)
	if !ok {
		that2, ok := that.(StructWithoutSizeCache)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Field1 != that1.Field1 {
		return false
	}
	if this.Field2 != that1.Field2 {
		return false
	}
	if !this.Field3.Equal(&that1.Field3) {
		return false
	}
	if this.Field6 != nil && that1.Field6 != nil {
		if *this.Field6 != *that1.Field6 {
			return false
		}
	} else if this.Field6 != nil {
		return false
	} else if that1.Field6 != nil {
		return false
	}
	if this.Field7 != that1.Field7 {
		return false
	}
	if !this.Field8.Equal(&that1.Field8) {
		return false
	}
	if len(this.Field13) != len(that1.Field13) {
		return false
	}
	for i := range this.Field13 {
		if this.Field13[i] != that1.Field13[i] {
			return false
		}
	}
	if this.Field14 != nil && that1.Field14 != nil {
		if *this.Field14 != *that1.Field14 {
			return false
		}
	} else if this.Field14 != nil {
		return false
	} else if that1.Field14 != nil {
		return false
	}
	if !bytes.Equal(this.Field15, that1.Field15) {
		return false
	}
	return true
}
func NewPopulatedNativeWithSizeCache(r randyXxxfields, easy bool) *NativeWithSizeCache {
	this := &NativeWithSizeCache{}
	if r.Intn(5) != 0 {
		v1 := float64(r.Float64())
		if r.Intn(2) == 0 {
			v1 *= -1
		}
		this.Field1 = &v1
	}
	if r.Intn(5) != 0 {
		v2 := float32(r.Float32())
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		this.Field2 = &v2
	}
	if r.Intn(5) != 0 {
		v3 := int32(r.Int31())
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		this.Field3 = &v3
	}
	if r.Intn(5) != 0 {
		v4 := int64(r.Int63())
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		this.Field4 = &v4
	}
	this.Field11 = uint64(uint64(r.Uint32()))
	if r.Intn(5) != 0 {
		v5 := bool(bool(r.Intn(2) == 0))
		this.Field13 = &v5
	}
	if r.Intn(5) != 0 {
		v6 := string(randStringXxxfields(r))
		this.Field14 = &v6
	}
	if r.Intn(5) != 0 {
		v7 := r.Intn(100)
		this.Field15 = make([]byte, v7)
		for i := 0; i < v7; i++ {
			this.Field15[i] = byte(r.Intn(256))
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedStructWithSizeCache(r randyXxxfields, easy bool) *StructWithSizeCache {
	this := &StructWithSizeCache{}
	this.Field1 = float64(r.Float64())
	if r.Intn(2) == 0 {
		this.Field1 *= -1
	}
	this.Field2 = float32(r.Float32())
	if r.Intn(2) == 0 {
		this.Field2 *= -1
	}
	v8 := NewPopulatedNativeWithSizeCache(r, easy)
	this.Field3 = *v8
	if r.Intn(5) != 0 {
		v9 := uint64(uint64(r.Uint32()))
		this.Field6 = &v9
	}
	this.Field7 = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Field7 *= -1
	}
	v10 := NewPopulatedNativeWithoutSizeCache(r, easy)
	this.Field8 = *v10
	if r.Intn(5) != 0 {
		v11 := r.Intn(10)
		this.Field13 = make([]bool, v11)
		for i := 0; i < v11; i++ {
			this.Field13[i] = bool(bool(r.Intn(2) == 0))
		}
	}
	if r.Intn(5) != 0 {
		v12 := string(randStringXxxfields(r))
		this.Field14 = &v12
	}
	v13 := r.Intn(100)
	this.Field15 = make([]byte, v13)
	for i := 0; i < v13; i++ {
		this.Field15[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedNativeWithoutSizeCache(r randyXxxfields, easy bool) *NativeWithoutSizeCache {
	this := &NativeWithoutSizeCache{}
	if r.Intn(5) != 0 {
		v14 := float64(r.Float64())
		if r.Intn(2) == 0 {
			v14 *= -1
		}
		this.Field1 = &v14
	}
	if r.Intn(5) != 0 {
		v15 := float32(r.Float32())
		if r.Intn(2) == 0 {
			v15 *= -1
		}
		this.Field2 = &v15
	}
	if r.Intn(5) != 0 {
		v16 := int32(r.Int31())
		if r.Intn(2) == 0 {
			v16 *= -1
		}
		this.Field3 = &v16
	}
	if r.Intn(5) != 0 {
		v17 := int64(r.Int63())
		if r.Intn(2) == 0 {
			v17 *= -1
		}
		this.Field4 = &v17
	}
	this.Field11 = uint64(uint64(r.Uint32()))
	if r.Intn(5) != 0 {
		v18 := bool(bool(r.Intn(2) == 0))
		this.Field13 = &v18
	}
	if r.Intn(5) != 0 {
		v19 := string(randStringXxxfields(r))
		this.Field14 = &v19
	}
	if r.Intn(5) != 0 {
		v20 := r.Intn(100)
		this.Field15 = make([]byte, v20)
		for i := 0; i < v20; i++ {
			this.Field15[i] = byte(r.Intn(256))
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedStructWithoutSizeCache(r randyXxxfields, easy bool) *StructWithoutSizeCache {
	this := &StructWithoutSizeCache{}
	this.Field1 = float64(r.Float64())
	if r.Intn(2) == 0 {
		this.Field1 *= -1
	}
	this.Field2 = float32(r.Float32())
	if r.Intn(2) == 0 {
		this.Field2 *= -1
	}
	v21 := NewPopulatedNativeWithSizeCache(r, easy)
	this.Field3 = *v21
	if r.Intn(5) != 0 {
		v22 := uint64(uint64(r.Uint32()))
		this.Field6 = &v22
	}
	this.Field7 = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Field7 *= -1
	}
	v23 := NewPopulatedNativeWithoutSizeCache(r, easy)
	this.Field8 = *v23
	if r.Intn(5) != 0 {
		v24 := r.Intn(10)
		this.Field13 = make([]bool, v24)
		for i := 0; i < v24; i++ {
			this.Field13[i] = bool(bool(r.Intn(2) == 0))
		}
	}
	if r.Intn(5) != 0 {
		v25 := string(randStringXxxfields(r))
		this.Field14 = &v25
	}
	v26 := r.Intn(100)
	this.Field15 = make([]byte, v26)
	for i := 0; i < v26; i++ {
		this.Field15[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyXxxfields interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneXxxfields(r randyXxxfields) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringXxxfields(r randyXxxfields) string {
	v27 := r.Intn(100)
	tmps := make([]rune, v27)
	for i := 0; i < v27; i++ {
		tmps[i] = randUTF8RuneXxxfields(r)
	}
	return string(tmps)
}
func randUnrecognizedXxxfields(r randyXxxfields, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldXxxfields(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldXxxfields(dAtA []byte, r randyXxxfields, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateXxxfields(dAtA, uint64(key))
		v28 := r.Int63()
		if r.Intn(2) == 0 {
			v28 *= -1
		}
		dAtA = encodeVarintPopulateXxxfields(dAtA, uint64(v28))
	case 1:
		dAtA = encodeVarintPopulateXxxfields(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateXxxfields(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateXxxfields(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateXxxfields(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateXxxfields(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
