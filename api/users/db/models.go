package db

import (
	"github.com/frouioui/tagenal/api/users/pb"
)

type User struct {
	ID              int64  `json:"ID"`
	Timestamp       string `json:"Timestamp"`
	ID2             string `json:"ID2"`
	UID             string `json:"UID"`
	Name            string `json:"Name"`
	Gender          string `json:"Gender"`
	Email           string `json:"Email"`
	Phone           string `json:"Phone"`
	Dept            string `json:"Dept"`
	Grade           string `json:"Grade"`
	Language        string `json:"Language"`
	Region          string `json:"Region"`
	Role            string `json:"Role"`
	PreferTags      string `json:"PreferTags"`
	ObtainedCredits string `json:"ObtainedCredits"`
}

func (u *User) ProtoUser() *pb.User {
	return &pb.User{
		ID:              u.ID,
		Timestamp:       u.Timestamp,
		ID2:             u.ID2,
		UID:             u.UID,
		Name:            u.Name,
		Gender:          u.Gender,
		Email:           u.Email,
		Phone:           u.Phone,
		Dept:            u.Dept,
		Grade:           u.Grade,
		Language:        u.Language,
		Region:          u.Region,
		Role:            u.Role,
		PreferTags:      u.PreferTags,
		ObtainedCredits: u.ObtainedCredits,
	}
}

func UsersToProtoUsers(users []User) *pb.Users {
	pbusers := &pb.Users{Users: make([]*pb.User, len(users))}
	for i, u := range users {
		pbusers.Users[i] = u.ProtoUser()
	}
	return pbusers
}
