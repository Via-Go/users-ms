package api

import (
	"context"
	"errors"
	"github.com/wzslr321/road_runner/server/users/src/domain"
	"github.com/wzslr321/road_runner/server/users/src/logic"
	pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"
)

func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	serviceResponse := s.service.ExecuteAndLogGetUser(ctx, req)
	if serviceResponse.Status != logic.SUCCESS {
		return nil, errors.New(serviceResponse.Message)
	}

	response := &pb.GetUserResponse{
		Id:       serviceResponse.Body[0].(*domain.User).Id,
		Username: serviceResponse.Body[0].(*domain.User).Username,
		Email:    serviceResponse.Body[0].(*domain.User).Email,
		Code:     "200",
		Message:  serviceResponse.Message,
	}

	return response, nil
}

func (s *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	serviceResponse := s.service.ExecuteAndLogUpdateUser(ctx, req)
	if serviceResponse.Status != logic.SUCCESS {
		return nil, errors.New(serviceResponse.Message)
	}

	response := &pb.UpdateUserResponse{
		Code:    "200",
		Message: serviceResponse.Message,
	}

	return response, nil
}

func (s *Server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	serviceResponse := s.service.ExecuteAndLogDeleteUser(ctx, req)
	if serviceResponse.Status != logic.SUCCESS {
		return nil, errors.New(serviceResponse.Message)
	}

	response := &pb.DeleteUserResponse{
		Code:    "200",
		Message: serviceResponse.Message,
	}

	return response, nil
}

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	serviceResponse := s.service.ExecuteAndLogCreateUser(ctx, req)
	if serviceResponse.Status != logic.SUCCESS {
		return nil, errors.New(serviceResponse.Message)
	}

	response := &pb.CreateUserResponse{
		Code:    "200",
		Message: serviceResponse.Message,
	}

	return response, nil
}

func (s *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	serviceResponse := s.service.ExecuteAndLogLoginUser(ctx, req)
	if serviceResponse.Status != logic.SUCCESS {
		return &pb.LoginUserResponse{
			Success: false,
			Message: serviceResponse.Message,
		}, errors.New(serviceResponse.Message)
	}

	tokens := serviceResponse.Body[0].(map[string]string)
	userDTO := serviceResponse.Body[1].(*pb.UserDTO)

	response := &pb.LoginUserResponse{
		Success:      true,
		Message:      serviceResponse.Message,
		JwtToken:     tokens["Base token"],
		RefreshToken: tokens["Refresh token"],
		UserDTO:      userDTO,
	}

	return response, nil
}

func (s *Server) LogoutUser(ctx context.Context, req *pb.LogoutUserRequest) (*pb.LogoutUserResponse, error) {
	serviceResponse := s.service.ExecuteAndLogLogoutUser(ctx, req)
	if serviceResponse.Status != logic.SUCCESS {
		return nil, errors.New(serviceResponse.Message)
	}

	response := &pb.LogoutUserResponse{
		Success: false,
		Message: serviceResponse.Message,
	}

	return response, nil
}
