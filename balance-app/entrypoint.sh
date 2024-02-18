#!/bin/sh
set -e

# Run the main.go file
go run balance-app/main.go

# Execute the CMD from the Dockerfile
exec "$@"
