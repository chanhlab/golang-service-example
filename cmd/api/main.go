package main

import (
	"context"
	"fmt"
	"os"

	"github.com/chanhteam/go-utils/logger"
	"github.com/chanhteam/golang-service-example/config"
	"github.com/chanhteam/golang-service-example/internal/grpc"
	"github.com/chanhteam/golang-service-example/internal/rest"
)

// Server ...
type Server struct {
	Config *config.Config
}

// RunServer ... runs gRPC server and HTTP gateway
func (srv *Server) RunServer() error {
	ctx := context.Background()

	// initialize logger
	logger.Init(srv.Config.Logger.LogLevel, srv.Config.Logger.LogTimeFormat)

	// run HTTP API Gateway
	go func() {
		_ = rest.RunRestServer(ctx, srv.Config.Server.GRPCPort, srv.Config.Server.HTTPPort)
	}()

	return grpc.RunGrpcServer(ctx, srv.Config)
}

func main() {
	config.NewConfig()
	server := &Server{
		Config: config.AppConfig,
	}
	if err := server.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
