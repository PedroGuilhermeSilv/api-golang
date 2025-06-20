package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/PedroGuilhermeSilv/api-golang/internal/dto"
	"github.com/PedroGuilhermeSilv/api-golang/internal/entity"
	"github.com/PedroGuilhermeSilv/api-golang/internal/infra/database"
)

type UserHandler struct {
	UserDB database.UserRepositoryInterface
}

func NewUserHandler(db database.UserRepositoryInterface) *UserHandler {
	return &UserHandler{UserDB: db}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.UserCreateInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userEntity, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.UserDB.Create(userEntity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userEntity)
}
