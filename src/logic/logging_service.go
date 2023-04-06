package logic

import (
	"context"
	"fmt"
	pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"
	"go.uber.org/zap"
	"time"
)

type ILoggingService interface {
	LogHandleGetUser(ctx context.Context, req *pb.GetUserRequest) (serviceResponse *ServiceResponse)
	LogHandleUpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (ServiceResponse *ServiceResponse)
	LogHandleDeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (serviceResponse *ServiceResponse)
	LogHandleCreateUser(ctx context.Context, req *pb.CreateUserRequest) (serviceResponse *ServiceResponse)
	LogHandleLoginUser(ctx context.Context, req *pb.LoginUserRequest) (serviceResponse *ServiceResponse)
	LogHandleLogoutUser(ctx context.Context, req *pb.LogoutUserRequest) (serviceResponse *ServiceResponse)
}

type LoggingService struct {
	logger *zap.Logger
	child  IUserService
}

func NewLoggingService(child IUserService) ILoggingService {
	logger, _ := zap.NewProduction()
	defer logger.Sync() //nolint:errcheck
	return &LoggingService{child: child, logger: logger}
}

func (s *LoggingService) LogHandleGetUser(ctx context.Context, req *pb.GetUserRequest) (serviceResponse *ServiceResponse) {
	defer func(start time.Time) {
		s.logger.Info(fmt.Sprintf("user=%v\nerr=%s\ntook=%v", req.Username, serviceResponse.message, time.Since(start)))

	}(time.Now())
	return s.child.HandleGetUser(ctx, req)
}

func (s *LoggingService) LogHandleUpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (serviceResponse *ServiceResponse) {
	defer func(start time.Time) {
		s.logger.Info(fmt.Sprintf("user=%v\nerr=%s\ntime\ntook=%v", req.Username, serviceResponse.message, time.Since(start)))

	}(time.Now())
	return s.child.HandleUpdateUser(ctx, req)
}

func (s *LoggingService) LogHandleDeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (serviceResponse *ServiceResponse) {
	defer func(start time.Time) {
		s.logger.Info(fmt.Sprintf("user=%v\nerr=%s\ntime\ntook=%v", req.Id, serviceResponse.message, time.Since(start)))
	}(time.Now())
	return s.child.HandleDeleteUser(ctx, req)
}

func (s *LoggingService) LogHandleCreateUser(ctx context.Context, req *pb.CreateUserRequest) (serviceResponse *ServiceResponse) {
	defer func(start time.Time) {
		s.logger.Info(fmt.Sprintf("user=%v\nerr=%s\ntime\ntook=%v", req.Username, serviceResponse.message, time.Since(start)))
	}(time.Now())
	return s.child.HandleCreateUser(ctx, req)
}

func (s *LoggingService) LogHandleLoginUser(ctx context.Context, req *pb.LoginUserRequest) (serviceResponse *ServiceResponse) {
	defer func(start time.Time) {
		s.logger.Info(fmt.Sprintf("user=%v\nerr=%s\ntime\ntook=%v", req.Username, serviceResponse.message, time.Since(start)))
	}(time.Now())
	return s.child.HandleLoginUser(ctx, req)
}

func (s *LoggingService) LogHandleLogoutUser(ctx context.Context, req *pb.LogoutUserRequest) (serviceResponse *ServiceResponse) {
	defer func(start time.Time) {
		s.logger.Info(fmt.Sprintf("user=%v\nerr=%s\ntime\ntook=%v", req.Id, serviceResponse.message, time.Since(start)))
	}(time.Now())
	return s.child.HandleLogoutUser(ctx, req)
}
