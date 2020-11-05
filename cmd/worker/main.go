package main

import (
	"fmt"
	"strconv"

	"github.com/chanhteam/golang-service-example/internal/workers"
	"github.com/chanhteam/golang-service-example/pkg"

	JobWorkers "github.com/digitalocean/go-workers2"
)

// main ...
func main() {
	redisHost := pkg.GetEnv("REDIS_HOST", "127.0.0.1")
	redisPort := pkg.GetEnv("REDIS_PORT", "6379")
	redisDatabase, _ := strconv.ParseInt(pkg.GetEnv("REDIS_DATABASE", "0"), 10, 64)
	workerPool, _ := strconv.ParseInt(pkg.GetEnv("WORKER_CONCURRENCY", "10"), 10, 64)

	manager, err := JobWorkers.NewManager(JobWorkers.Options{
		// location of redis instance
		ServerAddr: fmt.Sprintf("%s:%s", redisHost, redisPort),
		// instance of the database
		Database: int(redisDatabase),
		// number of connections to keep open with redis
		PoolSize: int(workerPool),
		// unique process id for this instance of workers (for proper recovery of inprogress jobs on crash)
		ProcessID: "1",
	})

	if err != nil {
		JobWorkers.Logger.Fatal(err)
	}

	// Register Credential Worker
	workers.NewCredentialWorker(manager, "credential", int(workerPool)).Register()

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
}
