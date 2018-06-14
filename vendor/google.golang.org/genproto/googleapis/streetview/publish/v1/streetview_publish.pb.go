// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/streetview/publish/v1/streetview_publish.proto

package publish // import "google.golang.org/genproto/googleapis/streetview/publish/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StreetViewPublishServiceClient is the client API for StreetViewPublishService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreetViewPublishServiceClient interface {
	// Creates an upload session to start uploading photo data. The upload URL of
	// the returned `UploadRef` is used to upload the data for the photo.
	//
	// After the upload is complete, the `UploadRef` is used with
	// `StreetViewPublishService:CreatePhoto()` to create the `Photo` object
	// entry.
	StartUpload(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UploadRef, error)
	// After the client finishes uploading the photo with the returned
	// `UploadRef`, `photo.create` publishes the uploaded photo to Street View on
	// Google Maps.
	//
	// This method returns the following error codes:
	//
	// * `INVALID_ARGUMENT` if the request is malformed.
	// * `NOT_FOUND` if the upload reference does not exist.
	CreatePhoto(ctx context.Context, in *CreatePhotoRequest, opts ...grpc.CallOption) (*Photo, error)
	// Gets the metadata of the specified `Photo`.
	//
	// This method returns the following error codes:
	//
	// * `PERMISSION_DENIED` if the requesting user did not create the requested
	// photo.
	// * `NOT_FOUND` if the requested photo does not exist.
	GetPhoto(ctx context.Context, in *GetPhotoRequest, opts ...grpc.CallOption) (*Photo, error)
	// Gets the metadata of the specified `Photo` batch.
	//
	// Note that if `photos.batchGet` fails, either critical fields are
	// missing or there was an authentication error.
	// Even if `photos.batchGet` succeeds, there may have been failures
	// for single photos in the batch. These failures will be specified in
	// `BatchGetPhotosResponse.results.status`.
	// See `photo.get` for specific failures that will occur per photo.
	BatchGetPhotos(ctx context.Context, in *BatchGetPhotosRequest, opts ...grpc.CallOption) (*BatchGetPhotosResponse, error)
	// Lists all the photos that belong to the user.
	ListPhotos(ctx context.Context, in *ListPhotosRequest, opts ...grpc.CallOption) (*ListPhotosResponse, error)
	// Updates the metadata of a photo, such as pose, place association, etc.
	// Changing the pixels of a photo is not supported.
	//
	// This method returns the following error codes:
	//
	// * `PERMISSION_DENIED` if the requesting user did not create the requested
	// photo.
	// * `INVALID_ARGUMENT` if the request is malformed.
	// * `NOT_FOUND` if the photo ID does not exist.
	UpdatePhoto(ctx context.Context, in *UpdatePhotoRequest, opts ...grpc.CallOption) (*Photo, error)
	// Updates the metadata of photos, such as pose, place association, etc.
	// Changing the pixels of a photo is not supported.
	//
	// Note that if `photos.batchUpdate` fails, either critical fields
	// are missing or there was an authentication error.
	// Even if `photos.batchUpdate` succeeds, there may have been
	// failures for single photos in the batch. These failures will be specified
	// in `BatchUpdatePhotosResponse.results.status`.
	// See `UpdatePhoto` for specific failures that will occur per photo.
	BatchUpdatePhotos(ctx context.Context, in *BatchUpdatePhotosRequest, opts ...grpc.CallOption) (*BatchUpdatePhotosResponse, error)
	// Deletes a photo and its metadata.
	//
	// This method returns the following error codes:
	//
	// * `PERMISSION_DENIED` if the requesting user did not create the requested
	// photo.
	// * `NOT_FOUND` if the photo ID does not exist.
	DeletePhoto(ctx context.Context, in *DeletePhotoRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Deletes a list of photos and their metadata.
	//
	// Note that if `photos.batchDelete` fails, either critical fields
	// are missing or there was an authentication error.
	// Even if `photos.batchDelete` succeeds, there may have been
	// failures for single photos in the batch. These failures will be specified
	// in `BatchDeletePhotosResponse.status`.
	// See `photo.update` for specific failures that will occur per photo.
	BatchDeletePhotos(ctx context.Context, in *BatchDeletePhotosRequest, opts ...grpc.CallOption) (*BatchDeletePhotosResponse, error)
}

type streetViewPublishServiceClient struct {
	cc *grpc.ClientConn
}

func NewStreetViewPublishServiceClient(cc *grpc.ClientConn) StreetViewPublishServiceClient {
	return &streetViewPublishServiceClient{cc}
}

func (c *streetViewPublishServiceClient) StartUpload(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UploadRef, error) {
	out := new(UploadRef)
	err := c.cc.Invoke(ctx, "/google.streetview.publish.v1.StreetViewPublishService/StartUpload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streetViewPublishServiceClient) CreatePhoto(ctx context.Context, in *CreatePhotoRequest, opts ...grpc.CallOption) (*Photo, error) {
	out := new(Photo)
	err := c.cc.Invoke(ctx, "/google.streetview.publish.v1.StreetViewPublishService/CreatePhoto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streetViewPublishServiceClient) GetPhoto(ctx context.Context, in *GetPhotoRequest, opts ...grpc.CallOption) (*Photo, error) {
	out := new(Photo)
	err := c.cc.Invoke(ctx, "/google.streetview.publish.v1.StreetViewPublishService/GetPhoto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streetViewPublishServiceClient) BatchGetPhotos(ctx context.Context, in *BatchGetPhotosRequest, opts ...grpc.CallOption) (*BatchGetPhotosResponse, error) {
	out := new(BatchGetPhotosResponse)
	err := c.cc.Invoke(ctx, "/google.streetview.publish.v1.StreetViewPublishService/BatchGetPhotos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streetViewPublishServiceClient) ListPhotos(ctx context.Context, in *ListPhotosRequest, opts ...grpc.CallOption) (*ListPhotosResponse, error) {
	out := new(ListPhotosResponse)
	err := c.cc.Invoke(ctx, "/google.streetview.publish.v1.StreetViewPublishService/ListPhotos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streetViewPublishServiceClient) UpdatePhoto(ctx context.Context, in *UpdatePhotoRequest, opts ...grpc.CallOption) (*Photo, error) {
	out := new(Photo)
	err := c.cc.Invoke(ctx, "/google.streetview.publish.v1.StreetViewPublishService/UpdatePhoto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streetViewPublishServiceClient) BatchUpdatePhotos(ctx context.Context, in *BatchUpdatePhotosRequest, opts ...grpc.CallOption) (*BatchUpdatePhotosResponse, error) {
	out := new(BatchUpdatePhotosResponse)
	err := c.cc.Invoke(ctx, "/google.streetview.publish.v1.StreetViewPublishService/BatchUpdatePhotos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streetViewPublishServiceClient) DeletePhoto(ctx context.Context, in *DeletePhotoRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.streetview.publish.v1.StreetViewPublishService/DeletePhoto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streetViewPublishServiceClient) BatchDeletePhotos(ctx context.Context, in *BatchDeletePhotosRequest, opts ...grpc.CallOption) (*BatchDeletePhotosResponse, error) {
	out := new(BatchDeletePhotosResponse)
	err := c.cc.Invoke(ctx, "/google.streetview.publish.v1.StreetViewPublishService/BatchDeletePhotos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StreetViewPublishServiceServer is the server API for StreetViewPublishService service.
type StreetViewPublishServiceServer interface {
	// Creates an upload session to start uploading photo data. The upload URL of
	// the returned `UploadRef` is used to upload the data for the photo.
	//
	// After the upload is complete, the `UploadRef` is used with
	// `StreetViewPublishService:CreatePhoto()` to create the `Photo` object
	// entry.
	StartUpload(context.Context, *empty.Empty) (*UploadRef, error)
	// After the client finishes uploading the photo with the returned
	// `UploadRef`, `photo.create` publishes the uploaded photo to Street View on
	// Google Maps.
	//
	// This method returns the following error codes:
	//
	// * `INVALID_ARGUMENT` if the request is malformed.
	// * `NOT_FOUND` if the upload reference does not exist.
	CreatePhoto(context.Context, *CreatePhotoRequest) (*Photo, error)
	// Gets the metadata of the specified `Photo`.
	//
	// This method returns the following error codes:
	//
	// * `PERMISSION_DENIED` if the requesting user did not create the requested
	// photo.
	// * `NOT_FOUND` if the requested photo does not exist.
	GetPhoto(context.Context, *GetPhotoRequest) (*Photo, error)
	// Gets the metadata of the specified `Photo` batch.
	//
	// Note that if `photos.batchGet` fails, either critical fields are
	// missing or there was an authentication error.
	// Even if `photos.batchGet` succeeds, there may have been failures
	// for single photos in the batch. These failures will be specified in
	// `BatchGetPhotosResponse.results.status`.
	// See `photo.get` for specific failures that will occur per photo.
	BatchGetPhotos(context.Context, *BatchGetPhotosRequest) (*BatchGetPhotosResponse, error)
	// Lists all the photos that belong to the user.
	ListPhotos(context.Context, *ListPhotosRequest) (*ListPhotosResponse, error)
	// Updates the metadata of a photo, such as pose, place association, etc.
	// Changing the pixels of a photo is not supported.
	//
	// This method returns the following error codes:
	//
	// * `PERMISSION_DENIED` if the requesting user did not create the requested
	// photo.
	// * `INVALID_ARGUMENT` if the request is malformed.
	// * `NOT_FOUND` if the photo ID does not exist.
	UpdatePhoto(context.Context, *UpdatePhotoRequest) (*Photo, error)
	// Updates the metadata of photos, such as pose, place association, etc.
	// Changing the pixels of a photo is not supported.
	//
	// Note that if `photos.batchUpdate` fails, either critical fields
	// are missing or there was an authentication error.
	// Even if `photos.batchUpdate` succeeds, there may have been
	// failures for single photos in the batch. These failures will be specified
	// in `BatchUpdatePhotosResponse.results.status`.
	// See `UpdatePhoto` for specific failures that will occur per photo.
	BatchUpdatePhotos(context.Context, *BatchUpdatePhotosRequest) (*BatchUpdatePhotosResponse, error)
	// Deletes a photo and its metadata.
	//
	// This method returns the following error codes:
	//
	// * `PERMISSION_DENIED` if the requesting user did not create the requested
	// photo.
	// * `NOT_FOUND` if the photo ID does not exist.
	DeletePhoto(context.Context, *DeletePhotoRequest) (*empty.Empty, error)
	// Deletes a list of photos and their metadata.
	//
	// Note that if `photos.batchDelete` fails, either critical fields
	// are missing or there was an authentication error.
	// Even if `photos.batchDelete` succeeds, there may have been
	// failures for single photos in the batch. These failures will be specified
	// in `BatchDeletePhotosResponse.status`.
	// See `photo.update` for specific failures that will occur per photo.
	BatchDeletePhotos(context.Context, *BatchDeletePhotosRequest) (*BatchDeletePhotosResponse, error)
}

func RegisterStreetViewPublishServiceServer(s *grpc.Server, srv StreetViewPublishServiceServer) {
	s.RegisterService(&_StreetViewPublishService_serviceDesc, srv)
}

func _StreetViewPublishService_StartUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreetViewPublishServiceServer).StartUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.streetview.publish.v1.StreetViewPublishService/StartUpload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreetViewPublishServiceServer).StartUpload(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreetViewPublishService_CreatePhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreetViewPublishServiceServer).CreatePhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.streetview.publish.v1.StreetViewPublishService/CreatePhoto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreetViewPublishServiceServer).CreatePhoto(ctx, req.(*CreatePhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreetViewPublishService_GetPhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreetViewPublishServiceServer).GetPhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.streetview.publish.v1.StreetViewPublishService/GetPhoto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreetViewPublishServiceServer).GetPhoto(ctx, req.(*GetPhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreetViewPublishService_BatchGetPhotos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetPhotosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreetViewPublishServiceServer).BatchGetPhotos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.streetview.publish.v1.StreetViewPublishService/BatchGetPhotos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreetViewPublishServiceServer).BatchGetPhotos(ctx, req.(*BatchGetPhotosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreetViewPublishService_ListPhotos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPhotosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreetViewPublishServiceServer).ListPhotos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.streetview.publish.v1.StreetViewPublishService/ListPhotos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreetViewPublishServiceServer).ListPhotos(ctx, req.(*ListPhotosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreetViewPublishService_UpdatePhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreetViewPublishServiceServer).UpdatePhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.streetview.publish.v1.StreetViewPublishService/UpdatePhoto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreetViewPublishServiceServer).UpdatePhoto(ctx, req.(*UpdatePhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreetViewPublishService_BatchUpdatePhotos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchUpdatePhotosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreetViewPublishServiceServer).BatchUpdatePhotos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.streetview.publish.v1.StreetViewPublishService/BatchUpdatePhotos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreetViewPublishServiceServer).BatchUpdatePhotos(ctx, req.(*BatchUpdatePhotosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreetViewPublishService_DeletePhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreetViewPublishServiceServer).DeletePhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.streetview.publish.v1.StreetViewPublishService/DeletePhoto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreetViewPublishServiceServer).DeletePhoto(ctx, req.(*DeletePhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreetViewPublishService_BatchDeletePhotos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchDeletePhotosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreetViewPublishServiceServer).BatchDeletePhotos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.streetview.publish.v1.StreetViewPublishService/BatchDeletePhotos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreetViewPublishServiceServer).BatchDeletePhotos(ctx, req.(*BatchDeletePhotosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StreetViewPublishService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.streetview.publish.v1.StreetViewPublishService",
	HandlerType: (*StreetViewPublishServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartUpload",
			Handler:    _StreetViewPublishService_StartUpload_Handler,
		},
		{
			MethodName: "CreatePhoto",
			Handler:    _StreetViewPublishService_CreatePhoto_Handler,
		},
		{
			MethodName: "GetPhoto",
			Handler:    _StreetViewPublishService_GetPhoto_Handler,
		},
		{
			MethodName: "BatchGetPhotos",
			Handler:    _StreetViewPublishService_BatchGetPhotos_Handler,
		},
		{
			MethodName: "ListPhotos",
			Handler:    _StreetViewPublishService_ListPhotos_Handler,
		},
		{
			MethodName: "UpdatePhoto",
			Handler:    _StreetViewPublishService_UpdatePhoto_Handler,
		},
		{
			MethodName: "BatchUpdatePhotos",
			Handler:    _StreetViewPublishService_BatchUpdatePhotos_Handler,
		},
		{
			MethodName: "DeletePhoto",
			Handler:    _StreetViewPublishService_DeletePhoto_Handler,
		},
		{
			MethodName: "BatchDeletePhotos",
			Handler:    _StreetViewPublishService_BatchDeletePhotos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/streetview/publish/v1/streetview_publish.proto",
}

func init() {
	proto.RegisterFile("google/streetview/publish/v1/streetview_publish.proto", fileDescriptor_streetview_publish_c124bcab571c3e8a)
}

var fileDescriptor_streetview_publish_c124bcab571c3e8a = []byte{
	// 533 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x15, 0x24, 0x10, 0xb8, 0x08, 0x69, 0x86, 0x55, 0x53, 0x3a, 0x24, 0x08, 0x12, 0xa0,
	0x6a, 0xd8, 0x1b, 0xe3, 0x8f, 0x54, 0x6e, 0x1d, 0x88, 0x0b, 0x87, 0x69, 0xd5, 0x38, 0x70, 0x99,
	0xdc, 0xf4, 0x5d, 0x6a, 0x29, 0x8d, 0x4d, 0xec, 0x74, 0x42, 0x30, 0x0e, 0xe3, 0xc8, 0x0d, 0x2e,
	0x7c, 0x03, 0x3e, 0x10, 0x5f, 0x81, 0x0f, 0x82, 0xea, 0xd8, 0x4d, 0x36, 0x8a, 0x49, 0x4e, 0x69,
	0xf3, 0x3e, 0xcf, 0xfb, 0xfc, 0xfa, 0xbe, 0xae, 0xd1, 0xd3, 0x44, 0x88, 0x24, 0x05, 0xaa, 0x74,
	0x0e, 0xa0, 0xe7, 0x1c, 0x4e, 0xa8, 0x2c, 0xc6, 0x29, 0x57, 0x53, 0x3a, 0xdf, 0xa9, 0xbd, 0x3d,
	0xb2, 0x6f, 0x89, 0xcc, 0x85, 0x16, 0x78, 0xb3, 0xb4, 0x91, 0x4a, 0x40, 0x9c, 0x60, 0xbe, 0x13,
	0xda, 0x2a, 0x65, 0x92, 0x53, 0x96, 0x65, 0x42, 0x33, 0xcd, 0x45, 0xa6, 0x4a, 0x6f, 0xd8, 0xb3,
	0x55, 0xf3, 0x6d, 0x5c, 0x1c, 0x53, 0x98, 0x49, 0xfd, 0xc1, 0x16, 0xb7, 0xbc, 0x3c, 0x39, 0x28,
	0x51, 0xe4, 0x31, 0xb8, 0x56, 0xc4, 0xaf, 0x96, 0xf1, 0x0c, 0x94, 0x62, 0x89, 0xd3, 0x3f, 0xfe,
	0x8a, 0xd0, 0xc6, 0xc8, 0x68, 0xdf, 0x72, 0x38, 0xd9, 0x2f, 0xa5, 0x23, 0xc8, 0xe7, 0x3c, 0x06,
	0x2c, 0x51, 0x67, 0xa4, 0x59, 0xae, 0x0f, 0x65, 0x2a, 0xd8, 0x04, 0x77, 0x6d, 0x73, 0xe2, 0x38,
	0xc9, 0xab, 0x05, 0x67, 0xf8, 0x80, 0xf8, 0x7e, 0x3b, 0x29, 0xdd, 0x07, 0x70, 0x1c, 0xdd, 0x39,
	0xfb, 0xf5, 0xfb, 0xfb, 0xa5, 0x30, 0x5a, 0x5f, 0xb0, 0xc8, 0xa9, 0xd0, 0x62, 0xa0, 0xaa, 0xfe,
	0x83, 0xa0, 0x8f, 0x3f, 0xa3, 0xce, 0x5e, 0x0e, 0x4c, 0xc3, 0xfe, 0xa2, 0x8a, 0xb7, 0xfd, 0x9d,
	0x6b, 0xd2, 0x03, 0x78, 0x5f, 0x80, 0xd2, 0xe1, 0x3d, 0xbf, 0xc3, 0x68, 0xa3, 0x0d, 0xc3, 0x81,
	0xa3, 0x6b, 0x15, 0xc7, 0x65, 0xf3, 0xc0, 0x9f, 0xd0, 0xd5, 0xd7, 0xa0, 0xcb, 0xf0, 0x47, 0xfe,
	0x56, 0x4e, 0xd7, 0x2a, 0x79, 0xd3, 0x24, 0x77, 0xf1, 0xad, 0x65, 0x32, 0xfd, 0x68, 0x1e, 0x47,
	0x7c, 0x72, 0x8a, 0x7f, 0x04, 0xe8, 0xc6, 0x90, 0xe9, 0x78, 0xea, 0x7a, 0x2b, 0xbc, 0xeb, 0xef,
	0x7a, 0x5e, 0xed, 0x50, 0x9e, 0xb4, 0x33, 0x29, 0x29, 0x32, 0x05, 0x51, 0xcf, 0xb0, 0xad, 0xe3,
	0x9b, 0x4b, 0x36, 0x35, 0x18, 0x5b, 0x29, 0xfe, 0x12, 0x20, 0xf4, 0x86, 0x2b, 0x87, 0x45, 0xfd,
	0x09, 0x95, 0xd2, 0x21, 0x6d, 0x37, 0x37, 0x58, 0x1c, 0x6c, 0x70, 0xae, 0x63, 0x54, 0xe1, 0xe0,
	0x6f, 0x01, 0xea, 0x1c, 0xca, 0x49, 0xd3, 0xf3, 0x51, 0x93, 0xb6, 0xda, 0xd2, 0x96, 0x89, 0xbe,
	0x1f, 0xde, 0xbe, 0xb8, 0x25, 0xe2, 0x76, 0x45, 0xf8, 0xe4, 0xd4, 0x9d, 0x99, 0x9f, 0x01, 0x5a,
	0x33, 0x23, 0xad, 0xc5, 0x29, 0xfc, 0xac, 0xc1, 0x0e, 0xea, 0x06, 0x07, 0xf8, 0xbc, 0xb5, 0xcf,
	0xce, 0xeb, 0xae, 0x81, 0xee, 0x45, 0xdd, 0x8b, 0xeb, 0x2b, 0xd5, 0x8b, 0x7f, 0x57, 0x81, 0x3a,
	0x2f, 0x21, 0x85, 0x86, 0xd3, 0xab, 0x49, 0x1d, 0xdc, 0x3f, 0x6e, 0x00, 0x77, 0xac, 0xfb, 0xab,
	0x8f, 0xf5, 0x72, 0x40, 0xb5, 0x8e, 0xcd, 0x06, 0x54, 0x37, 0xb4, 0x19, 0xd0, 0x79, 0xdf, 0xff,
	0x06, 0x54, 0xaa, 0x07, 0x41, 0x7f, 0x78, 0x16, 0xa0, 0x87, 0xb1, 0x98, 0xb9, 0x84, 0x04, 0x04,
	0x29, 0x92, 0x78, 0x75, 0xd2, 0x70, 0xed, 0xaf, 0x7b, 0xf3, 0xdd, 0x9e, 0x33, 0x8a, 0x94, 0x65,
	0x09, 0x11, 0x79, 0x42, 0x13, 0xc8, 0xcc, 0xb0, 0x68, 0x59, 0x62, 0x92, 0xab, 0xd5, 0x97, 0xf3,
	0x0b, 0xfb, 0x71, 0x7c, 0xc5, 0xe8, 0x77, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x9d, 0xfe,
	0x1c, 0x89, 0x06, 0x00, 0x00,
}
