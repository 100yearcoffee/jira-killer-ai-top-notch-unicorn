# Go Tasks API

Small REST API for tasks using Go and PostgreSQL.

## Requirements

- Go
- PostgreSQL

## Database

Create database:

```bash
docker run --name postgres-latest -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=postgres -p 5432:5432 -d postgres:latest

Run migration:

psql
"postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -f
migrations/001_create_tasks.sql

## Run API

DATABASE_URL="postgres://postgres:postgres@localhost:5432/postgres?
sslmode=disable" go run ./cmd/api

## Endpoints

GET    /health
POST   /tasks
GET    /tasks
GET    /tasks/{id}
PATCH  /tasks/{id}/complete

## Example Requests

curl -i http://localhost:8080/health

curl -i -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Learn Go","description":"Build the task API"}'

curl -i http://localhost:8080/tasks

curl -i http://localhost:8080/tasks/YOUR_TASK_ID

curl -i -X PATCH http://localhost:8080/tasks/YOUR_TASK_ID/complete
