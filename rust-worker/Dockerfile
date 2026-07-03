FROM rust:1.88 AS builder

RUN apt-get update \
    && apt-get install -y protobuf-compiler \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app/rust-worker

COPY rust-worker/Cargo.toml rust-worker/Cargo.lock ./
COPY rust-worker/build.rs ./build.rs
COPY rust-worker/src ./src
COPY proto /app/proto

RUN cargo build --release

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/rust-worker/target/release/rust-worker /app/rust-worker

CMD ["/app/rust-worker"]
