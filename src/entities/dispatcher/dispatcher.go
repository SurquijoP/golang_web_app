package entities

import (
	. "github.com/SurquijoP/golang_web_app/src/entities/jobs"
)

type Dispatcher struct {
	WorkerPool chan chan Job
	MaxWorkers int
	JobQueue   chan Job
}

func NewWorker(jobQueue chan Job, maxWorkers int) *Dispatcher {
	return &Dispatcher{WorkerPool: make(chan chan Job), JobQueue: jobQueue, MaxWorkers: maxWorkers}
}

func (d *Dispatcher) Dispatch() {
	for {
		select {
		case job := <-d.JobQueue:
			go func() {
				workerjobQueue := <-d.WorkerPool
				workerjobQueue <- job
			}()
		}
	}
}
