package main

import (
	"context"
	"fmt"
	"net"

	user "github.com/vadymbarabanov/go-microservices/proto/user"
	"google.golang.org/grpc"
)

type UserService struct {
	user.UnimplementedUserServer
}

func (UserService) Create(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	fmt.Println("Password:", req.Password)
	return &user.CreateUserResponse{
		Uid:   "1",
		Email: req.Email,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":4000")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	userService := &UserService{}

	user.RegisterUserServer(grpcServer, userService)

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
