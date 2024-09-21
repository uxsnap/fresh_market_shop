#!/bin/bash
source .env

export MIGRATION_DSN="host=pg port=5432 dbname=$PG_DB user=$PG_USER password=$PG_PASSWORD sslmode=disable"

goose -dir "${MIGRATION_DIR}" postgres "${MIGRATION_DSN}" up -v