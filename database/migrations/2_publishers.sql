-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE publishers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(256),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
)

-- +migrate StatementEnd