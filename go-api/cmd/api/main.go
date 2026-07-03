package main

import (
	"context"
	"fmt"
	"go-api/internal/db"
	"go-api/internal/events"
	httpserver "go-api/internal/http"
	"go-api/internal/tasks"
	"net/http"
	"os"

	"github.com/nats-io/nats.go"
)

func main() {
	ctx := context.Background()

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		panic("DATABASE_URL is required")
	}

	pool, err := db.Connect(ctx, databaseURL)
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	natsURL := os.Getenv("NATS_URL")
	if natsURL == "" {
		panic("NATS_URL is required")
	}

	natsConn, err := nats.Connect(natsURL)
	if err != nil {
		panic(err)
	}
	defer natsConn.Close()

	taskRepo := tasks.NewRepository(pool)
	eventPublisher := events.NewPublisher(natsConn)
	router := httpserver.NewRouter(taskRepo, eventPublisher)

	fmt.Println("API listening on :8080")

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
