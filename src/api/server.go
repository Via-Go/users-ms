package api

import (
	"github.com/wzslr321/road_runner/server/users/src/logic"
	pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"
)

type Server struct {
	service logic.Service
	pb.UnimplementedUsersServer
}

func NewServer(service logic.Service) *Server {
	s := &Server{
		service: service,
	}
	return s
}
