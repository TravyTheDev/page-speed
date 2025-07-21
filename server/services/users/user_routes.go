package users

import (
	"encoding/json"
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
