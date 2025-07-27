package pets

import (
	"database/sql"
	"fmt"
	"page-speed-server/utility"
	"strings"
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
		pet, err := scanRowsIntoPet(rows)
		if err != nil {
			return nil, err
		}
		pets = append(pets, pet)
	}
	return pets, nil
}

func (s *PetStore) GetFavoriteFoodFromPetID(id int) (*PetFavoriteFood, error) {
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

func (s *PetStore) GetFavoriteFoodFromPetIDs(ids []int) ([]*PetFavoriteFood, error) {
	if len(ids) == 0 {
		return []*PetFavoriteFood{}, nil
	}
	placeHolders, args := utility.GenerateWhereInIDPlaceHolders(ids)
	petFavoriteFoods := make([]*PetFavoriteFood, 0)
	stmt := fmt.Sprintf(`SELECT food, pet_id FROM pets_favorite_food WHERE pet_id IN (%s)`, strings.Join(placeHolders, ","))
	rows, err := s.db.Query(stmt, args...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		petFavoriteFood := new(PetFavoriteFood)
		err = rows.Scan(
			&petFavoriteFood.Food,
			&petFavoriteFood.PetID,
		)
		if err != nil {
			return nil, err
		}
		petFavoriteFoods = append(petFavoriteFoods, petFavoriteFood)
	}
	return petFavoriteFoods, nil
}

func (s *PetStore) GetPetFromUserID(id int) (*Pet, error) {
	pet := new(Pet)
	stmt := `SELECT id, name, animal, user_id FROM pets WHERE user_id = ?`
	rows, err := s.db.Query(stmt, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		pet, err = scanRowsIntoPet(rows)
		if err != nil {
			return nil, err
		}
	}
	return pet, nil
}

func (s *PetStore) GetPetFromUserIDs(ids []int) ([]*Pet, error) {
	if len(ids) == 0 {
		return []*Pet{}, nil
	}
	placeHolders, args := utility.GenerateWhereInIDPlaceHolders(ids)
	pets := make([]*Pet, 0)
	stmt := fmt.Sprintf(`SELECT id, name, animal, user_id FROM pets WHERE user_id IN (%s)`, strings.Join(placeHolders, ","))
	rows, err := s.db.Query(stmt, args...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		pet, err := scanRowsIntoPet(rows)
		if err != nil {
			return nil, err
		}
		pets = append(pets, pet)
	}
	return pets, nil
}

func scanRowsIntoPet(rows *sql.Rows) (*Pet, error) {
	Pet := new(Pet)
	err := rows.Scan(
		&Pet.ID,
		&Pet.Name,
		&Pet.Aninal,
		&Pet.UserID,
	)

	if err != nil {
		return nil, err
	}
	return Pet, nil
}
