DO $$
BEGIN
    -- Check if 'username' column *does not* exist in the 'users' table
    IF NOT EXISTS (
        SELECT 1
        FROM information_schema.columns
        WHERE table_name = 'users'
          AND column_name = 'username'
    ) THEN
        -- If 'username' doesn't exist, perform the rename
        ALTER TABLE users
        RENAME COLUMN name TO username;
    END IF;
END
$$;
