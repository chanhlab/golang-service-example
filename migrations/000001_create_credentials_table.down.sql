-- Drop trigger first
DROP TRIGGER IF EXISTS update_credentials_updated_at ON credentials;

-- Drop the function
DROP FUNCTION IF EXISTS update_updated_at_column();

-- Drop indexes
DROP INDEX IF EXISTS idx_credentials_status;
DROP INDEX IF EXISTS idx_credentials_key;

-- Drop the table
DROP TABLE IF EXISTS credentials;