-- +goose Up
-- +goose StatementBegin
CREATE TABLE users_pets (
    id integer primary key autoincrement,
    user_id integer references users(id),
    pet_id integer references pets(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users_pets;
-- +goose StatementEnd
