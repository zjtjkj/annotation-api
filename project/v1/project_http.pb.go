// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             v3.20.0
// source: api/project/v1/project.proto

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

const OperationProjectCreateProject = "/api.project.v1.Project/CreateProject"
const OperationProjectDeleteProject = "/api.project.v1.Project/DeleteProject"
const OperationProjectGetProject = "/api.project.v1.Project/GetProject"
const OperationProjectListProject = "/api.project.v1.Project/ListProject"
const OperationProjectUpdateProject = "/api.project.v1.Project/UpdateProject"

type ProjectHTTPServer interface {
	CreateProject(context.Context, *CreateProjectRequest) (*CreateProjectReply, error)
	DeleteProject(context.Context, *DeleteProjectRequest) (*DeleteProjectReply, error)
	GetProject(context.Context, *GetProjectRequest) (*GetProjectReply, error)
	ListProject(context.Context, *ListProjectRequest) (*ListProjectReply, error)
	UpdateProject(context.Context, *UpdateProjectRequest) (*UpdateProjectReply, error)
}

func RegisterProjectHTTPServer(s *http.Server, srv ProjectHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/project/create", _Project_CreateProject0_HTTP_Handler(srv))
	r.POST("/api/v1/project/update", _Project_UpdateProject0_HTTP_Handler(srv))
	r.POST("/api/v1/project/delete", _Project_DeleteProject0_HTTP_Handler(srv))
	r.POST("/api/v1/project/get", _Project_GetProject0_HTTP_Handler(srv))
	r.POST("/api/v1/project/list", _Project_ListProject0_HTTP_Handler(srv))
}

func _Project_CreateProject0_HTTP_Handler(srv ProjectHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateProjectRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProjectCreateProject)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateProject(ctx, req.(*CreateProjectRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateProjectReply)
		return ctx.Result(200, reply)
	}
}

func _Project_UpdateProject0_HTTP_Handler(srv ProjectHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateProjectRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProjectUpdateProject)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateProject(ctx, req.(*UpdateProjectRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateProjectReply)
		return ctx.Result(200, reply)
	}
}

func _Project_DeleteProject0_HTTP_Handler(srv ProjectHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteProjectRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProjectDeleteProject)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteProject(ctx, req.(*DeleteProjectRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteProjectReply)
		return ctx.Result(200, reply)
	}
}

func _Project_GetProject0_HTTP_Handler(srv ProjectHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetProjectRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProjectGetProject)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetProject(ctx, req.(*GetProjectRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetProjectReply)
		return ctx.Result(200, reply)
	}
}

func _Project_ListProject0_HTTP_Handler(srv ProjectHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListProjectRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProjectListProject)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListProject(ctx, req.(*ListProjectRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListProjectReply)
		return ctx.Result(200, reply)
	}
}

type ProjectHTTPClient interface {
	CreateProject(ctx context.Context, req *CreateProjectRequest, opts ...http.CallOption) (rsp *CreateProjectReply, err error)
	DeleteProject(ctx context.Context, req *DeleteProjectRequest, opts ...http.CallOption) (rsp *DeleteProjectReply, err error)
	GetProject(ctx context.Context, req *GetProjectRequest, opts ...http.CallOption) (rsp *GetProjectReply, err error)
	ListProject(ctx context.Context, req *ListProjectRequest, opts ...http.CallOption) (rsp *ListProjectReply, err error)
	UpdateProject(ctx context.Context, req *UpdateProjectRequest, opts ...http.CallOption) (rsp *UpdateProjectReply, err error)
}

type ProjectHTTPClientImpl struct {
	cc *http.Client
}

func NewProjectHTTPClient(client *http.Client) ProjectHTTPClient {
	return &ProjectHTTPClientImpl{client}
}

func (c *ProjectHTTPClientImpl) CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...http.CallOption) (*CreateProjectReply, error) {
	var out CreateProjectReply
	pattern := "/api/v1/project/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationProjectCreateProject))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ProjectHTTPClientImpl) DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...http.CallOption) (*DeleteProjectReply, error) {
	var out DeleteProjectReply
	pattern := "/api/v1/project/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationProjectDeleteProject))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ProjectHTTPClientImpl) GetProject(ctx context.Context, in *GetProjectRequest, opts ...http.CallOption) (*GetProjectReply, error) {
	var out GetProjectReply
	pattern := "/api/v1/project/get"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationProjectGetProject))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ProjectHTTPClientImpl) ListProject(ctx context.Context, in *ListProjectRequest, opts ...http.CallOption) (*ListProjectReply, error) {
	var out ListProjectReply
	pattern := "/api/v1/project/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationProjectListProject))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ProjectHTTPClientImpl) UpdateProject(ctx context.Context, in *UpdateProjectRequest, opts ...http.CallOption) (*UpdateProjectReply, error) {
	var out UpdateProjectReply
	pattern := "/api/v1/project/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationProjectUpdateProject))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}