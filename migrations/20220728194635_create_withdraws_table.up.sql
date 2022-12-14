CREATE TABLE IF NOT EXISTS withdraws
(
    id UUID NOT NULL PRIMARY KEY,
    order_number VARCHAR(255) NOT NULL,
    user_id UUID NOT NULL CONSTRAINT withdraws_users_fk REFERENCES users(id),
    sum NUMERIC(6,2) NOT NULL,
    processed_at TIMESTAMP  NOT NULL
);