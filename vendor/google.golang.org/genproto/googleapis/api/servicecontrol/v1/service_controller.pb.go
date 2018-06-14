// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/servicecontrol/v1/service_controller.proto

package servicecontrol // import "google.golang.org/genproto/googleapis/api/servicecontrol/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import status "google.golang.org/genproto/googleapis/rpc/status"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request message for the Check method.
type CheckRequest struct {
	// The service name as specified in its service configuration. For example,
	// `"pubsub.googleapis.com"`.
	//
	// See
	// [google.api.Service](https://cloud.google.com/service-management/reference/rpc/google.api#google.api.Service)
	// for the definition of a service name.
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// The operation to be checked.
	Operation *Operation `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	// Specifies which version of service configuration should be used to process
	// the request.
	//
	// If unspecified or no matching version can be found, the
	// latest one will be used.
	ServiceConfigId      string   `protobuf:"bytes,4,opt,name=service_config_id,json=serviceConfigId,proto3" json:"service_config_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckRequest) Reset()         { *m = CheckRequest{} }
func (m *CheckRequest) String() string { return proto.CompactTextString(m) }
func (*CheckRequest) ProtoMessage()    {}
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_controller_525d7997df40b2d7, []int{0}
}
func (m *CheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckRequest.Unmarshal(m, b)
}
func (m *CheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckRequest.Marshal(b, m, deterministic)
}
func (dst *CheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckRequest.Merge(dst, src)
}
func (m *CheckRequest) XXX_Size() int {
	return xxx_messageInfo_CheckRequest.Size(m)
}
func (m *CheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckRequest proto.InternalMessageInfo

func (m *CheckRequest) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *CheckRequest) GetOperation() *Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *CheckRequest) GetServiceConfigId() string {
	if m != nil {
		return m.ServiceConfigId
	}
	return ""
}

// Response message for the Check method.
type CheckResponse struct {
	// The same operation_id value used in the [CheckRequest][google.api.servicecontrol.v1.CheckRequest].
	// Used for logging and diagnostics purposes.
	OperationId string `protobuf:"bytes,1,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	// Indicate the decision of the check.
	//
	// If no check errors are present, the service should process the operation.
	// Otherwise the service should use the list of errors to determine the
	// appropriate action.
	CheckErrors []*CheckError `protobuf:"bytes,2,rep,name=check_errors,json=checkErrors,proto3" json:"check_errors,omitempty"`
	// The actual config id used to process the request.
	ServiceConfigId string `protobuf:"bytes,5,opt,name=service_config_id,json=serviceConfigId,proto3" json:"service_config_id,omitempty"`
	// Feedback data returned from the server during processing a Check request.
	CheckInfo            *CheckResponse_CheckInfo `protobuf:"bytes,6,opt,name=check_info,json=checkInfo,proto3" json:"check_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *CheckResponse) Reset()         { *m = CheckResponse{} }
func (m *CheckResponse) String() string { return proto.CompactTextString(m) }
func (*CheckResponse) ProtoMessage()    {}
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_controller_525d7997df40b2d7, []int{1}
}
func (m *CheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckResponse.Unmarshal(m, b)
}
func (m *CheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckResponse.Marshal(b, m, deterministic)
}
func (dst *CheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckResponse.Merge(dst, src)
}
func (m *CheckResponse) XXX_Size() int {
	return xxx_messageInfo_CheckResponse.Size(m)
}
func (m *CheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckResponse proto.InternalMessageInfo

func (m *CheckResponse) GetOperationId() string {
	if m != nil {
		return m.OperationId
	}
	return ""
}

func (m *CheckResponse) GetCheckErrors() []*CheckError {
	if m != nil {
		return m.CheckErrors
	}
	return nil
}

func (m *CheckResponse) GetServiceConfigId() string {
	if m != nil {
		return m.ServiceConfigId
	}
	return ""
}

func (m *CheckResponse) GetCheckInfo() *CheckResponse_CheckInfo {
	if m != nil {
		return m.CheckInfo
	}
	return nil
}

type CheckResponse_CheckInfo struct {
	// Consumer info of this check.
	ConsumerInfo         *CheckResponse_ConsumerInfo `protobuf:"bytes,2,opt,name=consumer_info,json=consumerInfo,proto3" json:"consumer_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *CheckResponse_CheckInfo) Reset()         { *m = CheckResponse_CheckInfo{} }
func (m *CheckResponse_CheckInfo) String() string { return proto.CompactTextString(m) }
func (*CheckResponse_CheckInfo) ProtoMessage()    {}
func (*CheckResponse_CheckInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_controller_525d7997df40b2d7, []int{1, 0}
}
func (m *CheckResponse_CheckInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckResponse_CheckInfo.Unmarshal(m, b)
}
func (m *CheckResponse_CheckInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckResponse_CheckInfo.Marshal(b, m, deterministic)
}
func (dst *CheckResponse_CheckInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckResponse_CheckInfo.Merge(dst, src)
}
func (m *CheckResponse_CheckInfo) XXX_Size() int {
	return xxx_messageInfo_CheckResponse_CheckInfo.Size(m)
}
func (m *CheckResponse_CheckInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckResponse_CheckInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CheckResponse_CheckInfo proto.InternalMessageInfo

func (m *CheckResponse_CheckInfo) GetConsumerInfo() *CheckResponse_ConsumerInfo {
	if m != nil {
		return m.ConsumerInfo
	}
	return nil
}

// `ConsumerInfo` provides information about the consumer project.
type CheckResponse_ConsumerInfo struct {
	// The Google cloud project number, e.g. 1234567890. A value of 0 indicates
	// no project number is found.
	ProjectNumber        int64    `protobuf:"varint,1,opt,name=project_number,json=projectNumber,proto3" json:"project_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckResponse_ConsumerInfo) Reset()         { *m = CheckResponse_ConsumerInfo{} }
func (m *CheckResponse_ConsumerInfo) String() string { return proto.CompactTextString(m) }
func (*CheckResponse_ConsumerInfo) ProtoMessage()    {}
func (*CheckResponse_ConsumerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_controller_525d7997df40b2d7, []int{1, 1}
}
func (m *CheckResponse_ConsumerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckResponse_ConsumerInfo.Unmarshal(m, b)
}
func (m *CheckResponse_ConsumerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckResponse_ConsumerInfo.Marshal(b, m, deterministic)
}
func (dst *CheckResponse_ConsumerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckResponse_ConsumerInfo.Merge(dst, src)
}
func (m *CheckResponse_ConsumerInfo) XXX_Size() int {
	return xxx_messageInfo_CheckResponse_ConsumerInfo.Size(m)
}
func (m *CheckResponse_ConsumerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckResponse_ConsumerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CheckResponse_ConsumerInfo proto.InternalMessageInfo

func (m *CheckResponse_ConsumerInfo) GetProjectNumber() int64 {
	if m != nil {
		return m.ProjectNumber
	}
	return 0
}

// Request message for the Report method.
type ReportRequest struct {
	// The service name as specified in its service configuration. For example,
	// `"pubsub.googleapis.com"`.
	//
	// See
	// [google.api.Service](https://cloud.google.com/service-management/reference/rpc/google.api#google.api.Service)
	// for the definition of a service name.
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// Operations to be reported.
	//
	// Typically the service should report one operation per request.
	// Putting multiple operations into a single request is allowed, but should
	// be used only when multiple operations are natually available at the time
	// of the report.
	//
	// If multiple operations are in a single request, the total request size
	// should be no larger than 1MB. See [ReportResponse.report_errors][google.api.servicecontrol.v1.ReportResponse.report_errors] for
	// partial failure behavior.
	Operations []*Operation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// Specifies which version of service config should be used to process the
	// request.
	//
	// If unspecified or no matching version can be found, the
	// latest one will be used.
	ServiceConfigId      string   `protobuf:"bytes,3,opt,name=service_config_id,json=serviceConfigId,proto3" json:"service_config_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportRequest) Reset()         { *m = ReportRequest{} }
func (m *ReportRequest) String() string { return proto.CompactTextString(m) }
func (*ReportRequest) ProtoMessage()    {}
func (*ReportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_controller_525d7997df40b2d7, []int{2}
}
func (m *ReportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportRequest.Unmarshal(m, b)
}
func (m *ReportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportRequest.Marshal(b, m, deterministic)
}
func (dst *ReportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportRequest.Merge(dst, src)
}
func (m *ReportRequest) XXX_Size() int {
	return xxx_messageInfo_ReportRequest.Size(m)
}
func (m *ReportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReportRequest proto.InternalMessageInfo

func (m *ReportRequest) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *ReportRequest) GetOperations() []*Operation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *ReportRequest) GetServiceConfigId() string {
	if m != nil {
		return m.ServiceConfigId
	}
	return ""
}

// Response message for the Report method.
type ReportResponse struct {
	// Partial failures, one for each `Operation` in the request that failed
	// processing. There are three possible combinations of the RPC status:
	//
	// 1. The combination of a successful RPC status and an empty `report_errors`
	//    list indicates a complete success where all `Operations` in the
	//    request are processed successfully.
	// 2. The combination of a successful RPC status and a non-empty
	//    `report_errors` list indicates a partial success where some
	//    `Operations` in the request succeeded. Each
	//    `Operation` that failed processing has a corresponding item
	//    in this list.
	// 3. A failed RPC status indicates a general non-deterministic failure.
	//    When this happens, it's impossible to know which of the
	//    'Operations' in the request succeeded or failed.
	ReportErrors []*ReportResponse_ReportError `protobuf:"bytes,1,rep,name=report_errors,json=reportErrors,proto3" json:"report_errors,omitempty"`
	// The actual config id used to process the request.
	ServiceConfigId      string   `protobuf:"bytes,2,opt,name=service_config_id,json=serviceConfigId,proto3" json:"service_config_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportResponse) Reset()         { *m = ReportResponse{} }
func (m *ReportResponse) String() string { return proto.CompactTextString(m) }
func (*ReportResponse) ProtoMessage()    {}
func (*ReportResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_controller_525d7997df40b2d7, []int{3}
}
func (m *ReportResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportResponse.Unmarshal(m, b)
}
func (m *ReportResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportResponse.Marshal(b, m, deterministic)
}
func (dst *ReportResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportResponse.Merge(dst, src)
}
func (m *ReportResponse) XXX_Size() int {
	return xxx_messageInfo_ReportResponse.Size(m)
}
func (m *ReportResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReportResponse proto.InternalMessageInfo

func (m *ReportResponse) GetReportErrors() []*ReportResponse_ReportError {
	if m != nil {
		return m.ReportErrors
	}
	return nil
}

func (m *ReportResponse) GetServiceConfigId() string {
	if m != nil {
		return m.ServiceConfigId
	}
	return ""
}

// Represents the processing error of one [Operation][google.api.servicecontrol.v1.Operation] in the request.
type ReportResponse_ReportError struct {
	// The [Operation.operation_id][google.api.servicecontrol.v1.Operation.operation_id] value from the request.
	OperationId string `protobuf:"bytes,1,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	// Details of the error when processing the [Operation][google.api.servicecontrol.v1.Operation].
	Status               *status.Status `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReportResponse_ReportError) Reset()         { *m = ReportResponse_ReportError{} }
func (m *ReportResponse_ReportError) String() string { return proto.CompactTextString(m) }
func (*ReportResponse_ReportError) ProtoMessage()    {}
func (*ReportResponse_ReportError) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_controller_525d7997df40b2d7, []int{3, 0}
}
func (m *ReportResponse_ReportError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportResponse_ReportError.Unmarshal(m, b)
}
func (m *ReportResponse_ReportError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportResponse_ReportError.Marshal(b, m, deterministic)
}
func (dst *ReportResponse_ReportError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportResponse_ReportError.Merge(dst, src)
}
func (m *ReportResponse_ReportError) XXX_Size() int {
	return xxx_messageInfo_ReportResponse_ReportError.Size(m)
}
func (m *ReportResponse_ReportError) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportResponse_ReportError.DiscardUnknown(m)
}

var xxx_messageInfo_ReportResponse_ReportError proto.InternalMessageInfo

func (m *ReportResponse_ReportError) GetOperationId() string {
	if m != nil {
		return m.OperationId
	}
	return ""
}

func (m *ReportResponse_ReportError) GetStatus() *status.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*CheckRequest)(nil), "google.api.servicecontrol.v1.CheckRequest")
	proto.RegisterType((*CheckResponse)(nil), "google.api.servicecontrol.v1.CheckResponse")
	proto.RegisterType((*CheckResponse_CheckInfo)(nil), "google.api.servicecontrol.v1.CheckResponse.CheckInfo")
	proto.RegisterType((*CheckResponse_ConsumerInfo)(nil), "google.api.servicecontrol.v1.CheckResponse.ConsumerInfo")
	proto.RegisterType((*ReportRequest)(nil), "google.api.servicecontrol.v1.ReportRequest")
	proto.RegisterType((*ReportResponse)(nil), "google.api.servicecontrol.v1.ReportResponse")
	proto.RegisterType((*ReportResponse_ReportError)(nil), "google.api.servicecontrol.v1.ReportResponse.ReportError")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceControllerClient is the client API for ServiceController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceControllerClient interface {
	// Checks an operation with Google Service Control to decide whether
	// the given operation should proceed. It should be called before the
	// operation is executed.
	//
	// If feasible, the client should cache the check results and reuse them for
	// 60 seconds. In case of server errors, the client can rely on the cached
	// results for longer time.
	//
	// NOTE: the [CheckRequest][google.api.servicecontrol.v1.CheckRequest] has the size limit of 64KB.
	//
	// This method requires the `servicemanagement.services.check` permission
	// on the specified service. For more information, see
	// [Google Cloud IAM](https://cloud.google.com/iam).
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
	// Reports operation results to Google Service Control, such as logs and
	// metrics. It should be called after an operation is completed.
	//
	// If feasible, the client should aggregate reporting data for up to 5
	// seconds to reduce API traffic. Limiting aggregation to 5 seconds is to
	// reduce data loss during client crashes. Clients should carefully choose
	// the aggregation time window to avoid data loss risk more than 0.01%
	// for business and compliance reasons.
	//
	// NOTE: the [ReportRequest][google.api.servicecontrol.v1.ReportRequest] has the size limit of 1MB.
	//
	// This method requires the `servicemanagement.services.report` permission
	// on the specified service. For more information, see
	// [Google Cloud IAM](https://cloud.google.com/iam).
	Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error)
}

type serviceControllerClient struct {
	cc *grpc.ClientConn
}

func NewServiceControllerClient(cc *grpc.ClientConn) ServiceControllerClient {
	return &serviceControllerClient{cc}
}

func (c *serviceControllerClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, "/google.api.servicecontrol.v1.ServiceController/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceControllerClient) Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error) {
	out := new(ReportResponse)
	err := c.cc.Invoke(ctx, "/google.api.servicecontrol.v1.ServiceController/Report", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceControllerServer is the server API for ServiceController service.
type ServiceControllerServer interface {
	// Checks an operation with Google Service Control to decide whether
	// the given operation should proceed. It should be called before the
	// operation is executed.
	//
	// If feasible, the client should cache the check results and reuse them for
	// 60 seconds. In case of server errors, the client can rely on the cached
	// results for longer time.
	//
	// NOTE: the [CheckRequest][google.api.servicecontrol.v1.CheckRequest] has the size limit of 64KB.
	//
	// This method requires the `servicemanagement.services.check` permission
	// on the specified service. For more information, see
	// [Google Cloud IAM](https://cloud.google.com/iam).
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
	// Reports operation results to Google Service Control, such as logs and
	// metrics. It should be called after an operation is completed.
	//
	// If feasible, the client should aggregate reporting data for up to 5
	// seconds to reduce API traffic. Limiting aggregation to 5 seconds is to
	// reduce data loss during client crashes. Clients should carefully choose
	// the aggregation time window to avoid data loss risk more than 0.01%
	// for business and compliance reasons.
	//
	// NOTE: the [ReportRequest][google.api.servicecontrol.v1.ReportRequest] has the size limit of 1MB.
	//
	// This method requires the `servicemanagement.services.report` permission
	// on the specified service. For more information, see
	// [Google Cloud IAM](https://cloud.google.com/iam).
	Report(context.Context, *ReportRequest) (*ReportResponse, error)
}

func RegisterServiceControllerServer(s *grpc.Server, srv ServiceControllerServer) {
	s.RegisterService(&_ServiceController_serviceDesc, srv)
}

func _ServiceController_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceControllerServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.api.servicecontrol.v1.ServiceController/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceControllerServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceController_Report_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceControllerServer).Report(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.api.servicecontrol.v1.ServiceController/Report",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceControllerServer).Report(ctx, req.(*ReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceController_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.api.servicecontrol.v1.ServiceController",
	HandlerType: (*ServiceControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _ServiceController_Check_Handler,
		},
		{
			MethodName: "Report",
			Handler:    _ServiceController_Report_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/api/servicecontrol/v1/service_controller.proto",
}

func init() {
	proto.RegisterFile("google/api/servicecontrol/v1/service_controller.proto", fileDescriptor_service_controller_525d7997df40b2d7)
}

var fileDescriptor_service_controller_525d7997df40b2d7 = []byte{
	// 619 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0xd6, 0x3a, 0x6d, 0xa4, 0x4c, 0x9c, 0xfe, 0xea, 0x1e, 0x7e, 0x22, 0xab, 0x87, 0xd4, 0x12,
	0x34, 0x4a, 0x8b, 0xad, 0x16, 0x55, 0x42, 0xe1, 0x44, 0xa3, 0xaa, 0x0a, 0x48, 0xa5, 0x72, 0x38,
	0x21, 0xaa, 0xc8, 0xdd, 0x6c, 0x8c, 0x4b, 0xb2, 0x6b, 0xd6, 0x4e, 0x2e, 0x88, 0x0b, 0x0f, 0xc0,
	0xa1, 0xbc, 0x01, 0xaa, 0xc4, 0x33, 0xf0, 0x1c, 0xbc, 0x02, 0x0f, 0x01, 0x37, 0x94, 0xdd, 0xb5,
	0xeb, 0x08, 0x63, 0x92, 0x9b, 0xf7, 0xdb, 0x99, 0xf9, 0xbe, 0x9d, 0xf9, 0x3c, 0x70, 0x1c, 0x70,
	0x1e, 0x4c, 0xa8, 0xeb, 0x47, 0xa1, 0x1b, 0x53, 0x31, 0x0f, 0x09, 0x25, 0x9c, 0x25, 0x82, 0x4f,
	0xdc, 0xf9, 0x61, 0x8a, 0x0c, 0x35, 0x34, 0xa1, 0xc2, 0x89, 0x04, 0x4f, 0x38, 0xde, 0x51, 0x69,
	0x8e, 0x1f, 0x85, 0xce, 0x72, 0x9a, 0x33, 0x3f, 0xb4, 0x76, 0x72, 0x45, 0x7d, 0xc6, 0x78, 0xe2,
	0x27, 0x21, 0x67, 0xb1, 0xca, 0xb5, 0x9c, 0x52, 0x4a, 0xf2, 0x86, 0x92, 0xb7, 0x43, 0x2a, 0x04,
	0xd7, 0x5c, 0xd6, 0x41, 0x69, 0x3c, 0x8f, 0xa8, 0x90, 0xe5, 0x75, 0xf4, 0x3d, 0x1d, 0x2d, 0x22,
	0xe2, 0xc6, 0x89, 0x9f, 0xcc, 0x34, 0xad, 0x7d, 0x8b, 0xc0, 0xec, 0x2d, 0x8a, 0x7b, 0xf4, 0xdd,
	0x8c, 0xc6, 0x09, 0xde, 0x05, 0x33, 0x7d, 0x1f, 0xf3, 0xa7, 0xb4, 0x89, 0x5a, 0xa8, 0x5d, 0xf3,
	0xea, 0x1a, 0x3b, 0xf7, 0xa7, 0x14, 0x9f, 0x42, 0x2d, 0xab, 0xdf, 0x34, 0x5a, 0xa8, 0x5d, 0x3f,
	0xda, 0x73, 0xca, 0x9e, 0xee, 0xbc, 0x48, 0xc3, 0xbd, 0xbb, 0x4c, 0xdc, 0x81, 0xed, 0x5c, 0x27,
	0xc7, 0x61, 0x30, 0x0c, 0x47, 0xcd, 0x0d, 0x49, 0xf7, 0x9f, 0xbe, 0xe8, 0x49, 0xbc, 0x3f, 0xb2,
	0x6f, 0x2b, 0xd0, 0xd0, 0x32, 0xe3, 0x88, 0xb3, 0x98, 0x2e, 0x74, 0x66, 0xa5, 0x16, 0x89, 0x5a,
	0x67, 0x86, 0xf5, 0x47, 0xf8, 0x39, 0x98, 0xb9, 0xbe, 0xc5, 0x4d, 0xa3, 0x55, 0x69, 0xd7, 0x8f,
	0xda, 0xe5, 0x52, 0x25, 0xcb, 0xe9, 0x22, 0xc1, 0xab, 0x93, 0xec, 0x3b, 0x2e, 0x56, 0xbb, 0x59,
	0xa8, 0x16, 0xbf, 0x04, 0x50, 0xc4, 0x21, 0x1b, 0xf3, 0x66, 0x55, 0x76, 0xe8, 0x78, 0x05, 0xda,
	0xf4, 0x71, 0xea, 0xd4, 0x67, 0x63, 0xee, 0xd5, 0x48, 0xfa, 0x69, 0x5d, 0x43, 0x2d, 0xc3, 0xf1,
	0x25, 0x34, 0x08, 0x67, 0xf1, 0x6c, 0x4a, 0x85, 0x62, 0x51, 0x73, 0x78, 0xbc, 0x16, 0x8b, 0x2e,
	0x20, 0x89, 0x4c, 0x92, 0x3b, 0x59, 0xc7, 0x60, 0xe6, 0x6f, 0xf1, 0x7d, 0xd8, 0x8a, 0x04, 0xbf,
	0xa6, 0x24, 0x19, 0xb2, 0xd9, 0xf4, 0x8a, 0x0a, 0xd9, 0xef, 0x8a, 0xd7, 0xd0, 0xe8, 0xb9, 0x04,
	0xed, 0xaf, 0x08, 0x1a, 0x1e, 0x8d, 0xb8, 0x48, 0xd6, 0xb0, 0xd3, 0x19, 0x40, 0x36, 0xb5, 0x74,
	0x48, 0x2b, 0xfb, 0x29, 0x97, 0x5a, 0x3c, 0xa2, 0x4a, 0xb1, 0xa1, 0x7e, 0x21, 0xd8, 0x4a, 0x95,
	0x6a, 0x47, 0x5d, 0x42, 0x43, 0x48, 0x24, 0xf5, 0x0b, 0x92, 0x52, 0xfe, 0xd1, 0xd2, 0xe5, 0x22,
	0xfa, 0xa8, 0xfc, 0x63, 0x8a, 0xbb, 0xc3, 0x5f, 0xd4, 0x19, 0x85, 0xea, 0xac, 0xd7, 0x50, 0xcf,
	0x15, 0x5a, 0xc5, 0xeb, 0x1d, 0xa8, 0xaa, 0xff, 0x5a, 0x1b, 0x01, 0xa7, 0xaa, 0x45, 0x44, 0x9c,
	0x81, 0xbc, 0xf1, 0x74, 0xc4, 0xd1, 0x37, 0x03, 0xb6, 0x07, 0x19, 0xa3, 0x5e, 0x61, 0xf8, 0x13,
	0x82, 0x4d, 0xe9, 0x0f, 0xdc, 0x59, 0xc9, 0x44, 0x72, 0xbe, 0xd6, 0xfe, 0x1a, 0x86, 0xb3, 0x0f,
	0x3e, 0x7e, 0xff, 0xf1, 0xd9, 0x78, 0x60, 0xef, 0xe6, 0xb6, 0x68, 0xec, 0xbe, 0xcf, 0x1b, 0xe4,
	0x43, 0x57, 0x1a, 0xbe, 0x8b, 0x3a, 0xf8, 0x06, 0x41, 0x55, 0x75, 0x01, 0xef, 0xaf, 0x36, 0x03,
	0x25, 0xe9, 0x60, 0x9d, 0x81, 0xd9, 0x0f, 0xa5, 0xa6, 0x3d, 0xdb, 0x2e, 0xd3, 0xa4, 0x06, 0xd9,
	0x45, 0x9d, 0x93, 0x1b, 0x04, 0x2d, 0xc2, 0xa7, 0xa5, 0x14, 0x27, 0xff, 0xff, 0xd1, 0xdd, 0x8b,
	0xc5, 0xb2, 0xbd, 0x40, 0xaf, 0x9e, 0xe9, 0xbc, 0x80, 0x4f, 0x7c, 0x16, 0x38, 0x5c, 0x04, 0x6e,
	0x40, 0x99, 0x5c, 0xc5, 0xae, 0xba, 0xf2, 0xa3, 0x30, 0x2e, 0x5e, 0xea, 0x4f, 0x96, 0x91, 0x9f,
	0x08, 0x7d, 0x31, 0x36, 0xce, 0x9e, 0x0e, 0x7a, 0x57, 0x55, 0x59, 0xe0, 0xd1, 0xef, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x5e, 0x28, 0x7b, 0xe6, 0xb7, 0x06, 0x00, 0x00,
}
