CREATE TABLE types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(191) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);