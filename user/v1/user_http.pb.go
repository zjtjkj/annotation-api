// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             v3.20.0
// source: user/v1/user.proto

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

const OperationUserCreateUser = "/api.user.v1.User/CreateUser"
const OperationUserDeleteUser = "/api.user.v1.User/DeleteUser"
const OperationUserDistributeGroup = "/api.user.v1.User/DistributeGroup"
const OperationUserGetDeletedUser = "/api.user.v1.User/GetDeletedUser"
const OperationUserGetStatus = "/api.user.v1.User/GetStatus"
const OperationUserGetUser = "/api.user.v1.User/GetUser"
const OperationUserGetUserByUsername = "/api.user.v1.User/GetUserByUsername"
const OperationUserListUser = "/api.user.v1.User/ListUser"
const OperationUserUpdatePwd = "/api.user.v1.User/UpdatePwd"
const OperationUserUpdateStatus = "/api.user.v1.User/UpdateStatus"
const OperationUserUpdateUser = "/api.user.v1.User/UpdateUser"

type UserHTTPServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserReply, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserReply, error)
	DistributeGroup(context.Context, *DistributeGroupRequest) (*DistributeGroupReply, error)
	GetDeletedUser(context.Context, *GetDeletedUserRequest) (*GetDeletedUserReply, error)
	GetStatus(context.Context, *GetStatusRequest) (*GetStatusReply, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserReply, error)
	GetUserByUsername(context.Context, *GetUserByUsernameRequest) (*GetUserByUsernameReply, error)
	ListUser(context.Context, *ListUserRequest) (*ListUserReply, error)
	UpdatePwd(context.Context, *UpdatePwdRequest) (*UpdatePwdReply, error)
	UpdateStatus(context.Context, *UpdateStatusRequest) (*UpdateStatusReply, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserReply, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/user/create", _User_CreateUser0_HTTP_Handler(srv))
	r.POST("/api/v1/user/update", _User_UpdateUser0_HTTP_Handler(srv))
	r.POST("/api/v1/user/pwd/update", _User_UpdatePwd0_HTTP_Handler(srv))
	r.POST("/api/v1/user/status/update", _User_UpdateStatus0_HTTP_Handler(srv))
	r.POST("/api/v1/user/delete", _User_DeleteUser0_HTTP_Handler(srv))
	r.POST("/api/v1/user/get", _User_GetUser0_HTTP_Handler(srv))
	r.POST("/api/v1/user/deleted/get", _User_GetDeletedUser0_HTTP_Handler(srv))
	r.POST("/api/v1/user/list", _User_ListUser0_HTTP_Handler(srv))
	r.POST("/api/v1/user/get/username", _User_GetUserByUsername0_HTTP_Handler(srv))
	r.POST("/api/v1/status/get", _User_GetStatus0_HTTP_Handler(srv))
	r.POST("/api/v1/user/distribute", _User_DistributeGroup0_HTTP_Handler(srv))
}

func _User_CreateUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserCreateUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateUser(ctx, req.(*CreateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_UpdateUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserUpdateUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUser(ctx, req.(*UpdateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_UpdatePwd0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdatePwdRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserUpdatePwd)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdatePwd(ctx, req.(*UpdatePwdRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdatePwdReply)
		return ctx.Result(200, reply)
	}
}

func _User_UpdateStatus0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateStatusRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserUpdateStatus)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateStatus(ctx, req.(*UpdateStatusRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateStatusReply)
		return ctx.Result(200, reply)
	}
}

func _User_DeleteUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserDeleteUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteUser(ctx, req.(*DeleteUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_GetUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUser(ctx, req.(*GetUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_GetDeletedUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetDeletedUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetDeletedUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetDeletedUser(ctx, req.(*GetDeletedUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetDeletedUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_ListUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserListUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUser(ctx, req.(*ListUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_GetUserByUsername0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserByUsernameRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetUserByUsername)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserByUsername(ctx, req.(*GetUserByUsernameRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserByUsernameReply)
		return ctx.Result(200, reply)
	}
}

func _User_GetStatus0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetStatusRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetStatus)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetStatus(ctx, req.(*GetStatusRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetStatusReply)
		return ctx.Result(200, reply)
	}
}

func _User_DistributeGroup0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DistributeGroupRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserDistributeGroup)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DistributeGroup(ctx, req.(*DistributeGroupRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DistributeGroupReply)
		return ctx.Result(200, reply)
	}
}

type UserHTTPClient interface {
	CreateUser(ctx context.Context, req *CreateUserRequest, opts ...http.CallOption) (rsp *CreateUserReply, err error)
	DeleteUser(ctx context.Context, req *DeleteUserRequest, opts ...http.CallOption) (rsp *DeleteUserReply, err error)
	DistributeGroup(ctx context.Context, req *DistributeGroupRequest, opts ...http.CallOption) (rsp *DistributeGroupReply, err error)
	GetDeletedUser(ctx context.Context, req *GetDeletedUserRequest, opts ...http.CallOption) (rsp *GetDeletedUserReply, err error)
	GetStatus(ctx context.Context, req *GetStatusRequest, opts ...http.CallOption) (rsp *GetStatusReply, err error)
	GetUser(ctx context.Context, req *GetUserRequest, opts ...http.CallOption) (rsp *GetUserReply, err error)
	GetUserByUsername(ctx context.Context, req *GetUserByUsernameRequest, opts ...http.CallOption) (rsp *GetUserByUsernameReply, err error)
	ListUser(ctx context.Context, req *ListUserRequest, opts ...http.CallOption) (rsp *ListUserReply, err error)
	UpdatePwd(ctx context.Context, req *UpdatePwdRequest, opts ...http.CallOption) (rsp *UpdatePwdReply, err error)
	UpdateStatus(ctx context.Context, req *UpdateStatusRequest, opts ...http.CallOption) (rsp *UpdateStatusReply, err error)
	UpdateUser(ctx context.Context, req *UpdateUserRequest, opts ...http.CallOption) (rsp *UpdateUserReply, err error)
}

type UserHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHTTPClient(client *http.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...http.CallOption) (*CreateUserReply, error) {
	var out CreateUserReply
	pattern := "/api/v1/user/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserCreateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...http.CallOption) (*DeleteUserReply, error) {
	var out DeleteUserReply
	pattern := "/api/v1/user/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserDeleteUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) DistributeGroup(ctx context.Context, in *DistributeGroupRequest, opts ...http.CallOption) (*DistributeGroupReply, error) {
	var out DistributeGroupReply
	pattern := "/api/v1/user/distribute"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserDistributeGroup))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) GetDeletedUser(ctx context.Context, in *GetDeletedUserRequest, opts ...http.CallOption) (*GetDeletedUserReply, error) {
	var out GetDeletedUserReply
	pattern := "/api/v1/user/deleted/get"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserGetDeletedUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) GetStatus(ctx context.Context, in *GetStatusRequest, opts ...http.CallOption) (*GetStatusReply, error) {
	var out GetStatusReply
	pattern := "/api/v1/status/get"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserGetStatus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) GetUser(ctx context.Context, in *GetUserRequest, opts ...http.CallOption) (*GetUserReply, error) {
	var out GetUserReply
	pattern := "/api/v1/user/get"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserGetUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) GetUserByUsername(ctx context.Context, in *GetUserByUsernameRequest, opts ...http.CallOption) (*GetUserByUsernameReply, error) {
	var out GetUserByUsernameReply
	pattern := "/api/v1/user/get/username"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserGetUserByUsername))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) ListUser(ctx context.Context, in *ListUserRequest, opts ...http.CallOption) (*ListUserReply, error) {
	var out ListUserReply
	pattern := "/api/v1/user/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserListUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) UpdatePwd(ctx context.Context, in *UpdatePwdRequest, opts ...http.CallOption) (*UpdatePwdReply, error) {
	var out UpdatePwdReply
	pattern := "/api/v1/user/pwd/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserUpdatePwd))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) UpdateStatus(ctx context.Context, in *UpdateStatusRequest, opts ...http.CallOption) (*UpdateStatusReply, error) {
	var out UpdateStatusReply
	pattern := "/api/v1/user/status/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserUpdateStatus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...http.CallOption) (*UpdateUserReply, error) {
	var out UpdateUserReply
	pattern := "/api/v1/user/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserUpdateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
