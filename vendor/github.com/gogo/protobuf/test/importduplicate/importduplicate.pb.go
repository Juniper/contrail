// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: importduplicate.proto

package importduplicate

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
	proto1 "github.com/gogo/protobuf/test/importduplicate/proto"
	sortkeys "github.com/gogo/protobuf/test/importduplicate/sortkeys"
	math "math"
	reflect "reflect"
	strings "strings"
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

type MapAndSortKeys struct {
	Key                  *sortkeys.Object `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	KeyValue             map[int32]string `protobuf:"bytes,2,rep,name=keyValue,proto3" json:"keyValue,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Value                *proto1.Subject  `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *MapAndSortKeys) Reset()         { *m = MapAndSortKeys{} }
func (m *MapAndSortKeys) String() string { return proto.CompactTextString(m) }
func (*MapAndSortKeys) ProtoMessage()    {}
func (*MapAndSortKeys) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3b420b76fd5209f, []int{0}
}
func (m *MapAndSortKeys) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MapAndSortKeys.Unmarshal(m, b)
}
func (m *MapAndSortKeys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MapAndSortKeys.Marshal(b, m, deterministic)
}
func (m *MapAndSortKeys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapAndSortKeys.Merge(m, src)
}
func (m *MapAndSortKeys) XXX_Size() int {
	return xxx_messageInfo_MapAndSortKeys.Size(m)
}
func (m *MapAndSortKeys) XXX_DiscardUnknown() {
	xxx_messageInfo_MapAndSortKeys.DiscardUnknown(m)
}

var xxx_messageInfo_MapAndSortKeys proto.InternalMessageInfo

func (m *MapAndSortKeys) GetKey() *sortkeys.Object {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *MapAndSortKeys) GetKeyValue() map[int32]string {
	if m != nil {
		return m.KeyValue
	}
	return nil
}

func (m *MapAndSortKeys) GetValue() *proto1.Subject {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*MapAndSortKeys)(nil), "importduplicate.MapAndSortKeys")
	proto.RegisterMapType((map[int32]string)(nil), "importduplicate.MapAndSortKeys.KeyValueEntry")
}

func init() { proto.RegisterFile("importduplicate.proto", fileDescriptor_f3b420b76fd5209f) }

var fileDescriptor_f3b420b76fd5209f = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcd, 0xcc, 0x2d, 0xc8,
	0x2f, 0x2a, 0x49, 0x29, 0x2d, 0xc8, 0xc9, 0x4c, 0x4e, 0x2c, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x47, 0x13, 0x96, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce,
	0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0xab, 0x4b, 0x2a, 0x4d, 0x03, 0xf3, 0xc0, 0x1c,
	0x30, 0x0b, 0xa2, 0x5f, 0xca, 0x15, 0xa7, 0xf2, 0x92, 0xd4, 0xe2, 0x12, 0x7d, 0x34, 0xd3, 0xf5,
	0x8b, 0xf3, 0x8b, 0x4a, 0xb2, 0x53, 0x2b, 0x8b, 0xc1, 0x8c, 0xc4, 0xa4, 0x1c, 0xa8, 0x33, 0xa4,
	0xec, 0x49, 0x33, 0x06, 0xe2, 0x0c, 0x30, 0x09, 0x31, 0x40, 0xe9, 0x11, 0x23, 0x17, 0x9f, 0x6f,
	0x62, 0x81, 0x63, 0x5e, 0x4a, 0x70, 0x7e, 0x51, 0x89, 0x77, 0x6a, 0x65, 0xb1, 0x90, 0x12, 0x17,
	0x73, 0x76, 0x6a, 0xa5, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0x80, 0x1e, 0xcc, 0x6a, 0x3d,
	0xff, 0xa4, 0xac, 0xd4, 0xe4, 0x92, 0x20, 0x90, 0xa4, 0x90, 0x27, 0x17, 0x47, 0x76, 0x6a, 0x65,
	0x58, 0x62, 0x4e, 0x69, 0xaa, 0x04, 0x93, 0x02, 0xb3, 0x06, 0xb7, 0x91, 0xae, 0x1e, 0x7a, 0x40,
	0xa1, 0x1a, 0xab, 0xe7, 0x0d, 0x55, 0xef, 0x9a, 0x57, 0x52, 0x54, 0x19, 0x04, 0xd7, 0x2e, 0xa4,
	0xc2, 0xc5, 0x5a, 0x06, 0x36, 0x87, 0x19, 0x6c, 0x21, 0x1f, 0xc4, 0x61, 0x7a, 0xc1, 0xa5, 0x10,
	0xeb, 0x20, 0x92, 0x52, 0xd6, 0x5c, 0xbc, 0x28, 0x06, 0x08, 0x09, 0x20, 0x5c, 0xc9, 0x0a, 0x71,
	0x93, 0x08, 0xcc, 0x20, 0x26, 0x05, 0x46, 0x0d, 0x4e, 0xa8, 0x46, 0x2b, 0x26, 0x0b, 0x46, 0x27,
	0x81, 0x0f, 0x0f, 0xe5, 0x18, 0x7f, 0x3c, 0x94, 0x63, 0x5c, 0xf1, 0x48, 0x8e, 0x71, 0xc7, 0x23,
	0x39, 0xc6, 0x24, 0x36, 0xb0, 0x25, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x50, 0xcc, 0xec,
	0x38, 0xde, 0x01, 0x00, 0x00,
}

func (this *MapAndSortKeys) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MapAndSortKeys)
	if !ok {
		that2, ok := that.(MapAndSortKeys)
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
	if !this.Key.Equal(that1.Key) {
		return false
	}
	if len(this.KeyValue) != len(that1.KeyValue) {
		return false
	}
	for i := range this.KeyValue {
		if this.KeyValue[i] != that1.KeyValue[i] {
			return false
		}
	}
	if !this.Value.Equal(that1.Value) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MapAndSortKeys) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&importduplicate.MapAndSortKeys{")
	if this.Key != nil {
		s = append(s, "Key: "+fmt.Sprintf("%#v", this.Key)+",\n")
	}
	keysForKeyValue := make([]int32, 0, len(this.KeyValue))
	for k := range this.KeyValue {
		keysForKeyValue = append(keysForKeyValue, k)
	}
	github_com_gogo_protobuf_sortkeys.Int32s(keysForKeyValue)
	mapStringForKeyValue := "map[int32]string{"
	for _, k := range keysForKeyValue {
		mapStringForKeyValue += fmt.Sprintf("%#v: %#v,", k, this.KeyValue[k])
	}
	mapStringForKeyValue += "}"
	if this.KeyValue != nil {
		s = append(s, "KeyValue: "+mapStringForKeyValue+",\n")
	}
	if this.Value != nil {
		s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringImportduplicate(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func NewPopulatedMapAndSortKeys(r randyImportduplicate, easy bool) *MapAndSortKeys {
	this := &MapAndSortKeys{}
	if r.Intn(5) != 0 {
		this.Key = sortkeys.NewPopulatedObject(r, easy)
	}
	if r.Intn(5) != 0 {
		v1 := r.Intn(10)
		this.KeyValue = make(map[int32]string)
		for i := 0; i < v1; i++ {
			this.KeyValue[int32(r.Int31())] = randStringImportduplicate(r)
		}
	}
	if r.Intn(5) != 0 {
		this.Value = proto1.NewPopulatedSubject(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedImportduplicate(r, 4)
	}
	return this
}

type randyImportduplicate interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneImportduplicate(r randyImportduplicate) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringImportduplicate(r randyImportduplicate) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneImportduplicate(r)
	}
	return string(tmps)
}
func randUnrecognizedImportduplicate(r randyImportduplicate, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldImportduplicate(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldImportduplicate(dAtA []byte, r randyImportduplicate, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateImportduplicate(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateImportduplicate(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateImportduplicate(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateImportduplicate(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateImportduplicate(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateImportduplicate(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateImportduplicate(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
