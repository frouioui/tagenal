package server

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/frouioui/tagenal/api/articles/pb"
	otgrpc "github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPort = 9090
	httpPort = 8080
)

var (
	ready bool
)

// ArticleServerAPI mainly servers as a configuration holder.
// It contains two servers: http.Server and grpc.Server,
// both servers can be automatically initialized using the
// NewArticleServerAPI package's method.
type ArticleServerAPI struct {
	ServerHTTP *http.Server
	ServerGRPC *grpc.Server

	tracingCloser io.Closer
}

func SetReady(readiness bool) {
	ready = readiness
}

func (artsrv *ArticleServerAPI) setServerHTTP() (err error) {
	httpservice, err := newServiceHTTP()
	if err != nil {
		return err
	}
	artsrv.ServerHTTP = &http.Server{
		Addr:         fmt.Sprintf(":%d", httpPort),
		Handler:      httpservice.getRouter(),
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return nil
}

func (artsrv *ArticleServerAPI) setServerGRPC() (err error) {
	grpcService, err := newServiceGRPC()
	if err != nil {
		return err
	}
	tracer := opentracing.GlobalTracer()
	artsrv.ServerGRPC = grpc.NewServer(
		grpc.UnaryInterceptor(otgrpc.OpenTracingServerInterceptor(tracer)),
		grpc.StreamInterceptor(otgrpc.OpenTracingStreamServerInterceptor(tracer)),
	)
	pb.RegisterArticleServiceServer(artsrv.ServerGRPC, &grpcService)
	reflection.Register(artsrv.ServerGRPC)
	return nil
}

// NewArticleServerAPI will create a new ArticleServerAPI
// struct and initiate both HTTP and GRPC servers.
// Once returned, the ArticleServerAPI is ready to be used,
// both gRPC and HTTP endpoint can be started using namely:
// artsrv.RunServerGRPC() and artsrv.RunServerHTTP()
func NewArticleServerAPI() (artsrv ArticleServerAPI, err error) {
	artsrv.tracingCloser, err = newTracer()
	if err != nil {
		return artsrv, err
	}

	err = artsrv.setServerHTTP()
	if err != nil {
		return artsrv, err
	}

	err = artsrv.setServerGRPC()
	if err != nil {
		return artsrv, err
	}
	initRedisClusterClient()
	return artsrv, nil
}

// RunServerHTTP starts the http service.
func (artsrv *ArticleServerAPI) RunServerHTTP() (err error) {
	err = artsrv.ServerHTTP.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

// RunServerGRPC starts the grpc service.
func (artsrv *ArticleServerAPI) RunServerGRPC() (err error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		return err
	}
	err = artsrv.ServerGRPC.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}
