-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id bigserial PRIMARY KEY,
    first_name text NOT NULL,
    last_name text NOT NULL,
    email citext UNIQUE NOT NULL,
    password_hash text NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
