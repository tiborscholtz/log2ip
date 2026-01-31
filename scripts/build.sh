#!/usr/bin/env sh
set -e

APP_NAME="log2ip"
CMD_PATH="./cmd/log2ip"
OUT_DIR="./bin"

echo "▶ Building $APP_NAME..."

# Ensure output directory exists
mkdir -p "$OUT_DIR"

# Clean previous build
rm -f "$OUT_DIR/$APP_NAME"

# Build (static, optimized)
CGO_ENABLED=0 GOOS=$(go env GOOS) GOARCH=$(go env GOARCH) \
  go build -trimpath -ldflags="-s -w" \
  -o "$OUT_DIR/$APP_NAME" \
  "$CMD_PATH"

echo "✅ Build complete: $OUT_DIR/$APP_NAME"
