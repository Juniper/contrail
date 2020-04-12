// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: issue438.proto

package issue438

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	math "math"
	math_bits "math/bits"
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

type Types struct {
	Any                  *types.Any           `protobuf:"bytes,1,opt,name=any,proto3" json:"any,omitempty"`
	Api                  *types.Api           `protobuf:"bytes,2,opt,name=api,proto3" json:"api,omitempty"`
	Met                  *types.Method        `protobuf:"bytes,3,opt,name=met,proto3" json:"met,omitempty"`
	Mx                   *types.Mixin         `protobuf:"bytes,4,opt,name=mx,proto3" json:"mx,omitempty"`
	Dur                  *types.Duration      `protobuf:"bytes,5,opt,name=dur,proto3" json:"dur,omitempty"`
	Em                   *types.Empty         `protobuf:"bytes,6,opt,name=em,proto3" json:"em,omitempty"`
	Fm                   *types.FieldMask     `protobuf:"bytes,7,opt,name=fm,proto3" json:"fm,omitempty"`
	Sc                   *types.SourceContext `protobuf:"bytes,8,opt,name=sc,proto3" json:"sc,omitempty"`
	St                   *types.Struct        `protobuf:"bytes,9,opt,name=st,proto3" json:"st,omitempty"`
	Val                  *types.Value         `protobuf:"bytes,10,opt,name=val,proto3" json:"val,omitempty"`
	Nlval                types.NullValue      `protobuf:"varint,11,opt,name=nlval,proto3,enum=google.protobuf.NullValue" json:"nlval,omitempty"`
	Stval                *types.StringValue   `protobuf:"bytes,12,opt,name=stval,proto3" json:"stval,omitempty"`
	Bval                 *types.BoolValue     `protobuf:"bytes,13,opt,name=bval,proto3" json:"bval,omitempty"`
	Strval               *types.Struct        `protobuf:"bytes,14,opt,name=strval,proto3" json:"strval,omitempty"`
	Lstv                 *types.ListValue     `protobuf:"bytes,15,opt,name=lstv,proto3" json:"lstv,omitempty"`
	Ts                   *types.Timestamp     `protobuf:"bytes,16,opt,name=ts,proto3" json:"ts,omitempty"`
	T                    *types.Type          `protobuf:"bytes,17,opt,name=t,proto3" json:"t,omitempty"`
	F                    *types.Field         `protobuf:"bytes,18,opt,name=f,proto3" json:"f,omitempty"`
	En                   *types.Enum          `protobuf:"bytes,19,opt,name=en,proto3" json:"en,omitempty"`
	Enval                *types.EnumValue     `protobuf:"bytes,20,opt,name=enval,proto3" json:"enval,omitempty"`
	Opt                  *types.Option        `protobuf:"bytes,21,opt,name=opt,proto3" json:"opt,omitempty"`
	Dbl                  *types.DoubleValue   `protobuf:"bytes,22,opt,name=dbl,proto3" json:"dbl,omitempty"`
	Flt                  *types.FloatValue    `protobuf:"bytes,23,opt,name=flt,proto3" json:"flt,omitempty"`
	I64                  *types.Int64Value    `protobuf:"bytes,24,opt,name=i64,proto3" json:"i64,omitempty"`
	U64                  *types.UInt64Value   `protobuf:"bytes,25,opt,name=u64,proto3" json:"u64,omitempty"`
	I32                  *types.Int32Value    `protobuf:"bytes,26,opt,name=i32,proto3" json:"i32,omitempty"`
	U32                  *types.UInt32Value   `protobuf:"bytes,27,opt,name=u32,proto3" json:"u32,omitempty"`
	Bool                 *types.BoolValue     `protobuf:"bytes,28,opt,name=bool,proto3" json:"bool,omitempty"`
	Str                  *types.StringValue   `protobuf:"bytes,29,opt,name=str,proto3" json:"str,omitempty"`
	Bytes                *types.BytesValue    `protobuf:"bytes,30,opt,name=bytes,proto3" json:"bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Types) Reset()         { *m = Types{} }
func (m *Types) String() string { return proto.CompactTextString(m) }
func (*Types) ProtoMessage()    {}
func (*Types) Descriptor() ([]byte, []int) {
	return fileDescriptor_43147f0c8dedbac4, []int{0}
}
func (m *Types) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Types.Unmarshal(m, b)
}
func (m *Types) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Types.Marshal(b, m, deterministic)
}
func (m *Types) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Types.Merge(m, src)
}
func (m *Types) XXX_Size() int {
	return xxx_messageInfo_Types.Size(m)
}
func (m *Types) XXX_DiscardUnknown() {
	xxx_messageInfo_Types.DiscardUnknown(m)
}

var xxx_messageInfo_Types proto.InternalMessageInfo

func (m *Types) GetAny() *types.Any {
	if m != nil {
		return m.Any
	}
	return nil
}

func (m *Types) GetApi() *types.Api {
	if m != nil {
		return m.Api
	}
	return nil
}

func (m *Types) GetMet() *types.Method {
	if m != nil {
		return m.Met
	}
	return nil
}

func (m *Types) GetMx() *types.Mixin {
	if m != nil {
		return m.Mx
	}
	return nil
}

func (m *Types) GetDur() *types.Duration {
	if m != nil {
		return m.Dur
	}
	return nil
}

func (m *Types) GetEm() *types.Empty {
	if m != nil {
		return m.Em
	}
	return nil
}

func (m *Types) GetFm() *types.FieldMask {
	if m != nil {
		return m.Fm
	}
	return nil
}

func (m *Types) GetSc() *types.SourceContext {
	if m != nil {
		return m.Sc
	}
	return nil
}

func (m *Types) GetSt() *types.Struct {
	if m != nil {
		return m.St
	}
	return nil
}

func (m *Types) GetVal() *types.Value {
	if m != nil {
		return m.Val
	}
	return nil
}

func (m *Types) GetNlval() types.NullValue {
	if m != nil {
		return m.Nlval
	}
	return types.NullValue_NULL_VALUE
}

func (m *Types) GetStval() *types.StringValue {
	if m != nil {
		return m.Stval
	}
	return nil
}

func (m *Types) GetBval() *types.BoolValue {
	if m != nil {
		return m.Bval
	}
	return nil
}

func (m *Types) GetStrval() *types.Struct {
	if m != nil {
		return m.Strval
	}
	return nil
}

func (m *Types) GetLstv() *types.ListValue {
	if m != nil {
		return m.Lstv
	}
	return nil
}

func (m *Types) GetTs() *types.Timestamp {
	if m != nil {
		return m.Ts
	}
	return nil
}

func (m *Types) GetT() *types.Type {
	if m != nil {
		return m.T
	}
	return nil
}

func (m *Types) GetF() *types.Field {
	if m != nil {
		return m.F
	}
	return nil
}

func (m *Types) GetEn() *types.Enum {
	if m != nil {
		return m.En
	}
	return nil
}

func (m *Types) GetEnval() *types.EnumValue {
	if m != nil {
		return m.Enval
	}
	return nil
}

func (m *Types) GetOpt() *types.Option {
	if m != nil {
		return m.Opt
	}
	return nil
}

func (m *Types) GetDbl() *types.DoubleValue {
	if m != nil {
		return m.Dbl
	}
	return nil
}

func (m *Types) GetFlt() *types.FloatValue {
	if m != nil {
		return m.Flt
	}
	return nil
}

func (m *Types) GetI64() *types.Int64Value {
	if m != nil {
		return m.I64
	}
	return nil
}

func (m *Types) GetU64() *types.UInt64Value {
	if m != nil {
		return m.U64
	}
	return nil
}

func (m *Types) GetI32() *types.Int32Value {
	if m != nil {
		return m.I32
	}
	return nil
}

func (m *Types) GetU32() *types.UInt32Value {
	if m != nil {
		return m.U32
	}
	return nil
}

func (m *Types) GetBool() *types.BoolValue {
	if m != nil {
		return m.Bool
	}
	return nil
}

func (m *Types) GetStr() *types.StringValue {
	if m != nil {
		return m.Str
	}
	return nil
}

func (m *Types) GetBytes() *types.BytesValue {
	if m != nil {
		return m.Bytes
	}
	return nil
}

func init() {
	proto.RegisterType((*Types)(nil), "issue438.Types")
}

func init() { proto.RegisterFile("issue438.proto", fileDescriptor_43147f0c8dedbac4) }

var fileDescriptor_43147f0c8dedbac4 = []byte{
	// 682 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd3, 0xef, 0x4e, 0x14, 0x3b,
	0x14, 0x00, 0xf0, 0x3b, 0x33, 0x2c, 0x97, 0x5b, 0xae, 0xa8, 0x95, 0x3f, 0x65, 0xc1, 0x95, 0x28,
	0x22, 0x6a, 0x18, 0x74, 0x77, 0xb3, 0xf1, 0xab, 0x08, 0x26, 0x26, 0xa2, 0x49, 0x41, 0xbf, 0x92,
	0x99, 0xdd, 0xce, 0xd2, 0xd0, 0x99, 0x36, 0xd3, 0x16, 0x77, 0xde, 0xc6, 0x47, 0xf2, 0x11, 0x8c,
	0xbe, 0x88, 0x69, 0xa7, 0xab, 0x66, 0x87, 0x12, 0xbe, 0xcd, 0xe4, 0xfc, 0xce, 0x69, 0x7b, 0x4e,
	0x0b, 0x96, 0xa8, 0x94, 0x9a, 0xf4, 0x7b, 0xaf, 0x62, 0x51, 0x72, 0xc5, 0xe1, 0xc2, 0xf4, 0xbf,
	0xbd, 0x37, 0xa6, 0xea, 0x5c, 0xa7, 0xf1, 0x90, 0xe7, 0xfb, 0x63, 0x3e, 0xe6, 0xfb, 0x16, 0xa4,
	0x3a, 0xb3, 0x7f, 0xf6, 0xc7, 0x7e, 0xd5, 0x89, 0xed, 0xf5, 0x31, 0xe7, 0x63, 0x46, 0xfe, 0xa8,
	0xa4, 0xa8, 0xbc, 0x21, 0x41, 0x5d, 0xa8, 0x33, 0x1b, 0x1a, 0xe9, 0x32, 0x51, 0x94, 0x17, 0x2e,
	0xbe, 0x31, 0x1b, 0x27, 0xb9, 0x50, 0xd3, 0xba, 0x5b, 0xb3, 0xc1, 0x8c, 0x12, 0x36, 0x3a, 0xcb,
	0x13, 0x79, 0xe1, 0xc4, 0xf6, 0xac, 0x90, 0x5c, 0x97, 0x43, 0x72, 0x36, 0xe4, 0x85, 0x22, 0x13,
	0xe5, 0xd4, 0x66, 0x43, 0xa9, 0x52, 0x0f, 0xa7, 0xd1, 0x07, 0xb3, 0x51, 0x45, 0x73, 0x22, 0x55,
	0x92, 0x0b, 0x07, 0xda, 0x0d, 0x50, 0x09, 0xe2, 0x3b, 0xdf, 0x97, 0x32, 0x11, 0x82, 0x94, 0xb2,
	0x8e, 0x3f, 0xfc, 0x06, 0x40, 0xeb, 0xb4, 0x12, 0x44, 0xc2, 0x1d, 0x10, 0x25, 0x45, 0x85, 0x82,
	0xad, 0x60, 0x77, 0xb1, 0xbb, 0x1c, 0xd7, 0x79, 0xf1, 0x34, 0x2f, 0x7e, 0x5d, 0x54, 0xd8, 0x00,
	0xeb, 0x04, 0x45, 0xa1, 0xcf, 0x09, 0x8a, 0x0d, 0x80, 0x4f, 0x41, 0x94, 0x13, 0x85, 0x22, 0xeb,
	0xd6, 0x1a, 0xee, 0x98, 0xa8, 0x73, 0x3e, 0xc2, 0xc6, 0xc0, 0x1d, 0x10, 0xe6, 0x13, 0x34, 0x67,
	0xe5, 0x6a, 0x53, 0xd2, 0x09, 0x2d, 0x70, 0x98, 0x4f, 0xe0, 0x73, 0x10, 0x8d, 0x74, 0x89, 0x5a,
	0x16, 0xae, 0x37, 0xe0, 0xa1, 0x1b, 0x1d, 0x36, 0xca, 0x14, 0x25, 0x39, 0x9a, 0xf7, 0x14, 0x3d,
	0x32, 0x63, 0xc4, 0x21, 0xc9, 0xe1, 0x33, 0x10, 0x66, 0x39, 0xfa, 0xd7, 0xba, 0x76, 0xc3, 0xbd,
	0x35, 0x13, 0x3d, 0x4e, 0xe4, 0x05, 0x0e, 0xb3, 0x1c, 0xc6, 0x20, 0x94, 0x43, 0xb4, 0x60, 0x6d,
	0xa7, 0x61, 0x4f, 0xec, 0x6c, 0xdf, 0xd4, 0xa3, 0xc5, 0xa1, 0x1c, 0xc2, 0x27, 0x20, 0x94, 0x0a,
	0xfd, 0xe7, 0x69, 0xc1, 0x89, 0x9d, 0x32, 0x0e, 0xa5, 0x82, 0xbb, 0x20, 0xba, 0x4c, 0x18, 0x02,
	0x9e, 0xdd, 0x7e, 0x4e, 0x98, 0x26, 0xd8, 0x10, 0xf8, 0x02, 0xb4, 0x0a, 0x66, 0xec, 0xe2, 0x56,
	0xb0, 0xbb, 0x74, 0xc5, 0x8e, 0x3f, 0x68, 0xc6, 0x6a, 0x5f, 0x43, 0xd8, 0x05, 0x2d, 0xa9, 0x4c,
	0xc6, 0xff, 0xb6, 0xfa, 0xe6, 0x55, 0xfb, 0xa0, 0xc5, 0xd8, 0xe5, 0x58, 0x0a, 0x63, 0x30, 0x97,
	0x9a, 0x94, 0x5b, 0x9e, 0xb6, 0x1c, 0x70, 0xee, 0x16, 0xb1, 0x0e, 0xee, 0x83, 0x79, 0xa9, 0x4a,
	0x93, 0xb1, 0x74, 0xfd, 0x61, 0x1d, 0x33, 0x0b, 0x30, 0xa9, 0x2e, 0xd1, 0x6d, 0xcf, 0x02, 0xef,
	0xa9, 0x54, 0x6e, 0x01, 0xe3, 0xcc, 0x94, 0x94, 0x44, 0x77, 0x3c, 0xfa, 0x74, 0xfa, 0x22, 0x70,
	0xa8, 0x24, 0x7c, 0x04, 0x02, 0x85, 0xee, 0x5a, 0xba, 0xd2, 0xa4, 0x95, 0x20, 0x38, 0x50, 0x70,
	0x1b, 0x04, 0x19, 0x82, 0x9e, 0x7e, 0xdb, 0xa9, 0xe3, 0x20, 0x83, 0x8f, 0x41, 0x48, 0x0a, 0x74,
	0xcf, 0x53, 0xeb, 0xa8, 0xd0, 0x39, 0x0e, 0x49, 0x61, 0x86, 0x42, 0x0a, 0x73, 0xfa, 0x65, 0xcf,
	0x06, 0x8d, 0x74, 0x0d, 0xb6, 0xd0, 0xbc, 0x0e, 0x2e, 0x14, 0x5a, 0xf1, 0x74, 0xeb, 0xa3, 0xa8,
	0x2f, 0x32, 0x17, 0x0a, 0xc6, 0x20, 0x1a, 0xa5, 0x0c, 0xad, 0x7a, 0xa6, 0x77, 0xc8, 0x75, 0xca,
	0x88, 0xbb, 0x21, 0xa3, 0x94, 0xc1, 0x3d, 0x10, 0x65, 0x4c, 0xa1, 0x35, 0xeb, 0x37, 0x9a, 0x67,
	0x63, 0x3c, 0x71, 0xad, 0x35, 0xce, 0x70, 0x3a, 0xe8, 0x23, 0xe4, 0xe1, 0xef, 0x0a, 0x35, 0xe8,
	0x3b, 0x4e, 0x07, 0x7d, 0xb3, 0x1b, 0x3d, 0xe8, 0xa3, 0x75, 0xcf, 0x6e, 0x3e, 0xfd, 0xed, 0xf5,
	0xa0, 0x6f, 0xcb, 0xf7, 0xba, 0xa8, 0xed, 0x2f, 0xdf, 0xeb, 0x4e, 0xcb, 0xf7, 0xba, 0xb6, 0x7c,
	0xaf, 0x8b, 0x36, 0xae, 0x29, 0xff, 0xdb, 0x6b, 0xeb, 0xe7, 0x52, 0xce, 0x19, 0xda, 0xbc, 0xc1,
	0x45, 0xe5, 0xdc, 0xdc, 0xbb, 0x48, 0xaa, 0x12, 0xdd, 0xbf, 0xc1, 0x53, 0x30, 0x10, 0xbe, 0x04,
	0xad, 0xb4, 0x52, 0x44, 0xa2, 0x8e, 0xe7, 0x00, 0x07, 0x26, 0xea, 0x46, 0x6b, 0xe5, 0xc1, 0xc2,
	0xf7, 0x1f, 0x9d, 0x7f, 0xbe, 0xfe, 0xec, 0x04, 0xe9, 0xbc, 0x55, 0xbd, 0x5f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x99, 0xa8, 0x9a, 0xae, 0xe4, 0x06, 0x00, 0x00,
}

func (m *Types) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Any != nil {
		l = m.Any.ProtoSize()
		n += 1 + l + sovIssue438(uint64(l))
	}
	if m.Api != nil {
		l = m.Api.ProtoSize()
		n += 1 + l + sovIssue438(uint64(l))
	}
	if m.Met != nil {
		l = m.Met.ProtoSize()
		n += 1 + l + sovIssue438(uint64(l))
	}
	if m.Mx != nil {
		l = m.Mx.ProtoSize()
		n += 1 + l + sovIssue438(uint64(l))
	}
	if m.Dur != nil {
		l = m.Dur.ProtoSize()
		n += 1 + l + sovIssue438(uint64(l))
	}
	if m.Em != nil {
		l = m.Em.ProtoSize()
		n += 1 + l + sovIssue438(uint64(l))
	}
	if m.Fm != nil {
		l = m.Fm.ProtoSize()
		n += 1 + l + sovIssue438(uint64(l))
	}
	if m.Sc != nil {
		l = m.Sc.ProtoSize()
		n += 1 + l + sovIssue438(uint64(l))
	}
	if m.St != nil {
		l = m.St.ProtoSize()
		n += 1 + l + sovIssue438(uint64(l))
	}
	if m.Val != nil {
		l = m.Val.ProtoSize()
		n += 1 + l + sovIssue438(uint64(l))
	}
	if m.Nlval != 0 {
		n += 1 + sovIssue438(uint64(m.Nlval))
	}
	if m.Stval != nil {
		l = m.Stval.ProtoSize()
		n += 1 + l + sovIssue438(uint64(l))
	}
	if m.Bval != nil {
		l = m.Bval.ProtoSize()
		n += 1 + l + sovIssue438(uint64(l))
	}
	if m.Strval != nil {
		l = m.Strval.ProtoSize()
		n += 1 + l + sovIssue438(uint64(l))
	}
	if m.Lstv != nil {
		l = m.Lstv.ProtoSize()
		n += 1 + l + sovIssue438(uint64(l))
	}
	if m.Ts != nil {
		l = m.Ts.ProtoSize()
		n += 2 + l + sovIssue438(uint64(l))
	}
	if m.T != nil {
		l = m.T.ProtoSize()
		n += 2 + l + sovIssue438(uint64(l))
	}
	if m.F != nil {
		l = m.F.ProtoSize()
		n += 2 + l + sovIssue438(uint64(l))
	}
	if m.En != nil {
		l = m.En.ProtoSize()
		n += 2 + l + sovIssue438(uint64(l))
	}
	if m.Enval != nil {
		l = m.Enval.ProtoSize()
		n += 2 + l + sovIssue438(uint64(l))
	}
	if m.Opt != nil {
		l = m.Opt.ProtoSize()
		n += 2 + l + sovIssue438(uint64(l))
	}
	if m.Dbl != nil {
		l = m.Dbl.ProtoSize()
		n += 2 + l + sovIssue438(uint64(l))
	}
	if m.Flt != nil {
		l = m.Flt.ProtoSize()
		n += 2 + l + sovIssue438(uint64(l))
	}
	if m.I64 != nil {
		l = m.I64.ProtoSize()
		n += 2 + l + sovIssue438(uint64(l))
	}
	if m.U64 != nil {
		l = m.U64.ProtoSize()
		n += 2 + l + sovIssue438(uint64(l))
	}
	if m.I32 != nil {
		l = m.I32.ProtoSize()
		n += 2 + l + sovIssue438(uint64(l))
	}
	if m.U32 != nil {
		l = m.U32.ProtoSize()
		n += 2 + l + sovIssue438(uint64(l))
	}
	if m.Bool != nil {
		l = m.Bool.ProtoSize()
		n += 2 + l + sovIssue438(uint64(l))
	}
	if m.Str != nil {
		l = m.Str.ProtoSize()
		n += 2 + l + sovIssue438(uint64(l))
	}
	if m.Bytes != nil {
		l = m.Bytes.ProtoSize()
		n += 2 + l + sovIssue438(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovIssue438(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIssue438(x uint64) (n int) {
	return sovIssue438(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
