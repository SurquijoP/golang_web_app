package usecases

import (
	entities "github.com/SurquijoP/golang_web_app/src/entities/jobs"
	wor "github.com/SurquijoP/golang_web_app/src/entities/workers"
)

func NewWorker(id int, workerPool chan chan entities.Job) *wor.Worker {
	return &wor.Worker{Id: id, JobQueue: make(chan entities.Job), WorkerPool: workerPool, QuitChan: make(chan bool)}
}
