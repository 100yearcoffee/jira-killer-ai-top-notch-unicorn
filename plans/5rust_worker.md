# 5. Rust Event Processor

## Goal

Create a Rust worker that consumes task events and writes audit/activity data.

This gives you Rust experience without building a complicated Rust web service yet.

## Language

Rust

## Technologies

- Rust
- Tokio
- NATS client
- PostgreSQL
- sqlx
- Serde
- Docker

## Minimal Features

Consume:

```text
task.created
task.completed
```

Write to:

```text
audit_logs
task_stats
```

## Database Tables

### audit_logs

```sql
id UUID PRIMARY KEY
event_id UUID NOT NULL UNIQUE
event_type TEXT NOT NULL
task_id UUID NOT NULL
created_at TIMESTAMPTZ NOT NULL
```

### task_stats

```sql
id INT PRIMARY KEY
total_tasks INT NOT NULL
completed_tasks INT NOT NULL
updated_at TIMESTAMPTZ NOT NULL
```

Use one row in `task_stats`, for example `id = 1`.

## Implementation Steps

1. Create `rust-worker/` project.
2. Add Rust dependencies.
3. Connect to NATS.
4. Subscribe to task events.
5. Parse event payloads with Serde.
6. Connect to PostgreSQL.
7. Insert audit log for each unique event.
8. Update stats counters.
9. Add duplicate protection using `event_id UNIQUE`.
10. Add Rust worker to Docker Compose.

## Done Criteria

- Rust worker starts in Docker Compose.
- Worker receives `task.created` and `task.completed`.
- Worker writes audit logs.
- Worker updates stats.
- Duplicate event IDs do not double-count stats.
- Worker logs bad events without crashing.

## What You Learn

- Rust service basics
- Tokio async runtime
- Event processing
- Idempotency
- SQL from Rust
- Safe, typed backend code
