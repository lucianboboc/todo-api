-- +goose Up
-- +goose StatementBegin
CREATE TABLE todos (
    id bigserial PRIMARY KEY,
    text text NOT NULL,
    completed bool NOT NULL DEFAULT false,
    user_id int,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE todos;
-- +goose StatementEnd
