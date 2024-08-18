CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL,
    email TEXT NOT NULL UNIQUE,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    PRIMARY KEY (id)
);