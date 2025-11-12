-- Assuming 'users' table already exists with PK 'user_id'

-- 1. Create Orders Table (1:N with users)
CREATE TABLE IF NOT EXISTS orders (
    -- CHANGE: Using SERIAL for auto-increment in PostgreSQL
    order_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,  -- FK: The customer who placed the order
    order_date TIMESTAMP WITHOUT TIME ZONE NOT NULL, -- CHANGED DATETIME to TIMESTAMP
    status VARCHAR(50) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT
);

-- 2. Create Products Table
CREATE TABLE IF NOT EXISTS products (
    -- CHANGE: Using SERIAL for auto-increment in PostgreSQL
    product_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    price NUMERIC(10, 2) NOT NULL -- CHANGED DECIMAL to NUMERIC
);

-- 3. Create Order_Product Junction Table (N:N Resolver)
CREATE TABLE IF NOT EXISTS order_product (
    order_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL DEFAULT 1,
    unit_price NUMERIC(10, 2) NOT NULL,
    
    PRIMARY KEY (order_id, product_id),
    
    FOREIGN KEY (order_id) REFERENCES orders(order_id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE RESTRICT
);
