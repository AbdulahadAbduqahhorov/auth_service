// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: article.proto

package article_service

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

// ArticleServiceClient is the client API for ArticleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArticleServiceClient interface {
	CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*CreateArticleResponse, error)
	GetArticleList(ctx context.Context, in *GetArticleListRequest, opts ...grpc.CallOption) (*GetArticleListResponse, error)
	GetArticleById(ctx context.Context, in *GetArticleByIdRequest, opts ...grpc.CallOption) (*GetArticleByIdResponse, error)
	UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...grpc.CallOption) (*UpdateArticleResponse, error)
	DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*DeleteArticleResponse, error)
}

type articleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArticleServiceClient(cc grpc.ClientConnInterface) ArticleServiceClient {
	return &articleServiceClient{cc}
}

func (c *articleServiceClient) CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*CreateArticleResponse, error) {
	out := new(CreateArticleResponse)
	err := c.cc.Invoke(ctx, "/genproto.ArticleService/CreateArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) GetArticleList(ctx context.Context, in *GetArticleListRequest, opts ...grpc.CallOption) (*GetArticleListResponse, error) {
	out := new(GetArticleListResponse)
	err := c.cc.Invoke(ctx, "/genproto.ArticleService/GetArticleList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) GetArticleById(ctx context.Context, in *GetArticleByIdRequest, opts ...grpc.CallOption) (*GetArticleByIdResponse, error) {
	out := new(GetArticleByIdResponse)
	err := c.cc.Invoke(ctx, "/genproto.ArticleService/GetArticleById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...grpc.CallOption) (*UpdateArticleResponse, error) {
	out := new(UpdateArticleResponse)
	err := c.cc.Invoke(ctx, "/genproto.ArticleService/UpdateArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*DeleteArticleResponse, error) {
	out := new(DeleteArticleResponse)
	err := c.cc.Invoke(ctx, "/genproto.ArticleService/DeleteArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticleServiceServer is the server API for ArticleService service.
// All implementations must embed UnimplementedArticleServiceServer
// for forward compatibility
type ArticleServiceServer interface {
	CreateArticle(context.Context, *CreateArticleRequest) (*CreateArticleResponse, error)
	GetArticleList(context.Context, *GetArticleListRequest) (*GetArticleListResponse, error)
	GetArticleById(context.Context, *GetArticleByIdRequest) (*GetArticleByIdResponse, error)
	UpdateArticle(context.Context, *UpdateArticleRequest) (*UpdateArticleResponse, error)
	DeleteArticle(context.Context, *DeleteArticleRequest) (*DeleteArticleResponse, error)
	mustEmbedUnimplementedArticleServiceServer()
}

// UnimplementedArticleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedArticleServiceServer struct {
}

func (UnimplementedArticleServiceServer) CreateArticle(context.Context, *CreateArticleRequest) (*CreateArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticle not implemented")
}
func (UnimplementedArticleServiceServer) GetArticleList(context.Context, *GetArticleListRequest) (*GetArticleListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticleList not implemented")
}
func (UnimplementedArticleServiceServer) GetArticleById(context.Context, *GetArticleByIdRequest) (*GetArticleByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticleById not implemented")
}
func (UnimplementedArticleServiceServer) UpdateArticle(context.Context, *UpdateArticleRequest) (*UpdateArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArticle not implemented")
}
func (UnimplementedArticleServiceServer) DeleteArticle(context.Context, *DeleteArticleRequest) (*DeleteArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArticle not implemented")
}
func (UnimplementedArticleServiceServer) mustEmbedUnimplementedArticleServiceServer() {}

// UnsafeArticleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArticleServiceServer will
// result in compilation errors.
type UnsafeArticleServiceServer interface {
	mustEmbedUnimplementedArticleServiceServer()
}

func RegisterArticleServiceServer(s grpc.ServiceRegistrar, srv ArticleServiceServer) {
	s.RegisterService(&ArticleService_ServiceDesc, srv)
}

func _ArticleService_CreateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).CreateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ArticleService/CreateArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).CreateArticle(ctx, req.(*CreateArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_GetArticleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).GetArticleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ArticleService/GetArticleList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).GetArticleList(ctx, req.(*GetArticleListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_GetArticleById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).GetArticleById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ArticleService/GetArticleById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).GetArticleById(ctx, req.(*GetArticleByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_UpdateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).UpdateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ArticleService/UpdateArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).UpdateArticle(ctx, req.(*UpdateArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_DeleteArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).DeleteArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ArticleService/DeleteArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).DeleteArticle(ctx, req.(*DeleteArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ArticleService_ServiceDesc is the grpc.ServiceDesc for ArticleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArticleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.ArticleService",
	HandlerType: (*ArticleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateArticle",
			Handler:    _ArticleService_CreateArticle_Handler,
		},
		{
			MethodName: "GetArticleList",
			Handler:    _ArticleService_GetArticleList_Handler,
		},
		{
			MethodName: "GetArticleById",
			Handler:    _ArticleService_GetArticleById_Handler,
		},
		{
			MethodName: "UpdateArticle",
			Handler:    _ArticleService_UpdateArticle_Handler,
		},
		{
			MethodName: "DeleteArticle",
			Handler:    _ArticleService_DeleteArticle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "article.proto",
}