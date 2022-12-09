// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             v3.20.0
// source: preference/v1/preference.proto

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

const OperationPreferenceGetPreference = "/api.preference.v1.Preference/GetPreference"
const OperationPreferenceGetTagPreference = "/api.preference.v1.Preference/GetTagPreference"
const OperationPreferenceUpdatePreference = "/api.preference.v1.Preference/UpdatePreference"
const OperationPreferenceUpdateTagPreference = "/api.preference.v1.Preference/UpdateTagPreference"

type PreferenceHTTPServer interface {
	GetPreference(context.Context, *GetPreferenceRequest) (*GetPreferenceReply, error)
	GetTagPreference(context.Context, *GetTagPreferenceRequest) (*GetTagPreferenceReply, error)
	UpdatePreference(context.Context, *UpdatePreferenceRequest) (*UpdatePreferenceReply, error)
	UpdateTagPreference(context.Context, *UpdateTagPreferenceRequest) (*UpdateTagPreferenceReply, error)
}

func RegisterPreferenceHTTPServer(s *http.Server, srv PreferenceHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/preference/update", _Preference_UpdatePreference0_HTTP_Handler(srv))
	r.POST("/api/v1/preference/get", _Preference_GetPreference0_HTTP_Handler(srv))
	r.POST("/api/v1/preference/tag/update", _Preference_UpdateTagPreference0_HTTP_Handler(srv))
	r.POST("/api/v1/preference/tag/get", _Preference_GetTagPreference0_HTTP_Handler(srv))
}

func _Preference_UpdatePreference0_HTTP_Handler(srv PreferenceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdatePreferenceRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPreferenceUpdatePreference)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdatePreference(ctx, req.(*UpdatePreferenceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdatePreferenceReply)
		return ctx.Result(200, reply)
	}
}

func _Preference_GetPreference0_HTTP_Handler(srv PreferenceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetPreferenceRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPreferenceGetPreference)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPreference(ctx, req.(*GetPreferenceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetPreferenceReply)
		return ctx.Result(200, reply)
	}
}

func _Preference_UpdateTagPreference0_HTTP_Handler(srv PreferenceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateTagPreferenceRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPreferenceUpdateTagPreference)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateTagPreference(ctx, req.(*UpdateTagPreferenceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateTagPreferenceReply)
		return ctx.Result(200, reply)
	}
}

func _Preference_GetTagPreference0_HTTP_Handler(srv PreferenceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTagPreferenceRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPreferenceGetTagPreference)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTagPreference(ctx, req.(*GetTagPreferenceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetTagPreferenceReply)
		return ctx.Result(200, reply)
	}
}

type PreferenceHTTPClient interface {
	GetPreference(ctx context.Context, req *GetPreferenceRequest, opts ...http.CallOption) (rsp *GetPreferenceReply, err error)
	GetTagPreference(ctx context.Context, req *GetTagPreferenceRequest, opts ...http.CallOption) (rsp *GetTagPreferenceReply, err error)
	UpdatePreference(ctx context.Context, req *UpdatePreferenceRequest, opts ...http.CallOption) (rsp *UpdatePreferenceReply, err error)
	UpdateTagPreference(ctx context.Context, req *UpdateTagPreferenceRequest, opts ...http.CallOption) (rsp *UpdateTagPreferenceReply, err error)
}

type PreferenceHTTPClientImpl struct {
	cc *http.Client
}

func NewPreferenceHTTPClient(client *http.Client) PreferenceHTTPClient {
	return &PreferenceHTTPClientImpl{client}
}

func (c *PreferenceHTTPClientImpl) GetPreference(ctx context.Context, in *GetPreferenceRequest, opts ...http.CallOption) (*GetPreferenceReply, error) {
	var out GetPreferenceReply
	pattern := "/api/v1/preference/get"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPreferenceGetPreference))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PreferenceHTTPClientImpl) GetTagPreference(ctx context.Context, in *GetTagPreferenceRequest, opts ...http.CallOption) (*GetTagPreferenceReply, error) {
	var out GetTagPreferenceReply
	pattern := "/api/v1/preference/tag/get"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPreferenceGetTagPreference))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PreferenceHTTPClientImpl) UpdatePreference(ctx context.Context, in *UpdatePreferenceRequest, opts ...http.CallOption) (*UpdatePreferenceReply, error) {
	var out UpdatePreferenceReply
	pattern := "/api/v1/preference/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPreferenceUpdatePreference))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PreferenceHTTPClientImpl) UpdateTagPreference(ctx context.Context, in *UpdateTagPreferenceRequest, opts ...http.CallOption) (*UpdateTagPreferenceReply, error) {
	var out UpdateTagPreferenceReply
	pattern := "/api/v1/preference/tag/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPreferenceUpdateTagPreference))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
