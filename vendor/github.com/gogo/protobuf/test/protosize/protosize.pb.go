// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protosize.proto

package protosize

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
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

type SizeMessage struct {
	Size                 *int64   `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
	ProtoSize_           *int64   `protobuf:"varint,2,opt,name=proto_size,json=protoSize" json:"proto_size,omitempty"`
	Equal_               *bool    `protobuf:"varint,3,opt,name=Equal" json:"Equal,omitempty"`
	String_              *string  `protobuf:"bytes,4,opt,name=String" json:"String,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SizeMessage) Reset()         { *m = SizeMessage{} }
func (m *SizeMessage) String() string { return proto.CompactTextString(m) }
func (*SizeMessage) ProtoMessage()    {}
func (*SizeMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_16520eec64ed25c4, []int{0}
}
func (m *SizeMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SizeMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SizeMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SizeMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SizeMessage.Merge(m, src)
}
func (m *SizeMessage) XXX_Size() int {
	return m.ProtoSize()
}
func (m *SizeMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SizeMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SizeMessage proto.InternalMessageInfo

func (m *SizeMessage) GetSize() int64 {
	if m != nil && m.Size != nil {
		return *m.Size
	}
	return 0
}

func (m *SizeMessage) GetProtoSize_() int64 {
	if m != nil && m.ProtoSize_ != nil {
		return *m.ProtoSize_
	}
	return 0
}

func (m *SizeMessage) GetEqual_() bool {
	if m != nil && m.Equal_ != nil {
		return *m.Equal_
	}
	return false
}

func (m *SizeMessage) GetString_() string {
	if m != nil && m.String_ != nil {
		return *m.String_
	}
	return ""
}

func init() {
	proto.RegisterType((*SizeMessage)(nil), "protosize.SizeMessage")
}

func init() { proto.RegisterFile("protosize.proto", fileDescriptor_16520eec64ed25c4) }

var fileDescriptor_16520eec64ed25c4 = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xce, 0xac, 0x4a, 0xd5, 0x03, 0xb3, 0x84, 0x38, 0xe1, 0x02, 0x52, 0xba, 0xe9, 0x99,
	0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xe9, 0xf9, 0xe9, 0xf9, 0xfa, 0x60, 0xa9,
	0xa4, 0xd2, 0x34, 0x30, 0x0f, 0xcc, 0x01, 0xb3, 0x20, 0x3a, 0x95, 0xf2, 0xb8, 0xb8, 0x83, 0x33,
	0xab, 0x52, 0x7d, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x85, 0x84, 0xb8, 0x58, 0x40, 0xa6, 0x48,
	0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0x81, 0xd9, 0x42, 0xb2, 0x5c, 0x5c, 0x60, 0xb5, 0xf1, 0x60,
	0x19, 0x26, 0xb0, 0x0c, 0xc4, 0x42, 0x90, 0x4e, 0x21, 0x11, 0x2e, 0x56, 0xd7, 0xc2, 0xd2, 0xc4,
	0x1c, 0x09, 0x66, 0x05, 0x46, 0x0d, 0x8e, 0x20, 0x08, 0x47, 0x48, 0x8c, 0x8b, 0x2d, 0xb8, 0xa4,
	0x28, 0x33, 0x2f, 0x5d, 0x82, 0x45, 0x81, 0x51, 0x83, 0x33, 0x08, 0xca, 0x73, 0x92, 0xf8, 0xf1,
	0x50, 0x8e, 0x71, 0xc5, 0x23, 0x39, 0xc6, 0x1d, 0x8f, 0xe4, 0x18, 0x4f, 0x3c, 0x92, 0x63, 0xbc,
	0xf0, 0x48, 0x8e, 0x71, 0xc1, 0x63, 0x39, 0x46, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8b, 0xf7,
	0x87, 0xb3, 0xd5, 0x00, 0x00, 0x00,
}

func (this *SizeMessage) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SizeMessage)
	if !ok {
		that2, ok := that.(SizeMessage)
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
	if this.Size != nil && that1.Size != nil {
		if *this.Size != *that1.Size {
			return false
		}
	} else if this.Size != nil {
		return false
	} else if that1.Size != nil {
		return false
	}
	if this.ProtoSize_ != nil && that1.ProtoSize_ != nil {
		if *this.ProtoSize_ != *that1.ProtoSize_ {
			return false
		}
	} else if this.ProtoSize_ != nil {
		return false
	} else if that1.ProtoSize_ != nil {
		return false
	}
	if this.Equal_ != nil && that1.Equal_ != nil {
		if *this.Equal_ != *that1.Equal_ {
			return false
		}
	} else if this.Equal_ != nil {
		return false
	} else if that1.Equal_ != nil {
		return false
	}
	if this.String_ != nil && that1.String_ != nil {
		if *this.String_ != *that1.String_ {
			return false
		}
	} else if this.String_ != nil {
		return false
	} else if that1.String_ != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (m *SizeMessage) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SizeMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.ProtoSize()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SizeMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.String_ != nil {
		i -= len(*m.String_)
		copy(dAtA[i:], *m.String_)
		i = encodeVarintProtosize(dAtA, i, uint64(len(*m.String_)))
		i--
		dAtA[i] = 0x22
	}
	if m.Equal_ != nil {
		i--
		if *m.Equal_ {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.ProtoSize_ != nil {
		i = encodeVarintProtosize(dAtA, i, uint64(*m.ProtoSize_))
		i--
		dAtA[i] = 0x10
	}
	if m.Size != nil {
		i = encodeVarintProtosize(dAtA, i, uint64(*m.Size))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintProtosize(dAtA []byte, offset int, v uint64) int {
	offset -= sovProtosize(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedSizeMessage(r randyProtosize, easy bool) *SizeMessage {
	this := &SizeMessage{}
	if r.Intn(5) != 0 {
		v1 := int64(r.Int63())
		if r.Intn(2) == 0 {
			v1 *= -1
		}
		this.Size = &v1
	}
	if r.Intn(5) != 0 {
		v2 := int64(r.Int63())
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		this.ProtoSize_ = &v2
	}
	if r.Intn(5) != 0 {
		v3 := bool(bool(r.Intn(2) == 0))
		this.Equal_ = &v3
	}
	if r.Intn(5) != 0 {
		v4 := string(randStringProtosize(r))
		this.String_ = &v4
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedProtosize(r, 5)
	}
	return this
}

type randyProtosize interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneProtosize(r randyProtosize) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringProtosize(r randyProtosize) string {
	v5 := r.Intn(100)
	tmps := make([]rune, v5)
	for i := 0; i < v5; i++ {
		tmps[i] = randUTF8RuneProtosize(r)
	}
	return string(tmps)
}
func randUnrecognizedProtosize(r randyProtosize, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldProtosize(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldProtosize(dAtA []byte, r randyProtosize, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateProtosize(dAtA, uint64(key))
		v6 := r.Int63()
		if r.Intn(2) == 0 {
			v6 *= -1
		}
		dAtA = encodeVarintPopulateProtosize(dAtA, uint64(v6))
	case 1:
		dAtA = encodeVarintPopulateProtosize(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateProtosize(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateProtosize(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateProtosize(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateProtosize(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *SizeMessage) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Size != nil {
		n += 1 + sovProtosize(uint64(*m.Size))
	}
	if m.ProtoSize_ != nil {
		n += 1 + sovProtosize(uint64(*m.ProtoSize_))
	}
	if m.Equal_ != nil {
		n += 2
	}
	if m.String_ != nil {
		l = len(*m.String_)
		n += 1 + l + sovProtosize(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovProtosize(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProtosize(x uint64) (n int) {
	return sovProtosize(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SizeMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtosize
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SizeMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SizeMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtosize
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Size = &v
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProtoSize_", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtosize
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ProtoSize_ = &v
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Equal_", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtosize
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Equal_ = &b
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field String_", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtosize
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProtosize
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtosize
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.String_ = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtosize(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtosize
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProtosize
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipProtosize(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProtosize
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProtosize
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProtosize
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthProtosize
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProtosize
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProtosize
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProtosize        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProtosize          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProtosize = fmt.Errorf("proto: unexpected end of group")
)
