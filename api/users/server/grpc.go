package server

import (
	"context"
	"fmt"
	"log"
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

func (grpcsrv *grpcServer) Compute(cxt context.Context, r *pb.UserRequest) (*pb.UserResponse, error) {
	result := &pb.UserResponse{}
	result.Result = r.A + r.B

	logMessage := fmt.Sprintf("A: %d   B: %d     sum: %d", r.A, r.B, result.Result)
	log.Println(logMessage)

	return result, nil
}
