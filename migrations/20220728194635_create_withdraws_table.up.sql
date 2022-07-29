CREATE TABLE IF NOT EXISTS withdraws
(
    id UUID NOT NULL PRIMARY KEY,
    order_number VARCHAR(255) NOT NULL CONSTRAINT withdraws_orders_fk REFERENCES orders(number),
    user_id UUID NOT NULL CONSTRAINT withdraws_users_fk REFERENCES users(id),
    sum INTEGER NOT NULL,
    processed_at TIMESTAMP  NOT NULL
);