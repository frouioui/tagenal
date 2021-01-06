package server

import (
	"context"
	"database/sql"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

func (s *userServiceGRPC) ServiceInformation(ctx context.Context, r *pb.InformationRequest) (*pb.InformationResponse, error) {
	resp := &pb.InformationResponse{}
	resp.IP = getHostIP()
	resp.Host = getHostName()
	return resp, nil
}

func (s *userServiceGRPC) GetSingleUser(ctx context.Context, r *pb.ID) (*pb.User, error) {
	user, err := s.dbm.GetUserByID(ctx, "", uint64(r.ID))
	if err != nil {
		if err == sql.ErrNoRows {
			st := status.New(codes.NotFound, "Not found.")
			return nil, st.Err()
		}
		st, _ := status.FromError(err)
		return nil, st.Err()
	}
	return user.ProtoUser(), nil
}

func (s *userServiceGRPC) GetRegionUsers(ctx context.Context, r *pb.Region) (*pb.Users, error) {
	users, err := s.dbm.GetUsersOfRegion(ctx, "", r.Region)
	if err != nil {
		st, _ := status.FromError(err)
		return nil, st.Err()
	}
	return db.UsersToProtoUsers(users), nil
}

func (s *userServiceGRPC) NewUser(ctx context.Context, r *pb.User) (*pb.ID, error) {
	user := db.ProtoUserToUser(r)
	id, err := s.dbm.InsertUser(user)
	if err != nil {
		st, _ := status.FromError(err)
		return nil, st.Err()
	}
	pbid := &pb.ID{ID: int64(id)}
	return pbid, nil
}

func (s *userServiceGRPC) NewUsers(ctx context.Context, r *pb.Users) (*pb.IDs, error) {
	ids := &pb.IDs{IDs: make([]*pb.ID, 0)}
	for _, u := range r.Users {
		user := db.ProtoUserToUser(u)
		id, err := s.dbm.InsertUser(user)
		if err != nil {
			st, _ := status.FromError(err)
			return nil, st.Err()
		}
		ids.IDs = append(ids.IDs, &pb.ID{ID: int64(id)})
	}
	return ids, nil
}
