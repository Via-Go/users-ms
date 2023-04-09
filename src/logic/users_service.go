package logic

import (
	"context"
	"github.com/google/uuid"
	"github.com/wzslr321/road_runner/server/users/src/domain"
	"github.com/wzslr321/road_runner/server/users/src/mapper"
	pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"
	"github.com/wzslr321/road_runner/server/users/src/storage"
	"github.com/wzslr321/road_runner/server/users/src/util"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	GetUser(ctx context.Context, req *pb.GetUserRequest) *ServiceResponse
	UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) *ServiceResponse
	DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) *ServiceResponse
	CreateUser(ctx context.Context, req *pb.CreateUserRequest) *ServiceResponse
	LoginUser(ctx context.Context, req *pb.LoginUserRequest) *ServiceResponse
	LogoutUser(ctx context.Context, req *pb.LogoutUserRequest) *ServiceResponse
}

type UsersService struct {
	db            storage.IUserStorage
	validator     util.IValidator
	mapper        mapper.IUserMapper
	authenticator IAuthenticator
	authorizer    IAuthorizer
}

func NewUsersService() *UsersService {
	db := storage.New()
	userMapper := mapper.NewUserMapper()
	return &UsersService{
		validator:     util.NewValidator(),
		authenticator: NewAuthenticator(db, userMapper),
		authorizer:    NewAuthorizer(),
	}
}

func (s *UsersService) GetUser(ctx context.Context, req *pb.GetUserRequest) *ServiceResponse {
	user, err := s.db.FindUserByUsername(req.GetUsername())
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
func (s *UsersService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) *ServiceResponse {
	err := s.db.UpdateUserByID(req)
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
func (s *UsersService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) *ServiceResponse {
	err := s.db.DeleteUserByID(req.Id)
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

func (s *UsersService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) *ServiceResponse {
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

	err = s.db.SaveUser(user)
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

func (s *UsersService) LoginUser(ctx context.Context, req *pb.LoginUserRequest) *ServiceResponse {
	authenticationResult := s.authenticator.AuthenticateUser(req.Username, req.Password)
	if authenticationResult.Status != SUCCESS {
		return authenticationResult
	}

	userDTO := authenticationResult.Body[0].(*pb.UserDTO)
	tokens, err := s.authorizer.GenerateTokens(userDTO)

	if err != nil {
		return &ServiceResponse{
			Status:  FAILED,
			Message: err.Error(),
			Body:    []interface{}{},
		}
	}

	return &ServiceResponse{
		Status:  SUCCESS,
		Message: "Logged in",
		Body:    []interface{}{tokens, userDTO},
	}
}

func (s *UsersService) LogoutUser(ctx context.Context, req *pb.LogoutUserRequest) *ServiceResponse {
	return &ServiceResponse{
		Status:  0,
		Message: "",
		Body:    nil,
	}
}
