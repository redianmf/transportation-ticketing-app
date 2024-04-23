-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE wallets (
    id INT GENERATED ALWAYS AS IDENTITY,
    user_id INT,
    amount INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY(id),
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
            REFERENCES users(id)
)

-- +migrate StatementEnd