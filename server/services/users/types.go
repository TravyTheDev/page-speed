package users

import "page-speed-server/services/pets"

type User struct {
	ID       int
	UserName string
	Email    string
}

type UserWithPet struct {
	ID       int                      `json:"id"`
	UserName string                   `json:"username"`
	Pet      pets.PetWithFavoriteFood `json:"pet"`
}
