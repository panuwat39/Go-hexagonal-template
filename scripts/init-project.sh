#!/usr/bin/env bash

set -euo pipefail

if [ "$#" -lt 1 ]; then
  echo "Usage: ./scripts/init-project.sh <new-module-path> [new-project-name]"
  echo "Example: ./scripts/init-project.sh github.com/panuwat39/shop-api shop-api"
  exit 1
fi

NEW_MODULE="$1"
NEW_NAME="${2:-$(basename "$NEW_MODULE")}"

if [ ! -f "go.mod" ]; then
  echo "Error: go.mod not found. Run this script from project root."
  exit 1
fi

OLD_MODULE="$(grep '^module ' go.mod | awk '{print $2}')"
OLD_NAME="$(basename "$OLD_MODULE")"

echo "Old module: $OLD_MODULE"
echo "New module: $NEW_MODULE"
echo "Old name:   $OLD_NAME"
echo "New name:   $NEW_NAME"

go mod edit -module "$NEW_MODULE"

find . \
  -type f \
  ! -path "./.git/*" \
  ! -path "./bin/*" \
  ! -path "./scripts/init-project.sh" \
  -print0 |
  while IFS= read -r -d '' file; do
    perl -0pi -e "s|\Q$OLD_MODULE\E|$NEW_MODULE|g" "$file"
    perl -0pi -e "s|\Q$OLD_NAME\E|$NEW_NAME|g" "$file"
  done

go mod tidy

echo "Project initialized successfully"
