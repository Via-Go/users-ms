package api

import (
	"buf.build/gen/go/viago/auth/bufbuild/connect-go/_goconnect"
	"github.com/wzslr321/road_runner/server/users/logic"
)

type Server struct {
	service logic.ILoggingService
	_goconnect.UnimplementedAuthHandler
}

func NewServer(service logic.ILoggingService) *Server {
	return &Server{
		service: service,
	}
}
