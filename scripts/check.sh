#!/usr/bin/env bash

set -euo pipefail

echo "==> Checking go.mod / go.sum"
go mod tidy

if ! git diff --exit-code -- go.mod go.sum >/dev/null; then
  echo "go mod tidy changed go.mod or go.sum"
  echo "Please run: go mod tidy"
  git diff -- go.mod go.sum
  exit 1
fi

echo "==> Checking gofmt"
UNFORMATTED_FILES="$(gofmt -l .)"

if [ -n "$UNFORMATTED_FILES" ]; then
  echo "The following files are not formatted:"
  echo "$UNFORMATTED_FILES"
  echo "Please run: make fmt"
  exit 1
fi

echo "==> Running go vet"
go vet ./...

echo "==> Running tests"
go test -race -cover ./...

echo "==> Building all packages"
go build ./...

echo "All checks passed"
