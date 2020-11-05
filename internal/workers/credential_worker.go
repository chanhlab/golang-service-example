package workers

import (
	"context"
	"fmt"
	"os"
	"time"

	JobWorkers "github.com/digitalocean/go-workers2"
	"github.com/google/uuid"
	credentail_pb "github.com/chanhteam/golang-service-example/protobuf/v1/credential"
	"google.golang.org/grpc"
)

// CredentialWorker ...
type CredentialWorker struct {
	Manager     *JobWorkers.Manager
	queue       string
	concurrency int
}

// NewCredentialWorker create instance
func NewCredentialWorker(manager *JobWorkers.Manager, queue string, concurrency int) *CredentialWorker {
	return &CredentialWorker{
		Manager:     manager,
		queue:       queue,
		concurrency: concurrency,
	}
}

// Register is a registration worker
func (worker *CredentialWorker) Register() {
	workerMiddleware := NewWorkerMiddleware()
	worker.Manager.AddWorker(worker.queue, worker.concurrency, worker.Execute, workerMiddleware...)
}

// Execute is a execution job of worker
func (worker *CredentialWorker) Execute(message *JobWorkers.Msg) error {
	fmt.Printf("\nJob ID: %s", message.Jid())
	fmt.Printf("\nID: %v\n", message.Args().Interface())

	// ================== Example Worker calls GRPC service to create a new Credential ====================
	grpcPort := os.Getenv("GRPC_PORT")
	connection, err := grpc.Dial(fmt.Sprintf(":%s", grpcPort), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		JobWorkers.Logger.Fatalf("Did not connect: %v", err)
	}
	defer connection.Close()
	client := credentail_pb.NewCredentialServiceClient(connection)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	credential, err := client.Create(ctx, &credentail_pb.CreateCredentialRequest{
		Key:   uuid.New().String(),
		Value: uuid.New().String(),
	})
	if err != nil {
		JobWorkers.Logger.Fatalf("Can not create Credential: %v", err)
	}
	JobWorkers.Logger.Printf(
		"Id: %s === Key: %s === Value: %s",
		credential.Credential.GetId(),
		credential.Credential.GetKey(),
		credential.Credential.GetValue(),
	)
	// ================== End ====================

	return nil
}
