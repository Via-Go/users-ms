package main

import (
	pb "buf.build/gen/go/viago/users-ms/grpc/go/v1/usersv1grpc"
	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/wzslr321/road_runner/server/users/api"
	"github.com/wzslr321/road_runner/server/users/logic"
	"github.com/wzslr321/road_runner/server/users/mapper"
	"github.com/wzslr321/road_runner/server/users/pkg/interceptors"
	"github.com/wzslr321/road_runner/server/users/pkg/metrics"
	"github.com/wzslr321/road_runner/server/users/storage"
	"github.com/wzslr321/road_runner/server/users/util"
	"go.uber.org/zap"
	"google.golang.org/grpc"
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

	// cert, err := tls.LoadX509KeyPair("./cert/server_cert.pem", "./cert/server_key.pem")
	if err != nil {
		log.Fatalf("Failed to load cert: %v", err)
	}

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer(
		// grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
		grpc.KeepaliveParams(keepalive.ServerParameters{}),
		grpc.UnaryInterceptor(intercs.Metrics),
		grpc.ChainUnaryInterceptor(grpcprometheus.UnaryServerInterceptor),
		grpc.ChainUnaryInterceptor(intercs.EnsureValidToken),
	)

	pb.RegisterUsersServer(server, api.NewServer(loggingService))
	grpcprometheus.Register(server)
	http.Handle("/metrics", promhttp.Handler())
	_ = server.Serve(listener)
}
