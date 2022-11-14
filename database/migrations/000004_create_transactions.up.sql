CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    date DATE NOT NULL DEFAULT CURRENT_DATE,
    product_id INT NOT NULL,
    nominal NUMERIC(12,2),
    user_id INT NOT NULL,
    type_id INT NOT NULL,
    description VARCHAR(191),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (type_id) REFERENCES types (id),
    FOREIGN KEY (product_id) REFERENCES products (id)
);