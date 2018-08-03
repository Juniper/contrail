// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/Juniper/contrail/pkg/services/baseservices/base.proto

/*
Package baseservices is a generated protocol buffer package.

It is generated from these files:
	github.com/Juniper/contrail/pkg/services/baseservices/base.proto

It has these top-level messages:
	ListSpec
	Filter
*/
package baseservices

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import reflect "reflect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ListSpec struct {
	Filters      []*Filter `protobuf:"bytes,1,rep,name=filters" json:"filters,omitempty"`
	Limit        int64     `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset       int64     `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Detail       bool      `protobuf:"varint,4,opt,name=detail,proto3" json:"detail,omitempty"`
	Count        bool      `protobuf:"varint,5,opt,name=count,proto3" json:"count,omitempty"`
	Shared       bool      `protobuf:"varint,6,opt,name=shared,proto3" json:"shared,omitempty"`
	ExcludeHrefs bool      `protobuf:"varint,7,opt,name=exclude_hrefs,json=excludeHrefs,proto3" json:"exclude_hrefs,omitempty"`
	ParentFQName []string  `protobuf:"bytes,8,rep,name=parent_fq_name,json=parentFqName" json:"parent_fq_name,omitempty"`
	ParentType   string    `protobuf:"bytes,9,opt,name=parent_type,json=parentType,proto3" json:"parent_type,omitempty"`
	ParentUUIDs  []string  `protobuf:"bytes,10,rep,name=parent_uuids,json=parentUuids" json:"parent_uuids,omitempty"`
	BackRefUUIDs []string  `protobuf:"bytes,11,rep,name=backref_uuids,json=backrefUuids" json:"backref_uuids,omitempty"`
	ObjectUUIDs  []string  `protobuf:"bytes,12,rep,name=object_uuids,json=objectUuids" json:"object_uuids,omitempty"`
	Fields       []string  `protobuf:"bytes,13,rep,name=fields" json:"fields,omitempty"`
}

func (m *ListSpec) Reset()                    { *m = ListSpec{} }
func (*ListSpec) ProtoMessage()               {}
func (*ListSpec) Descriptor() ([]byte, []int) { return fileDescriptorBase, []int{0} }

func (m *ListSpec) GetFilters() []*Filter {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *ListSpec) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListSpec) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListSpec) GetDetail() bool {
	if m != nil {
		return m.Detail
	}
	return false
}

func (m *ListSpec) GetCount() bool {
	if m != nil {
		return m.Count
	}
	return false
}

func (m *ListSpec) GetShared() bool {
	if m != nil {
		return m.Shared
	}
	return false
}

func (m *ListSpec) GetExcludeHrefs() bool {
	if m != nil {
		return m.ExcludeHrefs
	}
	return false
}

func (m *ListSpec) GetParentFQName() []string {
	if m != nil {
		return m.ParentFQName
	}
	return nil
}

func (m *ListSpec) GetParentType() string {
	if m != nil {
		return m.ParentType
	}
	return ""
}

func (m *ListSpec) GetParentUUIDs() []string {
	if m != nil {
		return m.ParentUUIDs
	}
	return nil
}

func (m *ListSpec) GetBackRefUUIDs() []string {
	if m != nil {
		return m.BackRefUUIDs
	}
	return nil
}

func (m *ListSpec) GetObjectUUIDs() []string {
	if m != nil {
		return m.ObjectUUIDs
	}
	return nil
}

func (m *ListSpec) GetFields() []string {
	if m != nil {
		return m.Fields
	}
	return nil
}

type Filter struct {
	Key    string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values []string `protobuf:"bytes,2,rep,name=values" json:"values,omitempty"`
}

func (m *Filter) Reset()                    { *m = Filter{} }
func (*Filter) ProtoMessage()               {}
func (*Filter) Descriptor() ([]byte, []int) { return fileDescriptorBase, []int{1} }

func (m *Filter) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Filter) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*ListSpec)(nil), "github.com.Juniper.contrail.pkg.services.baseservices.ListSpec")
	proto.RegisterType((*Filter)(nil), "github.com.Juniper.contrail.pkg.services.baseservices.Filter")
}
func (m *ListSpec) Size() (n int) {
	var l int
	_ = l
	if len(m.Filters) > 0 {
		for _, e := range m.Filters {
			l = e.Size()
			n += 1 + l + sovBase(uint64(l))
		}
	}
	if m.Limit != 0 {
		n += 1 + sovBase(uint64(m.Limit))
	}
	if m.Offset != 0 {
		n += 1 + sovBase(uint64(m.Offset))
	}
	if m.Detail {
		n += 2
	}
	if m.Count {
		n += 2
	}
	if m.Shared {
		n += 2
	}
	if m.ExcludeHrefs {
		n += 2
	}
	if len(m.ParentFQName) > 0 {
		for _, s := range m.ParentFQName {
			l = len(s)
			n += 1 + l + sovBase(uint64(l))
		}
	}
	l = len(m.ParentType)
	if l > 0 {
		n += 1 + l + sovBase(uint64(l))
	}
	if len(m.ParentUUIDs) > 0 {
		for _, s := range m.ParentUUIDs {
			l = len(s)
			n += 1 + l + sovBase(uint64(l))
		}
	}
	if len(m.BackRefUUIDs) > 0 {
		for _, s := range m.BackRefUUIDs {
			l = len(s)
			n += 1 + l + sovBase(uint64(l))
		}
	}
	if len(m.ObjectUUIDs) > 0 {
		for _, s := range m.ObjectUUIDs {
			l = len(s)
			n += 1 + l + sovBase(uint64(l))
		}
	}
	if len(m.Fields) > 0 {
		for _, s := range m.Fields {
			l = len(s)
			n += 1 + l + sovBase(uint64(l))
		}
	}
	return n
}

func (m *Filter) Size() (n int) {
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovBase(uint64(l))
	}
	if len(m.Values) > 0 {
		for _, s := range m.Values {
			l = len(s)
			n += 1 + l + sovBase(uint64(l))
		}
	}
	return n
}

func sovBase(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozBase(x uint64) (n int) {
	return sovBase(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ListSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ListSpec{`,
		`Filters:` + strings.Replace(fmt.Sprintf("%v", this.Filters), "Filter", "Filter", 1) + `,`,
		`Limit:` + fmt.Sprintf("%v", this.Limit) + `,`,
		`Offset:` + fmt.Sprintf("%v", this.Offset) + `,`,
		`Detail:` + fmt.Sprintf("%v", this.Detail) + `,`,
		`Count:` + fmt.Sprintf("%v", this.Count) + `,`,
		`Shared:` + fmt.Sprintf("%v", this.Shared) + `,`,
		`ExcludeHrefs:` + fmt.Sprintf("%v", this.ExcludeHrefs) + `,`,
		`ParentFQName:` + fmt.Sprintf("%v", this.ParentFQName) + `,`,
		`ParentType:` + fmt.Sprintf("%v", this.ParentType) + `,`,
		`ParentUUIDs:` + fmt.Sprintf("%v", this.ParentUUIDs) + `,`,
		`BackRefUUIDs:` + fmt.Sprintf("%v", this.BackRefUUIDs) + `,`,
		`ObjectUUIDs:` + fmt.Sprintf("%v", this.ObjectUUIDs) + `,`,
		`Fields:` + fmt.Sprintf("%v", this.Fields) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Filter) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Filter{`,
		`Key:` + fmt.Sprintf("%v", this.Key) + `,`,
		`Values:` + fmt.Sprintf("%v", this.Values) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringBase(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}

func init() {
	proto.RegisterFile("github.com/Juniper/contrail/pkg/services/baseservices/base.proto", fileDescriptorBase)
}

var fileDescriptorBase = []byte{
	// 558 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x3d, 0x6f, 0xd3, 0x40,
	0x18, 0xb6, 0x6b, 0xe2, 0x34, 0x17, 0xb7, 0x20, 0x4f, 0x07, 0xc3, 0x39, 0xca, 0x94, 0x05, 0x5b,
	0x2a, 0x2a, 0x1b, 0x08, 0x85, 0x12, 0xbe, 0x2a, 0x0a, 0x57, 0xc2, 0xc0, 0x12, 0xf9, 0xe3, 0x75,
	0x62, 0xe2, 0xc4, 0xae, 0xcf, 0xae, 0xc8, 0x06, 0xff, 0x84, 0x91, 0x9f, 0xc1, 0xc8, 0xd8, 0x91,
	0xc9, 0x6a, 0x8e, 0x3f, 0xc0, 0xc8, 0x88, 0x7c, 0x77, 0x21, 0x61, 0x65, 0x7b, 0xdf, 0xe7, 0xf3,
	0xce, 0x3a, 0xa3, 0x47, 0xd3, 0xa4, 0x9c, 0x55, 0x81, 0x1b, 0x66, 0x0b, 0xef, 0x45, 0xb5, 0x4c,
	0x72, 0x28, 0xbc, 0x30, 0x5b, 0x96, 0x85, 0x9f, 0xa4, 0x5e, 0x3e, 0x9f, 0x7a, 0x0c, 0x8a, 0xcb,
	0x24, 0x04, 0xe6, 0x05, 0x3e, 0x83, 0x7f, 0x16, 0x37, 0x2f, 0xb2, 0x32, 0xb3, 0x8f, 0xb7, 0x09,
	0xae, 0x4a, 0x70, 0x37, 0x09, 0x6e, 0x3e, 0x9f, 0xba, 0x1b, 0x93, 0xbb, 0x9b, 0x70, 0xe7, 0xee,
	0x4e, 0xf1, 0x34, 0x9b, 0x66, 0x9e, 0x48, 0x0b, 0xaa, 0x58, 0x6c, 0x62, 0x11, 0x93, 0x6c, 0xe9,
	0x7f, 0x6e, 0xa1, 0xfd, 0xd3, 0x84, 0x95, 0xe7, 0x39, 0x84, 0x76, 0x84, 0xda, 0x71, 0x92, 0x96,
	0x50, 0x30, 0xac, 0xf7, 0x8c, 0x41, 0xf7, 0xe8, 0x81, 0xfb, 0x5f, 0x87, 0x70, 0x47, 0x22, 0x65,
	0xd8, 0xe5, 0xb5, 0xd3, 0x96, 0x33, 0xa3, 0x9b, 0x68, 0xdb, 0x41, 0xad, 0x34, 0x59, 0x24, 0x25,
	0xde, 0xeb, 0xe9, 0x03, 0x63, 0xd8, 0xe1, 0xb5, 0xd3, 0x3a, 0x6d, 0x00, 0x2a, 0x71, 0xbb, 0x8f,
	0xcc, 0x2c, 0x8e, 0x19, 0x94, 0xd8, 0x10, 0x0a, 0xc4, 0x6b, 0xc7, 0x3c, 0x13, 0x08, 0x55, 0x4c,
	0xa3, 0x89, 0xa0, 0xf4, 0x93, 0x14, 0xdf, 0xe8, 0xe9, 0x83, 0x7d, 0xa9, 0x39, 0x11, 0x08, 0x55,
	0x4c, 0x53, 0x14, 0x66, 0xd5, 0xb2, 0xc4, 0x2d, 0x21, 0x11, 0x45, 0x8f, 0x1b, 0x80, 0x4a, 0xbc,
	0x09, 0x61, 0x33, 0xbf, 0x80, 0x08, 0x9b, 0xdb, 0x90, 0x73, 0x81, 0x50, 0xc5, 0xd8, 0xc7, 0xe8,
	0x00, 0x3e, 0x86, 0x69, 0x15, 0xc1, 0x64, 0x56, 0x40, 0xcc, 0x70, 0x5b, 0x48, 0x6f, 0xf1, 0xda,
	0xb1, 0x9e, 0x48, 0xe2, 0x59, 0x83, 0x53, 0x0b, 0x76, 0x36, 0xfb, 0x3e, 0x3a, 0xcc, 0xfd, 0x02,
	0x96, 0xe5, 0x24, 0xbe, 0x98, 0x2c, 0xfd, 0x05, 0xe0, 0xfd, 0x9e, 0x31, 0xe8, 0x48, 0xdf, 0x6b,
	0xc1, 0x8c, 0xde, 0xbc, 0xf2, 0x17, 0x40, 0x2d, 0xa9, 0x1b, 0x5d, 0x34, 0x9b, 0xed, 0xa1, 0xae,
	0xf2, 0x95, 0xab, 0x1c, 0x70, 0xa7, 0xa7, 0x0f, 0x3a, 0xc3, 0x43, 0x5e, 0x3b, 0x48, 0x9a, 0xde,
	0xae, 0x72, 0xa0, 0x28, 0xff, 0x3b, 0xdb, 0x47, 0x48, 0x05, 0x4c, 0xaa, 0x2a, 0x89, 0x18, 0x46,
	0xa2, 0xe6, 0x26, 0xaf, 0x9d, 0xae, 0x74, 0x8c, 0xc7, 0xcf, 0x4f, 0x18, 0x55, 0xa9, 0xe3, 0x46,
	0xd3, 0xdc, 0x29, 0xf0, 0xc3, 0x79, 0x01, 0xb1, 0x32, 0x75, 0xb7, 0x67, 0x1b, 0xfa, 0xe1, 0x9c,
	0x42, 0x2c, 0x5d, 0x96, 0x92, 0x49, 0xdb, 0x11, 0xb2, 0xb2, 0xe0, 0x03, 0x84, 0x9b, 0x2a, 0x6b,
	0x5b, 0x75, 0x26, 0x70, 0x55, 0x25, 0x45, 0xd2, 0xd3, 0x47, 0x66, 0x9c, 0x40, 0x1a, 0x31, 0x7c,
	0x20, 0xd4, 0xe2, 0x13, 0x8f, 0x04, 0x42, 0x15, 0xd3, 0x7f, 0x8a, 0x4c, 0xf9, 0x48, 0xec, 0xdb,
	0xc8, 0x98, 0xc3, 0x0a, 0xeb, 0xe2, 0xd6, 0x6d, 0x5e, 0x3b, 0xc6, 0x4b, 0x58, 0xd1, 0x06, 0x6b,
	0x82, 0x2e, 0xfd, 0xb4, 0x02, 0x86, 0xf7, 0xb6, 0x41, 0xef, 0x04, 0x42, 0x15, 0x33, 0x7c, 0x78,
	0xb5, 0x26, 0xda, 0x8f, 0x35, 0xd1, 0xae, 0xd7, 0x44, 0xfb, 0xb5, 0x26, 0xda, 0xef, 0x35, 0xd1,
	0x3e, 0x71, 0xa2, 0x7f, 0xe5, 0x44, 0xfb, 0xc6, 0x89, 0xf6, 0x9d, 0x13, 0xed, 0x8a, 0x13, 0xed,
	0x9a, 0x13, 0xfd, 0xcb, 0x4f, 0xa2, 0xbd, 0xb7, 0x76, 0x9f, 0x6d, 0x60, 0x8a, 0x7f, 0xe2, 0xde,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x82, 0xdf, 0x79, 0xbd, 0x03, 0x00, 0x00,
}
