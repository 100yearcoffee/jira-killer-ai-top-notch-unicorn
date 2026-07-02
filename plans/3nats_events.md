# 3. NATS Events

## Goal

Add asynchronous event publishing from the Go API.

When a task is created or completed, the API should publish an event to NATS.

## Language

Go + configuration

## Technologies

- NATS
- Event-driven architecture
- JSON events
- Event IDs
- Basic idempotency thinking

## Minimal Events

```text
task.created
task.completed
```

## Event Shape

```json
{
  "event_id": "uuid",
  "type": "task.created",
  "task_id": "uuid",
  "created_at": "2026-01-01T12:00:00Z"
}
```

## Implementation Steps

1. Add NATS service to Docker Compose.
2. Add NATS client to Go API.
3. Create event publisher package.
4. Publish `task.created` after task creation.
5. Publish `task.completed` after task completion.
6. Add event IDs.
7. Log event publishing success/failure.
8. Add a simple NATS CLI/manual subscriber for testing.
9. Document event names and schemas.

## Important Rule

For this minimal version, keep it simple:

```text
database write first
then publish event
```

Later you can learn the transactional outbox pattern.

## Done Criteria

- NATS starts in Docker Compose.
- Creating a task publishes `task.created`.
- Completing a task publishes `task.completed`.
- You can observe events using a subscriber.
- Event payloads include event ID, event type, task ID, and timestamp.

## What You Learn

- Queues/event buses
- Async communication
- Event schemas
- Service decoupling
- Eventual consistency basics
