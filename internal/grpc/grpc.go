package grpc

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/chanhteam/go-utils/database/mysql"
	"github.com/chanhteam/go-utils/grpc/middleware"
	"github.com/chanhteam/go-utils/logger"
	"github.com/chanhteam/golang-service-example/config"
	"github.com/chanhteam/golang-service-example/internal/models"
	"github.com/chanhteam/golang-service-example/internal/services"

	credentail_v1_pb "github.com/chanhteam/golang-service-example/protobuf/v1/credential"

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

	db := mysql.GetConnection(
		appConfig.MySQL.Host,
		appConfig.MySQL.DBName,
		appConfig.MySQL.Username,
		appConfig.MySQL.Password,
		appConfig.MySQL.MaxIDLEConnection,
		appConfig.MySQL.MaxOpenConnection)
	credentialRepository := models.NewCredentialRepository(db)
	credentialService := services.NewCredentialService(credentialRepository)

	credentail_v1_pb.RegisterCredentialServiceServer(server, credentialService)

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
