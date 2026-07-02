package http

import (
	"encoding/json"
	"go-api/internal/tasks"
	"net/http"
	"strings"
)

type TaskHandler struct {
	repo *tasks.Repository
}

func NewTaskHandler(repo *tasks.Repository) *TaskHandler {
	return &TaskHandler{repo: repo}
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

	task, err := h.repo.GetByID(r.Context(), id)
	if err != nil {
		writeError(w, http.StatusBadRequest, "cannot get the task by id")
		return
	}

	writeJSON(w, http.StatusOK, task)
}

func (h *TaskHandler) Complete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	task, err := h.repo.Complete(r.Context(), id)
	if err != nil {
		writeError(w, http.StatusBadRequest, "could not update the task to completed")
		return
	}

	writeJSON(w, http.StatusCreated, task)
}
