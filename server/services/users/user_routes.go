package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"page-speed-server/services/pets"
)

type UserHandler struct {
	userStore UserStore
	petStore  pets.PetStore
}

func NewHandler(userStore UserStore, petStore pets.PetStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
		petStore:  petStore,
	}
}

func (h *UserHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("/get_users", h.getUsers)
	router.HandleFunc("/get_users_with_pets_bad", h.getUsersWithPetsBadQuery)
	router.HandleFunc("/get_users_with_pets_good", h.getUsersWithPetsGoodQuery)
	router.HandleFunc("/search_users", h.searchUser)
}

func (h *UserHandler) getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userStore.GetUsers()
	if err != nil {
		http.Error(w, "error getting user", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "error getting user", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) getUsersWithPetsBadQuery(w http.ResponseWriter, r *http.Request) {
	usersWithPets := make([]UserWithPet, 0)
	users, err := h.userStore.GetUsers()
	if err != nil {
		http.Error(w, "error getting user", http.StatusInternalServerError)
		return
	}

	for _, user := range users {
		pet, err := h.petStore.GetPetFromUserID(user.ID)
		if err != nil {
			http.Error(w, "error getting pet", http.StatusInternalServerError)
			return
		}
		petFavoriteFood, err := h.petStore.GetFavoriteFoodFromPetID(pet.ID)
		if err != nil {
			http.Error(w, "error getting favorite food", http.StatusInternalServerError)
			return
		}
		userWithPet := UserWithPet{
			UserName: user.UserName,
			Pet: pets.PetWithFavoriteFood{
				Name:         pet.Name,
				Animal:       pet.Aninal,
				FavoriteFood: petFavoriteFood.Food,
			},
		}
		usersWithPets = append(usersWithPets, userWithPet)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(usersWithPets); err != nil {
		http.Error(w, "error getting user", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) getUsersWithPetsGoodQuery(w http.ResponseWriter, r *http.Request) {
	usersWithPets := make([]UserWithPet, 0)
	userMap := make(map[int]*User)
	userPetsMap := make(map[int][]*pets.Pet)
	users, err := h.userStore.GetUsers()
	if err != nil {
		http.Error(w, "error getting user", http.StatusInternalServerError)
		return
	}
	userIDs := make([]int, 0, len(users))
	for _, user := range users {
		userMap[user.ID] = user
		userIDs = append(userIDs, user.ID)
	}
	userPets, err := h.petStore.GetPetFromUserIDs(userIDs)
	if err != nil {
		http.Error(w, "error getting pets", http.StatusInternalServerError)
		return
	}
	petIDs := make([]int, 0, len(userPets))
	for _, pet := range userPets {
		userPetsMap[pet.UserID] = append(userPetsMap[pet.UserID], pet)
		petIDs = append(petIDs, pet.ID)
	}
	petFavoriteFoods, err := h.petStore.GetFavoriteFoodFromPetIDs(petIDs)
	if err != nil {
		http.Error(w, "error getting favorite foods", http.StatusInternalServerError)
		return
	}
	petFoodMap := make(map[int]string)
	for _, pf := range petFavoriteFoods {
		petFoodMap[pf.PetID] = pf.Food
	}
	for userID, userPets := range userPetsMap {
		user := userMap[userID]
		for _, pet := range userPets {
			userWithPet := UserWithPet{
				UserName: user.UserName,
				Pet: pets.PetWithFavoriteFood{
					Name:         pet.Name,
					Animal:       pet.Aninal,
					FavoriteFood: petFoodMap[pet.ID],
				},
			}
			usersWithPets = append(usersWithPets, userWithPet)
		}
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(usersWithPets); err != nil {
		http.Error(w, "error getting user", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) searchUser(w http.ResponseWriter, r *http.Request) {
	userName := r.URL.Query().Get("userName")
	usersWithPets, err := h.userStore.SearchUser(userName)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "error getting users", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(usersWithPets); err != nil {
		http.Error(w, "error getting user", http.StatusInternalServerError)
		return
	}
}
