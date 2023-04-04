package api

import (
	"context"
	pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"
)

func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return s.service.HandleGetUser(ctx, req)
}

func (s *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return s.service.HandleUpdateUser(ctx, req)
}

func (s *Server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return s.service.HandleDeleteUser(ctx, req)
}

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return s.service.HandleCreateUser(ctx, req)
}

func (s *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	return s.service.HandleLoginUser(ctx, req)
}

func (s *Server) LogoutUser(ctx context.Context, req *pb.LogoutUserRequest) (*pb.LogoutUserResponse, error) {
	return s.service.HandleLogoutUser(ctx, req)
}
