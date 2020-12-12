package db

import (
	"encoding/json"

	"github.com/frouioui/tagenal/api/users/pb"
)

// UserArray represents an array of User.
// Implements encoding.BinaryMarshaler and encoding.BinaryUnmarshaler.
type UserArray []User

// User model maps to the user table of Vitess MySQL cluster.
// Implements encoding.BinaryMarshaler and encoding.BinaryUnmarshaler.
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

// ProtoUser transforms an User into a the auto-generated
// pb.User structure from protobuf.
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

// ProtoUserToUser transforms an auto-generated pb.User from
// protobuf into the package implementation of User.
func ProtoUserToUser(pbuser *pb.User) (user User) {
	return User{
		ID:              pbuser.ID,
		Timestamp:       pbuser.Timestamp,
		ID2:             pbuser.ID2,
		UID:             pbuser.UID,
		Name:            pbuser.Name,
		Gender:          pbuser.Gender,
		Email:           pbuser.Email,
		Phone:           pbuser.Phone,
		Dept:            pbuser.Dept,
		Grade:           pbuser.Grade,
		Language:        pbuser.Language,
		Region:          pbuser.Region,
		Role:            pbuser.Role,
		PreferTags:      pbuser.PreferTags,
		ObtainedCredits: pbuser.ObtainedCredits,
	}
}

// UsersToProtoUsers transforms an array of User into
// an array of pb.Users which are auto-generated from protobuf.
func UsersToProtoUsers(users []User) *pb.Users {
	pbusers := &pb.Users{Users: make([]*pb.User, len(users))}
	for i, u := range users {
		pbusers.Users[i] = u.ProtoUser()
	}
	return pbusers
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (u *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (u *User) MarshalBinary() (data []byte, err error) {
	return json.Marshal(u)
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (ua *UserArray) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, ua)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (ua *UserArray) MarshalBinary() (data []byte, err error) {
	return json.Marshal(ua)
}
