package server

import (
	"context"

	"github.com/frouioui/tagenal/api/users/pb"
)

type grpcService struct {
	pb.UnimplementedUserServiceServer
	// TODO: add a copy of the mysql client
}

func (grpcsrv *grpcService) Home(cxt context.Context, r *pb.UserHomeRequest) (*pb.UserHomeResponse, error) {
	resp := &pb.UserHomeResponse{}
	resp.IP = getHostIP()
	resp.Host = getHostName()
	return resp, nil
}
