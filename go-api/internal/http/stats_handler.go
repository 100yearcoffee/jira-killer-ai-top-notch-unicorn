package http

import (
	"context"
	"go-api/internal/stats"
	nethttp "net/http"
	"time"
)

type StatsHandler struct {
	client *stats.Client
}

func NewStatsHandler(client *stats.Client) *StatsHandler {
	return &StatsHandler{
		client: client,
	}
}

func (h *StatsHandler) Get(w nethttp.ResponseWriter, r *nethttp.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()
	stats, err := h.client.GetStats(ctx)

	if err != nil {
		writeError(w, nethttp.StatusBadGateway, "failed to fetch stats")
		return
	}

	writeJSON(w, nethttp.StatusOK, map[string]int32{
		"total_tasks":     stats.TotalTasks,
		"completed_tasks": stats.CompletedTasks,
	})
}
