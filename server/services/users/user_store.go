package users

import "database/sql"

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
