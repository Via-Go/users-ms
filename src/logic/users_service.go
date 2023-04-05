package logic

import (
	"context"
	pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"
	"github.com/wzslr321/road_runner/server/users/src/storage"
)

type UsersService struct {
	db storage.Database
}

func NewUsersService() Service {
	return &UsersService{
		db: storage.New(),
	}
}

func (s *UsersService) HandleGetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {

	user, err := s.db.GetUser(req)
	if err != nil {
		return &pb.GetUserResponse{Code: "404", Message: err.Error()}, nil
	}

	return user, nil
}
func (s *UsersService) HandleUpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	resp, err := s.db.UpdateUser(req)
	if err != nil {
		return &pb.UpdateUserResponse{Code: "500", Message: err.Error()}, nil
	}
	return resp, nil
}
func (s *UsersService) HandleDeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	resp, err := s.db.DeleteUser(req)
	if err != nil {
		return &pb.DeleteUserResponse{Code: "500", Message: err.Error()}, nil
	}
	return resp, nil
}

// i dont like this naming but im not sure if create will be better, gotta consider TODO
func (s *UsersService) HandleCreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user, err := s.db.CreateUser(req)
	if err != nil {
		return &pb.CreateUserResponse{Code: "500", Message: err.Error()}, nil
	}

	return user, nil
}

func (s *UsersService) HandleLoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	return nil, nil
}
func (s *UsersService) HandleLogoutUser(ctx context.Context, req *pb.LogoutUserRequest) (*pb.LogoutUserResponse, error) {
	return nil, nil
}
