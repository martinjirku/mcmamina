#!/bin/sh

# Set default values if not provided
DB_USER=${POSTGRES_USER:-user}
DB_PASS=${POSTGRES_PASSWORD:-password}
DB_NAME=${POSTGRES_DB:-dbname}
DB_HOST=${DB_HOST:-pg}
DB_PORT=${DB_PORT:-5432}

# Run migrations
./migrate -path /migrations -database "postgres://$DB_USER:$DB_PASS@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable" up
