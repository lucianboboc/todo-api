# Include the variables from the .env file
include .env

.PHONY: start
start:
	@echo "Starting server on port $(PORT)"
	@trap 'exit 0' SIGINT; \
	go run ./cmd/api

.PHONY: new-migration
new-migration:
	@echo "Usage: 'make new-migration name=<migration name>'"
	@echo "Creating migration file for ${name}"
	@goose postgres ${DB_URL} create ${name} sql -dir=migrations

.PHONY: migration-up
migration-up:
	@echo "Usage: make migration-up"
	@echo "Apply new migrations..."
	@goose postgres ${DB_URL} up -dir=migrations

.PHONY: migration-down
migration-down:
	@echo "Usage: make migration-down"
	@echo "Revert migration..."
	@goose postgres ${DB_URL} down -dir=migrations

.PHONY: migration-status
migration-status:
	@echo "Usage: make migration-status"
	@echo "Checking migration status..."
	@goose postgres ${DB_URL} status -dir=migrations