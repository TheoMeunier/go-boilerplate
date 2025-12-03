# ===============================
# CONFIGURATION
# ===============================

APP_NAME=go-uptime
MIGRATE_CMD=go run cmd/migrate/main.go

# ===============================
# MIGRATIONS
# ===============================

# Create migration :
#   make migration name=create_users_table
migration:
	@if [ -z "$(name)" ]; then \
		echo "Usage: make migration name=create_xxx"; \
		exit 1; \
	fi
	go run cmd/migration/create_migration.go new $(name)


migrate:
	go run cmd/migration/migration.go

# Roollback last migration :
rollback:
	$(MIGRATE_CMD) down

# ===============================
# APP
# ===============================

# Lancer l'app :
#   make run
run:
	go run main.go