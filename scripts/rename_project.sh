#!/usr/bin/env sh
set -eu

usage() {
  cat <<'EOF'
Usage: ./scripts/rename_project.sh --project-name <name> --module-name <module> [--frontend-package-name <name>] [--dry-run]
EOF
}

PROJECT_NAME=""
MODULE_NAME=""
FRONTEND_PACKAGE_NAME=""
DRY_RUN=0

while [ "$#" -gt 0 ]; do
  case "$1" in
    --project-name)
      PROJECT_NAME="$2"
      shift 2
      ;;
    --module-name)
      MODULE_NAME="$2"
      shift 2
      ;;
    --frontend-package-name)
      FRONTEND_PACKAGE_NAME="$2"
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

if [ -z "$PROJECT_NAME" ] || [ -z "$MODULE_NAME" ]; then
  usage >&2
  exit 1
fi

if [ -z "$FRONTEND_PACKAGE_NAME" ]; then
  FRONTEND_PACKAGE_NAME="${PROJECT_NAME}-frontend"
fi

SCRIPT_DIR=$(CDPATH= cd -- "$(dirname "$0")" && pwd)
REPO_ROOT=$(CDPATH= cd -- "$SCRIPT_DIR/.." && pwd)

GO_MOD_PATH="$REPO_ROOT/go.mod"
PACKAGE_JSON_PATH="$REPO_ROOT/frontend/package.json"
CONFIG_PATH="$REPO_ROOT/internal/config/config.go"

CURRENT_MODULE_NAME=$(sed -n 's/^module[[:space:]]\+//p' "$GO_MOD_PATH" | head -n 1)
CURRENT_FRONTEND_PACKAGE_NAME=$(node -p "JSON.parse(require('fs').readFileSync(process.argv[1], 'utf8')).name" "$PACKAGE_JSON_PATH")
CURRENT_PROJECT_NAME=$(sed -n 's/.*getEnv("SERVICE_NAME", "\([^"]*\)").*/\1/p' "$CONFIG_PATH" | head -n 1)

if [ -z "$CURRENT_MODULE_NAME" ] || [ -z "$CURRENT_FRONTEND_PACKAGE_NAME" ] || [ -z "$CURRENT_PROJECT_NAME" ]; then
  echo "Failed to detect current template identifiers." >&2
  exit 1
fi

UPDATED_FILES=""

while IFS= read -r file; do
  original=$(cat "$file")
  updated=$(printf '%s' "$original" | node -e '
    const fs = require("fs");
    const data = fs.readFileSync(0, "utf8");
    const replacements = [
      [process.argv[1], process.argv[2]],
      [process.argv[3], process.argv[4]],
      [process.argv[5], process.argv[6]],
    ];
    let result = data;
    for (const [from, to] of replacements) {
      result = result.split(from).join(to);
    }
    process.stdout.write(result);
  ' "$CURRENT_MODULE_NAME" "$MODULE_NAME" "$CURRENT_FRONTEND_PACKAGE_NAME" "$FRONTEND_PACKAGE_NAME" "$CURRENT_PROJECT_NAME" "$PROJECT_NAME")

  if [ "$updated" != "$original" ]; then
    UPDATED_FILES="${UPDATED_FILES}${file}\n"
    if [ "$DRY_RUN" -eq 0 ]; then
      printf '%s' "$updated" > "$file"
    fi
  fi
done <<EOF
$(find "$REPO_ROOT" -type f \
  ! -path "$REPO_ROOT/.git/*" \
  ! -path "$REPO_ROOT/frontend/node_modules/*" \
  ! -path "$REPO_ROOT/frontend/dist/*" \
  ! -path "$REPO_ROOT/docs/plans/*" \
  ! -name "go.sum" \
  ! -name "pnpm-lock.yaml" \
  ! -name "rename_project.ps1" \
  ! -name "rename_project.sh")
EOF

if [ -z "$UPDATED_FILES" ]; then
  echo "No files needed updates."
  exit 0
fi

printf 'Updated files:\n'
printf '%b' "$UPDATED_FILES" | while IFS= read -r file; do
  [ -n "$file" ] && printf ' - %s\n' "$file"
done

if [ "$DRY_RUN" -eq 1 ]; then
  echo "Dry run complete. No files were changed."
fi
