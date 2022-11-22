// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             v3.20.0
// source: group/v1/group.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationGroupCreateGroup = "/api.group.Group/CreateGroup"
const OperationGroupDeleteGroup = "/api.group.Group/DeleteGroup"
const OperationGroupGetGroup = "/api.group.Group/GetGroup"
const OperationGroupGetUsers = "/api.group.Group/GetUsers"
const OperationGroupListGroup = "/api.group.Group/ListGroup"
const OperationGroupUpdateGroup = "/api.group.Group/UpdateGroup"

type GroupHTTPServer interface {
	CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupReply, error)
	DeleteGroup(context.Context, *DeleteGroupRequest) (*DeleteGroupReply, error)
	GetGroup(context.Context, *GetGroupRequest) (*GetGroupReply, error)
	GetUsers(context.Context, *GetUsersRequest) (*GetUsersReply, error)
	ListGroup(context.Context, *ListGroupRequest) (*ListGroupReply, error)
	UpdateGroup(context.Context, *UpdateGroupRequest) (*UpdateGroupReply, error)
}

func RegisterGroupHTTPServer(s *http.Server, srv GroupHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/group/create", _Group_CreateGroup0_HTTP_Handler(srv))
	r.POST("/api/v1/group/update", _Group_UpdateGroup0_HTTP_Handler(srv))
	r.POST("/api/v1/group/delete", _Group_DeleteGroup0_HTTP_Handler(srv))
	r.POST("/api/v1/group/get", _Group_GetGroup0_HTTP_Handler(srv))
	r.GET("/api/v1/group/list", _Group_ListGroup0_HTTP_Handler(srv))
	r.POST("/api/v1/group/user/get", _Group_GetUsers0_HTTP_Handler(srv))
}

func _Group_CreateGroup0_HTTP_Handler(srv GroupHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateGroupRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGroupCreateGroup)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateGroup(ctx, req.(*CreateGroupRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateGroupReply)
		return ctx.Result(200, reply)
	}
}

func _Group_UpdateGroup0_HTTP_Handler(srv GroupHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateGroupRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGroupUpdateGroup)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateGroup(ctx, req.(*UpdateGroupRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateGroupReply)
		return ctx.Result(200, reply)
	}
}

func _Group_DeleteGroup0_HTTP_Handler(srv GroupHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteGroupRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGroupDeleteGroup)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteGroup(ctx, req.(*DeleteGroupRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteGroupReply)
		return ctx.Result(200, reply)
	}
}

func _Group_GetGroup0_HTTP_Handler(srv GroupHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetGroupRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGroupGetGroup)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetGroup(ctx, req.(*GetGroupRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetGroupReply)
		return ctx.Result(200, reply)
	}
}

func _Group_ListGroup0_HTTP_Handler(srv GroupHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListGroupRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGroupListGroup)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListGroup(ctx, req.(*ListGroupRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListGroupReply)
		return ctx.Result(200, reply)
	}
}

func _Group_GetUsers0_HTTP_Handler(srv GroupHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUsersRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGroupGetUsers)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUsers(ctx, req.(*GetUsersRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUsersReply)
		return ctx.Result(200, reply)
	}
}

type GroupHTTPClient interface {
	CreateGroup(ctx context.Context, req *CreateGroupRequest, opts ...http.CallOption) (rsp *CreateGroupReply, err error)
	DeleteGroup(ctx context.Context, req *DeleteGroupRequest, opts ...http.CallOption) (rsp *DeleteGroupReply, err error)
	GetGroup(ctx context.Context, req *GetGroupRequest, opts ...http.CallOption) (rsp *GetGroupReply, err error)
	GetUsers(ctx context.Context, req *GetUsersRequest, opts ...http.CallOption) (rsp *GetUsersReply, err error)
	ListGroup(ctx context.Context, req *ListGroupRequest, opts ...http.CallOption) (rsp *ListGroupReply, err error)
	UpdateGroup(ctx context.Context, req *UpdateGroupRequest, opts ...http.CallOption) (rsp *UpdateGroupReply, err error)
}

type GroupHTTPClientImpl struct {
	cc *http.Client
}

func NewGroupHTTPClient(client *http.Client) GroupHTTPClient {
	return &GroupHTTPClientImpl{client}
}

func (c *GroupHTTPClientImpl) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...http.CallOption) (*CreateGroupReply, error) {
	var out CreateGroupReply
	pattern := "/api/v1/group/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationGroupCreateGroup))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GroupHTTPClientImpl) DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...http.CallOption) (*DeleteGroupReply, error) {
	var out DeleteGroupReply
	pattern := "/api/v1/group/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationGroupDeleteGroup))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GroupHTTPClientImpl) GetGroup(ctx context.Context, in *GetGroupRequest, opts ...http.CallOption) (*GetGroupReply, error) {
	var out GetGroupReply
	pattern := "/api/v1/group/get"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationGroupGetGroup))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GroupHTTPClientImpl) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...http.CallOption) (*GetUsersReply, error) {
	var out GetUsersReply
	pattern := "/api/v1/group/user/get"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationGroupGetUsers))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GroupHTTPClientImpl) ListGroup(ctx context.Context, in *ListGroupRequest, opts ...http.CallOption) (*ListGroupReply, error) {
	var out ListGroupReply
	pattern := "/api/v1/group/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationGroupListGroup))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GroupHTTPClientImpl) UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...http.CallOption) (*UpdateGroupReply, error) {
	var out UpdateGroupReply
	pattern := "/api/v1/group/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationGroupUpdateGroup))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
