import json
from uuid import UUID

from pydantic import BaseModel


class TaskCreatedEvent(BaseModel):
    task_id: UUID


def parse_task_created(data: bytes) -> TaskCreatedEvent:
    payload = json.loads(data.decode())
    return TaskCreatedEvent(**payload)
