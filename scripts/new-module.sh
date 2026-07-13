#!/usr/bin/env bash

set -euo pipefail

if [ "$#" -ne 1 ]; then
  echo "Usage: ./scripts/new-module.sh <module-name>"
  echo "Example: ./scripts/new-module.sh user"
  exit 1
fi

MODULE="$1"

if [[ ! "$MODULE" =~ ^[a-z][a-z0-9]*$ ]]; then
  echo "Error: module name must be lowercase alphanumeric and start with a letter."
  echo "Valid examples: user, order, payment, inventory"
  exit 1
fi

BASE="internal/modules/$MODULE"

if [ -d "$BASE" ]; then
  echo "Error: module already exists: $BASE"
  exit 1
fi

mkdir -p "$BASE/domain/entity"
mkdir -p "$BASE/domain/valueobject"
mkdir -p "$BASE/domain/repository"
mkdir -p "$BASE/domain/service"

mkdir -p "$BASE/application/command"
mkdir -p "$BASE/application/query"
mkdir -p "$BASE/application/usecase"
mkdir -p "$BASE/application/port"

mkdir -p "$BASE/adapter/inbound/http"
mkdir -p "$BASE/adapter/inbound/consumer"

mkdir -p "$BASE/adapter/outbound/persistence"
mkdir -p "$BASE/adapter/outbound/external"

cat > "$BASE/module.go" <<MODULE_EOF
package $MODULE
MODULE_EOF

find "$BASE" -type d -empty -exec touch "{}/.gitkeep" \;

echo "Module created: $BASE"
