package api

import (
	"context"

	"github.com/chanhlab/go-kit/logger"

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
	// initialize logger (convert string level to int - temporary until logger is updated)
	logLevel := 0 // default info
	switch srv.Config.Logging.Level {
	case "debug":
		logLevel = -1
	case "info":
		logLevel = 0
	case "warn":
		logLevel = 1
	case "error":
		logLevel = 2
	}
	logger.Init(int8(logLevel), srv.Config.Logging.TimeFormat)

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
