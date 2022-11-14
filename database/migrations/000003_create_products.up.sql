CREATE TABLE products (
    id SERIAL PRIMARY KEY, 
    name VARCHAR(191) NOT NULL UNIQUE,
    description VARCHAR(191),
    stock INT DEFAULT 0,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);