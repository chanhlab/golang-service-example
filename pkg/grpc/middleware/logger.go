package middleware

import (
	"context"
	"fmt"
	"strings"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"github.com/chanhteam/golang-service-example/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

// codeToLevel redirects OK to DEBUG level logging instead of INFO
// This is example how you can log several gRPC code results
func codeToLevel(code codes.Code) zapcore.Level {
	if code == codes.OK {
		// It is DEBUG
		return zap.DebugLevel
	}
	return grpc_zap.DefaultCodeToLevel(code)
}

// AddLogging returns grpc.Server config option that turn on logging.
func AddLogging(logger *zap.Logger, options []grpc.ServerOption) []grpc.ServerOption {
	// Shared options for the logger, with a custom gRPC code to log level function.
	zapOptions := []grpc_zap.Option{
		grpc_zap.WithLevels(codeToLevel),
	}

	// Make sure that log statements internal to gRPC library are logged using the zapLogger as well.
	grpc_zap.ReplaceGrpcLogger(logger)

	// Add unary interceptor
	options = append(options, grpc_middleware.WithUnaryServerChain(
		grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
		grpc_zap.UnaryServerInterceptor(logger, zapOptions...),
	))

	// Add stream interceptor (added as an example here)
	options = append(options, grpc_middleware.WithStreamServerChain(
		grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
		grpc_zap.StreamServerInterceptor(logger, zapOptions...),
	))
	return options
}

// AddCustomerToctx ...
func AddCustomerToctx(ctx context.Context) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if requestID, ok := md["x-request-id"]; ok {
			headerRequestID := strings.Join(requestID, ",")
			logger.Log.Debug(fmt.Sprintf("Request-ID: %s", headerRequestID))
			AddFields(ctx, map[string]interface{}{"request-id": headerRequestID})
		}
	}
}
