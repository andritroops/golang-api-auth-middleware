CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(191) NOT NULL,
    email VARCHAR(191) NOT NULL UNIQUE,
    phone_number VARCHAR(13) NOT NULL UNIQUE,
    password VARCHAR(191),
    google_id VARCHAR(191),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);