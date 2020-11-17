package server

import (
	"fmt"
	"net"

	"github.com/frouioui/tagenal/api/users/pb"
)

type grpcServer struct {
	pb.UnimplementedUserServiceServer
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
