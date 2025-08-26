package workers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/chanhlab/golang-service-example/config"
	credentialv1 "github.com/chanhlab/golang-service-example/generated/go/credential/v1"
)

// CredentialWorker handles credential processing tasks
type CredentialWorker struct {
	config *config.Config
}

// TaskPayload represents the task payload structure
type TaskPayload struct {
	ID    int    `json:"id"`
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

// NewCredentialWorker creates a new credential worker instance
func NewCredentialWorker() *CredentialWorker {
	return &CredentialWorker{
		config: config.AppConfig,
	}
}

// ProcessTask processes a credential task
func (w *CredentialWorker) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var payload TaskPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal task payload: %w", err)
	}

	log.Printf("Processing credential task: ID=%d", payload.ID)

	// ================== Example Worker calls GRPC service to create a new Credential ====================
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = fmt.Sprintf("%d", w.config.Server.GRPCPort)
	}

	connection, err := grpc.NewClient(fmt.Sprintf(":%s", grpcPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect to gRPC server: %w", err)
	}
	defer func() {
		if closeErr := connection.Close(); closeErr != nil {
			log.Printf("Warning: failed to close gRPC connection: %v", closeErr)
		}
	}()

	client := credentialv1.NewCredentialServiceClient(connection)

	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	key := payload.Key
	value := payload.Value
	if key == "" {
		key = uuid.New().String()
	}
	if value == "" {
		value = uuid.New().String()
	}

	credential, err := client.Create(ctx, &credentialv1.CreateRequest{
		Key:   key,
		Value: value,
	})
	if err != nil {
		return fmt.Errorf("failed to create credential: %w", err)
	}

	log.Printf("Credential created - Id: %s, Key: %s, Value: %s",
		credential.Credential.GetId(),
		credential.Credential.GetKey(),
		credential.Credential.GetValue(),
	)
	// ================== End ====================

	return nil
}
