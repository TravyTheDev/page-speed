-- +goose Up
-- +goose StatementBegin
CREATE TABLE pets (
    id integer primary key autoincrement,
    name varchar(255),
    animal varchar(255),
    user_id integer references users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE pets;
-- +goose StatementEnd
