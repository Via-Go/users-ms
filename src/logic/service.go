package logic

import (
	"context"
	pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"
)

type Service interface {
	HandleGetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error)
	HandleUpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error)
	HandleDeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error)
	HandleCreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error)
	HandleLoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error)
	HandleLogoutUser(ctx context.Context, req *pb.LogoutUserRequest) (*pb.LogoutUserResponse, error)
}
