package server

import (
	"context"

	"github.com/frouioui/tagenal/api/users/pb"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) Home(cxt context.Context, r *pb.UserHomeRequest) (*pb.UserHomeResponse, error) {
	// resp := &pb.UserHomeResponse{}
	// resp.IP = getHostIP()
	// resp.Host = getHostName()
	// return resp, nil
	return &pb.UserHomeResponse{IP: "test", Host: "toto"}, nil
}
