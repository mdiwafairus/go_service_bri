# Default values (bisa override pakai: make migrate-up DB_NAME=testdb)
DB_USER ?= postgres
DB_PASS ?= password
DB_HOST ?= localhost
DB_PORT ?= 5432
DB_NAME ?= mydb

# Path ke migration
MIGRATE_CMD = go run cmd/migrate/main.go

# Migration Commands
migrate-up:
	@echo "Running migrations UP..."
	DB_USER=$(DB_USER) DB_PASS=$(DB_PASS) DB_HOST=$(DB_HOST) DB_PORT=$(DB_PORT) DB_NAME=$(DB_NAME) $(MIGRATE_CMD) -action=up

migrate-down:
	@echo "Rolling back migration (1 step)..."
	DB_USER=$(DB_USER) DB_PASS=$(DB_PASS) DB_HOST=$(DB_HOST) DB_PORT=$(DB_PORT) DB_NAME=$(DB_NAME) $(MIGRATE_CMD) -action=down

migrate-drop:
	@echo "Dropping all tables..."
	DB_USER=$(DB_USER) DB_PASS=$(DB_PASS) DB_HOST=$(DB_HOST) DB_PORT=$(DB_PORT) DB_NAME=$(DB_NAME) $(MIGRATE_CMD) -action=drop
