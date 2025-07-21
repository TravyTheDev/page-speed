package pets

import (
	"database/sql"

	"github.com/brianvoe/gofakeit/v7"
)

func generateFakePet() Pet {
	return Pet{
		Name:   gofakeit.LastName(),
		Aminal: gofakeit.Animal(),
	}
}

func generatePetFavoriteFood(petID int) PetFavoriteFood {
	return PetFavoriteFood{
		Food:  gofakeit.Snack(),
		PetID: petID,
	}
}

func SeedPets(db *sql.DB) error {
	count := 10
	for i := 1; i <= 10; i++ {
		fakePet := generateFakePet()
		stmt := `INSERT INTO pets (name, animal) VALUES (?, ?)`

		_, err := db.Exec(stmt, fakePet.Name, fakePet.Aminal)
		if err != nil {
			return err
		}
		seedFavoriteFood(db, count)
		count--
	}

	return nil
}

func seedFavoriteFood(db *sql.DB, count int) error {
	//could directly use count as the pet_id, doesn't matter because this isn't a serious project
	favoriteFood := generatePetFavoriteFood(count)
	stmt := `INSERT INTO pets_favorite_food (food, pet_id) VALUES (?, ?)`

	_, err := db.Exec(stmt, favoriteFood.Food, favoriteFood.PetID)
	if err != nil {
		return err
	}
	return nil
}
