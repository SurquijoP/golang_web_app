package usecases

import entities "github.com/SurquijoP/golang_web_app/src/entities/jobs"

type Dispatcher struct {
	WorkerPool chan chan entities.Job
	MaxWorkers int
	JobQueue   chan entities.Job
}
