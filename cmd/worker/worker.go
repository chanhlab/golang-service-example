package worker

import (
	"context"
	"fmt"
	"log"

	"github.com/hibiken/asynq"

	"github.com/chanhlab/golang-service-example/config"
	"github.com/chanhlab/golang-service-example/internal/workers"
)

// RunWorker starts the worker server
func RunWorker(ctx context.Context) error {
	config.NewConfig()
	cfg := config.AppConfig

	// Create Redis connection config for asynq
	redisConnOpt := asynq.RedisClientOpt{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	}

	// Create asynq server
	srv := asynq.NewServer(redisConnOpt, asynq.Config{
		Concurrency: 10,
		Queues: map[string]int{
			"default": 6,
			"low":     3,
			"high":    1,
		},
	})

	// Create mux and register handlers
	mux := asynq.NewServeMux()

	// Register credential worker handler
	credentialWorker := workers.NewCredentialWorker()
	mux.HandleFunc("credential:process", credentialWorker.ProcessTask)

	log.Println("Starting worker server...")

	// Start the server and blocks until receiving termination signal
	return srv.Run(mux)
}
