package logic

import (
	"context"
	"fmt"
	pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"
	"go.uber.org/zap"
	"time"
)

type LoggingService struct {
	logger *zap.Logger
	child  Service
}

func NewLoggingService(child Service) Service {
	logger, _ := zap.NewProduction()
	defer logger.Sync() //nolint:errcheck
	return &LoggingService{child: child, logger: logger}
}

// TODO make these message better, placeholders for now

func (s *LoggingService) HandleGetUser(ctx context.Context, req *pb.GetUserRequest) (user *pb.GetUserResponse, err error) {
	defer func(start time.Time) {
		s.logger.Info(fmt.Sprintf("user=%v\nerr=%s\ntook=%v", user, err, time.Since(start)))

	}(time.Now())
	return s.child.HandleGetUser(ctx, req)
}

func (s *LoggingService) HandleUpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (user *pb.UpdateUserResponse, err error) {
	defer func(start time.Time) {
		s.logger.Info(fmt.Sprintf("user=%v\nerr=%s\ntime\ntook=%v", user, err, time.Since(start)))

	}(time.Now())
	return s.child.HandleUpdateUser(ctx, req)
}

func (s *LoggingService) HandleDeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (user *pb.DeleteUserResponse, err error) {
	defer func(start time.Time) {
		s.logger.Info(fmt.Sprintf("user=%v\nerr=%s\ntime\ntook=%v", user, err, time.Since(start)))
	}(time.Now())
	return s.child.HandleDeleteUser(ctx, req)
}

func (s *LoggingService) HandleCreateUser(ctx context.Context, req *pb.CreateUserRequest) (user *pb.CreateUserResponse, err error) {
	defer func(start time.Time) {
		s.logger.Info(fmt.Sprintf("user=%v\nerr=%s\ntime\ntook=%v", user, err, time.Since(start)))
	}(time.Now())
	return s.child.HandleCreateUser(ctx, req)
}

func (s *LoggingService) HandleLoginUser(ctx context.Context, req *pb.LoginUserRequest) (user *pb.LoginUserResponse, err error) {
	defer func(start time.Time) {
		s.logger.Info(fmt.Sprintf("user=%v\nerr=%s\ntime\ntook=%v", user, err, time.Since(start)))
	}(time.Now())
	return s.child.HandleLoginUser(ctx, req)
}

func (s *LoggingService) HandleLogoutUser(ctx context.Context, req *pb.LogoutUserRequest) (user *pb.LogoutUserResponse, err error) {
	defer func(start time.Time) {
		s.logger.Info(fmt.Sprintf("user=%v\nerr=%s\ntime\ntook=%v", user, err, time.Since(start)))
	}(time.Now())
	return s.child.HandleLogoutUser(ctx, req)
}
