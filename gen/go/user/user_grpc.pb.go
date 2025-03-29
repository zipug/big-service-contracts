// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: user/user.proto

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AuthService_Login_FullMethodName         = "/user.AuthService/Login"
	AuthService_Register_FullMethodName      = "/user.AuthService/Register"
	AuthService_ValidatePerms_FullMethodName = "/user.AuthService/ValidatePerms"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	ValidatePerms(ctx context.Context, in *ValidatePermsRequest, opts ...grpc.CallOption) (*ValidatePermsResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, AuthService_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, AuthService_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ValidatePerms(ctx context.Context, in *ValidatePermsRequest, opts ...grpc.CallOption) (*ValidatePermsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ValidatePermsResponse)
	err := c.cc.Invoke(ctx, AuthService_ValidatePerms_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility.
type AuthServiceServer interface {
	Login(context.Context, *LoginRequest) (*AuthResponse, error)
	Register(context.Context, *RegisterRequest) (*AuthResponse, error)
	ValidatePerms(context.Context, *ValidatePermsRequest) (*ValidatePermsResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthServiceServer struct{}

func (UnimplementedAuthServiceServer) Login(context.Context, *LoginRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServiceServer) Register(context.Context, *RegisterRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAuthServiceServer) ValidatePerms(context.Context, *ValidatePermsRequest) (*ValidatePermsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatePerms not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}
func (UnimplementedAuthServiceServer) testEmbeddedByValue()                     {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	// If the following call pancis, it indicates UnimplementedAuthServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ValidatePerms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidatePermsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ValidatePerms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ValidatePerms_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ValidatePerms(ctx, req.(*ValidatePermsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _AuthService_Register_Handler,
		},
		{
			MethodName: "ValidatePerms",
			Handler:    _AuthService_ValidatePerms_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}

const (
	UserService_GetById_FullMethodName   = "/user.UserService/GetById"
	UserService_GetMe_FullMethodName     = "/user.UserService/GetMe"
	UserService_GetAll_FullMethodName    = "/user.UserService/GetAll"
	UserService_Update_FullMethodName    = "/user.UserService/Update"
	UserService_GrantRole_FullMethodName = "/user.UserService/GrantRole"
	UserService_Delete_FullMethodName    = "/user.UserService/Delete"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	GetById(ctx context.Context, in *ByIdRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GetMe(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserResponse, error)
	GetAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AllUsersResponse, error)
	Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	GrantRole(ctx context.Context, in *GrantRoleRequest, opts ...grpc.CallOption) (*GrantRoleResponse, error)
	Delete(ctx context.Context, in *ByIdRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetById(ctx context.Context, in *ByIdRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UserService_GetById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetMe(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UserService_GetMe_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AllUsersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AllUsersResponse)
	err := c.cc.Invoke(ctx, UserService_GetAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, UserService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GrantRole(ctx context.Context, in *GrantRoleRequest, opts ...grpc.CallOption) (*GrantRoleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GrantRoleResponse)
	err := c.cc.Invoke(ctx, UserService_GrantRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Delete(ctx context.Context, in *ByIdRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, UserService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility.
type UserServiceServer interface {
	GetById(context.Context, *ByIdRequest) (*UserResponse, error)
	GetMe(context.Context, *emptypb.Empty) (*UserResponse, error)
	GetAll(context.Context, *emptypb.Empty) (*AllUsersResponse, error)
	Update(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	GrantRole(context.Context, *GrantRoleRequest) (*GrantRoleResponse, error)
	Delete(context.Context, *ByIdRequest) (*DeleteUserResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServiceServer struct{}

func (UnimplementedUserServiceServer) GetById(context.Context, *ByIdRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedUserServiceServer) GetMe(context.Context, *emptypb.Empty) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMe not implemented")
}
func (UnimplementedUserServiceServer) GetAll(context.Context, *emptypb.Empty) (*AllUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedUserServiceServer) Update(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUserServiceServer) GrantRole(context.Context, *GrantRoleRequest) (*GrantRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GrantRole not implemented")
}
func (UnimplementedUserServiceServer) Delete(context.Context, *ByIdRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}
func (UnimplementedUserServiceServer) testEmbeddedByValue()                     {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetById(ctx, req.(*ByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetMe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetMe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetMe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetMe(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAll(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Update(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GrantRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrantRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GrantRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GrantRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GrantRole(ctx, req.(*GrantRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Delete(ctx, req.(*ByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetById",
			Handler:    _UserService_GetById_Handler,
		},
		{
			MethodName: "GetMe",
			Handler:    _UserService_GetMe_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _UserService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserService_Update_Handler,
		},
		{
			MethodName: "GrantRole",
			Handler:    _UserService_GrantRole_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}

const (
	RoleService_GetById_FullMethodName           = "/user.RoleService/GetById"
	RoleService_GetAll_FullMethodName            = "/user.RoleService/GetAll"
	RoleService_Create_FullMethodName            = "/user.RoleService/Create"
	RoleService_Update_FullMethodName            = "/user.RoleService/Update"
	RoleService_Delete_FullMethodName            = "/user.RoleService/Delete"
	RoleService_UpdatePermissions_FullMethodName = "/user.RoleService/UpdatePermissions"
)

// RoleServiceClient is the client API for RoleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoleServiceClient interface {
	GetById(ctx context.Context, in *ByIdRequest, opts ...grpc.CallOption) (*RoleResponse, error)
	GetAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AllRolesResponse, error)
	Create(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*CreateRoleResponse, error)
	Update(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*UpdateRoleResponse, error)
	Delete(ctx context.Context, in *ByIdRequest, opts ...grpc.CallOption) (*DeleteRoleResponse, error)
	UpdatePermissions(ctx context.Context, in *UpdatePermissionsRequest, opts ...grpc.CallOption) (*UpdatePermissionsResponse, error)
}

type roleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoleServiceClient(cc grpc.ClientConnInterface) RoleServiceClient {
	return &roleServiceClient{cc}
}

func (c *roleServiceClient) GetById(ctx context.Context, in *ByIdRequest, opts ...grpc.CallOption) (*RoleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoleResponse)
	err := c.cc.Invoke(ctx, RoleService_GetById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) GetAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AllRolesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AllRolesResponse)
	err := c.cc.Invoke(ctx, RoleService_GetAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) Create(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*CreateRoleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateRoleResponse)
	err := c.cc.Invoke(ctx, RoleService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) Update(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*UpdateRoleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateRoleResponse)
	err := c.cc.Invoke(ctx, RoleService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) Delete(ctx context.Context, in *ByIdRequest, opts ...grpc.CallOption) (*DeleteRoleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteRoleResponse)
	err := c.cc.Invoke(ctx, RoleService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) UpdatePermissions(ctx context.Context, in *UpdatePermissionsRequest, opts ...grpc.CallOption) (*UpdatePermissionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePermissionsResponse)
	err := c.cc.Invoke(ctx, RoleService_UpdatePermissions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoleServiceServer is the server API for RoleService service.
// All implementations must embed UnimplementedRoleServiceServer
// for forward compatibility.
type RoleServiceServer interface {
	GetById(context.Context, *ByIdRequest) (*RoleResponse, error)
	GetAll(context.Context, *emptypb.Empty) (*AllRolesResponse, error)
	Create(context.Context, *CreateRoleRequest) (*CreateRoleResponse, error)
	Update(context.Context, *UpdateRoleRequest) (*UpdateRoleResponse, error)
	Delete(context.Context, *ByIdRequest) (*DeleteRoleResponse, error)
	UpdatePermissions(context.Context, *UpdatePermissionsRequest) (*UpdatePermissionsResponse, error)
	mustEmbedUnimplementedRoleServiceServer()
}

// UnimplementedRoleServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRoleServiceServer struct{}

func (UnimplementedRoleServiceServer) GetById(context.Context, *ByIdRequest) (*RoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedRoleServiceServer) GetAll(context.Context, *emptypb.Empty) (*AllRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedRoleServiceServer) Create(context.Context, *CreateRoleRequest) (*CreateRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRoleServiceServer) Update(context.Context, *UpdateRoleRequest) (*UpdateRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRoleServiceServer) Delete(context.Context, *ByIdRequest) (*DeleteRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRoleServiceServer) UpdatePermissions(context.Context, *UpdatePermissionsRequest) (*UpdatePermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePermissions not implemented")
}
func (UnimplementedRoleServiceServer) mustEmbedUnimplementedRoleServiceServer() {}
func (UnimplementedRoleServiceServer) testEmbeddedByValue()                     {}

// UnsafeRoleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoleServiceServer will
// result in compilation errors.
type UnsafeRoleServiceServer interface {
	mustEmbedUnimplementedRoleServiceServer()
}

func RegisterRoleServiceServer(s grpc.ServiceRegistrar, srv RoleServiceServer) {
	// If the following call pancis, it indicates UnimplementedRoleServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RoleService_ServiceDesc, srv)
}

func _RoleService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).GetById(ctx, req.(*ByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).GetAll(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).Create(ctx, req.(*CreateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).Update(ctx, req.(*UpdateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).Delete(ctx, req.(*ByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_UpdatePermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).UpdatePermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_UpdatePermissions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).UpdatePermissions(ctx, req.(*UpdatePermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RoleService_ServiceDesc is the grpc.ServiceDesc for RoleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.RoleService",
	HandlerType: (*RoleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetById",
			Handler:    _RoleService_GetById_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _RoleService_GetAll_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _RoleService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _RoleService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RoleService_Delete_Handler,
		},
		{
			MethodName: "UpdatePermissions",
			Handler:    _RoleService_UpdatePermissions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}
