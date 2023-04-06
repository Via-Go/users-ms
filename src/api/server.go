package api

import (
	"github.com/wzslr321/road_runner/server/users/src/logic"
	pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"
)

type Server struct {
	service logic.IUserService
	pb.UnimplementedUsersServer
}

func NewServer(service logic.IUserService) *Server {
	return &Server{
		service: service,
	}
}
