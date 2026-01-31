#!/usr/bin/env sh
set -e

APP_NAME="log2ip"
CMD_PATH="./cmd/log2ip"

echo "▶ Starting $APP_NAME..."

# Check Go installation
if ! command -v go >/dev/null 2>&1; then
  echo "❌ Go is not installed or not in PATH"
  exit 1
fi

# Ensure dependencies are clean
go mod tidy

# Run the app
clear && go run "$CMD_PATH"
