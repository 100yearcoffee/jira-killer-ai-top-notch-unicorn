# 2. Docker Compose Local Environment

## Goal

Run the Go API and PostgreSQL locally with one command.

This stage makes the project easy to start and prepares it for adding more services later.

## Language

Configuration / YAML

## Technologies

- Docker
- Docker Compose
- PostgreSQL container
- Go API container
- Environment variables

## Minimal Services

```text
go-api
postgres
```

## Implementation Steps

1. Create a `Dockerfile` for the Go API.
2. Create `docker-compose.yml`.
3. Add PostgreSQL service.
4. Add Go API service.
5. Add environment variables.
6. Add a database volume.
7. Make the Go API wait/retry until PostgreSQL is ready.
8. Add a simple local run command to the README.

## Example Compose Services

```text
postgres:
  image: postgres

go-api:
  build: ./go-api
  depends_on:
    - postgres
```

## Required Environment Variables

```text
DATABASE_URL
PORT
```

## Done Criteria

- `docker compose up` starts the API and PostgreSQL.
- Go API connects to PostgreSQL inside Docker.
- You can call `/health`.
- You can create/list/complete tasks using the containerized API.
- Database data survives container restarts.

## What You Learn

- Containerized backend development
- Service networking
- Environment configuration
- Local development workflow
