// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: booking.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	BookingService_CreateBooking_FullMethodName         = "/booking.BookingService/CreateBooking"
	BookingService_GetBookingHistory_FullMethodName     = "/booking.BookingService/GetBookingHistory"
	BookingService_CancelBooking_FullMethodName         = "/booking.BookingService/CancelBooking"
	BookingService_CheckRoomAvailability_FullMethodName = "/booking.BookingService/CheckRoomAvailability"
)

// BookingServiceClient is the client API for BookingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookingServiceClient interface {
	CreateBooking(ctx context.Context, in *CreateBookingRequest, opts ...grpc.CallOption) (*CreateBookingResponse, error)
	GetBookingHistory(ctx context.Context, in *GetBookingHistoryRequest, opts ...grpc.CallOption) (*GetBookingHistoryResponse, error)
	CancelBooking(ctx context.Context, in *CancelBookingRequest, opts ...grpc.CallOption) (*CancelBookingResponse, error)
	CheckRoomAvailability(ctx context.Context, in *CheckRoomAvailabilityRequest, opts ...grpc.CallOption) (*CheckRoomAvailabilityResponse, error)
}

type bookingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookingServiceClient(cc grpc.ClientConnInterface) BookingServiceClient {
	return &bookingServiceClient{cc}
}

func (c *bookingServiceClient) CreateBooking(ctx context.Context, in *CreateBookingRequest, opts ...grpc.CallOption) (*CreateBookingResponse, error) {
	out := new(CreateBookingResponse)
	err := c.cc.Invoke(ctx, BookingService_CreateBooking_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetBookingHistory(ctx context.Context, in *GetBookingHistoryRequest, opts ...grpc.CallOption) (*GetBookingHistoryResponse, error) {
	out := new(GetBookingHistoryResponse)
	err := c.cc.Invoke(ctx, BookingService_GetBookingHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) CancelBooking(ctx context.Context, in *CancelBookingRequest, opts ...grpc.CallOption) (*CancelBookingResponse, error) {
	out := new(CancelBookingResponse)
	err := c.cc.Invoke(ctx, BookingService_CancelBooking_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) CheckRoomAvailability(ctx context.Context, in *CheckRoomAvailabilityRequest, opts ...grpc.CallOption) (*CheckRoomAvailabilityResponse, error) {
	out := new(CheckRoomAvailabilityResponse)
	err := c.cc.Invoke(ctx, BookingService_CheckRoomAvailability_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookingServiceServer is the server API for BookingService service.
// All implementations must embed UnimplementedBookingServiceServer
// for forward compatibility
type BookingServiceServer interface {
	CreateBooking(context.Context, *CreateBookingRequest) (*CreateBookingResponse, error)
	GetBookingHistory(context.Context, *GetBookingHistoryRequest) (*GetBookingHistoryResponse, error)
	CancelBooking(context.Context, *CancelBookingRequest) (*CancelBookingResponse, error)
	CheckRoomAvailability(context.Context, *CheckRoomAvailabilityRequest) (*CheckRoomAvailabilityResponse, error)
	mustEmbedUnimplementedBookingServiceServer()
}

// UnimplementedBookingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBookingServiceServer struct {
}

func (UnimplementedBookingServiceServer) CreateBooking(context.Context, *CreateBookingRequest) (*CreateBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBooking not implemented")
}
func (UnimplementedBookingServiceServer) GetBookingHistory(context.Context, *GetBookingHistoryRequest) (*GetBookingHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookingHistory not implemented")
}
func (UnimplementedBookingServiceServer) CancelBooking(context.Context, *CancelBookingRequest) (*CancelBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelBooking not implemented")
}
func (UnimplementedBookingServiceServer) CheckRoomAvailability(context.Context, *CheckRoomAvailabilityRequest) (*CheckRoomAvailabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckRoomAvailability not implemented")
}
func (UnimplementedBookingServiceServer) mustEmbedUnimplementedBookingServiceServer() {}

// UnsafeBookingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookingServiceServer will
// result in compilation errors.
type UnsafeBookingServiceServer interface {
	mustEmbedUnimplementedBookingServiceServer()
}

func RegisterBookingServiceServer(s grpc.ServiceRegistrar, srv BookingServiceServer) {
	s.RegisterService(&BookingService_ServiceDesc, srv)
}

func _BookingService_CreateBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).CreateBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingService_CreateBooking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).CreateBooking(ctx, req.(*CreateBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetBookingHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookingHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetBookingHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingService_GetBookingHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetBookingHistory(ctx, req.(*GetBookingHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_CancelBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).CancelBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingService_CancelBooking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).CancelBooking(ctx, req.(*CancelBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_CheckRoomAvailability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRoomAvailabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).CheckRoomAvailability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingService_CheckRoomAvailability_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).CheckRoomAvailability(ctx, req.(*CheckRoomAvailabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookingService_ServiceDesc is the grpc.ServiceDesc for BookingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "booking.BookingService",
	HandlerType: (*BookingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBooking",
			Handler:    _BookingService_CreateBooking_Handler,
		},
		{
			MethodName: "GetBookingHistory",
			Handler:    _BookingService_GetBookingHistory_Handler,
		},
		{
			MethodName: "CancelBooking",
			Handler:    _BookingService_CancelBooking_Handler,
		},
		{
			MethodName: "CheckRoomAvailability",
			Handler:    _BookingService_CheckRoomAvailability_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "booking.proto",
}
