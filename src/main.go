package main

import (
	"crypto/tls"
	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/wzslr321/road_runner/server/users/src/api"
	"github.com/wzslr321/road_runner/server/users/src/logic"
	"github.com/wzslr321/road_runner/server/users/src/mapper"
	"github.com/wzslr321/road_runner/server/users/src/pkg/interceptors"
	"github.com/wzslr321/road_runner/server/users/src/pkg/metrics"
	pb "github.com/wzslr321/road_runner/server/users/src/proto-gen"
	"github.com/wzslr321/road_runner/server/users/src/storage"
	"github.com/wzslr321/road_runner/server/users/src/util"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"log"
	"net"
	"net/http"
)

var db *storage.UserStorage
var validator *util.Validator
var loggingService *logic.LoggingService
var usersMapper *mapper.UserMapper
var usersService *logic.UsersService
var authenticator *logic.Authenticator
var authorizer *logic.Authorizer

func init() {
	logger, _ := zap.NewProduction()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			panic("Couldn't initialize logger")
		}
	}(logger) //nolint:errcheck

	db = storage.New()
	validator = util.NewValidator()
	usersMapper = mapper.NewUserMapper()
	authenticator = logic.NewAuthenticator(db, usersMapper)
	authorizer = logic.NewAuthorizer()
	usersService = logic.NewUsersService(db, usersMapper, validator, authenticator, authorizer)
	loggingService = logic.NewLoggingService(logger, usersService)
}

func main() {
	ms, err := metrics.Create("0.0.0.0:7070", "users")
	if err != nil {
		log.Fatalf("Failed to create metrics: %v", err)
	}
	intercs := interceptors.NewInterceptorManager(ms)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	cert, err := tls.LoadX509KeyPair("./cert/server_cert.pem", "./cert/server_key.pem")
	if err != nil {
		log.Fatalf("Failed to load cert: %v", err)
	}

	server := grpc.NewServer(
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
		grpc.KeepaliveParams(keepalive.ServerParameters{}),
		grpc.UnaryInterceptor(intercs.Metrics),
		grpc.ChainUnaryInterceptor(grpcprometheus.UnaryServerInterceptor),
		grpc.ChainUnaryInterceptor(intercs.EnsureValidToken),
	)

	pb.RegisterUsersServer(server, api.NewServer(loggingService))
	grpcprometheus.Register(server)
	http.Handle("/metrics", promhttp.Handler())
	server.Serve(listener) //nolint:errcheck
}
