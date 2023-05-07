package handler

import (
	"net/http"
	"strconv"
	"time"

	. "github.com/SurquijoP/golang_web_app/src/entities/jobs"
)

func RequestHandler(w http.ResponseWriter, r *http.Request, jobQueue chan Job) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "invalid Value delay", http.StatusBadRequest)
	}

	value, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		http.Error(w, "invalid Value value", http.StatusBadRequest)
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "invalid name", http.StatusBadRequest)
	}

	job := Job{Delay: delay, Number: value, Nombre: name}
	jobQueue <- job
}
