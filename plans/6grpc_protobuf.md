# 6. gRPC + Protocol Buffers

## Goal

Add one internal RPC call from the Go API to a Rust stats service.

This stage teaches typed service-to-service communication.

## Languages

Go + Rust

## Technologies

- gRPC
- Protocol Buffers
- Go gRPC client
- Rust gRPC server
- tonic
- prost
- protobuf code generation

## Minimal Feature

Expose this HTTP endpoint in Go:

```text
GET /stats
```

Internally:

```text
Go API -> gRPC -> Rust Stats Service
```

Rust returns:

```json
{
  "total_tasks": 10,
  "completed_tasks": 4
}
```

## Protobuf Service

```proto
syntax = "proto3";

package stats.v1;

service StatsService {
  rpc GetStats(GetStatsRequest) returns (GetStatsResponse);
}

message GetStatsRequest {}

message GetStatsResponse {
  int32 total_tasks = 1;
  int32 completed_tasks = 2;
}
```

## Implementation Steps

1. Create `proto/stats/v1/stats.proto`.
2. Generate Go protobuf/gRPC code.
3. Generate Rust protobuf/gRPC code.
4. Create Rust gRPC server.
5. Rust service reads stats from PostgreSQL.
6. Go API creates gRPC client.
7. Add Go `GET /stats` endpoint.
8. Go endpoint calls Rust service.
9. Add timeouts.
10. Add the Rust gRPC service to Docker Compose.

## Done Criteria

- Rust gRPC service starts.
- Go API can call Rust gRPC service.
- `GET /stats` returns stats from Rust.
- If Rust service is down, Go API returns a clean error.
- Protobuf files are committed and documented.

## What You Learn

- gRPC
- Protobuf contracts
- Internal APIs
- Typed service communication
- Timeouts and service failure handling
