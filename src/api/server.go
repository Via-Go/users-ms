package api

import (
	"github.com/wzslr321/road_runner/server/users/logic"
)

type Server struct {
	service logic.ILoggingService
}

func NewServer(service logic.ILoggingService) *Server {
	return &Server{
		service: service,
	}
}
