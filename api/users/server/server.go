package server

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/frouioui/tagenal/api/users/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPort = 9090
	httpPort = 10000
)

// UserServerAPI contains the HTTP and GRPC servers which will
// run the whole user api server.
// This structure is mainly used as a configuration component
// that uses a http.Server and grpc.Server objects, themself
// containing a router/handler for their requests.
type UserServerAPI struct {
	ServerHTTP *http.Server
	ServerGRPC *grpc.Server

	tracingCloser io.Closer
}

func (usersrv *UserServerAPI) setServerHTTP() (err error) {
	httpservice, err := newServiceHTTP()
	if err != nil {
		return err
	}
	usersrv.ServerHTTP = &http.Server{
		Addr:         fmt.Sprintf(":%d", httpPort),
		Handler:      httpservice.getRouter(),
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return nil
}

func (usersrv *UserServerAPI) setServerGRPC() (err error) {
	grpcService, err := newServiceGRPC()
	if err != nil {
		return err
	}
	usersrv.ServerGRPC = grpc.NewServer()
	pb.RegisterUserServiceServer(usersrv.ServerGRPC, &grpcService)
	reflection.Register(usersrv.ServerGRPC)
	return nil
}

// NewUserServerAPI will create a new UserServerAPI object.
// It will initiate the HTTP and GRPC servers by calling
// the setServerHTTP and serServerGRPC package functions.
// The function returns an UserServerAPI which will be ready
// to use, if no error is returned, by calling either
// usersrv.RunServerHTTP() or usersrv.RunServerGRPC().
//
// TODO: give parameters to the function in order to tune
// the servers.
func NewUserServerAPI() (usersrv UserServerAPI, err error) {
	usersrv.tracingCloser, err = newTracer()
	if err != nil {
		return usersrv, err
	}

	err = usersrv.setServerHTTP()
	if err != nil {
		return usersrv, err
	}

	err = usersrv.setServerGRPC()
	if err != nil {
		return usersrv, err
	}

	initRedisClusterClient()
	return usersrv, nil
}

// RunServerHTTP starts the http service.
func (usersrv *UserServerAPI) RunServerHTTP() (err error) {
	err = usersrv.ServerHTTP.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

// RunServerGRPC starts the grpc service.
func (usersrv *UserServerAPI) RunServerGRPC() (err error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		return err
	}
	err = usersrv.ServerGRPC.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}
