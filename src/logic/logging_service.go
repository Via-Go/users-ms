package logic

import (
	pb "buf.build/gen/go/viago/users-ms/protocolbuffers/go/v1"
	"context"
	"fmt"
	"go.uber.org/zap"
	"time"
)

type ILoggingService interface {
	ExecuteAndLogGetUser(ctx context.Context, req *pb.GetUserRequest) (serviceResponse *ServiceResponse)
	ExecuteAndLogUpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (ServiceResponse *ServiceResponse)
	ExecuteAndLogDeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (serviceResponse *ServiceResponse)
	ExecuteAndLogCreateUser(ctx context.Context, req *pb.CreateUserRequest) (serviceResponse *ServiceResponse)
	ExecuteAndLogLoginUser(ctx context.Context, req *pb.LoginUserRequest) (serviceResponse *ServiceResponse)
	ExecuteAndLogLogoutUser(ctx context.Context, req *pb.LogoutUserRequest) (serviceResponse *ServiceResponse)
}

type LoggingService struct {
	logger *zap.Logger
	child  IUserService
}

func NewLoggingService(logger *zap.Logger, child IUserService) *LoggingService {
	return &LoggingService{logger: logger, child: child}
}

func (s *LoggingService) ExecuteAndLogGetUser(ctx context.Context, req *pb.GetUserRequest) (serviceResponse *ServiceResponse) {
	defer func(start time.Time) {
		s.logger.Info(fmt.Sprintf("user=%v\nerr=%s\ntook=%v", req.Username, serviceResponse.Message, time.Since(start)))

	}(time.Now())
	return s.child.GetUser(ctx, req)
}

func (s *LoggingService) ExecuteAndLogUpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (serviceResponse *ServiceResponse) {
	defer func(start time.Time) {
		s.logger.Info(fmt.Sprintf("user=%v\nerr=%s\ntime\ntook=%v", req.Username, serviceResponse.Message, time.Since(start)))

	}(time.Now())
	return s.child.UpdateUser(ctx, req)
}

func (s *LoggingService) ExecuteAndLogDeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (serviceResponse *ServiceResponse) {
	defer func(start time.Time) {
		s.logger.Info(fmt.Sprintf("user=%v\nerr=%s\ntime\ntook=%v", req.Id, serviceResponse.Message, time.Since(start)))
	}(time.Now())
	return s.child.DeleteUser(ctx, req)
}

func (s *LoggingService) ExecuteAndLogCreateUser(ctx context.Context, req *pb.CreateUserRequest) (serviceResponse *ServiceResponse) {
	defer func(start time.Time) {
		s.logger.Info(fmt.Sprintf("user=%v\nerr=%s\ntime\ntook=%v", req.Username, serviceResponse.Message, time.Since(start)))
	}(time.Now())
	return s.child.CreateUser(ctx, req)
}

func (s *LoggingService) ExecuteAndLogLoginUser(ctx context.Context, req *pb.LoginUserRequest) (serviceResponse *ServiceResponse) {
	defer func(start time.Time) {
		s.logger.Info(fmt.Sprintf("user=%v\nerr=%s\ntime\ntook=%v", req.Username, serviceResponse.Message, time.Since(start)))
	}(time.Now())
	return s.child.LoginUser(ctx, req)
}

func (s *LoggingService) ExecuteAndLogLogoutUser(ctx context.Context, req *pb.LogoutUserRequest) (serviceResponse *ServiceResponse) {
	defer func(start time.Time) {
		s.logger.Info(fmt.Sprintf("user=%v\nerr=%s\ntime\ntook=%v", req.Id, serviceResponse.Message, time.Since(start)))
	}(time.Now())
	return s.child.LogoutUser(ctx, req)
}
