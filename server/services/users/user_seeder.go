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
	countUp := 1
	countDown := 10
	for i := 1; i <= 10; i++ {
		fakeUser := generateFakeUser()
		stmt := `INSERT INTO users (username, email) VALUES (?, ?)`

		_, err := db.Exec(stmt, fakeUser.Name, fakeUser.Email)
		if err != nil {
			return err
		}
		if err := seedUserPet(db, countUp, countDown); err != nil {
			return err
		}
		countUp++
		countDown--
	}
	return nil
}

func seedUserPet(db *sql.DB, countUp int, countDown int) error {
	stmt := `INSERT INTO users_pets (user_id, pet_id) VALUES (?, ?)`

	_, err := db.Exec(stmt, countUp, countDown)
	if err != nil {
		return err
	}
	return nil
}
