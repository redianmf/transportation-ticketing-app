-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE transportation_modes (
    id INT GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(20),
    base_price INT,
    additional_price INT,
    price_calculation VARCHAR(20),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY(id)
)

-- +migrate StatementEnd