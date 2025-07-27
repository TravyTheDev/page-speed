package users

import (
	"database/sql"

	"github.com/brianvoe/gofakeit/v7"
)

func generateFakeUser() User {
	return User{
		UserName: gofakeit.Name(),
		Email:    gofakeit.Email(),
	}
}

func SeedUsers(db *sql.DB) error {
	for i := 1; i <= 500; i++ {
		fakeUser := generateFakeUser()
		stmt := `INSERT INTO users (username, email) VALUES (?, ?)`

		_, err := db.Exec(stmt, fakeUser.UserName, fakeUser.Email)
		if err != nil {
			return err
		}
	}
	return nil
}
