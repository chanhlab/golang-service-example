# Database Migrations

This directory contains database migrations using [golang-migrate](https://github.com/golang-migrate/migrate).

## Migration Files

Migration files follow the pattern: `{version}_{description}.{direction}.sql`

- `version`: Sequential version number (e.g., 000001, 000002, etc.)
- `description`: Brief description of the migration
- `direction`: Either `up` or `down`

Examples:
- `000001_create_credentials_table.up.sql`
- `000001_create_credentials_table.down.sql`

## Commands

### Run all pending migrations
```bash
./main migration up
```

### Rollback the last migration
```bash
./main migration down
```

### Show current migration version
```bash
./main migration version
```

### Create a new migration
```bash
./scripts/create-migration.sh <migration_name>
```

Example:
```bash
./scripts/create-migration.sh add_user_table
```

This creates:
- `000002_add_user_table.up.sql`
- `000002_add_user_table.down.sql`

## Configuration

Migration settings are configured through environment variables or config file:

```yaml
database:
  host: localhost
  port: 5432
  db_name: your_database
  username: your_username
  password: your_password
```

## Best Practices

1. **Always create both up and down migrations**
2. **Keep migrations atomic** - each migration should do one logical thing
3. **Test your migrations** on a copy of production data
4. **Never edit existing migration files** - create a new migration instead
5. **Use transactions** where appropriate
6. **Add indexes in separate migrations** for large tables

## Troubleshooting

### Dirty Migration State
If a migration fails and leaves the database in a "dirty" state:

```bash
# Check the current state
./main migration version

# Fix manually, then force version
# Note: Be very careful with this!
```

### Rolling Back Multiple Migrations
Currently, the down command only rolls back one migration at a time. To rollback multiple:

```bash
./main migration down  # Roll back once
./main migration down  # Roll back again
# etc.
```