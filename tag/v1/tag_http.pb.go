// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             v3.20.0
// source: tag/v1/tag.proto

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

const OperationTagCreateTag = "/api.tag.v1.Tag/CreateTag"
const OperationTagDeleteImage = "/api.tag.v1.Tag/DeleteImage"
const OperationTagDeleteTag = "/api.tag.v1.Tag/DeleteTag"
const OperationTagGetTag = "/api.tag.v1.Tag/GetTag"
const OperationTagListTag = "/api.tag.v1.Tag/ListTag"
const OperationTagUpdateTag = "/api.tag.v1.Tag/UpdateTag"
const OperationTagUploadImage = "/api.tag.v1.Tag/UploadImage"

type TagHTTPServer interface {
	CreateTag(context.Context, *CreateTagRequest) (*CreateTagReply, error)
	DeleteImage(context.Context, *DeleteImageRequest) (*DeleteImageReply, error)
	DeleteTag(context.Context, *DeleteTagRequest) (*DeleteTagReply, error)
	GetTag(context.Context, *GetTagRequest) (*GetTagReply, error)
	ListTag(context.Context, *ListTagRequest) (*ListTagReply, error)
	UpdateTag(context.Context, *UpdateTagRequest) (*UpdateTagReply, error)
	UploadImage(context.Context, *UploadImageRequest) (*UploadImageReply, error)
}

func RegisterTagHTTPServer(s *http.Server, srv TagHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/tag/create", _Tag_CreateTag0_HTTP_Handler(srv))
	r.POST("/api/v1/tag/update", _Tag_UpdateTag0_HTTP_Handler(srv))
	r.POST("/api/v1/tag/delete", _Tag_DeleteTag0_HTTP_Handler(srv))
	r.POST("/api/v1/tag/list", _Tag_ListTag0_HTTP_Handler(srv))
	r.POST("/api/v1/tag/get", _Tag_GetTag0_HTTP_Handler(srv))
	r.POST("/api/v1/tag/image/upload", _Tag_UploadImage0_HTTP_Handler(srv))
	r.POST("/api/v1/tag/image/delete", _Tag_DeleteImage0_HTTP_Handler(srv))
}

func _Tag_CreateTag0_HTTP_Handler(srv TagHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateTagRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTagCreateTag)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateTag(ctx, req.(*CreateTagRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateTagReply)
		return ctx.Result(200, reply)
	}
}

func _Tag_UpdateTag0_HTTP_Handler(srv TagHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateTagRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTagUpdateTag)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateTag(ctx, req.(*UpdateTagRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateTagReply)
		return ctx.Result(200, reply)
	}
}

func _Tag_DeleteTag0_HTTP_Handler(srv TagHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteTagRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTagDeleteTag)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteTag(ctx, req.(*DeleteTagRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteTagReply)
		return ctx.Result(200, reply)
	}
}

func _Tag_ListTag0_HTTP_Handler(srv TagHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListTagRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTagListTag)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListTag(ctx, req.(*ListTagRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListTagReply)
		return ctx.Result(200, reply)
	}
}

func _Tag_GetTag0_HTTP_Handler(srv TagHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTagRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTagGetTag)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTag(ctx, req.(*GetTagRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetTagReply)
		return ctx.Result(200, reply)
	}
}

func _Tag_UploadImage0_HTTP_Handler(srv TagHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UploadImageRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTagUploadImage)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UploadImage(ctx, req.(*UploadImageRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UploadImageReply)
		return ctx.Result(200, reply)
	}
}

func _Tag_DeleteImage0_HTTP_Handler(srv TagHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteImageRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTagDeleteImage)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteImage(ctx, req.(*DeleteImageRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteImageReply)
		return ctx.Result(200, reply)
	}
}

type TagHTTPClient interface {
	CreateTag(ctx context.Context, req *CreateTagRequest, opts ...http.CallOption) (rsp *CreateTagReply, err error)
	DeleteImage(ctx context.Context, req *DeleteImageRequest, opts ...http.CallOption) (rsp *DeleteImageReply, err error)
	DeleteTag(ctx context.Context, req *DeleteTagRequest, opts ...http.CallOption) (rsp *DeleteTagReply, err error)
	GetTag(ctx context.Context, req *GetTagRequest, opts ...http.CallOption) (rsp *GetTagReply, err error)
	ListTag(ctx context.Context, req *ListTagRequest, opts ...http.CallOption) (rsp *ListTagReply, err error)
	UpdateTag(ctx context.Context, req *UpdateTagRequest, opts ...http.CallOption) (rsp *UpdateTagReply, err error)
	UploadImage(ctx context.Context, req *UploadImageRequest, opts ...http.CallOption) (rsp *UploadImageReply, err error)
}

type TagHTTPClientImpl struct {
	cc *http.Client
}

func NewTagHTTPClient(client *http.Client) TagHTTPClient {
	return &TagHTTPClientImpl{client}
}

func (c *TagHTTPClientImpl) CreateTag(ctx context.Context, in *CreateTagRequest, opts ...http.CallOption) (*CreateTagReply, error) {
	var out CreateTagReply
	pattern := "/api/v1/tag/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTagCreateTag))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TagHTTPClientImpl) DeleteImage(ctx context.Context, in *DeleteImageRequest, opts ...http.CallOption) (*DeleteImageReply, error) {
	var out DeleteImageReply
	pattern := "/api/v1/tag/image/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTagDeleteImage))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TagHTTPClientImpl) DeleteTag(ctx context.Context, in *DeleteTagRequest, opts ...http.CallOption) (*DeleteTagReply, error) {
	var out DeleteTagReply
	pattern := "/api/v1/tag/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTagDeleteTag))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TagHTTPClientImpl) GetTag(ctx context.Context, in *GetTagRequest, opts ...http.CallOption) (*GetTagReply, error) {
	var out GetTagReply
	pattern := "/api/v1/tag/get"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTagGetTag))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TagHTTPClientImpl) ListTag(ctx context.Context, in *ListTagRequest, opts ...http.CallOption) (*ListTagReply, error) {
	var out ListTagReply
	pattern := "/api/v1/tag/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTagListTag))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TagHTTPClientImpl) UpdateTag(ctx context.Context, in *UpdateTagRequest, opts ...http.CallOption) (*UpdateTagReply, error) {
	var out UpdateTagReply
	pattern := "/api/v1/tag/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTagUpdateTag))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TagHTTPClientImpl) UploadImage(ctx context.Context, in *UploadImageRequest, opts ...http.CallOption) (*UploadImageReply, error) {
	var out UploadImageReply
	pattern := "/api/v1/tag/image/upload"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTagUploadImage))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
