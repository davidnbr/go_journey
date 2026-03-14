# Project 01: REST API

## Overview
A full CRUD REST API with middleware, JSON serialization, and tests using `httptest`.

## Run Tests
```bash
go test ./phases/phase4-mastery/projects/01-rest-api/
```

## Skills Applied
- `net/http` routing with `http.NewServeMux()`
- JSON encode/decode
- `sync.RWMutex` for concurrent store
- Middleware: Logger, Recovery
- HTTP status codes (200, 201, 204, 404)
- `httptest` for request/response testing

## Routes to Implement in Handler()
| Method | Path | Status | Body |
|--------|------|--------|------|
| GET | /items | 200 | JSON array |
| POST | /items | 201 | created item |
| GET | /items/{id} | 200/404 | item or error |
| DELETE | /items/{id} | 204/404 | empty |

## Implementation Order
1. GET /items -> list
2. POST /items -> decode body, create, return 201
3. GET/DELETE /items/{id} -> parse ID, lookup, respond
