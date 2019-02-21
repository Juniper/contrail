// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: issue498.proto

package issue449

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import bytes "bytes"

import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Message struct {
	Uint8                *uint8   `protobuf:"varint,1,req,name=uint8,casttype=uint8" json:"uint8,omitempty"`
	Uint16               *uint16  `protobuf:"varint,2,req,name=uint16,casttype=uint16" json:"uint16,omitempty"`
	Int8                 *int8    `protobuf:"varint,3,req,name=int8,casttype=int8" json:"int8,omitempty"`
	Int16                *int16   `protobuf:"varint,4,req,name=int16,casttype=int16" json:"int16,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_issue498_7fda9e1d424f3446, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return m.Size()
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetUint8() uint8 {
	if m != nil && m.Uint8 != nil {
		return *m.Uint8
	}
	return 0
}

func (m *Message) GetUint16() uint16 {
	if m != nil && m.Uint16 != nil {
		return *m.Uint16
	}
	return 0
}

func (m *Message) GetInt8() int8 {
	if m != nil && m.Int8 != nil {
		return *m.Int8
	}
	return 0
}

func (m *Message) GetInt16() int16 {
	if m != nil && m.Int16 != nil {
		return *m.Int16
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "issue449.Message")
}
func (this *Message) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Message)
	if !ok {
		that2, ok := that.(Message)
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
	if this.Uint8 != nil && that1.Uint8 != nil {
		if *this.Uint8 != *that1.Uint8 {
			return false
		}
	} else if this.Uint8 != nil {
		return false
	} else if that1.Uint8 != nil {
		return false
	}
	if this.Uint16 != nil && that1.Uint16 != nil {
		if *this.Uint16 != *that1.Uint16 {
			return false
		}
	} else if this.Uint16 != nil {
		return false
	} else if that1.Uint16 != nil {
		return false
	}
	if this.Int8 != nil && that1.Int8 != nil {
		if *this.Int8 != *that1.Int8 {
			return false
		}
	} else if this.Int8 != nil {
		return false
	} else if that1.Int8 != nil {
		return false
	}
	if this.Int16 != nil && that1.Int16 != nil {
		if *this.Int16 != *that1.Int16 {
			return false
		}
	} else if this.Int16 != nil {
		return false
	} else if that1.Int16 != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (m *Message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Uint8 == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("uint8")
	} else {
		dAtA[i] = 0x8
		i++
		i = encodeVarintIssue498(dAtA, i, uint64(*m.Uint8))
	}
	if m.Uint16 == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("uint16")
	} else {
		dAtA[i] = 0x10
		i++
		i = encodeVarintIssue498(dAtA, i, uint64(*m.Uint16))
	}
	if m.Int8 == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("int8")
	} else {
		dAtA[i] = 0x18
		i++
		i = encodeVarintIssue498(dAtA, i, uint64(*m.Int8))
	}
	if m.Int16 == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("int16")
	} else {
		dAtA[i] = 0x20
		i++
		i = encodeVarintIssue498(dAtA, i, uint64(*m.Int16))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintIssue498(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedMessage(r randyIssue498, easy bool) *Message {
	this := &Message{}
	v1 := uint8(r.Uint32())
	this.Uint8 = &v1
	v2 := uint16(r.Uint32())
	this.Uint16 = &v2
	v3 := int8(r.Uint32())
	this.Int8 = &v3
	v4 := int16(r.Uint32())
	this.Int16 = &v4
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedIssue498(r, 5)
	}
	return this
}

type randyIssue498 interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneIssue498(r randyIssue498) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringIssue498(r randyIssue498) string {
	v5 := r.Intn(100)
	tmps := make([]rune, v5)
	for i := 0; i < v5; i++ {
		tmps[i] = randUTF8RuneIssue498(r)
	}
	return string(tmps)
}
func randUnrecognizedIssue498(r randyIssue498, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldIssue498(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldIssue498(dAtA []byte, r randyIssue498, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateIssue498(dAtA, uint64(key))
		v6 := r.Int63()
		if r.Intn(2) == 0 {
			v6 *= -1
		}
		dAtA = encodeVarintPopulateIssue498(dAtA, uint64(v6))
	case 1:
		dAtA = encodeVarintPopulateIssue498(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateIssue498(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateIssue498(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateIssue498(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateIssue498(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Uint8 != nil {
		n += 1 + sovIssue498(uint64(*m.Uint8))
	}
	if m.Uint16 != nil {
		n += 1 + sovIssue498(uint64(*m.Uint16))
	}
	if m.Int8 != nil {
		n += 1 + sovIssue498(uint64(*m.Int8))
	}
	if m.Int16 != nil {
		n += 1 + sovIssue498(uint64(*m.Int16))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovIssue498(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozIssue498(x uint64) (n int) {
	return sovIssue498(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Message) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIssue498
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uint8", wireType)
			}
			var v uint8
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue498
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint8(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Uint8 = &v
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uint16", wireType)
			}
			var v uint16
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue498
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint16(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Uint16 = &v
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Int8", wireType)
			}
			var v int8
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue498
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int8(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Int8 = &v
			hasFields[0] |= uint64(0x00000004)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Int16", wireType)
			}
			var v int16
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue498
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int16(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Int16 = &v
			hasFields[0] |= uint64(0x00000008)
		default:
			iNdEx = preIndex
			skippy, err := skipIssue498(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIssue498
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("uint8")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("uint16")
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("int8")
	}
	if hasFields[0]&uint64(0x00000008) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("int16")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipIssue498(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIssue498
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
					return 0, ErrIntOverflowIssue498
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowIssue498
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthIssue498
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowIssue498
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipIssue498(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthIssue498 = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIssue498   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("issue498.proto", fileDescriptor_issue498_7fda9e1d424f3446) }

var fileDescriptor_issue498_7fda9e1d424f3446 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x2c, 0x2e, 0x2e,
	0x4d, 0x35, 0xb1, 0xb4, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf0, 0x4d, 0x2c,
	0xa5, 0x74, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3, 0xf3, 0xd3,
	0xf3, 0xf5, 0xc1, 0x0a, 0x92, 0x4a, 0xd3, 0xc0, 0x3c, 0x30, 0x07, 0xcc, 0x82, 0x68, 0x54, 0xea,
	0x65, 0xe4, 0x62, 0xf7, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x92, 0xe7, 0x62, 0x2d, 0xcd,
	0xcc, 0x2b, 0xb1, 0x90, 0x60, 0x54, 0x60, 0xd2, 0xe0, 0x75, 0xe2, 0xfc, 0x75, 0x4f, 0x1e, 0x22,
	0x10, 0x04, 0xa1, 0x84, 0x94, 0xb8, 0xd8, 0x40, 0x0c, 0x43, 0x33, 0x09, 0x26, 0xb0, 0x0a, 0xae,
	0x5f, 0xf7, 0xe4, 0xa1, 0x22, 0x41, 0x50, 0x5a, 0x48, 0x86, 0x8b, 0x05, 0x6c, 0x06, 0x33, 0x58,
	0x05, 0xc7, 0xaf, 0x7b, 0xf2, 0x60, 0x7e, 0x10, 0x98, 0x04, 0x59, 0x01, 0x31, 0x80, 0x05, 0x61,
	0x05, 0x44, 0x3f, 0x84, 0x72, 0x92, 0xf8, 0xf1, 0x50, 0x8e, 0x71, 0xc5, 0x23, 0x39, 0xc6, 0x1d,
	0x8f, 0xe4, 0x18, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0x46,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x95, 0x2c, 0xad, 0xb7, 0xf3, 0x00, 0x00, 0x00,
}
