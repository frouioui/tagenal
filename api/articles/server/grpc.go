package server

import (
	"context"

	"github.com/frouioui/tagenal/api/articles/pb"
)

type articleServiceGRPC struct {
	pb.UnimplementedArticleServiceServer
}

func newServiceGRPC() (grpcsrv articleServiceGRPC, err error) {
	return grpcsrv, nil
}

func (s *articleServiceGRPC) ServiceInformation(cxt context.Context, r *pb.ArticleHomeRequest) (*pb.ArticleHomeResponse, error) {
	resp := &pb.ArticleHomeResponse{}
	resp.IP = getHostIP()
	resp.Host = getHostName()
	return resp, nil
}
