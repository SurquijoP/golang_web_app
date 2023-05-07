package entities

import (
	"fmt"
	"time"

	. "github.com/SurquijoP/golang_web_app/src/entities/jobs"
	"github.com/SurquijoP/golang_web_app/src/utils"
)

type Worker struct {
	Id         int
	JobQueue   chan Job
	WorkerPool chan chan Job
	QuitChan   chan bool
}

func (w *Worker) Start() {
	go func() { //Funcion Anonima auto ejecutada
		for {
			w.WorkerPool <- w.JobQueue
			select {
			case job := <-w.JobQueue:
				fmt.Printf("Worker with id %d Starded \n", w.Id)
				fib := utils.Fibonacci(job.Number)
				time.Sleep(job.Delay)
				fmt.Printf("Worker with id %d Finished with result %d", w.Id, fib)
			case <-w.QuitChan:
				fmt.Printf("Worker with %d Stopped \n", w.Id)
			}
		} //FOR indefinido
	}()

}

func (w *Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

func NewWorker(id int, workerPool chan chan Job) *Worker {
	return &Worker{Id: id, JobQueue: make(chan Job), WorkerPool: workerPool, QuitChan: make(chan bool)}
}
