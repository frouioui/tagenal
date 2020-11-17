package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/frouioui/tagenal/mysql/api/users/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8080"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) Compute(cxt context.Context, r *pb.UserRequest) (*pb.UserResponse, error) {
	result := &pb.UserResponse{}
	result.Result = r.A + r.B

	logMessage := fmt.Sprintf("A: %d   B: %d     sum: %d", r.A, r.B, result.Result)
	log.Println(logMessage)

	return result, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen:  %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
