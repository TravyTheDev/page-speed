-- +goose Up
-- +goose StatementBegin
CREATE TABLE pets_favorite_food (
    id integer primary key autoincrement,
    food varchar(255),
    pet_id integer references pets(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE pets_favorite_food;
-- +goose StatementEnd