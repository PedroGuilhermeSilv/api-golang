package database

import (
	"testing"

	"github.com/PedroGuilhermeSilv/api-golang/internal/entity"
	"github.com/PedroGuilhermeSilv/api-golang/pkg/testutils"
	"github.com/stretchr/testify/assert"
)

func TestUserDBCreateSuccess(t *testing.T) {
	db := testutils.SetupTestDB(t, &entity.User{})
	userDb := NewUserDB(db)

	user, _ := entity.NewUser("John Doe", "john.doe@example.com", "123456")

	err := userDb.Create(user)
	assert.Nil(t, err)

	userFound := &entity.User{}
	err = db.First(&userFound, "email = ?", user.Email).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotEmpty(t, userFound.Password)
}

func TestUserDBFindByEmailSuccess(t *testing.T) {
	db := testutils.SetupTestDB(t, &entity.User{})
	userDb := NewUserDB(db)

	user, _ := entity.NewUser("John Doe", "john.doe@example.com", "123456")

	err := userDb.Create(user)
	assert.Nil(t, err)

	userFound, err := userDb.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
}
