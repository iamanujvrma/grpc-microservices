package main

import (
	"context"
	"fmt"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rebelITT/mobile_core_service/proto"
	"github.com/rebelITT/user_service/config"
	"google.golang.org/grpc"
)

type userService struct {
}

type User struct {
	ID   int64
	Name string
}

type UserList struct {
	UserInfo []User
}

func (s *userService) GetUsers(ctx context.Context, emp *empty.Empty) (*proto.UserListResponse, error) {
	usersInfo := UserList{
		UserInfo: []User{
			{
				ID:   1,
				Name: "test user1",
			},
			{
				ID:   2,
				Name: "test user2",
			},
		},
	}

	var users []*proto.User
	var user *proto.User
	for _, u := range usersInfo.UserInfo {
		user = &proto.User{
			ID:   u.ID,
			Name: u.Name,
		}
		users = append(users, user)
	}

	response := &proto.UserListResponse{
		Users: users,
	}
	return response, nil
}

func main() {
	err := config.Init()
	if err != nil {
		fmt.Printf("error in init. error: %s", err.Error())
	}

	addr := config.GetGRPCPort()
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("error in listening at grpc port. error: %s", err.Error())
	}

	srv := grpc.NewServer()

	proto.RegisterUserServiceServer(srv, &userService{})
	srv.Serve(lis)
}
