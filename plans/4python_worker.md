# 4. Python Worker

## Goal

Create a Python worker that consumes task events and updates the database.

For the first version, do not use real AI. Generate a fake summary from the task description.

## Language

Python

## Technologies

- Python
- NATS client
- PostgreSQL client
- Pydantic
- Worker process
- Docker

## Minimal Behavior

When `task.created` is received:

1. Load the task from PostgreSQL.
2. Generate a simple summary.
3. Save the summary to `tasks.ai_summary`.

Example summary:

```text
Summary: first 100 characters of the task description
```

## Implementation Steps

1. Create `python-worker/` service.
2. Add Python dependencies.
3. Connect to NATS.
4. Subscribe to `task.created`.
5. Connect to PostgreSQL.
6. Fetch task by ID.
7. Generate simple summary.
8. Update `tasks.ai_summary`.
9. Add worker service to Docker Compose.
10. Add logs for received/processed/failed events.

## Suggested Folder Structure

```text
python-worker/
  app/main.py
  app/db.py
  app/events.py
  app/summarizer.py
  requirements.txt
  Dockerfile
```

## Done Criteria

- Python worker starts in Docker Compose.
- Worker receives `task.created`.
- Worker updates the task with `ai_summary`.
- If the task does not exist, the worker logs an error and continues.
- Worker does not crash on bad event payloads.

## What You Learn

- Background workers
- Queue consumers
- Async processing
- Python backend service structure
- Event-driven database updates
