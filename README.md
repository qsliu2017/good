# Good

A monolith application template.

## Quick Start

### Prerequisites

- Golang with CGO enabled
- [Bun](https://bun.sh) installed

```sh
git clone https://github.com/qsliu2017/good.git
cd good
go generate ./...
go run main.go
```

## Tech Stack

- Golang + Sqlite
  - [echo](github.com/labstack/echo) for HTTP server
- React + TypeScript
  - Tailwind CSS
  - Bun + Vite

## Principles

- DX(Developer Experience) first
- One binary (with [`go:embed`](https://pkg.go.dev/embed) and [sqlite](github.com/mattn/go-sqlite3))
