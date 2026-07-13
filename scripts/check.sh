#!/usr/bin/env bash

set -euo pipefail

TMP_DIR="$(mktemp -d)"
trap 'rm -rf "$TMP_DIR"' EXIT

echo "==> Checking go.mod / go.sum"

cp go.mod "$TMP_DIR/go.mod.before"

if [ -f go.sum ]; then
  cp go.sum "$TMP_DIR/go.sum.before"
else
  touch "$TMP_DIR/go.sum.before"
fi

go mod tidy

cp go.mod "$TMP_DIR/go.mod.after"

if [ -f go.sum ]; then
  cp go.sum "$TMP_DIR/go.sum.after"
else
  touch "$TMP_DIR/go.sum.after"
fi

MOD_CHANGED=0

if ! diff -u "$TMP_DIR/go.mod.before" "$TMP_DIR/go.mod.after" >/dev/null; then
  MOD_CHANGED=1
  echo "go.mod changed after running go mod tidy:"
  diff -u "$TMP_DIR/go.mod.before" "$TMP_DIR/go.mod.after" || true
fi

if ! diff -u "$TMP_DIR/go.sum.before" "$TMP_DIR/go.sum.after" >/dev/null; then
  MOD_CHANGED=1
  echo "go.sum changed after running go mod tidy:"
  diff -u "$TMP_DIR/go.sum.before" "$TMP_DIR/go.sum.after" || true
fi

if [ "$MOD_CHANGED" -eq 1 ]; then
  echo "Please run: go mod tidy"
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
