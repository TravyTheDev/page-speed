package users

import (
	"database/sql"
	"strings"
)

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (s *UserStore) GetUsers() ([]*User, error) {
	users := make([]*User, 0)
	stmt := `SELECT id, username, email FROM USERS`
	rows, err := s.db.Query(stmt)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		u, err := scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (s *UserStore) SearchUser(text string) ([]*UserWithPet, error) {
	search := "%" + strings.ReplaceAll(text, " ", "%") + "%"
	stmt := `
    SELECT
		users.id,
        users.username,
        pets.name,
        pets.animal,
        pets_favorite_food.food
    FROM users
    JOIN pets ON pets.user_id = users.id
    JOIN pets_favorite_food ON pets_favorite_food.pet_id = pets.id
    WHERE users.username LIKE ?
`
	rows, err := s.db.Query(stmt, search)
	if err != nil {
		return nil, err
	}
	var results []*UserWithPet

	for rows.Next() {
		uwp := new(UserWithPet)
		err := rows.Scan(
			&uwp.ID,
			&uwp.UserName,
			&uwp.Pet.Name,
			&uwp.Pet.Animal,
			&uwp.Pet.FavoriteFood,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, uwp)
	}

	return results, nil
}

func scanRowsIntoUser(rows *sql.Rows) (*User, error) {
	user := new(User)
	err := rows.Scan(
		&user.ID,
		&user.UserName,
		&user.Email,
	)

	if err != nil {
		return nil, err
	}
	return user, nil
}
