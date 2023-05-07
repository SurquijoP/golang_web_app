package entities

import (
	. "github.com/SurquijoP/golang_web_app/src/entities/jobs"
)

type Worker struct {
	Id         int
	JobQueue   chan Job
	WorkerPool chan chan Job
	QuitChan   chan bool
}
