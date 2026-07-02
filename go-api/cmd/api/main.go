package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"go-api/internal/db"
	httpserver "go-api/internal/http"
	"go-api/internal/tasks"
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

	taskRepo := tasks.NewRepository(pool)
	router := httpserver.NewRouter(taskRepo)

	fmt.Println("API listening on :8080")

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
