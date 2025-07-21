package pets

import (
	"database/sql"
	"fmt"
)

type PetStore struct {
	db *sql.DB
}

func NewPetStore(db *sql.DB) *PetStore {
	return &PetStore{
		db: db,
	}
}

func (s *PetStore) GetPets() ([]*Pet, error) {
	pets := make([]*Pet, 0)
	stmt := `SELECT id, name, animal FROM pets`
	rows, err := s.db.Query(stmt)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		u, err := scanRowsIntoPet(rows)
		if err != nil {
			return nil, err
		}
		pets = append(pets, u)
	}
	return pets, nil
}

func (s *PetStore) GetFavoriteFoodFromPetID(id int) (*PetFavoriteFood, error) {
	fmt.Println(id)
	favoriteFood := new(PetFavoriteFood)
	stmt := `SELECT food FROM pets_favorite_food WHERE pet_id = ?`
	row, err := s.db.Query(stmt, id)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		err = row.Scan(
			&favoriteFood.Food,
		)
		if err != nil {
			return nil, err
		}
	}
	return favoriteFood, nil
}

func scanRowsIntoPet(rows *sql.Rows) (*Pet, error) {
	Pet := new(Pet)
	err := rows.Scan(
		&Pet.ID,
		&Pet.Name,
		&Pet.Aminal,
	)

	if err != nil {
		return nil, err
	}
	return Pet, nil
}
