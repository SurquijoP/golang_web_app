package usecases

import (
	"log"
	"net/http"

	handler "github.com/SurquijoP/golang_web_app/src/adapters/http"
	. "github.com/SurquijoP/golang_web_app/src/entities/dispatcher"
	. "github.com/SurquijoP/golang_web_app/src/entities/jobs"
)

func ApplyFibonacciPerWorkers() {
	const (
		maxWorkers   = 4
		maxQueueSize = 20
	)

	jobQueue := make(chan Job, maxQueueSize)
	dispatcher := NewDisPatcher(jobQueue, maxWorkers)

	dispatcher.Run()

	http.HandleFunc("/fib", func(w http.ResponseWriter, r *http.Request) {
		handler.RequestHandler(w, r, jobQueue)
	})
	log.Fatal(http.ListenAndServe("8080", nil))
}
