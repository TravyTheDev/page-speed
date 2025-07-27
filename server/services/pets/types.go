package pets

type Pet struct {
	ID     int
	Name   string
	Aninal string
	UserID int
}

type PetFavoriteFood struct {
	Food  string
	PetID int
}

type PetWithFavoriteFood struct {
	Name         string
	Animal       string
	FavoriteFood string
}
