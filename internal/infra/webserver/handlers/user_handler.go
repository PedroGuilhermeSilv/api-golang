package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/PedroGuilhermeSilv/api-golang/internal/dto"
	"github.com/PedroGuilhermeSilv/api-golang/internal/entity"
	"github.com/PedroGuilhermeSilv/api-golang/internal/infra/database"
	"github.com/go-chi/jwtauth"
)

type UserHandler struct {
	UserDB       database.UserRepositoryInterface
	Jwt          *jwtauth.JWTAuth
	JwtExpiresIn int
}

func NewUserHandler(db database.UserRepositoryInterface, jwt *jwtauth.JWTAuth, jwtExpiresIn int) *UserHandler {
	return &UserHandler{UserDB: db, Jwt: jwt, JwtExpiresIn: jwtExpiresIn}
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user dto.UserLoginInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userEntity, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !userEntity.ValidatePassword(user.Password) {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}
	mapClaims := map[string]interface{}{
		"sub": userEntity.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtExpiresIn)).Unix(),
	}
	_, tokenString, _ := h.Jwt.Encode(mapClaims)
	acessToken := struct {
		AccessToken string `json:"access_token"`
	}{AccessToken: tokenString}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(acessToken)
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
