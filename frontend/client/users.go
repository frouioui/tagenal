package client

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"google.golang.org/grpc/codes"

	pb "github.com/frouioui/tagenal/frontend/client/pb/users"

	"github.com/opentracing/opentracing-go"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	"github.com/frouioui/tagenal/frontend/models"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	otgrpc "github.com/opentracing-contrib/go-grpc"
	otlog "github.com/opentracing/opentracing-go/log"
)

type responseSingleUserHTTP struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	User   models.User `json:"data"`
}

type responseArrayUsersHTTP struct {
	Status string        `json:"status"`
	Code   int           `json:"code"`
	Users  []models.User `json:"data"`
}

var grpcUsersClient pb.UserServiceClient

func InitUsersGRPC() (err error) {
	tracer := opentracing.GlobalTracer()
	conn, err := grpc.Dial(os.Getenv("USERS_API_SERVICE_HOST")+":"+os.Getenv("USERS_API_SERVICE_PORT_GRPC"),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(tracer)),
		grpc.WithStreamInterceptor(otgrpc.OpenTracingStreamClientInterceptor(tracer)),
	)

	grpcUsersClient = pb.NewUserServiceClient(conn)
	return err
}

func UsersFromIDGRPC(c echo.Context, ID int) (user *models.User, err error) {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second)
	defer cancel()
	r, err := grpcUsersClient.GetSingleUser(ctx, &pb.ID{ID: int64(ID)})
	if err != nil {
		s := status.Convert(err)
		switch s.Code() {
		case codes.NotFound:
			log.Printf("User not found: %s", s.Message())
		default:
			log.Printf("Error: %s", s.Message())
		}
		return nil, err
	}
	return &models.User{
		ID:              r.ID,
		Timestamp:       r.Timestamp,
		UID:             r.UID,
		Name:            r.Name,
		Gender:          r.Gender,
		Email:           r.Email,
		Phone:           r.Phone,
		Dept:            r.Dept,
		Grade:           r.Grade,
		Language:        r.Language,
		Region:          r.Region,
		Role:            r.Role,
		PreferTags:      r.PreferTags,
		ObtainedCredits: r.ObtainedCredits,
	}, nil
}

func UserFromID(c echo.Context, ID int) (user *models.User, err error) {
	url := fmt.Sprintf("http://users-api:10000/id/%d", ID)
	method := "GET"

	client := &http.Client{Timeout: time.Second * 10}

	span := jaegertracing.CreateChildSpan(c, "UserFromID")
	defer span.Finish()
	req, err := jaegertracing.NewTracedRequest(method, url, nil, span)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer res.Body.Close()

	var response responseSingleUserHTTP
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	span.LogFields(
		otlog.String("event", "api call UserFromID"),
		otlog.String("value", response.Status),
	)
	return &response.User, nil
}

func UsersFromRegionGRPC(c echo.Context, region string) (users []models.User, err error) {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second)
	defer cancel()
	r, err := grpcUsersClient.GetRegionUsers(ctx, &pb.Region{Region: region})
	if err != nil {
		s := status.Convert(err)
		log.Printf("Error: %s", s.Message())
		return nil, err
	}
	for _, pbUser := range r.Users {
		users = append(users, models.User{
			ID:              pbUser.ID,
			Timestamp:       pbUser.Timestamp,
			UID:             pbUser.UID,
			Name:            pbUser.Name,
			Gender:          pbUser.Gender,
			Email:           pbUser.Email,
			Phone:           pbUser.Phone,
			Dept:            pbUser.Dept,
			Grade:           pbUser.Grade,
			Language:        pbUser.Language,
			Region:          pbUser.Region,
			Role:            pbUser.Role,
			PreferTags:      pbUser.PreferTags,
			ObtainedCredits: pbUser.ObtainedCredits,
		})
	}
	return users, nil
}

func UsersFromRegion(c echo.Context, region string) (users []models.User, err error) {
	url := fmt.Sprintf("http://users-api:10000/region/%s", region)
	method := "GET"

	client := &http.Client{Timeout: time.Second * 10}
	span := jaegertracing.CreateChildSpan(c, "UsersFromRegion")
	defer span.Finish()
	req, err := jaegertracing.NewTracedRequest(method, url, nil, span)

	if err != nil {
		span.LogFields(otlog.String("err", err.Error()))
		log.Println(err.Error())
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		span.LogFields(otlog.String("err", err.Error()))
		log.Println(err.Error())
		return nil, err
	}
	defer res.Body.Close()

	var response responseArrayUsersHTTP
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		span.LogFields(otlog.String("err", err.Error()))
		log.Println(err.Error())
		return nil, err
	}
	span.LogFields(
		otlog.String("event", "api call UsersFromRegion"),
		otlog.String("value", response.Status),
	)
	return response.Users, nil
}
