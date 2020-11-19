package server

import (
	"context"

	"github.com/frouioui/tagenal/api/users/pb"
)

type userServiceGRPC struct {
	pb.UnimplementedUserServiceServer
}

func (s *userServiceGRPC) ServiceInformation(cxt context.Context, r *pb.UserHomeRequest) (*pb.UserHomeResponse, error) {
	resp := &pb.UserHomeResponse{}
	resp.IP = getHostIP()
	resp.Host = getHostName()
	return resp, nil
}
