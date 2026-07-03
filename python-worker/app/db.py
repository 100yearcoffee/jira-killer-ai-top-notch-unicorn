import os
from uuid import UUID

import asyncpg


async def connect_db():
    database_url = os.getenv(
        "DATABASE_URL", "postgres://tasks_user:tasks_password@postgres:5432/tasks"
    )

    return await asyncpg.connect(database_url)


async def get_task_by_id(db, task_id: UUID):
    return await db.fetchrow(
        """
        SELECT id, description
        FROM tasks
        WHERE id = $1
        """,
        task_id,
    )


async def update_task_summary(db, task_id: UUID, summary: str):
    await db.execute(
        """
        UPDATE tasks
        SET ai_summary = $2
        WHERE id = $1
        """,
        task_id,
        summary,
    )
