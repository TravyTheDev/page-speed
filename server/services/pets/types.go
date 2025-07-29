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
	Name         string `json:"name"`
	Animal       string `json:"animal"`
	FavoriteFood string `json:"favoriteFood"`
}
