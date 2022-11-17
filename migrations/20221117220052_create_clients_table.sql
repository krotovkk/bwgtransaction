-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS clients (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS identity,
    balance FLOAT NOT NULL CONSTRAINT  price_chk CHECK ( balance >= 0 )
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS clients;
-- +goose StatementEnd
