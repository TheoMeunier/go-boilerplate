## 📘 Boilerplate Go

A minimal boilerplate to quickly start an API in Go with:

- Postgres
- Migrations
- Router HTTP
- Logger
- Utils / Helpers
- Authentification JWT

## 🚀 Installation

1. Configure `.env` file
2. Run `make setup`

## ▶️ Start project

- `make run` or `go run cmd/api/main.go`

Start migrations:
- `make migrate`
- `make migration name="create_users_table"`


## 📂 Structure (résumé)
```
cmd/               → Entry points (API servers, migrations, CLI tools)

internal/
  app/             → Application setup and initialization
  domain/          → Business logic layer (entities, services, use cases)
  infra/
    db/            → Database connection, repositories, and migrations
    http/          → HTTP layer: router, handlers, middlewares

pkg/               → Shared packages/utilities: logger, error handling, helpers
```

🎯 Objective

Provide a clean, simple, and scalable foundation to get started with a Go project quickly, without imposing too much structure.

📄 Licence
MIT