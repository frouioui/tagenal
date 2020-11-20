// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	ServiceInformation(ctx context.Context, in *UserHomeRequest, opts ...grpc.CallOption) (*UserHomeResponse, error)
	GetSingleUser(ctx context.Context, in *RequestID, opts ...grpc.CallOption) (*User, error)
	GetRegionUsers(ctx context.Context, in *RequestRegion, opts ...grpc.CallOption) (*Users, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) ServiceInformation(ctx context.Context, in *UserHomeRequest, opts ...grpc.CallOption) (*UserHomeResponse, error) {
	out := new(UserHomeResponse)
	err := c.cc.Invoke(ctx, "/pb.UserService/ServiceInformation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetSingleUser(ctx context.Context, in *RequestID, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/pb.UserService/GetSingleUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetRegionUsers(ctx context.Context, in *RequestRegion, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/pb.UserService/GetRegionUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	ServiceInformation(context.Context, *UserHomeRequest) (*UserHomeResponse, error)
	GetSingleUser(context.Context, *RequestID) (*User, error)
	GetRegionUsers(context.Context, *RequestRegion) (*Users, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) ServiceInformation(context.Context, *UserHomeRequest) (*UserHomeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServiceInformation not implemented")
}
func (UnimplementedUserServiceServer) GetSingleUser(context.Context, *RequestID) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSingleUser not implemented")
}
func (UnimplementedUserServiceServer) GetRegionUsers(context.Context, *RequestRegion) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegionUsers not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_ServiceInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserHomeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ServiceInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/ServiceInformation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ServiceInformation(ctx, req.(*UserHomeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetSingleUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetSingleUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/GetSingleUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetSingleUser(ctx, req.(*RequestID))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetRegionUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestRegion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetRegionUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/GetRegionUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetRegionUsers(ctx, req.(*RequestRegion))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ServiceInformation",
			Handler:    _UserService_ServiceInformation_Handler,
		},
		{
			MethodName: "GetSingleUser",
			Handler:    _UserService_GetSingleUser_Handler,
		},
		{
			MethodName: "GetRegionUsers",
			Handler:    _UserService_GetRegionUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/users.proto",
}
