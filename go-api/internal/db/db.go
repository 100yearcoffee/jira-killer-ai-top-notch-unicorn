package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context, databaseURL string) (*pgxpool.Pool, error) {
	var lastErr error

	for range 10 {
		pool, err := pgxpool.New(ctx, databaseURL)
		if err != nil {
			lastErr = err
			time.Sleep(time.Second)
			continue
		}

		err = pool.Ping(ctx)
		if err == nil {
			return pool, nil
		}

		lastErr = err
		pool.Close()
		time.Sleep(time.Second)
	}
	return nil, lastErr
}
