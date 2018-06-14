// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/appengine/v1/audit_data.proto

package appengine // import "google.golang.org/genproto/googleapis/appengine/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/iam/v1"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// App Engine admin service audit log.
type AuditData struct {
	// Detailed information about methods that require it. Does not include
	// simple Get, List or Delete methods because all significant information
	// (resource name, number of returned elements for List operations) is already
	// included in parent audit log message.
	//
	// Types that are valid to be assigned to Method:
	//	*AuditData_UpdateService
	//	*AuditData_CreateVersion
	Method               isAuditData_Method `protobuf_oneof:"method"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AuditData) Reset()         { *m = AuditData{} }
func (m *AuditData) String() string { return proto.CompactTextString(m) }
func (*AuditData) ProtoMessage()    {}
func (*AuditData) Descriptor() ([]byte, []int) {
	return fileDescriptor_audit_data_196ce8036024e2bd, []int{0}
}
func (m *AuditData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuditData.Unmarshal(m, b)
}
func (m *AuditData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuditData.Marshal(b, m, deterministic)
}
func (dst *AuditData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuditData.Merge(dst, src)
}
func (m *AuditData) XXX_Size() int {
	return xxx_messageInfo_AuditData.Size(m)
}
func (m *AuditData) XXX_DiscardUnknown() {
	xxx_messageInfo_AuditData.DiscardUnknown(m)
}

var xxx_messageInfo_AuditData proto.InternalMessageInfo

type isAuditData_Method interface {
	isAuditData_Method()
}

type AuditData_UpdateService struct {
	UpdateService *UpdateServiceMethod `protobuf:"bytes,1,opt,name=update_service,json=updateService,proto3,oneof"`
}
type AuditData_CreateVersion struct {
	CreateVersion *CreateVersionMethod `protobuf:"bytes,2,opt,name=create_version,json=createVersion,proto3,oneof"`
}

func (*AuditData_UpdateService) isAuditData_Method() {}
func (*AuditData_CreateVersion) isAuditData_Method() {}

func (m *AuditData) GetMethod() isAuditData_Method {
	if m != nil {
		return m.Method
	}
	return nil
}

func (m *AuditData) GetUpdateService() *UpdateServiceMethod {
	if x, ok := m.GetMethod().(*AuditData_UpdateService); ok {
		return x.UpdateService
	}
	return nil
}

func (m *AuditData) GetCreateVersion() *CreateVersionMethod {
	if x, ok := m.GetMethod().(*AuditData_CreateVersion); ok {
		return x.CreateVersion
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AuditData) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AuditData_OneofMarshaler, _AuditData_OneofUnmarshaler, _AuditData_OneofSizer, []interface{}{
		(*AuditData_UpdateService)(nil),
		(*AuditData_CreateVersion)(nil),
	}
}

func _AuditData_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AuditData)
	// method
	switch x := m.Method.(type) {
	case *AuditData_UpdateService:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UpdateService); err != nil {
			return err
		}
	case *AuditData_CreateVersion:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CreateVersion); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("AuditData.Method has unexpected type %T", x)
	}
	return nil
}

func _AuditData_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AuditData)
	switch tag {
	case 1: // method.update_service
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UpdateServiceMethod)
		err := b.DecodeMessage(msg)
		m.Method = &AuditData_UpdateService{msg}
		return true, err
	case 2: // method.create_version
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CreateVersionMethod)
		err := b.DecodeMessage(msg)
		m.Method = &AuditData_CreateVersion{msg}
		return true, err
	default:
		return false, nil
	}
}

func _AuditData_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AuditData)
	// method
	switch x := m.Method.(type) {
	case *AuditData_UpdateService:
		s := proto.Size(x.UpdateService)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AuditData_CreateVersion:
		s := proto.Size(x.CreateVersion)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Detailed information about UpdateService call.
type UpdateServiceMethod struct {
	// Update service request.
	Request              *UpdateServiceRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateServiceMethod) Reset()         { *m = UpdateServiceMethod{} }
func (m *UpdateServiceMethod) String() string { return proto.CompactTextString(m) }
func (*UpdateServiceMethod) ProtoMessage()    {}
func (*UpdateServiceMethod) Descriptor() ([]byte, []int) {
	return fileDescriptor_audit_data_196ce8036024e2bd, []int{1}
}
func (m *UpdateServiceMethod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateServiceMethod.Unmarshal(m, b)
}
func (m *UpdateServiceMethod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateServiceMethod.Marshal(b, m, deterministic)
}
func (dst *UpdateServiceMethod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateServiceMethod.Merge(dst, src)
}
func (m *UpdateServiceMethod) XXX_Size() int {
	return xxx_messageInfo_UpdateServiceMethod.Size(m)
}
func (m *UpdateServiceMethod) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateServiceMethod.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateServiceMethod proto.InternalMessageInfo

func (m *UpdateServiceMethod) GetRequest() *UpdateServiceRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

// Detailed information about CreateVersion call.
type CreateVersionMethod struct {
	// Create version request.
	Request              *CreateVersionRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateVersionMethod) Reset()         { *m = CreateVersionMethod{} }
func (m *CreateVersionMethod) String() string { return proto.CompactTextString(m) }
func (*CreateVersionMethod) ProtoMessage()    {}
func (*CreateVersionMethod) Descriptor() ([]byte, []int) {
	return fileDescriptor_audit_data_196ce8036024e2bd, []int{2}
}
func (m *CreateVersionMethod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateVersionMethod.Unmarshal(m, b)
}
func (m *CreateVersionMethod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateVersionMethod.Marshal(b, m, deterministic)
}
func (dst *CreateVersionMethod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateVersionMethod.Merge(dst, src)
}
func (m *CreateVersionMethod) XXX_Size() int {
	return xxx_messageInfo_CreateVersionMethod.Size(m)
}
func (m *CreateVersionMethod) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateVersionMethod.DiscardUnknown(m)
}

var xxx_messageInfo_CreateVersionMethod proto.InternalMessageInfo

func (m *CreateVersionMethod) GetRequest() *CreateVersionRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func init() {
	proto.RegisterType((*AuditData)(nil), "google.appengine.v1.AuditData")
	proto.RegisterType((*UpdateServiceMethod)(nil), "google.appengine.v1.UpdateServiceMethod")
	proto.RegisterType((*CreateVersionMethod)(nil), "google.appengine.v1.CreateVersionMethod")
}

func init() {
	proto.RegisterFile("google/appengine/v1/audit_data.proto", fileDescriptor_audit_data_196ce8036024e2bd)
}

var fileDescriptor_audit_data_196ce8036024e2bd = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x09, 0x43, 0x01, 0x23, 0x3a, 0xa4, 0x03, 0x55, 0x07, 0x84, 0x0a, 0x43, 0x59, 0x1c,
	0x15, 0x46, 0x58, 0x48, 0x19, 0x58, 0x90, 0x4a, 0x10, 0x0c, 0x5d, 0xa2, 0x23, 0x39, 0x19, 0x4b,
	0x49, 0x6c, 0x1c, 0x27, 0x12, 0xcf, 0xc6, 0xcb, 0x21, 0xdb, 0x21, 0xb4, 0xc8, 0x82, 0x8e, 0xb9,
	0xfb, 0xfe, 0x2f, 0xbf, 0x74, 0x26, 0xe7, 0x4c, 0x08, 0x56, 0x60, 0x04, 0x52, 0x62, 0xc5, 0x78,
	0x85, 0x51, 0x3b, 0x8f, 0xa0, 0xc9, 0xb9, 0x4e, 0x73, 0xd0, 0x40, 0xa5, 0x12, 0x5a, 0x84, 0x23,
	0x47, 0xd1, 0x9e, 0xa2, 0xed, 0x7c, 0x72, 0xe6, 0x8d, 0xf6, 0x84, 0x4d, 0x4e, 0x4e, 0x3a, 0x88,
	0x43, 0x69, 0xd6, 0x1c, 0xca, 0x54, 0x8a, 0x82, 0x67, 0x1f, 0x6e, 0x3f, 0xfd, 0x0c, 0xc8, 0xc1,
	0xad, 0xf9, 0xdd, 0x1d, 0x68, 0x08, 0x1f, 0xc9, 0xb0, 0x91, 0x39, 0x68, 0x4c, 0x6b, 0x54, 0x2d,
	0xcf, 0x70, 0x1c, 0x9c, 0x06, 0xb3, 0xc3, 0xcb, 0x19, 0xf5, 0x14, 0xa0, 0xcf, 0x16, 0x7d, 0x72,
	0xe4, 0x03, 0xea, 0x37, 0x91, 0xdf, 0xef, 0x24, 0x47, 0xcd, 0xfa, 0xd8, 0x28, 0x33, 0x85, 0x46,
	0xd9, 0xa2, 0xaa, 0xb9, 0xa8, 0xc6, 0xbb, 0x7f, 0x28, 0x17, 0x16, 0x7d, 0x71, 0xe4, 0x8f, 0x32,
	0x5b, 0x1f, 0xc7, 0xfb, 0x64, 0x50, 0xda, 0xd5, 0x74, 0x45, 0x46, 0x9e, 0x12, 0xe1, 0x82, 0xec,
	0x29, 0x7c, 0x6f, 0xb0, 0xd6, 0x5d, 0xff, 0x8b, 0xff, 0xfb, 0x27, 0x2e, 0x90, 0x7c, 0x27, 0x8d,
	0xdb, 0xd3, 0x66, 0x5b, 0xf7, 0x46, 0xf4, 0xb7, 0x3b, 0xe6, 0xe4, 0x38, 0x13, 0xa5, 0x2f, 0x18,
	0x0f, 0xfb, 0x6b, 0x2c, 0xcd, 0x81, 0x96, 0xc1, 0xea, 0xa6, 0xc3, 0x98, 0x28, 0xa0, 0x62, 0x54,
	0x28, 0x16, 0x31, 0xac, 0xec, 0xf9, 0x22, 0xb7, 0x02, 0xc9, 0xeb, 0x8d, 0x67, 0x70, 0xdd, 0x7f,
	0xbc, 0x0e, 0x2c, 0x78, 0xf5, 0x15, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x56, 0x80, 0x49, 0x69, 0x02,
	0x00, 0x00,
}
