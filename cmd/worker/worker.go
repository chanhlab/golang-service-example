package worker

import (
	"context"
	"fmt"

	"github.com/chanhlab/golang-service-example/config"
	"github.com/chanhlab/golang-service-example/internal/workers"
	JobWorkers "github.com/digitalocean/go-workers2"
)

// main ...
func RunWorker(_ context.Context) error {
	config.NewConfig()
	config := config.AppConfig

	workerPool := 10

	manager, err := JobWorkers.NewManager(JobWorkers.Options{
		// location of redis instance
		ServerAddr: fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
		// instance of the database
		Database: config.Redis.Database,
		// number of connections to keep open with redis
		PoolSize: workerPool,
		// unique process id for this instance of workers (for proper recovery of inprogress jobs on crash)
		ProcessID: "1",
	})

	if err != nil {
		JobWorkers.Logger.Fatal(err)
	}

	// Register Credential Worker
	workers.NewCredentialWorker(manager, "credential", workerPool).Register()

	// producer := manager.Producer()
	// Add a job to a queue
	// i := 0
	// for i < 10000 {
	// 	producer.Enqueue("credential", "CredentialWorker", []int{i})
	// 	i++
	// }

	// stats will be available at http://localhost:8080/stats
	// go JobWorkers.StatsServer(8080)

	// Blocks until process is told to exit via unix signal
	manager.Run()
	return nil
}
