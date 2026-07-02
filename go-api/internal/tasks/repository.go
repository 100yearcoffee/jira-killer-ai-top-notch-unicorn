package tasks

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, title string, description string) (*Task, error) {
	now := time.Now().UTC()

	task := &Task{
		ID:          uuid.NewString(),
		Title:       title,
		Description: description,
		Status:      "open",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	query := `
		INSERT INTO tasks (id, title, description, status, ai_summary, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		task.ID,
		task.Title,
		task.Description,
		task.Status,
		task.AISummary,
		task.CreatedAt,
		task.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r *Repository) List(ctx context.Context) ([]Task, error) {
	query := `
		SELECT id, title, description, status, ai_summary, created_at, updated_at
		FROM tasks
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var task Task

		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.AISummary,
			&task.CreatedAt,
			&task.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return tasks, nil
}

func (r *Repository) GetByID(ctx context.Context, id string) (*Task, error) {
	query := `
		SELECT id, title, description, status, ai_summary, created_at, updated_at
		FROM tasks WHERE id = $1
	`

	var task Task

	err := r.db.QueryRow(ctx, query, id).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.AISummary,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *Repository) Complete(ctx context.Context, id string) (*Task, error) {
	query := `
		UPDATE tasks
		SET
			status = 'completed',
			updated_at = $2
		WHERE id = $1
		RETURNING id, title, description, status, ai_summary, created_at, updated_at
	`

	var task Task

	err := r.db.QueryRow(ctx, query, id, time.Now().UTC()).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.AISummary,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &task, nil
}
