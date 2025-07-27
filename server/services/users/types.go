package users

import "page-speed-server/services/pets"

type User struct {
	ID       int
	UserName string
	Email    string
}

type UserWithPet struct {
	UserName string
	Pet      pets.PetWithFavoriteFood
}
