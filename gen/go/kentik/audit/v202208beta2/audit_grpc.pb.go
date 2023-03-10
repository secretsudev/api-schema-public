// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package audit

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

// AuditServiceClient is the client API for AuditService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuditServiceClient interface {
	CreateAuditEvents(ctx context.Context, in *CreateAuditEventsRequest, opts ...grpc.CallOption) (*CreateAuditEventsResponse, error)
	ListAuditEvents(ctx context.Context, in *ListAuditEventsRequest, opts ...grpc.CallOption) (*ListAuditEventsResponse, error)
	GetAuditEvent(ctx context.Context, in *GetAuditEventRequest, opts ...grpc.CallOption) (*GetAuditEventResponse, error)
}

type auditServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuditServiceClient(cc grpc.ClientConnInterface) AuditServiceClient {
	return &auditServiceClient{cc}
}

func (c *auditServiceClient) CreateAuditEvents(ctx context.Context, in *CreateAuditEventsRequest, opts ...grpc.CallOption) (*CreateAuditEventsResponse, error) {
	out := new(CreateAuditEventsResponse)
	err := c.cc.Invoke(ctx, "/kentik.audit.v202208beta2.AuditService/CreateAuditEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auditServiceClient) ListAuditEvents(ctx context.Context, in *ListAuditEventsRequest, opts ...grpc.CallOption) (*ListAuditEventsResponse, error) {
	out := new(ListAuditEventsResponse)
	err := c.cc.Invoke(ctx, "/kentik.audit.v202208beta2.AuditService/ListAuditEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auditServiceClient) GetAuditEvent(ctx context.Context, in *GetAuditEventRequest, opts ...grpc.CallOption) (*GetAuditEventResponse, error) {
	out := new(GetAuditEventResponse)
	err := c.cc.Invoke(ctx, "/kentik.audit.v202208beta2.AuditService/GetAuditEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuditServiceServer is the server API for AuditService service.
// All implementations should embed UnimplementedAuditServiceServer
// for forward compatibility
type AuditServiceServer interface {
	CreateAuditEvents(context.Context, *CreateAuditEventsRequest) (*CreateAuditEventsResponse, error)
	ListAuditEvents(context.Context, *ListAuditEventsRequest) (*ListAuditEventsResponse, error)
	GetAuditEvent(context.Context, *GetAuditEventRequest) (*GetAuditEventResponse, error)
}

// UnimplementedAuditServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAuditServiceServer struct {
}

func (UnimplementedAuditServiceServer) CreateAuditEvents(context.Context, *CreateAuditEventsRequest) (*CreateAuditEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAuditEvents not implemented")
}
func (UnimplementedAuditServiceServer) ListAuditEvents(context.Context, *ListAuditEventsRequest) (*ListAuditEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAuditEvents not implemented")
}
func (UnimplementedAuditServiceServer) GetAuditEvent(context.Context, *GetAuditEventRequest) (*GetAuditEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuditEvent not implemented")
}

// UnsafeAuditServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuditServiceServer will
// result in compilation errors.
type UnsafeAuditServiceServer interface {
	mustEmbedUnimplementedAuditServiceServer()
}

func RegisterAuditServiceServer(s grpc.ServiceRegistrar, srv AuditServiceServer) {
	s.RegisterService(&AuditService_ServiceDesc, srv)
}

func _AuditService_CreateAuditEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAuditEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuditServiceServer).CreateAuditEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kentik.audit.v202208beta2.AuditService/CreateAuditEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditServiceServer).CreateAuditEvents(ctx, req.(*CreateAuditEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuditService_ListAuditEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAuditEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuditServiceServer).ListAuditEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kentik.audit.v202208beta2.AuditService/ListAuditEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditServiceServer).ListAuditEvents(ctx, req.(*ListAuditEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuditService_GetAuditEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuditEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuditServiceServer).GetAuditEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kentik.audit.v202208beta2.AuditService/GetAuditEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditServiceServer).GetAuditEvent(ctx, req.(*GetAuditEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuditService_ServiceDesc is the grpc.ServiceDesc for AuditService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuditService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kentik.audit.v202208beta2.AuditService",
	HandlerType: (*AuditServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAuditEvents",
			Handler:    _AuditService_CreateAuditEvents_Handler,
		},
		{
			MethodName: "ListAuditEvents",
			Handler:    _AuditService_ListAuditEvents_Handler,
		},
		{
			MethodName: "GetAuditEvent",
			Handler:    _AuditService_GetAuditEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kentik/audit/v202208beta2/audit.proto",
}
