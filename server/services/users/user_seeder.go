package users

import (
	"database/sql"

	"github.com/brianvoe/gofakeit/v7"
)

func generateFakeUser() User {
	return User{
		Name:  gofakeit.Name(),
		Email: gofakeit.Email(),
	}
}

func SeedUsers(db *sql.DB) error {

	for i := 0; i < 10; i++ {
		fakeUser := generateFakeUser()
		stmt := `INSERT INTO users (username, email) VALUES (?, ?)`

		_, err := db.Exec(stmt, fakeUser.Name, fakeUser.Email)
		if err != nil {
			return err
		}
	}

	return nil
}
