CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email TEXT NULL UNIQUE,
    password_hash TEXT NULL,
    create_at TIMESTAMP NOT NULL DEFAULT NOW()
);