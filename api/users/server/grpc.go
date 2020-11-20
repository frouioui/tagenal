package server

import (
	"context"
	"log"

	"github.com/frouioui/tagenal/api/users/db"
	"github.com/frouioui/tagenal/api/users/pb"
)

type userServiceGRPC struct {
	pb.UnimplementedUserServiceServer
	dbm *db.DatabaseManager
}

func newServiceGRPC() (grpcsrv userServiceGRPC, err error) {
	grpcsrv.dbm, err = db.NewDatabaseManager()
	if err != nil {
		log.Println(err.Error())
		return grpcsrv, err
	}
	return grpcsrv, nil
}

func (s *userServiceGRPC) ServiceInformation(cxt context.Context, r *pb.UserHomeRequest) (*pb.UserHomeResponse, error) {
	resp := &pb.UserHomeResponse{}
	resp.IP = getHostIP()
	resp.Host = getHostName()
	return resp, nil
}

func (s *userServiceGRPC) GetSingleUser(cxt context.Context, r *pb.RequestID) (*pb.User, error) {
	user, err := s.dbm.GetUserByID(uint64(r.ID))
	if err != nil {
		return nil, err
	}
	resp := user.ProtoUser()
	return resp, nil
}

func (s *userServiceGRPC) GetRegionUsers(cxt context.Context, r *pb.RequestRegion) (*pb.Users, error) {
	users, err := s.dbm.GetUsersOfRegion(r.Region)
	if err != nil {
		return nil, err
	}
	resp := db.UsersToProtoUsers(users)
	return resp, nil
}
