# 1. Go API + PostgreSQL

## Goal

Build the smallest useful backend API for tasks.

This stage creates the core service that users interact with. Keep it simple: no auth, no projects, no complex permissions.

## Language

Go

## Technologies

- Go
- REST
- HTTP/JSON
- PostgreSQL
- SQL migrations
- pgx or sqlc
- Basic unit/integration tests

## Minimal Features

- Create task
- List tasks
- Get task by ID
- Complete task
- Store everything in PostgreSQL

## API Endpoints

```text
POST   /tasks
GET    /tasks
GET    /tasks/{id}
PATCH  /tasks/{id}/complete
GET    /health
```

## Database Tables

### tasks

```sql
id UUID PRIMARY KEY
title TEXT NOT NULL
description TEXT NOT NULL
status TEXT NOT NULL
ai_summary TEXT
created_at TIMESTAMPTZ NOT NULL
updated_at TIMESTAMPTZ NOT NULL
```

Allowed statuses:

```text
open
completed
```

## Implementation Steps

1. Create Go project structure.
2. Add PostgreSQL connection.
3. Add database migration for `tasks`.
4. Create task model/repository.
5. Implement REST handlers.
6. Add validation.
7. Add error responses.
8. Add basic tests.
9. Add `/health` endpoint.
10. Add README instructions for running the API locally.

## Suggested Folder Structure

```text
go-api/
  cmd/api/main.go
  internal/http/
  internal/tasks/
  internal/db/
  migrations/
  go.mod
```

## Done Criteria

- API starts locally.
- API connects to PostgreSQL.
- You can create a task.
- You can list tasks.
- You can complete a task.
- Data persists after API restart.
- Basic tests pass.

## What You Learn

- Go API structure
- REST design
- PostgreSQL usage
- SQL migrations
- Error handling
- Basic backend service design
