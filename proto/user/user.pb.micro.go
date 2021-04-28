// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/user/user.proto

package user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for User service

func NewUserEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for User service

type UserService interface {
	SendCaptcha(ctx context.Context, in *CaptchaRequest, opts ...client.CallOption) (*CaptchaResponse, error)
	UserLogin(ctx context.Context, in *UserLoginRequest, opts ...client.CallOption) (*UserLoginResponse, error)
	UserInfo(ctx context.Context, in *UserInfoRequest, opts ...client.CallOption) (*UserInfoResponse, error)
	UserEdit(ctx context.Context, in *UserEditRequest, opts ...client.CallOption) (*UserEditResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) SendCaptcha(ctx context.Context, in *CaptchaRequest, opts ...client.CallOption) (*CaptchaResponse, error) {
	req := c.c.NewRequest(c.name, "User.SendCaptcha", in)
	out := new(CaptchaResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UserLogin(ctx context.Context, in *UserLoginRequest, opts ...client.CallOption) (*UserLoginResponse, error) {
	req := c.c.NewRequest(c.name, "User.UserLogin", in)
	out := new(UserLoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UserInfo(ctx context.Context, in *UserInfoRequest, opts ...client.CallOption) (*UserInfoResponse, error) {
	req := c.c.NewRequest(c.name, "User.UserInfo", in)
	out := new(UserInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UserEdit(ctx context.Context, in *UserEditRequest, opts ...client.CallOption) (*UserEditResponse, error) {
	req := c.c.NewRequest(c.name, "User.UserEdit", in)
	out := new(UserEditResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	SendCaptcha(context.Context, *CaptchaRequest, *CaptchaResponse) error
	UserLogin(context.Context, *UserLoginRequest, *UserLoginResponse) error
	UserInfo(context.Context, *UserInfoRequest, *UserInfoResponse) error
	UserEdit(context.Context, *UserEditRequest, *UserEditResponse) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		SendCaptcha(ctx context.Context, in *CaptchaRequest, out *CaptchaResponse) error
		UserLogin(ctx context.Context, in *UserLoginRequest, out *UserLoginResponse) error
		UserInfo(ctx context.Context, in *UserInfoRequest, out *UserInfoResponse) error
		UserEdit(ctx context.Context, in *UserEditRequest, out *UserEditResponse) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) SendCaptcha(ctx context.Context, in *CaptchaRequest, out *CaptchaResponse) error {
	return h.UserHandler.SendCaptcha(ctx, in, out)
}

func (h *userHandler) UserLogin(ctx context.Context, in *UserLoginRequest, out *UserLoginResponse) error {
	return h.UserHandler.UserLogin(ctx, in, out)
}

func (h *userHandler) UserInfo(ctx context.Context, in *UserInfoRequest, out *UserInfoResponse) error {
	return h.UserHandler.UserInfo(ctx, in, out)
}

func (h *userHandler) UserEdit(ctx context.Context, in *UserEditRequest, out *UserEditResponse) error {
	return h.UserHandler.UserEdit(ctx, in, out)
}
