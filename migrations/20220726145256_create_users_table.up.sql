CREATE TABLE IF NOT EXISTS users
(
    id UUID NOT NULL PRIMARY KEY,
    login VARCHAR(255)  NOT NULL,
    password VARCHAR(255)  NOT NULL,
    current NUMERIC(6,2) DEFAULT 0 NOT NULL CONSTRAINT check_positive CHECK (current >= 0)
);