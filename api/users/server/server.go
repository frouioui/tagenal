package server

import (
	"fmt"
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
}

func (usersrv *UserServerAPI) setServerHTTP() (err error) {
	usersrv.ServerHTTP = &http.Server{
		Addr:         fmt.Sprintf(":%d", httpPort),
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return nil
}

func (usersrv *UserServerAPI) setServerGRPC() (err error) {
	usersrv.ServerGRPC = grpc.NewServer()
	pb.RegisterUserServiceServer(usersrv.ServerGRPC, &grpcServer{})
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