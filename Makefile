.PHONY: migrate-up migrate-down migrate-fresh seed run build tidy

# Database connection settings (can be overridden in environment)
POSTGRES_USER ?= postgres
POSTGRES_PASSWORD ?= postgres
POSTGRES_DB ?= app_db
POSTGRES_HOST ?= localhost
POSTGRES_PORT ?= 5432

MIGRATE_BIN ?= migrate

MIGRATE_DB_URL = postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable

## Run migrations up
migrate-up:
	@echo "-> migrate up: $(MIGRATE_DB_URL)"
	$(MIGRATE_BIN) -path ./migrations -database "$(MIGRATE_DB_URL)" up

## Run migrations down
migrate-down:
	@echo "-> migrate down: $(MIGRATE_DB_URL)"
	$(MIGRATE_BIN) -path ./migrations -database "$(MIGRATE_DB_URL)" down

## Drop and re-run all migrations (fresh)
migrate-fresh:
	@echo "-> migrate fresh (drop + up): $(MIGRATE_DB_URL)"
	-$(MIGRATE_BIN) -path ./migrations -database "$(MIGRATE_DB_URL)" drop -f || true
	$(MAKE) migrate-up

## Run DB seed (uses local Go seeder at cmd/seed)
seed:
	@echo "-> running seeder"
	go run ./cmd/seed

## Run the app with Air (live reload). Requires `air` in PATH.
run:
	air

## Build the server binary
build:
	go build -o bin/server ./cmd/server

## Tidy go.mod
tidy:
	go mod tidy
