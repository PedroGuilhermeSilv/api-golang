package database

import "github.com/PedroGuilhermeSilv/api-golang/internal/entity"

type UserRepositoryInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
