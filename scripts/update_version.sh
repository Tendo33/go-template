#!/usr/bin/env sh
set -eu

usage() {
  cat <<'EOF'
Usage: ./scripts/update_version.sh --version <version> [--dry-run]
EOF
}

VERSION=""
DRY_RUN=0

while [ "$#" -gt 0 ]; do
  case "$1" in
    --version)
      VERSION="$2"
      shift 2
      ;;
    --dry-run)
      DRY_RUN=1
      shift
      ;;
    -h|--help)
      usage
      exit 0
      ;;
    *)
      echo "Unknown argument: $1" >&2
      usage >&2
      exit 1
      ;;
  esac
done

if [ -z "$VERSION" ]; then
  usage >&2
  exit 1
fi

SCRIPT_DIR=$(CDPATH= cd -- "$(dirname "$0")" && pwd)
REPO_ROOT=$(CDPATH= cd -- "$SCRIPT_DIR/.." && pwd)
VERSION_FILE_PATH="$REPO_ROOT/VERSION"
PACKAGE_JSON_PATH="$REPO_ROOT/frontend/package.json"

CURRENT_TEMPLATE_VERSION=$(tr -d '\r\n' < "$VERSION_FILE_PATH")
CURRENT_VERSION=$(node -p "JSON.parse(require('fs').readFileSync(process.argv[1], 'utf8')).version" "$PACKAGE_JSON_PATH")

if [ "$CURRENT_TEMPLATE_VERSION" = "$VERSION" ] && [ "$CURRENT_VERSION" = "$VERSION" ]; then
  echo "Version is already $VERSION."
  exit 0
fi

if [ "$CURRENT_TEMPLATE_VERSION" != "$VERSION" ]; then
  echo "VERSION: $CURRENT_TEMPLATE_VERSION -> $VERSION"
fi

if [ "$CURRENT_VERSION" != "$VERSION" ]; then
  echo "frontend/package.json: $CURRENT_VERSION -> $VERSION"
fi

if [ "$DRY_RUN" -eq 1 ]; then
  echo "Dry run complete. No files were changed."
  exit 0
fi

printf '%s\n' "$VERSION" > "$VERSION_FILE_PATH"

node -e '
  const fs = require("fs");
  const path = process.argv[1];
  const version = process.argv[2];
  const data = JSON.parse(fs.readFileSync(path, "utf8"));
  data.version = version;
  fs.writeFileSync(path, `${JSON.stringify(data, null, 2)}\n`);
' "$PACKAGE_JSON_PATH" "$VERSION"
