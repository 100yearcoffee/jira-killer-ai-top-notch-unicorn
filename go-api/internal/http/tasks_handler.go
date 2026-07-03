package http

import (
	"encoding/json"
	"errors"
	"go-api/internal/events"
	"go-api/internal/tasks"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type TaskHandler struct {
	repo      *tasks.Repository
	publisher *events.Publisher
}

func NewTaskHandler(repo *tasks.Repository, publisher *events.Publisher) *TaskHandler {
	return &TaskHandler{
		repo:      repo,
		publisher: publisher,
	}
}

type createTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createTaskRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	req.Title = strings.TrimSpace(req.Title)
	req.Description = strings.TrimSpace(req.Description)

	if req.Title == "" {
		writeError(w, http.StatusBadRequest, "title is required")
		return
	}

	if req.Description == "" {
		writeError(w, http.StatusBadRequest, "description is required")
		return
	}

	task, err := h.repo.Create(r.Context(), req.Title, req.Description)
	if err != nil {
		writeError(w, http.StatusBadRequest, "failed to create a task")
		return
	}

	err = h.publisher.PublishTaskEvent(events.TaskCreated, task.ID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to publish task created event")
		return
	}

	writeJSON(w, http.StatusCreated, task)
}

func (h *TaskHandler) List(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.repo.List(r.Context())

	if err != nil {
		writeError(w, http.StatusBadRequest, "failed to list tasks")
		return
	}

	writeJSON(w, http.StatusOK, tasks)
}

func (h *TaskHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if _, err := uuid.Parse(id); err != nil {
		writeError(w, http.StatusBadRequest, "invalid task id")
		return
	}

	task, err := h.repo.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			writeError(w, http.StatusNotFound, "task not found")
			return
		}

		writeError(w, http.StatusBadRequest, "failed to get task")
		return
	}

	writeJSON(w, http.StatusOK, task)
}

func (h *TaskHandler) Complete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if _, err := uuid.Parse(id); err != nil {
		writeError(w, http.StatusBadRequest, "invalid task id")
		return
	}

	task, err := h.repo.Complete(r.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			writeError(w, http.StatusNotFound, "task not found")
			return
		}
		writeError(w, http.StatusBadRequest, "failed to complete task")
		return
	}

	err = h.publisher.PublishTaskEvent(events.TaskCompleted, task.ID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to publish task completed event")
		return
	}

	writeJSON(w, http.StatusCreated, task)
}
