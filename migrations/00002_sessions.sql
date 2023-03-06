-- +goose Up
-- +goose StatementBegin
CREATE TABLE sessions (
    id serial PRIMARY KEY,
    user_id int UNIQUE NOT NULL,
    token_hash text UNIQUE NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE sessions;
-- +goose StatementEnd