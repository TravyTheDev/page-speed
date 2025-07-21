package pets

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PetHandler struct {
	petStore PetStore
}

func NewHandler(petStore PetStore) *PetHandler {
	return &PetHandler{
		petStore: petStore,
	}
}

func (h *PetHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("/get_pets_with_favorite_food_bad", h.getPetsWithFavoriteFoodBadQuery)
}

// this is an example bad query
func (h *PetHandler) getPetsWithFavoriteFoodBadQuery(w http.ResponseWriter, r *http.Request) {
	petsWithFood := make([]PetWithFavoriteFood, 0)
	pets, err := h.petStore.GetPets()
	if err != nil {
		http.Error(w, "error getting pets", http.StatusInternalServerError)
		return
	}
	for _, pet := range pets {
		favoriteFood, err := h.petStore.GetFavoriteFoodFromPetID(pet.ID)
		fmt.Println(favoriteFood)
		if err != nil {
			http.Error(w, "error getting favorite food", http.StatusInternalServerError)
		}
		petWithFood := PetWithFavoriteFood{
			Name:         pet.Name,
			Animal:       pet.Aminal,
			FavoriteFood: favoriteFood.Food,
		}
		petsWithFood = append(petsWithFood, petWithFood)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(petsWithFood); err != nil {
		http.Error(w, "error getting user", http.StatusInternalServerError)
		return
	}
}

func (h *PetHandler) getPetsWithFavoriteFoodGoodQuery() {}
