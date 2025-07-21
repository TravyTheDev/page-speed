package users

type User struct {
	Name  string
	Email string
}

type UserPet struct {
	UserID int
	PetID  int
}
