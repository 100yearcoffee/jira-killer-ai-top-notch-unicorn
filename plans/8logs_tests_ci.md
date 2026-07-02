# 8. Logs, Tests, and CI

## Goal

Make the small system easier to debug and safer to change.

This stage adds lightweight production habits without overengineering.

## Languages

Go + Python + Rust + YAML

## Technologies

- Structured logs
- Request IDs
- Basic metrics
- Go tests
- pytest
- cargo test
- GitHub Actions
- Docker Compose test workflow

## Minimal Observability

Add logs for:

```text
API request started/completed
task created
task completed
event published
event consumed
worker success
worker failure
gRPC request success/failure
rate limit blocked
```

Add a request ID:

```text
X-Request-ID
```

If the client does not send one, generate one.

## Minimal Tests

### Go

- Task creation handler
- Task completion handler
- Repository tests
- Rate limit middleware test

### Python

- Summary function test
- Event parsing test

### Rust

- Event parsing test
- Stats update logic test
- gRPC response mapping test

## Minimal CI Pipeline

GitHub Actions should run:

```text
go test ./...
pytest
cargo test
docker compose config
```

Optional later:

```text
docker compose up -d
run integration tests
build Docker images
```

## Implementation Steps

1. Add structured logging to Go API.
2. Add structured logging to Python worker.
3. Add structured logging to Rust services.
4. Add request ID middleware to Go API.
5. Pass request/event correlation ID in event payloads.
6. Add minimal tests in each language.
7. Create GitHub Actions workflow.
8. Make CI run on pull requests and pushes.
9. Add status badge to README.
10. Document how to run tests locally.

## Done Criteria

- Logs are readable and structured.
- Request IDs appear in API logs.
- Event IDs appear in worker logs.
- Tests pass locally.
- GitHub Actions runs tests automatically.
- README explains how to run the system and tests.

## What You Learn

- Debugging distributed systems
- Structured logs
- Test discipline
- CI/CD basics
- Multi-language repository maintenance
