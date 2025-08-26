#!/bin/bash

# Script to create new migration files
# Usage: ./scripts/create-migration.sh <migration_name>

if [ $# -eq 0 ]; then
    echo "Usage: $0 <migration_name>"
    echo "Example: $0 add_user_table"
    exit 1
fi

MIGRATION_NAME="$1"
TIMESTAMP=$(date +%s)

# Get the latest migration number
LATEST_MIGRATION=$(ls migrations/ 2>/dev/null | grep -E '^[0-9]+_.*\.up\.sql$' | sort -n | tail -1 | cut -d'_' -f1)

if [ -z "$LATEST_MIGRATION" ]; then
    # No existing migrations
    MIGRATION_NUMBER="000001"
else
    # Increment the latest migration number
    MIGRATION_NUMBER=$(printf "%06d" $((10#$LATEST_MIGRATION + 1)))
fi

UP_FILE="migrations/${MIGRATION_NUMBER}_${MIGRATION_NAME}.up.sql"
DOWN_FILE="migrations/${MIGRATION_NUMBER}_${MIGRATION_NAME}.down.sql"

# Create the up migration file
cat > "$UP_FILE" << EOF
-- Migration: $MIGRATION_NAME
-- Created: $(date)

-- Add your up migration SQL here
EOF

# Create the down migration file  
cat > "$DOWN_FILE" << EOF
-- Rollback migration: $MIGRATION_NAME
-- Created: $(date)

-- Add your down migration SQL here
EOF

echo "Created migration files:"
echo "  Up:   $UP_FILE"
echo "  Down: $DOWN_FILE"