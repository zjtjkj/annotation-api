// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: project/v1/project.proto

package v1

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

// ProjectClient is the client API for Project service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectClient interface {
	CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*CreateProjectReply, error)
	UpdateProject(ctx context.Context, in *UpdateProjectRequest, opts ...grpc.CallOption) (*UpdateProjectReply, error)
	DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*DeleteProjectReply, error)
	GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*GetProjectReply, error)
	ListProject(ctx context.Context, in *ListProjectRequest, opts ...grpc.CallOption) (*ListProjectReply, error)
	DeleteProjectUser(ctx context.Context, in *DeleteProjectUserRequest, opts ...grpc.CallOption) (*DeleteProjectUserReply, error)
	DeleteProjectTag(ctx context.Context, in *DeleteProjectTagRequest, opts ...grpc.CallOption) (*DeleteProjectTagReply, error)
	GetProjectState(ctx context.Context, in *GetProjectStateRequest, opts ...grpc.CallOption) (*GetProjectStateReply, error)
	GetProjectUser(ctx context.Context, in *GetProjectUserRequest, opts ...grpc.CallOption) (*GetProjectUserReply, error)
	UpdateProjectUser(ctx context.Context, in *UpdateProjectUserRequest, opts ...grpc.CallOption) (*UpdateProjectUserReply, error)
	UpdateProjectTag(ctx context.Context, in *UpdateProjectTagRequest, opts ...grpc.CallOption) (*UpdateProjectTagReply, error)
	GetProjectTag(ctx context.Context, in *GetProjectTagRequest, opts ...grpc.CallOption) (*GetProjectTagReply, error)
}

type projectClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectClient(cc grpc.ClientConnInterface) ProjectClient {
	return &projectClient{cc}
}

func (c *projectClient) CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*CreateProjectReply, error) {
	out := new(CreateProjectReply)
	err := c.cc.Invoke(ctx, "/api.project.v1.Project/CreateProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectClient) UpdateProject(ctx context.Context, in *UpdateProjectRequest, opts ...grpc.CallOption) (*UpdateProjectReply, error) {
	out := new(UpdateProjectReply)
	err := c.cc.Invoke(ctx, "/api.project.v1.Project/UpdateProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectClient) DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*DeleteProjectReply, error) {
	out := new(DeleteProjectReply)
	err := c.cc.Invoke(ctx, "/api.project.v1.Project/DeleteProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectClient) GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*GetProjectReply, error) {
	out := new(GetProjectReply)
	err := c.cc.Invoke(ctx, "/api.project.v1.Project/GetProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectClient) ListProject(ctx context.Context, in *ListProjectRequest, opts ...grpc.CallOption) (*ListProjectReply, error) {
	out := new(ListProjectReply)
	err := c.cc.Invoke(ctx, "/api.project.v1.Project/ListProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectClient) DeleteProjectUser(ctx context.Context, in *DeleteProjectUserRequest, opts ...grpc.CallOption) (*DeleteProjectUserReply, error) {
	out := new(DeleteProjectUserReply)
	err := c.cc.Invoke(ctx, "/api.project.v1.Project/DeleteProjectUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectClient) DeleteProjectTag(ctx context.Context, in *DeleteProjectTagRequest, opts ...grpc.CallOption) (*DeleteProjectTagReply, error) {
	out := new(DeleteProjectTagReply)
	err := c.cc.Invoke(ctx, "/api.project.v1.Project/DeleteProjectTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectClient) GetProjectState(ctx context.Context, in *GetProjectStateRequest, opts ...grpc.CallOption) (*GetProjectStateReply, error) {
	out := new(GetProjectStateReply)
	err := c.cc.Invoke(ctx, "/api.project.v1.Project/GetProjectState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectClient) GetProjectUser(ctx context.Context, in *GetProjectUserRequest, opts ...grpc.CallOption) (*GetProjectUserReply, error) {
	out := new(GetProjectUserReply)
	err := c.cc.Invoke(ctx, "/api.project.v1.Project/GetProjectUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectClient) UpdateProjectUser(ctx context.Context, in *UpdateProjectUserRequest, opts ...grpc.CallOption) (*UpdateProjectUserReply, error) {
	out := new(UpdateProjectUserReply)
	err := c.cc.Invoke(ctx, "/api.project.v1.Project/UpdateProjectUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectClient) UpdateProjectTag(ctx context.Context, in *UpdateProjectTagRequest, opts ...grpc.CallOption) (*UpdateProjectTagReply, error) {
	out := new(UpdateProjectTagReply)
	err := c.cc.Invoke(ctx, "/api.project.v1.Project/UpdateProjectTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectClient) GetProjectTag(ctx context.Context, in *GetProjectTagRequest, opts ...grpc.CallOption) (*GetProjectTagReply, error) {
	out := new(GetProjectTagReply)
	err := c.cc.Invoke(ctx, "/api.project.v1.Project/GetProjectTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectServer is the server API for Project service.
// All implementations must embed UnimplementedProjectServer
// for forward compatibility
type ProjectServer interface {
	CreateProject(context.Context, *CreateProjectRequest) (*CreateProjectReply, error)
	UpdateProject(context.Context, *UpdateProjectRequest) (*UpdateProjectReply, error)
	DeleteProject(context.Context, *DeleteProjectRequest) (*DeleteProjectReply, error)
	GetProject(context.Context, *GetProjectRequest) (*GetProjectReply, error)
	ListProject(context.Context, *ListProjectRequest) (*ListProjectReply, error)
	DeleteProjectUser(context.Context, *DeleteProjectUserRequest) (*DeleteProjectUserReply, error)
	DeleteProjectTag(context.Context, *DeleteProjectTagRequest) (*DeleteProjectTagReply, error)
	GetProjectState(context.Context, *GetProjectStateRequest) (*GetProjectStateReply, error)
	GetProjectUser(context.Context, *GetProjectUserRequest) (*GetProjectUserReply, error)
	UpdateProjectUser(context.Context, *UpdateProjectUserRequest) (*UpdateProjectUserReply, error)
	UpdateProjectTag(context.Context, *UpdateProjectTagRequest) (*UpdateProjectTagReply, error)
	GetProjectTag(context.Context, *GetProjectTagRequest) (*GetProjectTagReply, error)
	mustEmbedUnimplementedProjectServer()
}

// UnimplementedProjectServer must be embedded to have forward compatible implementations.
type UnimplementedProjectServer struct {
}

func (UnimplementedProjectServer) CreateProject(context.Context, *CreateProjectRequest) (*CreateProjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}
func (UnimplementedProjectServer) UpdateProject(context.Context, *UpdateProjectRequest) (*UpdateProjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProject not implemented")
}
func (UnimplementedProjectServer) DeleteProject(context.Context, *DeleteProjectRequest) (*DeleteProjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProject not implemented")
}
func (UnimplementedProjectServer) GetProject(context.Context, *GetProjectRequest) (*GetProjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProject not implemented")
}
func (UnimplementedProjectServer) ListProject(context.Context, *ListProjectRequest) (*ListProjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProject not implemented")
}
func (UnimplementedProjectServer) DeleteProjectUser(context.Context, *DeleteProjectUserRequest) (*DeleteProjectUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProjectUser not implemented")
}
func (UnimplementedProjectServer) DeleteProjectTag(context.Context, *DeleteProjectTagRequest) (*DeleteProjectTagReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProjectTag not implemented")
}
func (UnimplementedProjectServer) GetProjectState(context.Context, *GetProjectStateRequest) (*GetProjectStateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProjectState not implemented")
}
func (UnimplementedProjectServer) GetProjectUser(context.Context, *GetProjectUserRequest) (*GetProjectUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProjectUser not implemented")
}
func (UnimplementedProjectServer) UpdateProjectUser(context.Context, *UpdateProjectUserRequest) (*UpdateProjectUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProjectUser not implemented")
}
func (UnimplementedProjectServer) UpdateProjectTag(context.Context, *UpdateProjectTagRequest) (*UpdateProjectTagReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProjectTag not implemented")
}
func (UnimplementedProjectServer) GetProjectTag(context.Context, *GetProjectTagRequest) (*GetProjectTagReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProjectTag not implemented")
}
func (UnimplementedProjectServer) mustEmbedUnimplementedProjectServer() {}

// UnsafeProjectServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectServer will
// result in compilation errors.
type UnsafeProjectServer interface {
	mustEmbedUnimplementedProjectServer()
}

func RegisterProjectServer(s grpc.ServiceRegistrar, srv ProjectServer) {
	s.RegisterService(&Project_ServiceDesc, srv)
}

func _Project_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.project.v1.Project/CreateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServer).CreateProject(ctx, req.(*CreateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Project_UpdateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServer).UpdateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.project.v1.Project/UpdateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServer).UpdateProject(ctx, req.(*UpdateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Project_DeleteProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServer).DeleteProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.project.v1.Project/DeleteProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServer).DeleteProject(ctx, req.(*DeleteProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Project_GetProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServer).GetProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.project.v1.Project/GetProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServer).GetProject(ctx, req.(*GetProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Project_ListProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServer).ListProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.project.v1.Project/ListProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServer).ListProject(ctx, req.(*ListProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Project_DeleteProjectUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProjectUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServer).DeleteProjectUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.project.v1.Project/DeleteProjectUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServer).DeleteProjectUser(ctx, req.(*DeleteProjectUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Project_DeleteProjectTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProjectTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServer).DeleteProjectTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.project.v1.Project/DeleteProjectTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServer).DeleteProjectTag(ctx, req.(*DeleteProjectTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Project_GetProjectState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServer).GetProjectState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.project.v1.Project/GetProjectState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServer).GetProjectState(ctx, req.(*GetProjectStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Project_GetProjectUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServer).GetProjectUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.project.v1.Project/GetProjectUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServer).GetProjectUser(ctx, req.(*GetProjectUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Project_UpdateProjectUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProjectUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServer).UpdateProjectUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.project.v1.Project/UpdateProjectUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServer).UpdateProjectUser(ctx, req.(*UpdateProjectUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Project_UpdateProjectTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProjectTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServer).UpdateProjectTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.project.v1.Project/UpdateProjectTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServer).UpdateProjectTag(ctx, req.(*UpdateProjectTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Project_GetProjectTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServer).GetProjectTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.project.v1.Project/GetProjectTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServer).GetProjectTag(ctx, req.(*GetProjectTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Project_ServiceDesc is the grpc.ServiceDesc for Project service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Project_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.project.v1.Project",
	HandlerType: (*ProjectServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProject",
			Handler:    _Project_CreateProject_Handler,
		},
		{
			MethodName: "UpdateProject",
			Handler:    _Project_UpdateProject_Handler,
		},
		{
			MethodName: "DeleteProject",
			Handler:    _Project_DeleteProject_Handler,
		},
		{
			MethodName: "GetProject",
			Handler:    _Project_GetProject_Handler,
		},
		{
			MethodName: "ListProject",
			Handler:    _Project_ListProject_Handler,
		},
		{
			MethodName: "DeleteProjectUser",
			Handler:    _Project_DeleteProjectUser_Handler,
		},
		{
			MethodName: "DeleteProjectTag",
			Handler:    _Project_DeleteProjectTag_Handler,
		},
		{
			MethodName: "GetProjectState",
			Handler:    _Project_GetProjectState_Handler,
		},
		{
			MethodName: "GetProjectUser",
			Handler:    _Project_GetProjectUser_Handler,
		},
		{
			MethodName: "UpdateProjectUser",
			Handler:    _Project_UpdateProjectUser_Handler,
		},
		{
			MethodName: "UpdateProjectTag",
			Handler:    _Project_UpdateProjectTag_Handler,
		},
		{
			MethodName: "GetProjectTag",
			Handler:    _Project_GetProjectTag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "project/v1/project.proto",
}
