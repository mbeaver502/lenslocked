-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id serial PRIMARY KEY,
    email text UNIQUE NOT NULL,
    password_hash text NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd