// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/datastore/v1beta3/entity.proto

package datastore // import "google.golang.org/genproto/googleapis/datastore/v1beta3"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _struct "github.com/golang/protobuf/ptypes/struct"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import latlng "google.golang.org/genproto/googleapis/type/latlng"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A partition ID identifies a grouping of entities. The grouping is always
// by project and namespace, however the namespace ID may be empty.
//
// A partition ID contains several dimensions:
// project ID and namespace ID.
//
// Partition dimensions:
//
// - May be `""`.
// - Must be valid UTF-8 bytes.
// - Must have values that match regex `[A-Za-z\d\.\-_]{1,100}`
// If the value of any dimension matches regex `__.*__`, the partition is
// reserved/read-only.
// A reserved/read-only partition ID is forbidden in certain documented
// contexts.
//
// Foreign partition IDs (in which the project ID does
// not match the context project ID ) are discouraged.
// Reads and writes of foreign partition IDs may fail if the project is not in an active state.
type PartitionId struct {
	// The ID of the project to which the entities belong.
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// If not empty, the ID of the namespace to which the entities belong.
	NamespaceId          string   `protobuf:"bytes,4,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PartitionId) Reset()         { *m = PartitionId{} }
func (m *PartitionId) String() string { return proto.CompactTextString(m) }
func (*PartitionId) ProtoMessage()    {}
func (*PartitionId) Descriptor() ([]byte, []int) {
	return fileDescriptor_entity_b4336e07cf80b29c, []int{0}
}
func (m *PartitionId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartitionId.Unmarshal(m, b)
}
func (m *PartitionId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartitionId.Marshal(b, m, deterministic)
}
func (dst *PartitionId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartitionId.Merge(dst, src)
}
func (m *PartitionId) XXX_Size() int {
	return xxx_messageInfo_PartitionId.Size(m)
}
func (m *PartitionId) XXX_DiscardUnknown() {
	xxx_messageInfo_PartitionId.DiscardUnknown(m)
}

var xxx_messageInfo_PartitionId proto.InternalMessageInfo

func (m *PartitionId) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *PartitionId) GetNamespaceId() string {
	if m != nil {
		return m.NamespaceId
	}
	return ""
}

// A unique identifier for an entity.
// If a key's partition ID or any of its path kinds or names are
// reserved/read-only, the key is reserved/read-only.
// A reserved/read-only key is forbidden in certain documented contexts.
type Key struct {
	// Entities are partitioned into subsets, currently identified by a project
	// ID and namespace ID.
	// Queries are scoped to a single partition.
	PartitionId *PartitionId `protobuf:"bytes,1,opt,name=partition_id,json=partitionId,proto3" json:"partition_id,omitempty"`
	// The entity path.
	// An entity path consists of one or more elements composed of a kind and a
	// string or numerical identifier, which identify entities. The first
	// element identifies a _root entity_, the second element identifies
	// a _child_ of the root entity, the third element identifies a child of the
	// second entity, and so forth. The entities identified by all prefixes of
	// the path are called the element's _ancestors_.
	//
	// An entity path is always fully complete: *all* of the entity's ancestors
	// are required to be in the path along with the entity identifier itself.
	// The only exception is that in some documented cases, the identifier in the
	// last path element (for the entity) itself may be omitted. For example,
	// the last path element of the key of `Mutation.insert` may have no
	// identifier.
	//
	// A path can never be empty, and a path can have at most 100 elements.
	Path                 []*Key_PathElement `protobuf:"bytes,2,rep,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}
func (*Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_entity_b4336e07cf80b29c, []int{1}
}
func (m *Key) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Key.Unmarshal(m, b)
}
func (m *Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Key.Marshal(b, m, deterministic)
}
func (dst *Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key.Merge(dst, src)
}
func (m *Key) XXX_Size() int {
	return xxx_messageInfo_Key.Size(m)
}
func (m *Key) XXX_DiscardUnknown() {
	xxx_messageInfo_Key.DiscardUnknown(m)
}

var xxx_messageInfo_Key proto.InternalMessageInfo

func (m *Key) GetPartitionId() *PartitionId {
	if m != nil {
		return m.PartitionId
	}
	return nil
}

func (m *Key) GetPath() []*Key_PathElement {
	if m != nil {
		return m.Path
	}
	return nil
}

// A (kind, ID/name) pair used to construct a key path.
//
// If either name or ID is set, the element is complete.
// If neither is set, the element is incomplete.
type Key_PathElement struct {
	// The kind of the entity.
	// A kind matching regex `__.*__` is reserved/read-only.
	// A kind must not contain more than 1500 bytes when UTF-8 encoded.
	// Cannot be `""`.
	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// The type of ID.
	//
	// Types that are valid to be assigned to IdType:
	//	*Key_PathElement_Id
	//	*Key_PathElement_Name
	IdType               isKey_PathElement_IdType `protobuf_oneof:"id_type"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *Key_PathElement) Reset()         { *m = Key_PathElement{} }
func (m *Key_PathElement) String() string { return proto.CompactTextString(m) }
func (*Key_PathElement) ProtoMessage()    {}
func (*Key_PathElement) Descriptor() ([]byte, []int) {
	return fileDescriptor_entity_b4336e07cf80b29c, []int{1, 0}
}
func (m *Key_PathElement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Key_PathElement.Unmarshal(m, b)
}
func (m *Key_PathElement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Key_PathElement.Marshal(b, m, deterministic)
}
func (dst *Key_PathElement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key_PathElement.Merge(dst, src)
}
func (m *Key_PathElement) XXX_Size() int {
	return xxx_messageInfo_Key_PathElement.Size(m)
}
func (m *Key_PathElement) XXX_DiscardUnknown() {
	xxx_messageInfo_Key_PathElement.DiscardUnknown(m)
}

var xxx_messageInfo_Key_PathElement proto.InternalMessageInfo

type isKey_PathElement_IdType interface {
	isKey_PathElement_IdType()
}

type Key_PathElement_Id struct {
	Id int64 `protobuf:"varint,2,opt,name=id,proto3,oneof"`
}
type Key_PathElement_Name struct {
	Name string `protobuf:"bytes,3,opt,name=name,proto3,oneof"`
}

func (*Key_PathElement_Id) isKey_PathElement_IdType()   {}
func (*Key_PathElement_Name) isKey_PathElement_IdType() {}

func (m *Key_PathElement) GetIdType() isKey_PathElement_IdType {
	if m != nil {
		return m.IdType
	}
	return nil
}

func (m *Key_PathElement) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *Key_PathElement) GetId() int64 {
	if x, ok := m.GetIdType().(*Key_PathElement_Id); ok {
		return x.Id
	}
	return 0
}

func (m *Key_PathElement) GetName() string {
	if x, ok := m.GetIdType().(*Key_PathElement_Name); ok {
		return x.Name
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Key_PathElement) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Key_PathElement_OneofMarshaler, _Key_PathElement_OneofUnmarshaler, _Key_PathElement_OneofSizer, []interface{}{
		(*Key_PathElement_Id)(nil),
		(*Key_PathElement_Name)(nil),
	}
}

func _Key_PathElement_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Key_PathElement)
	// id_type
	switch x := m.IdType.(type) {
	case *Key_PathElement_Id:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Id))
	case *Key_PathElement_Name:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Name)
	case nil:
	default:
		return fmt.Errorf("Key_PathElement.IdType has unexpected type %T", x)
	}
	return nil
}

func _Key_PathElement_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Key_PathElement)
	switch tag {
	case 2: // id_type.id
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.IdType = &Key_PathElement_Id{int64(x)}
		return true, err
	case 3: // id_type.name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.IdType = &Key_PathElement_Name{x}
		return true, err
	default:
		return false, nil
	}
}

func _Key_PathElement_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Key_PathElement)
	// id_type
	switch x := m.IdType.(type) {
	case *Key_PathElement_Id:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Id))
	case *Key_PathElement_Name:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Name)))
		n += len(x.Name)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// An array value.
type ArrayValue struct {
	// Values in the array.
	// The order of this array may not be preserved if it contains a mix of
	// indexed and unindexed values.
	Values               []*Value `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArrayValue) Reset()         { *m = ArrayValue{} }
func (m *ArrayValue) String() string { return proto.CompactTextString(m) }
func (*ArrayValue) ProtoMessage()    {}
func (*ArrayValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_entity_b4336e07cf80b29c, []int{2}
}
func (m *ArrayValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArrayValue.Unmarshal(m, b)
}
func (m *ArrayValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArrayValue.Marshal(b, m, deterministic)
}
func (dst *ArrayValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArrayValue.Merge(dst, src)
}
func (m *ArrayValue) XXX_Size() int {
	return xxx_messageInfo_ArrayValue.Size(m)
}
func (m *ArrayValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ArrayValue.DiscardUnknown(m)
}

var xxx_messageInfo_ArrayValue proto.InternalMessageInfo

func (m *ArrayValue) GetValues() []*Value {
	if m != nil {
		return m.Values
	}
	return nil
}

// A message that can hold any of the supported value types and associated
// metadata.
type Value struct {
	// Must have a value set.
	//
	// Types that are valid to be assigned to ValueType:
	//	*Value_NullValue
	//	*Value_BooleanValue
	//	*Value_IntegerValue
	//	*Value_DoubleValue
	//	*Value_TimestampValue
	//	*Value_KeyValue
	//	*Value_StringValue
	//	*Value_BlobValue
	//	*Value_GeoPointValue
	//	*Value_EntityValue
	//	*Value_ArrayValue
	ValueType isValue_ValueType `protobuf_oneof:"value_type"`
	// The `meaning` field should only be populated for backwards compatibility.
	Meaning int32 `protobuf:"varint,14,opt,name=meaning,proto3" json:"meaning,omitempty"`
	// If the value should be excluded from all indexes including those defined
	// explicitly.
	ExcludeFromIndexes   bool     `protobuf:"varint,19,opt,name=exclude_from_indexes,json=excludeFromIndexes,proto3" json:"exclude_from_indexes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_entity_b4336e07cf80b29c, []int{3}
}
func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (dst *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(dst, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

type isValue_ValueType interface {
	isValue_ValueType()
}

type Value_NullValue struct {
	NullValue _struct.NullValue `protobuf:"varint,11,opt,name=null_value,json=nullValue,proto3,enum=google.protobuf.NullValue,oneof"`
}
type Value_BooleanValue struct {
	BooleanValue bool `protobuf:"varint,1,opt,name=boolean_value,json=booleanValue,proto3,oneof"`
}
type Value_IntegerValue struct {
	IntegerValue int64 `protobuf:"varint,2,opt,name=integer_value,json=integerValue,proto3,oneof"`
}
type Value_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,3,opt,name=double_value,json=doubleValue,proto3,oneof"`
}
type Value_TimestampValue struct {
	TimestampValue *timestamp.Timestamp `protobuf:"bytes,10,opt,name=timestamp_value,json=timestampValue,proto3,oneof"`
}
type Value_KeyValue struct {
	KeyValue *Key `protobuf:"bytes,5,opt,name=key_value,json=keyValue,proto3,oneof"`
}
type Value_StringValue struct {
	StringValue string `protobuf:"bytes,17,opt,name=string_value,json=stringValue,proto3,oneof"`
}
type Value_BlobValue struct {
	BlobValue []byte `protobuf:"bytes,18,opt,name=blob_value,json=blobValue,proto3,oneof"`
}
type Value_GeoPointValue struct {
	GeoPointValue *latlng.LatLng `protobuf:"bytes,8,opt,name=geo_point_value,json=geoPointValue,proto3,oneof"`
}
type Value_EntityValue struct {
	EntityValue *Entity `protobuf:"bytes,6,opt,name=entity_value,json=entityValue,proto3,oneof"`
}
type Value_ArrayValue struct {
	ArrayValue *ArrayValue `protobuf:"bytes,9,opt,name=array_value,json=arrayValue,proto3,oneof"`
}

func (*Value_NullValue) isValue_ValueType()      {}
func (*Value_BooleanValue) isValue_ValueType()   {}
func (*Value_IntegerValue) isValue_ValueType()   {}
func (*Value_DoubleValue) isValue_ValueType()    {}
func (*Value_TimestampValue) isValue_ValueType() {}
func (*Value_KeyValue) isValue_ValueType()       {}
func (*Value_StringValue) isValue_ValueType()    {}
func (*Value_BlobValue) isValue_ValueType()      {}
func (*Value_GeoPointValue) isValue_ValueType()  {}
func (*Value_EntityValue) isValue_ValueType()    {}
func (*Value_ArrayValue) isValue_ValueType()     {}

func (m *Value) GetValueType() isValue_ValueType {
	if m != nil {
		return m.ValueType
	}
	return nil
}

func (m *Value) GetNullValue() _struct.NullValue {
	if x, ok := m.GetValueType().(*Value_NullValue); ok {
		return x.NullValue
	}
	return _struct.NullValue_NULL_VALUE
}

func (m *Value) GetBooleanValue() bool {
	if x, ok := m.GetValueType().(*Value_BooleanValue); ok {
		return x.BooleanValue
	}
	return false
}

func (m *Value) GetIntegerValue() int64 {
	if x, ok := m.GetValueType().(*Value_IntegerValue); ok {
		return x.IntegerValue
	}
	return 0
}

func (m *Value) GetDoubleValue() float64 {
	if x, ok := m.GetValueType().(*Value_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *Value) GetTimestampValue() *timestamp.Timestamp {
	if x, ok := m.GetValueType().(*Value_TimestampValue); ok {
		return x.TimestampValue
	}
	return nil
}

func (m *Value) GetKeyValue() *Key {
	if x, ok := m.GetValueType().(*Value_KeyValue); ok {
		return x.KeyValue
	}
	return nil
}

func (m *Value) GetStringValue() string {
	if x, ok := m.GetValueType().(*Value_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *Value) GetBlobValue() []byte {
	if x, ok := m.GetValueType().(*Value_BlobValue); ok {
		return x.BlobValue
	}
	return nil
}

func (m *Value) GetGeoPointValue() *latlng.LatLng {
	if x, ok := m.GetValueType().(*Value_GeoPointValue); ok {
		return x.GeoPointValue
	}
	return nil
}

func (m *Value) GetEntityValue() *Entity {
	if x, ok := m.GetValueType().(*Value_EntityValue); ok {
		return x.EntityValue
	}
	return nil
}

func (m *Value) GetArrayValue() *ArrayValue {
	if x, ok := m.GetValueType().(*Value_ArrayValue); ok {
		return x.ArrayValue
	}
	return nil
}

func (m *Value) GetMeaning() int32 {
	if m != nil {
		return m.Meaning
	}
	return 0
}

func (m *Value) GetExcludeFromIndexes() bool {
	if m != nil {
		return m.ExcludeFromIndexes
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Value) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Value_OneofMarshaler, _Value_OneofUnmarshaler, _Value_OneofSizer, []interface{}{
		(*Value_NullValue)(nil),
		(*Value_BooleanValue)(nil),
		(*Value_IntegerValue)(nil),
		(*Value_DoubleValue)(nil),
		(*Value_TimestampValue)(nil),
		(*Value_KeyValue)(nil),
		(*Value_StringValue)(nil),
		(*Value_BlobValue)(nil),
		(*Value_GeoPointValue)(nil),
		(*Value_EntityValue)(nil),
		(*Value_ArrayValue)(nil),
	}
}

func _Value_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Value)
	// value_type
	switch x := m.ValueType.(type) {
	case *Value_NullValue:
		b.EncodeVarint(11<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.NullValue))
	case *Value_BooleanValue:
		t := uint64(0)
		if x.BooleanValue {
			t = 1
		}
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *Value_IntegerValue:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.IntegerValue))
	case *Value_DoubleValue:
		b.EncodeVarint(3<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.DoubleValue))
	case *Value_TimestampValue:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TimestampValue); err != nil {
			return err
		}
	case *Value_KeyValue:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.KeyValue); err != nil {
			return err
		}
	case *Value_StringValue:
		b.EncodeVarint(17<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.StringValue)
	case *Value_BlobValue:
		b.EncodeVarint(18<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.BlobValue)
	case *Value_GeoPointValue:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GeoPointValue); err != nil {
			return err
		}
	case *Value_EntityValue:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.EntityValue); err != nil {
			return err
		}
	case *Value_ArrayValue:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ArrayValue); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Value.ValueType has unexpected type %T", x)
	}
	return nil
}

func _Value_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Value)
	switch tag {
	case 11: // value_type.null_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ValueType = &Value_NullValue{_struct.NullValue(x)}
		return true, err
	case 1: // value_type.boolean_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ValueType = &Value_BooleanValue{x != 0}
		return true, err
	case 2: // value_type.integer_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ValueType = &Value_IntegerValue{int64(x)}
		return true, err
	case 3: // value_type.double_value
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.ValueType = &Value_DoubleValue{math.Float64frombits(x)}
		return true, err
	case 10: // value_type.timestamp_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(timestamp.Timestamp)
		err := b.DecodeMessage(msg)
		m.ValueType = &Value_TimestampValue{msg}
		return true, err
	case 5: // value_type.key_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Key)
		err := b.DecodeMessage(msg)
		m.ValueType = &Value_KeyValue{msg}
		return true, err
	case 17: // value_type.string_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.ValueType = &Value_StringValue{x}
		return true, err
	case 18: // value_type.blob_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.ValueType = &Value_BlobValue{x}
		return true, err
	case 8: // value_type.geo_point_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(latlng.LatLng)
		err := b.DecodeMessage(msg)
		m.ValueType = &Value_GeoPointValue{msg}
		return true, err
	case 6: // value_type.entity_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Entity)
		err := b.DecodeMessage(msg)
		m.ValueType = &Value_EntityValue{msg}
		return true, err
	case 9: // value_type.array_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ArrayValue)
		err := b.DecodeMessage(msg)
		m.ValueType = &Value_ArrayValue{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Value_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Value)
	// value_type
	switch x := m.ValueType.(type) {
	case *Value_NullValue:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.NullValue))
	case *Value_BooleanValue:
		n += 1 // tag and wire
		n += 1
	case *Value_IntegerValue:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.IntegerValue))
	case *Value_DoubleValue:
		n += 1 // tag and wire
		n += 8
	case *Value_TimestampValue:
		s := proto.Size(x.TimestampValue)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Value_KeyValue:
		s := proto.Size(x.KeyValue)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Value_StringValue:
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(len(x.StringValue)))
		n += len(x.StringValue)
	case *Value_BlobValue:
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(len(x.BlobValue)))
		n += len(x.BlobValue)
	case *Value_GeoPointValue:
		s := proto.Size(x.GeoPointValue)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Value_EntityValue:
		s := proto.Size(x.EntityValue)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Value_ArrayValue:
		s := proto.Size(x.ArrayValue)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A Datastore data object.
//
// An entity is limited to 1 megabyte when stored. That _roughly_
// corresponds to a limit of 1 megabyte for the serialized form of this
// message.
type Entity struct {
	// The entity's key.
	//
	// An entity must have a key, unless otherwise documented (for example,
	// an entity in `Value.entity_value` may have no key).
	// An entity's kind is its key path's last element's kind,
	// or null if it has no key.
	Key *Key `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The entity's properties.
	// The map's keys are property names.
	// A property name matching regex `__.*__` is reserved.
	// A reserved property name is forbidden in certain documented contexts.
	// The name must not contain more than 500 characters.
	// The name cannot be `""`.
	Properties           map[string]*Value `protobuf:"bytes,3,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Entity) Reset()         { *m = Entity{} }
func (m *Entity) String() string { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()    {}
func (*Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_entity_b4336e07cf80b29c, []int{4}
}
func (m *Entity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entity.Unmarshal(m, b)
}
func (m *Entity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entity.Marshal(b, m, deterministic)
}
func (dst *Entity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entity.Merge(dst, src)
}
func (m *Entity) XXX_Size() int {
	return xxx_messageInfo_Entity.Size(m)
}
func (m *Entity) XXX_DiscardUnknown() {
	xxx_messageInfo_Entity.DiscardUnknown(m)
}

var xxx_messageInfo_Entity proto.InternalMessageInfo

func (m *Entity) GetKey() *Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Entity) GetProperties() map[string]*Value {
	if m != nil {
		return m.Properties
	}
	return nil
}

func init() {
	proto.RegisterType((*PartitionId)(nil), "google.datastore.v1beta3.PartitionId")
	proto.RegisterType((*Key)(nil), "google.datastore.v1beta3.Key")
	proto.RegisterType((*Key_PathElement)(nil), "google.datastore.v1beta3.Key.PathElement")
	proto.RegisterType((*ArrayValue)(nil), "google.datastore.v1beta3.ArrayValue")
	proto.RegisterType((*Value)(nil), "google.datastore.v1beta3.Value")
	proto.RegisterType((*Entity)(nil), "google.datastore.v1beta3.Entity")
	proto.RegisterMapType((map[string]*Value)(nil), "google.datastore.v1beta3.Entity.PropertiesEntry")
}

func init() {
	proto.RegisterFile("google/datastore/v1beta3/entity.proto", fileDescriptor_entity_b4336e07cf80b29c)
}

var fileDescriptor_entity_b4336e07cf80b29c = []byte{
	// 789 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0xdf, 0x8e, 0xdb, 0x44,
	0x14, 0xc6, 0xed, 0x64, 0xb3, 0x5d, 0x1f, 0xbb, 0xbb, 0x65, 0xda, 0x0b, 0x2b, 0x6a, 0xd9, 0x10,
	0x58, 0x29, 0xdc, 0xd8, 0xed, 0x56, 0x08, 0x44, 0xe9, 0x45, 0x03, 0xa1, 0x8e, 0x5a, 0x41, 0x34,
	0xaa, 0xf6, 0x02, 0xad, 0x88, 0x26, 0xf1, 0xd4, 0x1d, 0x62, 0xcf, 0x58, 0xf6, 0xb8, 0xaa, 0x5f,
	0x09, 0xee, 0x78, 0x0c, 0x9e, 0x83, 0x3b, 0x5e, 0x02, 0xcd, 0x1f, 0x3b, 0xab, 0x56, 0x29, 0xdc,
	0x79, 0xce, 0xf9, 0x9d, 0x6f, 0xbe, 0x99, 0x73, 0x26, 0x81, 0x8b, 0x4c, 0x88, 0x2c, 0xa7, 0x71,
	0x4a, 0x24, 0xa9, 0xa5, 0xa8, 0x68, 0xfc, 0xf6, 0xd1, 0x86, 0x4a, 0xf2, 0x38, 0xa6, 0x5c, 0x32,
	0xd9, 0x46, 0x65, 0x25, 0xa4, 0x40, 0xa1, 0xc1, 0xa2, 0x1e, 0x8b, 0x2c, 0x36, 0xbe, 0x6f, 0x05,
	0x48, 0xc9, 0x62, 0xc2, 0xb9, 0x90, 0x44, 0x32, 0xc1, 0x6b, 0x53, 0xd7, 0x67, 0xf5, 0x6a, 0xd3,
	0xbc, 0x8e, 0x6b, 0x59, 0x35, 0x5b, 0x69, 0xb3, 0xe7, 0xef, 0x67, 0x25, 0x2b, 0x68, 0x2d, 0x49,
	0x51, 0x5a, 0xc0, 0x6e, 0x1b, 0xcb, 0xb6, 0xa4, 0x71, 0x4e, 0x64, 0xce, 0x33, 0x93, 0x99, 0xfe,
	0x0c, 0xfe, 0x8a, 0x54, 0x92, 0xa9, 0xcd, 0x96, 0x29, 0x7a, 0x00, 0x50, 0x56, 0xe2, 0x37, 0xba,
	0x95, 0x6b, 0x96, 0x86, 0x83, 0x89, 0x3b, 0xf3, 0xb0, 0x67, 0x23, 0xcb, 0x14, 0x7d, 0x06, 0x01,
	0x27, 0x05, 0xad, 0x4b, 0xb2, 0xa5, 0x0a, 0x38, 0xd2, 0x80, 0xdf, 0xc7, 0x96, 0xe9, 0xf4, 0x6f,
	0x17, 0x86, 0x2f, 0x68, 0x8b, 0x12, 0x08, 0xca, 0x4e, 0x58, 0xa1, 0xee, 0xc4, 0x9d, 0xf9, 0x97,
	0x17, 0xd1, 0xa1, 0x0b, 0x88, 0x6e, 0xd8, 0xc0, 0x7e, 0x79, 0xc3, 0xd3, 0x53, 0x38, 0x2a, 0x89,
	0x7c, 0x13, 0x0e, 0x26, 0xc3, 0x99, 0x7f, 0xf9, 0xe5, 0x61, 0x85, 0x17, 0xb4, 0x8d, 0x56, 0x44,
	0xbe, 0x59, 0xe4, 0xb4, 0xa0, 0x5c, 0x62, 0x5d, 0x36, 0x7e, 0xa5, 0x4e, 0xd8, 0x07, 0x11, 0x82,
	0xa3, 0x1d, 0xe3, 0xc6, 0x8f, 0x87, 0xf5, 0x37, 0xba, 0x03, 0x03, 0x7b, 0xda, 0x61, 0xe2, 0xe0,
	0x01, 0x4b, 0xd1, 0x3d, 0x38, 0x52, 0x87, 0x0a, 0x87, 0x8a, 0x4a, 0x1c, 0xac, 0x57, 0x73, 0x0f,
	0x6e, 0xb1, 0x74, 0xad, 0x2e, 0x71, 0xba, 0x00, 0x78, 0x56, 0x55, 0xa4, 0xbd, 0x22, 0x79, 0x43,
	0xd1, 0xd7, 0x70, 0xfc, 0x56, 0x7d, 0xd4, 0xa1, 0xab, 0x4d, 0x9e, 0x1f, 0x36, 0xa9, 0x0b, 0xb0,
	0xc5, 0xa7, 0x7f, 0x8c, 0x60, 0x64, 0x24, 0x9e, 0x00, 0xf0, 0x26, 0xcf, 0xd7, 0x3a, 0x11, 0xfa,
	0x13, 0x77, 0x76, 0x7a, 0x39, 0xee, 0x64, 0xba, 0xc6, 0x46, 0x3f, 0x35, 0x79, 0xae, 0xf9, 0xc4,
	0xc1, 0x1e, 0xef, 0x16, 0xe8, 0x02, 0x6e, 0x6f, 0x84, 0xc8, 0x29, 0xe1, 0xb6, 0x5e, 0x9d, 0xee,
	0x24, 0x71, 0x70, 0x60, 0xc3, 0x3d, 0xc6, 0xb8, 0xa4, 0x19, 0xad, 0x2c, 0xd6, 0x1d, 0x39, 0xb0,
	0x61, 0x83, 0x7d, 0x0e, 0x41, 0x2a, 0x9a, 0x4d, 0x4e, 0x2d, 0xa5, 0x2e, 0xc1, 0x4d, 0x1c, 0xec,
	0x9b, 0xa8, 0x81, 0x16, 0x70, 0xd6, 0x4f, 0x99, 0xe5, 0x40, 0xb7, 0xf8, 0x43, 0xd3, 0xaf, 0x3a,
	0x2e, 0x71, 0xf0, 0x69, 0x5f, 0x64, 0x64, 0xbe, 0x03, 0x6f, 0x47, 0x5b, 0x2b, 0x30, 0xd2, 0x02,
	0x0f, 0x3e, 0xda, 0xe1, 0xc4, 0xc1, 0x27, 0x3b, 0xda, 0xf6, 0x4e, 0x6b, 0x59, 0x31, 0x9e, 0x59,
	0x81, 0x4f, 0x6c, 0xbb, 0x7c, 0x13, 0x35, 0xd0, 0x39, 0xc0, 0x26, 0x17, 0x1b, 0x8b, 0xa0, 0x89,
	0x3b, 0x0b, 0xd4, 0xed, 0xa9, 0x98, 0x01, 0x9e, 0xc2, 0x59, 0x46, 0xc5, 0xba, 0x14, 0x8c, 0x4b,
	0x4b, 0x9d, 0x68, 0x27, 0x77, 0x3b, 0x27, 0xaa, 0xe5, 0xd1, 0x4b, 0x22, 0x5f, 0xf2, 0x2c, 0x71,
	0xf0, 0xed, 0x8c, 0x8a, 0x95, 0x82, 0xbb, 0x9b, 0x08, 0xcc, 0x1b, 0xb7, 0xb5, 0xc7, 0xba, 0x76,
	0x72, 0xf8, 0x14, 0x0b, 0x4d, 0x2b, 0x9b, 0xa6, 0xce, 0xc8, 0x3c, 0x07, 0x9f, 0xa8, 0x89, 0xb2,
	0x2a, 0x9e, 0x56, 0xf9, 0xe2, 0xb0, 0xca, 0x7e, 0xfc, 0x12, 0x07, 0x03, 0xd9, 0x0f, 0x63, 0x08,
	0xb7, 0x0a, 0x4a, 0x38, 0xe3, 0x59, 0x78, 0x3a, 0x71, 0x67, 0x23, 0xdc, 0x2d, 0xd1, 0x43, 0xb8,
	0x47, 0xdf, 0x6d, 0xf3, 0x26, 0xa5, 0xeb, 0xd7, 0x95, 0x28, 0xd6, 0x8c, 0xa7, 0xf4, 0x1d, 0xad,
	0xc3, 0xbb, 0x6a, 0x5a, 0x30, 0xb2, 0xb9, 0x1f, 0x2b, 0x51, 0x2c, 0x4d, 0x66, 0x1e, 0x00, 0x68,
	0x3b, 0x66, 0xe8, 0xff, 0x71, 0xe1, 0xd8, 0x98, 0x47, 0x31, 0x0c, 0x77, 0xb4, 0xb5, 0xaf, 0xfa,
	0xe3, 0x1d, 0xc3, 0x8a, 0x44, 0x2b, 0xfd, 0xcb, 0x52, 0xd2, 0x4a, 0x32, 0x5a, 0x87, 0x43, 0xfd,
	0x4c, 0x1e, 0xfe, 0xd7, 0x1d, 0x45, 0xab, 0xbe, 0x64, 0xc1, 0x65, 0xd5, 0xe2, 0x1b, 0x1a, 0xe3,
	0x5f, 0xe1, 0xec, 0xbd, 0x34, 0xba, 0xb3, 0x77, 0xe5, 0x99, 0x6d, 0xbf, 0x82, 0xd1, 0x7e, 0xd4,
	0xff, 0xc7, 0xc3, 0x34, 0xf4, 0xb7, 0x83, 0x6f, 0xdc, 0xf9, 0x9f, 0x2e, 0xdc, 0xdf, 0x8a, 0xe2,
	0x60, 0xc5, 0xdc, 0x37, 0x26, 0x57, 0x6a, 0xce, 0x57, 0xee, 0x2f, 0xcf, 0x2c, 0x98, 0x89, 0x9c,
	0xf0, 0x2c, 0x12, 0x55, 0x16, 0x67, 0x94, 0xeb, 0x57, 0x10, 0x9b, 0x14, 0x29, 0x59, 0xfd, 0xe1,
	0x3f, 0xc4, 0x93, 0x3e, 0xf2, 0xfb, 0xe0, 0xd3, 0xe7, 0x46, 0xe3, 0xfb, 0x5c, 0x34, 0x69, 0xf4,
	0x43, 0xbf, 0xe5, 0xd5, 0xa3, 0xb9, 0x42, 0xff, 0xea, 0x80, 0x6b, 0x0d, 0x5c, 0xf7, 0xc0, 0xf5,
	0x95, 0xd1, 0xda, 0x1c, 0xeb, 0xfd, 0x1e, 0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x6d, 0x30,
	0xe2, 0x90, 0x06, 0x00, 0x00,
}
