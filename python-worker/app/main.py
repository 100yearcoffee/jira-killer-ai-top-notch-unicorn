import asyncio
import os

import nats
from app.db import connect_db, get_task_by_id, update_task_summary
from app.events import parse_task_created
from app.summarizer import generate_summary


async def handle_task_created(message):
    print("Received task.created event")

    try:
        event = parse_task_created(message.data)
        print(f"Parsed task.created event for task {event.task_id}")

        task = await get_task_by_id(handle_task_created.db, event.task_id)

        if task is None:
            print(f"Task not found: {event.task_id}")
            return

        summary = generate_summary(task["description"])
        await update_task_summary(handle_task_created.db, event.task_id, summary)

        print(f"Processed task.created event for task {event.task_id}")
    except Exception as error:
        print(f"Failed to process task.created event: {error}")


async def main():
    nats_url = os.getenv("NATS_URL", "nats://localhost:4222")

    db = await connect_db()
    handle_task_created.db = db
    print("Connected to PostgreSQL")

    nc = await nats.connect(nats_url)
    print(f"Connected to NATS at {nats_url}")

    await nc.subscribe("task.created", cb=handle_task_created)
    print("Subscribed to task.created")

    try:
        await asyncio.Event().wait()
    finally:
        await nc.drain()
        await db.close()


if __name__ == "__main__":
    asyncio.run(main())
