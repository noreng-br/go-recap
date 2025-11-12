-- UP Migration 7: Idempotently add description to products table

DO $$ 
BEGIN
    -- Check if the description column DOES NOT EXIST
    IF NOT EXISTS (
        SELECT 1 
        FROM information_schema.columns 
        WHERE table_name='products' AND column_name='description'
    ) THEN
        -- Add the description column as TEXT and NULLable
        ALTER TABLE products
            ADD COLUMN description TEXT NULL;
    END IF;
END $$;
