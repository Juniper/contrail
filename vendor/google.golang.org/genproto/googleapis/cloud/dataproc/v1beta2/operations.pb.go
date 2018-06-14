// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/dataproc/v1beta2/operations.proto

package dataproc // import "google.golang.org/genproto/googleapis/cloud/dataproc/v1beta2"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// The operation state.
type ClusterOperationStatus_State int32

const (
	// Unused.
	ClusterOperationStatus_UNKNOWN ClusterOperationStatus_State = 0
	// The operation has been created.
	ClusterOperationStatus_PENDING ClusterOperationStatus_State = 1
	// The operation is running.
	ClusterOperationStatus_RUNNING ClusterOperationStatus_State = 2
	// The operation is done; either cancelled or completed.
	ClusterOperationStatus_DONE ClusterOperationStatus_State = 3
)

var ClusterOperationStatus_State_name = map[int32]string{
	0: "UNKNOWN",
	1: "PENDING",
	2: "RUNNING",
	3: "DONE",
}
var ClusterOperationStatus_State_value = map[string]int32{
	"UNKNOWN": 0,
	"PENDING": 1,
	"RUNNING": 2,
	"DONE":    3,
}

func (x ClusterOperationStatus_State) String() string {
	return proto.EnumName(ClusterOperationStatus_State_name, int32(x))
}
func (ClusterOperationStatus_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_operations_54f6cccb1783f6c3, []int{0, 0}
}

// The status of the operation.
type ClusterOperationStatus struct {
	// Output only. A message containing the operation state.
	State ClusterOperationStatus_State `protobuf:"varint,1,opt,name=state,proto3,enum=google.cloud.dataproc.v1beta2.ClusterOperationStatus_State" json:"state,omitempty"`
	// Output only. A message containing the detailed operation state.
	InnerState string `protobuf:"bytes,2,opt,name=inner_state,json=innerState,proto3" json:"inner_state,omitempty"`
	// Output only. A message containing any operation metadata details.
	Details string `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`
	// Output only. The time this state was entered.
	StateStartTime       *timestamp.Timestamp `protobuf:"bytes,4,opt,name=state_start_time,json=stateStartTime,proto3" json:"state_start_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ClusterOperationStatus) Reset()         { *m = ClusterOperationStatus{} }
func (m *ClusterOperationStatus) String() string { return proto.CompactTextString(m) }
func (*ClusterOperationStatus) ProtoMessage()    {}
func (*ClusterOperationStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_operations_54f6cccb1783f6c3, []int{0}
}
func (m *ClusterOperationStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterOperationStatus.Unmarshal(m, b)
}
func (m *ClusterOperationStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterOperationStatus.Marshal(b, m, deterministic)
}
func (dst *ClusterOperationStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterOperationStatus.Merge(dst, src)
}
func (m *ClusterOperationStatus) XXX_Size() int {
	return xxx_messageInfo_ClusterOperationStatus.Size(m)
}
func (m *ClusterOperationStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterOperationStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterOperationStatus proto.InternalMessageInfo

func (m *ClusterOperationStatus) GetState() ClusterOperationStatus_State {
	if m != nil {
		return m.State
	}
	return ClusterOperationStatus_UNKNOWN
}

func (m *ClusterOperationStatus) GetInnerState() string {
	if m != nil {
		return m.InnerState
	}
	return ""
}

func (m *ClusterOperationStatus) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

func (m *ClusterOperationStatus) GetStateStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StateStartTime
	}
	return nil
}

// Metadata describing the operation.
type ClusterOperationMetadata struct {
	// Output only. Name of the cluster for the operation.
	ClusterName string `protobuf:"bytes,7,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	// Output only. Cluster UUID for the operation.
	ClusterUuid string `protobuf:"bytes,8,opt,name=cluster_uuid,json=clusterUuid,proto3" json:"cluster_uuid,omitempty"`
	// Output only. Current operation status.
	Status *ClusterOperationStatus `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	// Output only. The previous operation status.
	StatusHistory []*ClusterOperationStatus `protobuf:"bytes,10,rep,name=status_history,json=statusHistory,proto3" json:"status_history,omitempty"`
	// Output only. The operation type.
	OperationType string `protobuf:"bytes,11,opt,name=operation_type,json=operationType,proto3" json:"operation_type,omitempty"`
	// Output only. Short description of operation.
	Description string `protobuf:"bytes,12,opt,name=description,proto3" json:"description,omitempty"`
	// Output only. Labels associated with the operation
	Labels map[string]string `protobuf:"bytes,13,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Output only. Errors encountered during operation execution.
	Warnings             []string `protobuf:"bytes,14,rep,name=warnings,proto3" json:"warnings,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterOperationMetadata) Reset()         { *m = ClusterOperationMetadata{} }
func (m *ClusterOperationMetadata) String() string { return proto.CompactTextString(m) }
func (*ClusterOperationMetadata) ProtoMessage()    {}
func (*ClusterOperationMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_operations_54f6cccb1783f6c3, []int{1}
}
func (m *ClusterOperationMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterOperationMetadata.Unmarshal(m, b)
}
func (m *ClusterOperationMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterOperationMetadata.Marshal(b, m, deterministic)
}
func (dst *ClusterOperationMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterOperationMetadata.Merge(dst, src)
}
func (m *ClusterOperationMetadata) XXX_Size() int {
	return xxx_messageInfo_ClusterOperationMetadata.Size(m)
}
func (m *ClusterOperationMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterOperationMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterOperationMetadata proto.InternalMessageInfo

func (m *ClusterOperationMetadata) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ClusterOperationMetadata) GetClusterUuid() string {
	if m != nil {
		return m.ClusterUuid
	}
	return ""
}

func (m *ClusterOperationMetadata) GetStatus() *ClusterOperationStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ClusterOperationMetadata) GetStatusHistory() []*ClusterOperationStatus {
	if m != nil {
		return m.StatusHistory
	}
	return nil
}

func (m *ClusterOperationMetadata) GetOperationType() string {
	if m != nil {
		return m.OperationType
	}
	return ""
}

func (m *ClusterOperationMetadata) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ClusterOperationMetadata) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ClusterOperationMetadata) GetWarnings() []string {
	if m != nil {
		return m.Warnings
	}
	return nil
}

func init() {
	proto.RegisterType((*ClusterOperationStatus)(nil), "google.cloud.dataproc.v1beta2.ClusterOperationStatus")
	proto.RegisterType((*ClusterOperationMetadata)(nil), "google.cloud.dataproc.v1beta2.ClusterOperationMetadata")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.dataproc.v1beta2.ClusterOperationMetadata.LabelsEntry")
	proto.RegisterEnum("google.cloud.dataproc.v1beta2.ClusterOperationStatus_State", ClusterOperationStatus_State_name, ClusterOperationStatus_State_value)
}

func init() {
	proto.RegisterFile("google/cloud/dataproc/v1beta2/operations.proto", fileDescriptor_operations_54f6cccb1783f6c3)
}

var fileDescriptor_operations_54f6cccb1783f6c3 = []byte{
	// 537 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x5f, 0x8b, 0xd3, 0x4e,
	0x14, 0xfd, 0xa5, 0xff, 0x7b, 0xb3, 0xed, 0xaf, 0x0c, 0x22, 0x43, 0x51, 0x36, 0x5b, 0x10, 0xfa,
	0x94, 0x60, 0x45, 0x58, 0xdd, 0xb7, 0xdd, 0x16, 0x15, 0xdd, 0xb4, 0x66, 0xb7, 0x08, 0x2a, 0x94,
	0x69, 0x33, 0xc6, 0x60, 0x3a, 0x13, 0x66, 0x26, 0x2b, 0x7d, 0xf0, 0x0b, 0xf8, 0x41, 0xfc, 0x9c,
	0x32, 0x33, 0x49, 0xa9, 0xa2, 0x0b, 0xee, 0x53, 0xe7, 0xde, 0x7b, 0xce, 0x99, 0x73, 0xa6, 0x97,
	0x80, 0x9f, 0x70, 0x9e, 0x64, 0x34, 0xd8, 0x64, 0xbc, 0x88, 0x83, 0x98, 0x28, 0x92, 0x0b, 0xbe,
	0x09, 0x6e, 0x1e, 0xaf, 0xa9, 0x22, 0x93, 0x80, 0xe7, 0x54, 0x10, 0x95, 0x72, 0x26, 0xfd, 0x5c,
	0x70, 0xc5, 0xd1, 0x43, 0x8b, 0xf7, 0x0d, 0xde, 0xaf, 0xf0, 0x7e, 0x89, 0x1f, 0x3e, 0x28, 0xe5,
	0x48, 0x9e, 0x06, 0x84, 0x31, 0xae, 0x0e, 0xc9, 0xc3, 0xe3, 0x72, 0x6a, 0xaa, 0x75, 0xf1, 0x29,
	0x50, 0xe9, 0x96, 0x4a, 0x45, 0xb6, 0xb9, 0x05, 0x8c, 0x7e, 0xd4, 0xe0, 0xfe, 0x45, 0x56, 0x48,
	0x45, 0xc5, 0xbc, 0xba, 0xf9, 0x4a, 0x11, 0x55, 0x48, 0xf4, 0x16, 0x9a, 0x52, 0x11, 0x45, 0xb1,
	0xe3, 0x39, 0xe3, 0xfe, 0xe4, 0xcc, 0xbf, 0xd5, 0x88, 0xff, 0x67, 0x15, 0x5f, 0xff, 0xd0, 0xc8,
	0x2a, 0xa1, 0x63, 0x70, 0x53, 0xc6, 0xa8, 0x58, 0x59, 0xe1, 0x9a, 0xe7, 0x8c, 0xbb, 0x11, 0x98,
	0x96, 0xc1, 0x21, 0x0c, 0xed, 0x98, 0x2a, 0x92, 0x66, 0x12, 0xd7, 0xcd, 0xb0, 0x2a, 0xd1, 0x14,
	0x06, 0x86, 0xa4, 0xa9, 0x42, 0xad, 0x74, 0x0e, 0xdc, 0xf0, 0x9c, 0xb1, 0x3b, 0x19, 0x56, 0xc6,
	0xaa, 0x90, 0xfe, 0x75, 0x15, 0x32, 0xea, 0x1b, 0xce, 0x95, 0xa6, 0xe8, 0xe6, 0xe8, 0x14, 0x9a,
	0xf6, 0x22, 0x17, 0xda, 0xcb, 0xf0, 0x75, 0x38, 0x7f, 0x17, 0x0e, 0xfe, 0xd3, 0xc5, 0x62, 0x16,
	0x4e, 0x5f, 0x85, 0x2f, 0x06, 0x8e, 0x2e, 0xa2, 0x65, 0x18, 0xea, 0xa2, 0x86, 0x3a, 0xd0, 0x98,
	0xce, 0xc3, 0xd9, 0xa0, 0x3e, 0xfa, 0xde, 0x00, 0xfc, 0x7b, 0xc4, 0x4b, 0xaa, 0x88, 0x7e, 0x07,
	0x74, 0x02, 0x47, 0x1b, 0x3b, 0x5b, 0x31, 0xb2, 0xa5, 0xb8, 0x6d, 0xbc, 0xbb, 0x65, 0x2f, 0x24,
	0x5b, 0x7a, 0x08, 0x29, 0x8a, 0x34, 0xc6, 0x9d, 0x5f, 0x20, 0xcb, 0x22, 0x8d, 0xd1, 0x25, 0xb4,
	0xa4, 0x79, 0x34, 0xdc, 0x35, 0xc1, 0x9e, 0xde, 0xe9, 0xc5, 0xa3, 0x52, 0x04, 0x7d, 0x84, 0xbe,
	0x3d, 0xad, 0x3e, 0xa7, 0x52, 0x71, 0xb1, 0xc3, 0xe0, 0xd5, 0xef, 0x2e, 0xdb, 0xb3, 0x62, 0x2f,
	0xad, 0x16, 0x7a, 0x04, 0xfd, 0xfd, 0xaa, 0xae, 0xd4, 0x2e, 0xa7, 0xd8, 0x35, 0x89, 0x7a, 0xfb,
	0xee, 0xf5, 0x2e, 0xa7, 0xc8, 0x03, 0x37, 0xa6, 0x72, 0x23, 0xd2, 0x5c, 0xb7, 0xf0, 0x91, 0x4d,
	0x7d, 0xd0, 0x42, 0x1f, 0xa0, 0x95, 0x91, 0x35, 0xcd, 0x24, 0xee, 0x19, 0x7b, 0x17, 0xff, 0x68,
	0xaf, 0xfa, 0x13, 0xfc, 0x37, 0x46, 0x65, 0xc6, 0x94, 0xd8, 0x45, 0xa5, 0x24, 0x1a, 0x42, 0xe7,
	0x2b, 0x11, 0x2c, 0x65, 0x89, 0xc4, 0x7d, 0xaf, 0x3e, 0xee, 0x46, 0xfb, 0x7a, 0xf8, 0x0c, 0xdc,
	0x03, 0x0a, 0x1a, 0x40, 0xfd, 0x0b, 0xdd, 0x99, 0x65, 0xef, 0x46, 0xfa, 0x88, 0xee, 0x41, 0xf3,
	0x86, 0x64, 0x45, 0xb5, 0xa7, 0xb6, 0x78, 0x5e, 0x3b, 0x75, 0xce, 0xbf, 0xc1, 0xc9, 0x86, 0x6f,
	0x6f, 0x37, 0x7a, 0xfe, 0xff, 0xde, 0xa2, 0x5c, 0xe8, 0xcd, 0x5c, 0x38, 0xef, 0x67, 0x25, 0x23,
	0xe1, 0x19, 0x61, 0x89, 0xcf, 0x45, 0x12, 0x24, 0x94, 0x99, 0xbd, 0x0d, 0xec, 0x88, 0xe4, 0xa9,
	0xfc, 0xcb, 0xa7, 0xe1, 0xac, 0x6a, 0xac, 0x5b, 0x86, 0xf1, 0xe4, 0x67, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x83, 0x10, 0x95, 0x5e, 0x4b, 0x04, 0x00, 0x00,
}
