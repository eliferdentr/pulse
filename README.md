# Pulse 
Concurrent job processor built with Go using goroutines, channels, and worker pools.

Pulse is a small but production-style backend service that demonstrates how to build
a cancellable background job system in Go.

The goal of this project is to showcase:
- Go concurrency primitives
- worker pool architecture
- context cancellation
- thread-safe in-memory storage
- structured logging
- REST API design with Gin

---

## Architecture
```
Client
   |
   v
HTTP API (Gin)
   |
   v
Manager (worker pool)
   |
   v
Workers (goroutines)
   |
   v
Store (thread-safe job state storage)
```
---

## Features

- Worker pool implementation
- Buffered job queue
- Context-based job cancellation
- Progress tracking
- Thread-safe job store (RWMutex)
- Structured logging
- REST API

---

## API

### Create Job
POST /jobs

Example request:
```json
{
  "steps": 5,
  "sleepMs": 500,
  "timeoutMs": 0
}
```
Response:
```json
202 Accepted
{
  "id": "job-id",
  "status": "queued"
}
```

### Get Job Status
GET /jobs/:id


### Cancel Job
POST /jobs/:id/cancel

## Running the Project

### From The Terminal
```bash
go run ./cmd/pulse
```

### Using Docker
Build image:
```bash
docker build -t pulse .
```
Run container:
```bash
docker run -p 8080:8080 pulse
```
Server will be available at:
http://localhost:8080

## Example curl
### Create Job:
```bash
curl -X POST http://localhost:8080/jobs \
-H "Content-Type: application/json" \
-d '{"steps":5,"sleepMs":500,"timeoutMs":0}'
```
### Cancel job::
```bash
curl -X POST http://localhost:8080/jobs \
-H "Content-Type: application/json" \
-d '{"steps":5,"sleepMs":500,"timeoutMs":0}'
```

## Project Structure

cmd/pulse
internal/
  api/
  jobs/
  logger/

## Learning Goals
This project focuses on understanding:

- goroutines
- channels
- select
- context cancellation
- worker pools
- synchronization with mutexes
- API + background job coordination
