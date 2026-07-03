## Local Docker Environment

Start the API and PostgreSQL:

```sh
docker compose up --build
```
The API runs at:

- http://localhost:8080

PostgreSQL runs at:

- localhost:5432

Stop the containers:

```sh
docker compose down
```

Stop the containers and delete database data:

```sh
docker compose down -v
```

## Create schema
```sh
docker compose exec -T postgres psql -U tasks_user -d tasks < go-api/migrations/001_create_tasks.sql
```

## Task Events

The API published task lifecycle events to NATS

### `task.created`

Published after a task is created

```json
{
  "event_id": "uuid",
  "type": "task.created",
  "task_id": "uuid",
  "created_at": "2026-07-03T12:00:00Z"
}
```

### `task.completed`

Published after a task is marked completed.

```json
{
  "event_id": "uuid",
  "type": "task.completed",
  "task_id": "uuid",
  "created_at": "2026-07-03T12:00:00Z"
}
```

### Notes

- event_id is unique per published event
- type matches the NATS subject
- task_id identifies the task affected by the event
- created_at is UTS
- The database write happens before event publishing
