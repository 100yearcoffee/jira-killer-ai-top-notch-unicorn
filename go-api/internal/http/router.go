package http

import (
	"fmt"
	nethttp "net/http"

	"go-api/internal/tasks"
)

func NewRouter(taskRepo *tasks.Repository) nethttp.Handler {
	mux := nethttp.NewServeMux()

	taskHandler := NewTaskHandler(taskRepo)

	mux.HandleFunc("GET /health", func(w nethttp.ResponseWriter, r *nethttp.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(nethttp.StatusOK)
		fmt.Fprintln(w, `{"status":"ok"}`)
	})

	mux.HandleFunc("POST /tasks", taskHandler.Create)
	mux.HandleFunc("GET /tasks", taskHandler.List)
	mux.HandleFunc("GET /tasks/{id}", taskHandler.GetByID)
	mux.HandleFunc("PATCH /tasks/{id}/complete", taskHandler.Complete)

	return mux
}
