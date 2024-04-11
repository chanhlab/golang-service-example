package workers

import (
	JobWorkers "github.com/digitalocean/go-workers2"
)

// WorkerMiddleware is middleware of worker
type WorkerMiddleware struct{}

// NewWorkerMiddleware ...
func NewWorkerMiddleware() JobWorkers.Middlewares {
	workerMiddleware := &WorkerMiddleware{}
	return JobWorkers.DefaultMiddlewares().Append(workerMiddleware.Call)
}

// Call is callback
func (w *WorkerMiddleware) Call(_ string, mgr *JobWorkers.Manager, next JobWorkers.JobFunc) JobWorkers.JobFunc {
	return func(message *JobWorkers.Msg) (err error) {
		// do something before each message is processed
		err = next(message)
		// do something after each message is processed
		return
	}
}
