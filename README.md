# URL Shortener in Go

A simple REST API to shorten URLs using Go, Fiber, GORM, and PostgreSQL.

## Endpoints
- `POST /api/shorten` – Accepts a long URL
- `GET /:shortId` – Redirects to original URL

## Setup
1. Rename `.env.example` to `.env` and fill your PostgreSQL credentials.
2. Run:
```sh
go mod tidy
go run main.go
```