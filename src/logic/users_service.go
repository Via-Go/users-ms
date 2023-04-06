package logic

import (
	"context"
	"github.com/google/uuid"
	"github.com/wzslr321/road_runner/server/users/src/domain"
	pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"
	"github.com/wzslr321/road_runner/server/users/src/storage"
	"github.com/wzslr321/road_runner/server/users/src/util"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	HandleGetUser(ctx context.Context, req *pb.GetUserRequest) *ServiceResponse
	HandleUpdateUser(ctx context.Context, req *pb.UpdateUserRequest) *ServiceResponse
	HandleDeleteUser(ctx context.Context, req *pb.DeleteUserRequest) *ServiceResponse
	HandleCreateUser(ctx context.Context, req *pb.CreateUserRequest) *ServiceResponse
	HandleLoginUser(ctx context.Context, req *pb.LoginUserRequest) *ServiceResponse
	HandleLogoutUser(ctx context.Context, req *pb.LogoutUserRequest) *ServiceResponse
}

type UsersService struct {
	db        storage.IUserStorage
	validator util.IValidator
}

func NewUsersService() *UsersService {
	return &UsersService{
		db:        storage.New(),
		validator: util.NewValidator(),
	}
}

func (s *UsersService) HandleGetUser(ctx context.Context, req *pb.GetUserRequest) *ServiceResponse {
	user, err := s.db.GetUser(req.GetUsername())
	if err != nil {
		return &ServiceResponse{
			Status:  FAILED,
			Message: err.Error(),
			Body:    []interface{}{},
		}
	}

	return &ServiceResponse{
		Status:  SUCCESS,
		Message: "User found",
		Body:    []interface{}{user},
	}
}
func (s *UsersService) HandleUpdateUser(ctx context.Context, req *pb.UpdateUserRequest) *ServiceResponse {
	err := s.db.UpdateUser(req)
	if err != nil {
		return &ServiceResponse{
			Status:  FAILED,
			Message: err.Error(),
			Body:    []interface{}{},
		}
	}

	return &ServiceResponse{
		Status:  SUCCESS,
		Message: "User updated",
		Body:    []interface{}{},
	}
}
func (s *UsersService) HandleDeleteUser(ctx context.Context, req *pb.DeleteUserRequest) *ServiceResponse {
	err := s.db.DeleteUser(req.Id)
	if err != nil {
		return &ServiceResponse{
			Status:  FAILED,
			Message: err.Error(),
			Body:    []interface{}{},
		}
	}
	return &ServiceResponse{
		Status:  SUCCESS,
		Message: "Used deleted",
		Body:    []interface{}{},
	}
}

func (s *UsersService) HandleCreateUser(ctx context.Context, req *pb.CreateUserRequest) *ServiceResponse {
	ok := s.validator.ValidatePassword(req.Password)
	if !ok {
		return &ServiceResponse{
			Status:  FAILED,
			Message: "Invalid password, must 8-16 characters long, one uppercase letter, one lowercase letter, one digit and one special character",
			Body:    []interface{}{},
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 8)
	if err != nil {
		return &ServiceResponse{
			Status:  FAILED,
			Message: "Couldn't hash password",
			Body:    []interface{}{},
		}
	}

	user := &domain.User{
		Username: req.Username,
		Password: string(hashedPassword),
		Email:    req.Email,
		Id:       uuid.Must(uuid.NewRandom()).String(),
	}

	err = s.db.CreateUser(user)
	if err != nil {
		return &ServiceResponse{
			Status:  FAILED,
			Message: err.Error(),
			Body:    []interface{}{},
		}
	}

	return &ServiceResponse{
		Status:  SUCCESS,
		Message: "User created",
		Body:    []interface{}{user},
	}
}

func (s *UsersService) HandleLoginUser(ctx context.Context, req *pb.LoginUserRequest) *ServiceResponse {
	return &ServiceResponse{
		Status:  0,
		Message: "",
		Body:    nil,
	}
}
func (s *UsersService) HandleLogoutUser(ctx context.Context, req *pb.LogoutUserRequest) *ServiceResponse {
	return &ServiceResponse{
		Status:  0,
		Message: "",
		Body:    nil,
	}
}
