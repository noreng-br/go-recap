-- DOWN Migration 7: Idempotently drop description from products table

DO $$ 
BEGIN
    -- Check if the description column EXISTS
    IF EXISTS (
        SELECT 1 
        FROM information_schema.columns 
        WHERE table_name='products' AND column_name='description'
    ) THEN
        -- Drop the column
        ALTER TABLE products
            DROP COLUMN description;
    END IF;
END $$;
