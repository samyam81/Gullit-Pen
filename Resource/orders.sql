CREATE DATABASE ruud_pen;

-- Create a table for storing customer orders
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    phone_number VARCHAR(20),
    billing_address TEXT,
    email VARCHAR(100),
    product_quantity INTEGER,
    product_color VARCHAR(50),
    payment_method VARCHAR(50)
);
