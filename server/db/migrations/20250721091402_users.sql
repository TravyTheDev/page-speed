-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id integer primary key autoincrement,
    username varchar(255),
    email varchar(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (email)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
