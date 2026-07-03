package events

import (
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

const (
	TaskCreated   = "task.created"
	TaskCompleted = "task.completed"
)

type Publisher struct {
	nats *nats.Conn
}

func NewPublisher(natsConn *nats.Conn) *Publisher {
	return &Publisher{nats: natsConn}
}

type TaskEvent struct {
	EventID   string    `json:"event_id"`
	Type      string    `json:"type"`
	TaskID    string    `json:"task_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *Publisher) PublishTaskEvent(eventType string, taskID string) error {
	event := TaskEvent{
		EventID:   uuid.NewString(),
		Type:      eventType,
		TaskID:    taskID,
		CreatedAt: time.Now().UTC(),
	}

	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = p.nats.Publish(eventType, data)
	if err != nil {
		log.Printf("failed to publish event type=%s task_id=%s error=%v", eventType, taskID, err)
		return err
	}

	log.Printf("published event type=%s event_id=%s task_id=%s", event.Type, event.EventID, event.TaskID)
	return nil
}
