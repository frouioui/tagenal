package db

import (
	"reflect"
	"testing"

	"github.com/frouioui/tagenal/api/users/pb"
)

func TestUser_ProtoUser(t *testing.T) {
	type fields struct {
		ID              int64
		Timestamp       string
		ID2             string
		UID             string
		Name            string
		Gender          string
		Email           string
		Phone           string
		Dept            string
		Grade           string
		Language        string
		Region          string
		Role            string
		PreferTags      string
		ObtainedCredits string
	}
	tests := []struct {
		name   string
		fields fields
		want   *pb.User
	}{
		{
			"user to protobuf user",
			fields{ID: 1, Timestamp: "12345", ID2: "u1", UID: "u2", Name: "john", Gender: "male", Email: "u2", Phone: "john", Dept: "male", Language: "zh", Region: "Beijing", Role: "admin", PreferTags: "tag", ObtainedCredits: "1"},
			&pb.User{ID: 1, Timestamp: "12345", ID2: "u1", UID: "u2", Name: "john", Gender: "male", Email: "u2", Phone: "john", Dept: "male", Language: "zh", Region: "Beijing", Role: "admin", PreferTags: "tag", ObtainedCredits: "1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:              tt.fields.ID,
				Timestamp:       tt.fields.Timestamp,
				ID2:             tt.fields.ID2,
				UID:             tt.fields.UID,
				Name:            tt.fields.Name,
				Gender:          tt.fields.Gender,
				Email:           tt.fields.Email,
				Phone:           tt.fields.Phone,
				Dept:            tt.fields.Dept,
				Grade:           tt.fields.Grade,
				Language:        tt.fields.Language,
				Region:          tt.fields.Region,
				Role:            tt.fields.Role,
				PreferTags:      tt.fields.PreferTags,
				ObtainedCredits: tt.fields.ObtainedCredits,
			}
			if got := u.ProtoUser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.ProtoUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProtoUserToUser(t *testing.T) {
	type args struct {
		pbuser *pb.User
	}
	tests := []struct {
		name     string
		args     args
		wantUser User
	}{
		{
			"protobuf user to user",
			args{pbuser: &pb.User{ID: 1, Timestamp: "12345", ID2: "u1", UID: "u2", Name: "john", Gender: "male", Email: "u2", Phone: "john", Dept: "male", Language: "zh", Region: "Beijing", Role: "admin", PreferTags: "tag", ObtainedCredits: "1"}},
			User{ID: 1, Timestamp: "12345", ID2: "u1", UID: "u2", Name: "john", Gender: "male", Email: "u2", Phone: "john", Dept: "male", Language: "zh", Region: "Beijing", Role: "admin", PreferTags: "tag", ObtainedCredits: "1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotUser := ProtoUserToUser(tt.args.pbuser); !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("ProtoUserToUser() = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}

func TestUsersToProtoUsers(t *testing.T) {
	type args struct {
		users []User
	}
	tests := []struct {
		name string
		args args
		want *pb.Users
	}{
		{
			"protobuf user to user",
			args{users: []User{
				{ID: 1, Timestamp: "12345", ID2: "u1", UID: "u2", Name: "john", Gender: "male", Email: "u2", Phone: "john", Dept: "it", Language: "zh", Region: "Beijing", Role: "admin", PreferTags: "tag", ObtainedCredits: "1"},
				{ID: 2, Timestamp: "78910", ID2: "u2", UID: "u3", Name: "tela", Gender: "fem", Email: "u3", Phone: "tela", Dept: "cs", Language: "zh", Region: "Hong Kong", Role: "admin2", PreferTags: "tag2", ObtainedCredits: "5"},
			}},
			&pb.Users{Users: []*pb.User{
				{ID: 1, Timestamp: "12345", ID2: "u1", UID: "u2", Name: "john", Gender: "male", Email: "u2", Phone: "john", Dept: "it", Language: "zh", Region: "Beijing", Role: "admin", PreferTags: "tag", ObtainedCredits: "1"},
				{ID: 2, Timestamp: "78910", ID2: "u2", UID: "u3", Name: "tela", Gender: "fem", Email: "u3", Phone: "tela", Dept: "cs", Language: "zh", Region: "Hong Kong", Role: "admin2", PreferTags: "tag2", ObtainedCredits: "5"},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UsersToProtoUsers(tt.args.users); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsersToProtoUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}
