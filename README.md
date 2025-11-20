# be-test-logkar

Go backend scaffold with Fiber, fx (Uber), GORM, Postgres, Redis, golang-migrate, and Air. Uses Clean Architecture structure.

Quick start

1. Copy `.env.example` to `.env` and edit if needed.

2. Start DB and Redis:

```bash
docker-compose up -d
```

3. Install Air and migrate CLI (if not installed):

```bash
# Air
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
# migrate CLI (https://github.com/golang-migrate/migrate)
# use brew or download from releases
brew install golang-migrate
```

4. Run migrations:

```bash
migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5432/app_db?sslmode=disable" up
```

5. Run the app (live reload):

```bash
air
```

Project layout

- `cmd/server` — application entrypoint
- `internal/config` — env/config loader
- `internal/db` — GORM setup
- `internal/redis` — redis client provider
- `internal/server` — Fiber app provider
- `internal/user` — example package (model, repo, service, handler)

Additional docs

- `docs/postman_collection.json` — Postman collection (v2.1) for endpoints (health, users, products, customers, transactions).
- `docs/ERD.svg` — ERD diagram showing `products`, `customers`, `transactions`, and `users` and their relationships.

Change the module path in `go.mod` if desired.
