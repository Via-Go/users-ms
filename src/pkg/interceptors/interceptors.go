package interceptors

import (
	"context"
	"github.com/wzslr321/road_runner/server/users/pkg/metrics"
	"google.golang.org/grpc"
	"net/http"
	"time"
)

type InterceptorManager struct {
	metr metrics.Metrics
}

func NewInterceptorManager(metr metrics.Metrics) *InterceptorManager {
	return &InterceptorManager{metr: metr}
}

func (im *InterceptorManager) Metrics(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	resp, err := handler(ctx, req)
	if err != nil {
		return nil, err
	}
	var status = http.StatusOK
	im.metr.ObserveResponseTime(status, info.FullMethod, info.FullMethod, time.Since(start).Seconds())
	im.metr.IncHits(status, info.FullMethod, info.FullMethod)

	return resp, err
}

func (im *InterceptorManager) EnsureValidToken(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	return handler(ctx, req)
}
