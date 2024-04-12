package api

import (
	"context"

	"github.com/chanhlab/go-utils/logger"
	"github.com/chanhlab/golang-service-example/config"
	"github.com/chanhlab/golang-service-example/internal/grpc"
	"github.com/chanhlab/golang-service-example/internal/rest"
)

// Server ...
type Server struct {
	Config *config.Config
}

// RunServer ... runs gRPC server and HTTP gateway
func (srv *Server) RunServer(ctx context.Context) error {
	// initialize logger
	logger.Init(srv.Config.Logger.LogLevel, srv.Config.Logger.LogTimeFormat)

	// run HTTP API Gateway
	go func() {
		_ = rest.RunRestServer(ctx, srv.Config.Server.GRPCPort, srv.Config.Server.HTTPPort)
	}()

	return grpc.RunGrpcServer(ctx, srv.Config)
}

func RunAPI(ctx context.Context) error {
	config.NewConfig()
	server := &Server{
		Config: config.AppConfig,
	}
	return server.RunServer(ctx)
}
