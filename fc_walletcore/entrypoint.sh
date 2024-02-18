#!/bin/sh
set -e

# Run the main.go file
go run fc_walletcore/cmd/walletcore/main.go

# Execute the CMD from the Dockerfile
exec "$@"
