#!/bin/bash

APP_NAME="tfcli"

echo "🚀 Building $APP_NAME ..."

# macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o dist/${APP_NAME}-darwin-arm64-v1.2.0 .

# macOS (Intel)
GOOS=darwin GOARCH=amd64 go build -o dist/${APP_NAME}-darwin-amd64-v1.2.0 .

# Linux (common servers)
GOOS=linux GOARCH=amd64 go build -o dist/${APP_NAME}-linux-amd64-v1.2.0 .

echo "✅ Done. Files are in ./dist"