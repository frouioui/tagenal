package server

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/frouioui/tagenal/api/users/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPort = 9090
	httpPort = 8080
)

// UserServerAPI contains the HTTP and GRPC servers which will
// help running the users API.
type UserServerAPI struct {
	ServerHTTP *http.Server
	ServerGRPC *grpc.Server
	// TODO: add mysql client
}

func (usersrv *UserServerAPI) setServerHTTP() (err error) {
	httpservice := newServiceHTTP()
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
	usersrv.ServerGRPC = grpc.NewServer()
	// TODO: give a copy of mysql client to the new grpcService below
	pb.RegisterUserServiceServer(usersrv.ServerGRPC, &grpcService{})
	reflection.Register(usersrv.ServerGRPC)
	return nil
}

// NewUserServerAPI will create a new UserServerAPI object.
// It will initiate the HTTP and GRPC servers.
func NewUserServerAPI() (usersrv UserServerAPI, err error) {
	err = usersrv.setServerHTTP()
	if err != nil {
		return usersrv, err
	}

	err = usersrv.setServerGRPC()
	if err != nil {
		return usersrv, err
	}
	return usersrv, nil
}

func (usersrv *UserServerAPI) RunServerHTTP() (err error) {
	err = usersrv.ServerHTTP.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (usersrv *UserServerAPI) RunServerGRPC() (err error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		return err
	}
	err = usersrv.ServerGRPC.Serve(listener)
	if err != nil {
		return err
	}
	return nil
}
