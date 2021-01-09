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

// ArticleServiceClient is the client API for ArticleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArticleServiceClient interface {
	ServiceInformation(ctx context.Context, in *InformationRequest, opts ...grpc.CallOption) (*InformationResponse, error)
	GetSingleArticle(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Article, error)
	GetCategoryArticles(ctx context.Context, in *Category, opts ...grpc.CallOption) (*Articles, error)
	GetArticlesByRegion(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Articles, error)
	NewArticle(ctx context.Context, in *Article, opts ...grpc.CallOption) (*ID, error)
	NewArticles(ctx context.Context, in *Articles, opts ...grpc.CallOption) (*IDs, error)
}

type articleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArticleServiceClient(cc grpc.ClientConnInterface) ArticleServiceClient {
	return &articleServiceClient{cc}
}

func (c *articleServiceClient) ServiceInformation(ctx context.Context, in *InformationRequest, opts ...grpc.CallOption) (*InformationResponse, error) {
	out := new(InformationResponse)
	err := c.cc.Invoke(ctx, "/pb.ArticleService/ServiceInformation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) GetSingleArticle(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Article, error) {
	out := new(Article)
	err := c.cc.Invoke(ctx, "/pb.ArticleService/GetSingleArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) GetCategoryArticles(ctx context.Context, in *Category, opts ...grpc.CallOption) (*Articles, error) {
	out := new(Articles)
	err := c.cc.Invoke(ctx, "/pb.ArticleService/GetCategoryArticles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) GetArticlesByRegion(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Articles, error) {
	out := new(Articles)
	err := c.cc.Invoke(ctx, "/pb.ArticleService/GetArticlesByRegion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) NewArticle(ctx context.Context, in *Article, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/pb.ArticleService/NewArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) NewArticles(ctx context.Context, in *Articles, opts ...grpc.CallOption) (*IDs, error) {
	out := new(IDs)
	err := c.cc.Invoke(ctx, "/pb.ArticleService/NewArticles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticleServiceServer is the server API for ArticleService service.
// All implementations must embed UnimplementedArticleServiceServer
// for forward compatibility
type ArticleServiceServer interface {
	ServiceInformation(context.Context, *InformationRequest) (*InformationResponse, error)
	GetSingleArticle(context.Context, *ID) (*Article, error)
	GetCategoryArticles(context.Context, *Category) (*Articles, error)
	GetArticlesByRegion(context.Context, *ID) (*Articles, error)
	NewArticle(context.Context, *Article) (*ID, error)
	NewArticles(context.Context, *Articles) (*IDs, error)
	mustEmbedUnimplementedArticleServiceServer()
}

// UnimplementedArticleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedArticleServiceServer struct {
}

func (UnimplementedArticleServiceServer) ServiceInformation(context.Context, *InformationRequest) (*InformationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServiceInformation not implemented")
}
func (UnimplementedArticleServiceServer) GetSingleArticle(context.Context, *ID) (*Article, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSingleArticle not implemented")
}
func (UnimplementedArticleServiceServer) GetCategoryArticles(context.Context, *Category) (*Articles, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryArticles not implemented")
}
func (UnimplementedArticleServiceServer) GetArticlesByRegion(context.Context, *ID) (*Articles, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticlesByRegion not implemented")
}
func (UnimplementedArticleServiceServer) NewArticle(context.Context, *Article) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewArticle not implemented")
}
func (UnimplementedArticleServiceServer) NewArticles(context.Context, *Articles) (*IDs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewArticles not implemented")
}
func (UnimplementedArticleServiceServer) mustEmbedUnimplementedArticleServiceServer() {}

// UnsafeArticleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArticleServiceServer will
// result in compilation errors.
type UnsafeArticleServiceServer interface {
	mustEmbedUnimplementedArticleServiceServer()
}

func RegisterArticleServiceServer(s grpc.ServiceRegistrar, srv ArticleServiceServer) {
	s.RegisterService(&_ArticleService_serviceDesc, srv)
}

func _ArticleService_ServiceInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InformationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).ServiceInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ArticleService/ServiceInformation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).ServiceInformation(ctx, req.(*InformationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_GetSingleArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).GetSingleArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ArticleService/GetSingleArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).GetSingleArticle(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_GetCategoryArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Category)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).GetCategoryArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ArticleService/GetCategoryArticles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).GetCategoryArticles(ctx, req.(*Category))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_GetArticlesByRegion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).GetArticlesByRegion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ArticleService/GetArticlesByRegion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).GetArticlesByRegion(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_NewArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Article)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).NewArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ArticleService/NewArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).NewArticle(ctx, req.(*Article))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_NewArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Articles)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).NewArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ArticleService/NewArticles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).NewArticles(ctx, req.(*Articles))
	}
	return interceptor(ctx, in, info, handler)
}

var _ArticleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ArticleService",
	HandlerType: (*ArticleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ServiceInformation",
			Handler:    _ArticleService_ServiceInformation_Handler,
		},
		{
			MethodName: "GetSingleArticle",
			Handler:    _ArticleService_GetSingleArticle_Handler,
		},
		{
			MethodName: "GetCategoryArticles",
			Handler:    _ArticleService_GetCategoryArticles_Handler,
		},
		{
			MethodName: "GetArticlesByRegion",
			Handler:    _ArticleService_GetArticlesByRegion_Handler,
		},
		{
			MethodName: "NewArticle",
			Handler:    _ArticleService_NewArticle_Handler,
		},
		{
			MethodName: "NewArticles",
			Handler:    _ArticleService_NewArticles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/articles.proto",
}
