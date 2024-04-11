package rest

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/chanhlab/go-utils/logger"
	"github.com/chanhlab/go-utils/rest/middleware"
	credentialv1 "github.com/chanhlab/golang-service-example/generated/go/credential/v1"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	ContextTimeout = 5 * time.Second
)

// RunRestServer runs HTTP/REST gateway
func RunRestServer(ctx context.Context, grpcPort int, httpPort int) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{}),
		runtime.WithIncomingHeaderMatcher(CustomHeaderMatcher),
	)

	logger.Log.Info(fmt.Sprintf("HTTP Port: %d", httpPort))

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(
			grpc_opentracing.UnaryClientInterceptor(
				grpc_opentracing.WithTracer(opentracing.GlobalTracer()),
			),
		),
	}

	err := credentialv1.RegisterCredentialServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%d", grpcPort), opts)
	if err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
	}

	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", httpPort),
		// add handler with middleware
		Handler: middleware.TracingWrapper(middleware.RequestID(middleware.AddLogger(logger.Log, mux))),
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			logger.Log.Warn("shutting down gRPC server...")
		}
		_, cancel := context.WithTimeout(ctx, ContextTimeout)
		defer cancel()
		_ = srv.Shutdown(ctx)
	}()

	logger.Log.Info("starting HTTP/REST gateway...")
	return srv.ListenAndServe()
}

// CustomHeaderMatcher Mapping from HTTP request headers to gRPC client metadata
func CustomHeaderMatcher(key string) (string, bool) {
	switch key {
	case "X-Request-Id":
		return key, true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}
