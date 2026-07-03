use anyhow::Result;
use chrono::{DateTime, Utc};
use futures_util::StreamExt;
use serde::Deserialize;
use sqlx::postgres::PgPoolOptions;
use tracing::info;
use uuid::Uuid;

#[derive(Debug, Deserialize)]
struct TaskEvent {
    event_id: Uuid,
    task_id: Uuid,
    created_at: DateTime<Utc>,
}

#[tokio::main]
async fn main() -> Result<()> {
    tracing_subscriber::fmt::init();

    let nats_url =
        std::env::var("NATS_URL").unwrap_or_else(|_| "nats://localhost:4222".to_string());

    info!("connecting to NATS at {}", nats_url);

    let client = async_nats::connect(nats_url).await?;

    info!("connected to NATS");

    let database_url = std::env::var("DATABASE_URL")
        .unwrap_or_else(|_| "postgres://tasks_user:tasks_password@postgres:5432/tasks".to_string());

    info!("connecting to PostgreSQL");

    let db = PgPoolOptions::new()
        .max_connections(5)
        .connect(&database_url)
        .await?;

    info!("connected to PostgreSQL");

    let mut subscriber = client.subscribe("task.*").await?;

    while let Some(message) = subscriber.next().await {
        let payload = String::from_utf8_lossy(&message.payload);

        let event: TaskEvent = match serde_json::from_slice(&message.payload) {
            Ok(event) => event,
            Err(error) => {
                tracing::warn!(
                    subject = %message.subject,
                    error = %error,
                    payload = %payload,
                    "failed to parse event payload"
                );
                continue;
            }
        };

        info!(
            subject = %message.subject,
            event_id = %event.event_id,
            task_id = %event.task_id,
            created_at = %event.created_at,
            "parsed task event"
        );

        let result = sqlx::query(
            r#"
            INSERT INTO audit_logs (id, event_id, event_type, task_id, created_at)
            VALUES ($1, $2, $3, $4, $5)
            "#,
        )
        .bind(Uuid::new_v4())
        .bind(event.event_id)
        .bind(message.subject.as_str())
        .bind(event.task_id)
        .bind(event.created_at)
        .execute(&db)
        .await;

        match result {
            Ok(_) => {
                info!(
                    event_id = %event.event_id,
                    "inserted audit log"
                );

                let stats_result = match message.subject.as_str() {
                    "task.created" => {
                        sqlx::query(
                            r#"
                            UPDATE tasks_stats
                            SET total_tasks = total_tasks + 1,
                                updated_at = NOW()
                            WHERE id = 1
                            "#,
                        )
                        .execute(&db)
                        .await
                    }
                    "task.completed" => {
                        sqlx::query(
                            r#"
                            UPDATE tasks_stats
                            SET completed_tasks = completed_tasks + 1,
                                updated_at = NOW()
                            WHERE id = 1
                            "#,
                        )
                        .execute(&db)
                        .await
                    }
                    _ => {
                        info!(
                            subject = %message.subject,
                            "no stats update for subject"
                        );
                        continue;
                    }
                };

                match stats_result {
                    Ok(_) => {
                        info!(
                            event_id = %event.event_id,
                            subject = %message.subject,
                            "updated task stats"
                        )
                    }
                    Err(error) => {
                        tracing::error!(
                            event_id = %event.event_id,
                            error = %error,
                            "failed to update task stats"
                        )
                    }
                }
            }
            Err(error) => {
                if let Some(db_error) = error.as_database_error() {
                    if db_error.is_unique_violation() {
                        tracing::warn!(
                            event_id = %event.event_id,
                            "duplicate event ignore"
                        );
                        continue;
                    }
                }

                tracing::error!(
                    event_id = %event.event_id,
                    error = %error,
                    "failed to insert audit log"
                );
            }
        }
    }

    Ok(())
}
