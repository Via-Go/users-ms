package logic

import (
	pb "buf.build/gen/go/viago/users-ms/protocolbuffers/go/v1"
	"context"
	"github.com/google/uuid"
	"github.com/wzslr321/road_runner/server/users/domain"
	"github.com/wzslr321/road_runner/server/users/mapper"
	"github.com/wzslr321/road_runner/server/users/storage"
	"github.com/wzslr321/road_runner/server/users/util"
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

func NewUsersService(db storage.IUserStorage, userMapper mapper.IUserMapper, validator util.IValidator, authenticator IAuthenticator, authorizer IAuthorizer) *UsersService {
	return &UsersService{
		db:            db,
		validator:     validator,
		mapper:        userMapper,
		authenticator: authenticator,
		authorizer:    authorizer,
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
		Role:     int(domain.Common),
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
