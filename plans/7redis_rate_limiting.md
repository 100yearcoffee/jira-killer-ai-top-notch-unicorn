# 7. Redis Rate Limiting

## Goal

Use Redis to rate-limit task creation.

This gives you practical Redis experience without adding unnecessary complexity.

## Language

Go

## Technologies

- Redis
- TTLs
- Atomic counters
- Middleware
- Rate limiting
- Docker Compose

## Minimal Feature

Limit:

```text
POST /tasks
```

Rule:

```text
max 10 task creations per minute
```

For the no-auth version, rate-limit by IP address or a fake user ID header.

Example header:

```text
X-User-ID: demo-user
```

## Redis Key Example

```text
rate_limit:create_task:demo-user:2026-01-01T12:00
```

Simpler option:

```text
rate_limit:create_task:demo-user
```

with a 60-second TTL.

## Implementation Steps

1. Add Redis to Docker Compose.
2. Add Redis client to Go API.
3. Create rate-limiting middleware.
4. On `POST /tasks`, increment Redis counter.
5. Set TTL if key is new.
6. If counter exceeds limit, return `429 Too Many Requests`.
7. Add response headers if desired.
8. Add tests for allowed/blocked requests.
9. Document the rate limit behavior.

## Done Criteria

- Redis starts in Docker Compose.
- `POST /tasks` is rate-limited.
- After 10 requests/minute, API returns 429.
- Limit resets after TTL expires.
- Redis failure behavior is defined.

## Recommended Redis Failure Behavior

For learning, choose fail-open:

```text
If Redis is down, allow the request but log the error.
```

## What You Learn

- Redis commands
- TTLs
- Counters
- Middleware
- Rate limiting
- Production tradeoffs
