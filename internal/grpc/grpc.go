package grpc

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/chanhlab/go-kit/grpc/middleware"
	"github.com/chanhlab/go-kit/logger"
	"github.com/chanhlab/golang-service-example/config"
	"github.com/chanhlab/golang-service-example/internal/database"
	credentialv1 "github.com/chanhlab/golang-service-example/generated/go/credential/v1"
	"github.com/chanhlab/golang-service-example/internal/models"
	"github.com/chanhlab/golang-service-example/internal/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// RunGrpcServer ...
func RunGrpcServer(ctx context.Context, appConfig *config.Config) error {
	logger.Log.Sugar().Info(fmt.Sprintf("gRPC Port: %d", appConfig.Server.GRPCPort))
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", appConfig.Server.GRPCPort))
	if err != nil {
		logger.Log.Sugar().Error(err)
		return err
	}

	// gRPC server statup options
	options := []grpc.ServerOption{}

	// add middleware
	options = middleware.AddLogging(logger.Log, options)

	// register server
	server := grpc.NewServer(options...)

	db := database.GetConnection(
		appConfig.Database.Host,
		appConfig.Database.Port,
		appConfig.Database.DBName,
		appConfig.Database.Username,
		appConfig.Database.Password,
		appConfig.Database.MaxOpenConns,
		appConfig.Database.MaxIdleConns)
	credentialRepository := models.NewCredentialRepository(db)
	credentialService := services.NewCredentialService(credentialRepository)

	credentialv1.RegisterCredentialServiceServer(server, credentialService)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			logger.Log.Warn("shutting down gRPC server...")
			server.GracefulStop()
			<-ctx.Done()
		}
	}()

	reflection.Register(server)
	err = server.Serve(listen)
	if err != nil {
		logger.Log.Sugar().Error(err)
		return err
	}
	return server.Serve(listen)
}
