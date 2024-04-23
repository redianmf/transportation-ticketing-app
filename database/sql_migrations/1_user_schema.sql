-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE users (
    id INT GENERATED ALWAYS AS IDENTITY,
    username VARCHAR(50),
    email VARCHAR(30),
    password VARCHAR(50),
    birth_date DATE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY(id)
)

-- +migrate StatementEnd