-- Create Categories Table
CREATE TABLE IF NOT EXISTS categories (
    -- CHANGE: Using SERIAL for auto-increment in PostgreSQL
    category_id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);
