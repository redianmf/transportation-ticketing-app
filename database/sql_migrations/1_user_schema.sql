-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE users (
    id INT GENERATED ALWAYS AS IDENTITY,
    username VARCHAR(50),
    email VARCHAR(50),
    password VARCHAR(255),
    birth_date DATE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY(id)
)

-- +migrate StatementEnd