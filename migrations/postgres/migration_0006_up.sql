-- UP Migration X: Add date_delivered to orders table
-- UP Migration 6: Idempotent changes to orders table

-- 1. Conditionally Add 'delivered_date' (Only if it DOES NOT EXIST)
DO $$ 
BEGIN
    -- Check if the delivered_date column DOES NOT EXIST
    IF NOT EXISTS (
        SELECT 1 
        FROM information_schema.columns 
        WHERE table_name='orders' AND column_name='delivered_date'
    ) THEN
        -- Add the new column (It's a delivery date, so likely NULLable)
        ALTER TABLE orders
            ADD COLUMN delivered_date TIMESTAMP WITHOUT TIME ZONE NULL;
    END IF;
END $$;


-- 2. Conditionally Rename 'order_date' to 'ordered_date' (Only if 'order_date' still exists)
DO $$ 
BEGIN
    -- Check if the old column name 'order_date' EXISTS
    IF EXISTS (
        SELECT 1 
        FROM information_schema.columns 
        WHERE table_name='orders' AND column_name='order_date'
    ) THEN
        -- Rename the column
        ALTER TABLE orders
            RENAME COLUMN order_date TO ordered_date;
    END IF;
END $$;


-- 3. Increase the size of the 'status' column
-- This ALTER TYPE command will typically run safely even if the size is already 100 
-- or greater, as it's an enlargement, but it is not conditionalized to keep it simple.
ALTER TABLE orders
    ALTER COLUMN status TYPE VARCHAR(100);
