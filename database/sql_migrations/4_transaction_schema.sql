-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE transactions (
    id INT GENERATED ALWAYS AS IDENTITY,
    user_id INT,
    transportation_mode_id INT,
    transaction_reference VARCHAR(30),
    type VARCHAR(20),
    status VARCHAR(10),
    amount INT,
    start_point INT,
    end_point INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY(id),
    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
            REFERENCES users (id),
    CONSTRAINT fk_transportation_modes
        FOREIGN KEY (transportation_mode_id)
            REFERENCES transportation_modes (id)
)

-- +migrate StatementEnd