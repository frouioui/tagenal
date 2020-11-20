package server

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/frouioui/tagenal/api/articles/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPort = 9090
	httpPort = 8080
)

// ArticleServerAPI contains the HTTP and GRPC servers which will
// help running the articles API.
type ArticleServerAPI struct {
	ServerHTTP *http.Server
	ServerGRPC *grpc.Server
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
	artsrv.ServerGRPC = grpc.NewServer()
	pb.RegisterArticleServiceServer(artsrv.ServerGRPC, &grpcService)
	reflection.Register(artsrv.ServerGRPC)
	return nil
}

// NewArticleServerAPI will create a new ArticleServerAPI object.
// It will initiate the HTTP and GRPC servers.
func NewArticleServerAPI() (artsrv ArticleServerAPI, err error) {
	err = artsrv.setServerHTTP()
	if err != nil {
		return artsrv, err
	}

	err = artsrv.setServerGRPC()
	if err != nil {
		return artsrv, err
	}
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
