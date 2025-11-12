-- DOWN Migration 6: Idempotent reversal of schema changes

-- 1. Conditionally Reverse the status column size change (assuming original size was 50)
-- NOTE: This could lose data if values > 50 chars exist, but it's the required reversal.
ALTER TABLE orders
    ALTER COLUMN status TYPE VARCHAR(50);


-- 2. Conditionally Reverse the rename: 'ordered_date' back to 'order_date'
DO $$ 
BEGIN
    -- Check if the current column name 'ordered_date' EXISTS (to reverse the rename)
    IF EXISTS (
        SELECT 1 
        FROM information_schema.columns 
        WHERE table_name='orders' AND column_name='ordered_date'
    ) THEN
        -- Rename the column back
        ALTER TABLE orders
            RENAME COLUMN ordered_date TO order_date;
    END IF;
END $$;


-- 3. Conditionally Drop the added 'delivered_date' column
DO $$ 
BEGIN
    -- Check if the delivered_date column EXISTS
    IF EXISTS (
        SELECT 1 
        FROM information_schema.columns 
        WHERE table_name='orders' AND column_name='delivered_date'
    ) THEN
        -- Drop the column
        ALTER TABLE orders
            DROP COLUMN delivered_date;
    END IF;
END $$;
