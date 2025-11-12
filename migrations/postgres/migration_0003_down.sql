-- DOWN Migration 1 SQL

-- 1. Drop the Order_Product Junction Table first (to avoid FK dependency issues)
DROP TABLE IF EXISTS order_product;

-- 2. Drop the Orders Table
DROP TABLE IF EXISTS orders;

-- 3. Drop the Products Table
DROP TABLE IF EXISTS products;
