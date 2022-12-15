-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(256) NOT NULL,
    password VARCHAR(256) NOT NULL,
    name VARCHAR(256),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
)

-- +migrate StatementEnd