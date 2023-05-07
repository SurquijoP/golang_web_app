package entities

import (
	. "github.com/SurquijoP/golang_web_app/src/entities/jobs"
	. "github.com/SurquijoP/golang_web_app/src/entities/workers"
)

type Dispatcher struct {
	WorkerPool chan chan Job
	MaxWorkers int
	JobQueue   chan Job
}

func NewDisPatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {
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

func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(i, d.WorkerPool)
		worker.Start()
	}
	go d.Dispatch()
}
