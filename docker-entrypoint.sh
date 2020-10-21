#!/bin/bash

# Wait for MySQL server to come up
until mysql -u${DB_USERNAME} -p${DB_PASSWORD} -h${DB_HOST} -P${DB_PORT} -e "SELECT 1"; do sleep 1; done

# Start server
exec "$@"
