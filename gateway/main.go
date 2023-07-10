package main

import (
	"context"
	"net"

	"github.com/vadymbarabanov/go-microservices/proto/user"
	"google.golang.org/grpc"
)

type UserService struct {
}

func (UserService) Create(context.Context, *user.CreateUserRequest) (*user.CreateUserResponse, error) {

}

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

}
