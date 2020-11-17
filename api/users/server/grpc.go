package server

import (
	"context"
	"fmt"
	"log"

	"github.com/frouioui/tagenal/api/users/pb"
)

type grpcService struct {
	pb.UnimplementedUserServiceServer
	// TODO: add a copy of the mysql client
}

func (grpcsrv *grpcService) Compute(cxt context.Context, r *pb.UserRequest) (*pb.UserResponse, error) {
	result := &pb.UserResponse{}
	result.Result = r.A + r.B

	logMessage := fmt.Sprintf("A: %d   B: %d     sum: %d", r.A, r.B, result.Result)
	log.Println(logMessage)

	return result, nil
}
