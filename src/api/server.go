package api

import (
	"github.com/wzslr321/road_runner/server/users/src/logic"
	pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"
)

type Server struct {
	service logic.ILoggingService
	pb.UnimplementedUsersServer
}

func NewServer(service logic.ILoggingService) *Server {
	return &Server{
		service: service,
	}
}
