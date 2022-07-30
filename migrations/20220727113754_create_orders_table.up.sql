CREATE TABLE IF NOT EXISTS orders
(
    number VARCHAR(255) NOT NULL PRIMARY KEY,
    user_id UUID NOT NULL CONSTRAINT orders_users_fk REFERENCES users(id),
    status INTEGER NOT NULL,
    accrual NUMERIC(6,2) NOT NULL,
    uploaded_at  TIMESTAMP NOT NULL,
    processed_at TIMESTAMP
);