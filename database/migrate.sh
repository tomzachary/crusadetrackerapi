#!/usr/bin/env bash
set -euo pipefail

# ---------------------------------------------------------------------------
# PostgreSQL migration runner
#
# Usage:
#   ./migrate.sh [OPTIONS]
#
# Options:
#   -h, --host      DB host        (default: localhost)
#   -p, --port      DB port        (default: 5432)
#   -U, --user      DB user        (default: postgres)
#   -W, --password  DB password    (default: empty / uses PGPASSWORD env var)
#   -d, --dbname    DB name        (required)
#   -m, --migrations-dir  Directory containing .sql files  (default: ./sql)
#   --help          Show this help
#
# Environment variables can supply defaults for any option:
#   PGHOST, PGPORT, PGUSER, PGPASSWORD, PGDATABASE
# ---------------------------------------------------------------------------

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Defaults (fall back to standard PG env vars when set)
DB_HOST="${PGHOST:-localhost}"
DB_PORT="${PGPORT:-5432}"
DB_USER="${PGUSER:-postgres}"
DB_PASSWORD="${PGPASSWORD:-}"
DB_NAME="${PGDATABASE:-}"
MIGRATIONS_DIR="${SCRIPT_DIR}/sql"

usage() {
    grep '^#' "$0" | grep -v '#!/' | sed 's/^# \{0,1\}//'
    exit 0
}

# ── Argument parsing ────────────────────────────────────────────────────────
while [[ $# -gt 0 ]]; do
    case "$1" in
        -h|--host)        DB_HOST="$2";         shift 2 ;;
        -p|--port)        DB_PORT="$2";         shift 2 ;;
        -U|--user)        DB_USER="$2";         shift 2 ;;
        -W|--password)    DB_PASSWORD="$2";     shift 2 ;;
        -d|--dbname)      DB_NAME="$2";         shift 2 ;;
        -m|--migrations-dir) MIGRATIONS_DIR="$2"; shift 2 ;;
        --help)           usage ;;
        *) echo "Unknown option: $1" >&2; usage ;;
    esac
done

if [[ -z "$DB_NAME" ]]; then
    echo "ERROR: database name is required (-d / --dbname or PGDATABASE)" >&2
    exit 1
fi

if [[ ! -d "$MIGRATIONS_DIR" ]]; then
    echo "ERROR: migrations directory not found: $MIGRATIONS_DIR" >&2
    exit 1
fi

# Export password so psql picks it up without a prompt
export PGPASSWORD="$DB_PASSWORD"

# ── Helpers ─────────────────────────────────────────────────────────────────
psql_cmd() {
    psql \
        --host="$DB_HOST" \
        --port="$DB_PORT" \
        --username="$DB_USER" \
        --dbname="$DB_NAME" \
        --no-password \
        "$@"
}

psql_query() {
    psql_cmd --tuples-only --no-align --command="$1"
}

# ── Ensure migrations tracking table exists ──────────────────────────────────
echo "Connecting to ${DB_USER}@${DB_HOST}:${DB_PORT}/${DB_NAME} ..."

psql_cmd --command="
CREATE TABLE IF NOT EXISTS schema_migrations (
    id          SERIAL PRIMARY KEY,
    filename    TEXT        NOT NULL UNIQUE,
    applied_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    checksum    TEXT        NOT NULL
);" > /dev/null

# ── Collect and sort migration files ────────────────────────────────────────
mapfile -t SQL_FILES < <(find "$MIGRATIONS_DIR" -maxdepth 1 -name '*.sql' | sort)

if [[ ${#SQL_FILES[@]} -eq 0 ]]; then
    echo "No .sql migration files found in ${MIGRATIONS_DIR}."
    exit 0
fi

# ── Run pending migrations ───────────────────────────────────────────────────
applied=0
skipped=0

for filepath in "${SQL_FILES[@]}"; do
    filename="$(basename "$filepath")"

    # Compute a checksum of the file so we can detect drift
    checksum="$(sha256sum "$filepath" | awk '{print $1}')"

    # Check if already applied
    row="$(psql_query "SELECT checksum FROM schema_migrations WHERE filename = '$filename';" 2>/dev/null || true)"

    if [[ -n "$row" ]]; then
        stored_checksum="$(echo "$row" | tr -d '[:space:]')"
        if [[ "$stored_checksum" != "$checksum" ]]; then
            echo "WARNING: checksum mismatch for already-applied migration '${filename}'" >&2
            echo "         stored=${stored_checksum}" >&2
            echo "         current=${checksum}" >&2
            echo "         Skipping to avoid re-running a modified migration." >&2
        else
            echo "  [skip]  ${filename}"
        fi
        (( skipped++ )) || true
        continue
    fi

    echo "  [apply] ${filename}"

    # Run migration inside a transaction so failures leave the DB clean
    psql_cmd --single-transaction --file="$filepath"

    # Record successful application
    psql_cmd --command="
INSERT INTO schema_migrations (filename, checksum)
VALUES ('${filename}', '${checksum}');" > /dev/null

    (( applied++ )) || true
done

echo ""
echo "Done. Applied: ${applied}  Skipped (already run): ${skipped}"
