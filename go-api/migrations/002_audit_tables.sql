CREATE TABLE IF NOT EXISTS tasks_stats (
    id INT PRIMARY KEY,
    total_tasks INT NOT NULL,
    completed_tasks INT NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);

INSERT INTO tasks_stats (id, total_tasks, completed_tasks, updated_at)
VALUES (1, 0, 0, NOW())
ON CONFLICT (id) DO NOTHING;

CREATE TABLE IF NOT EXISTS audit_logs (
    id UUID PRIMARY KEY,
    event_id UUID NOT NULL UNIQUE,
    event_type TEXT NOT NULL,
    task_id UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL
);
