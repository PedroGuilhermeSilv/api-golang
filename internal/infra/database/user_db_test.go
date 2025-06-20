package database

import (
	"testing"

	"github.com/PedroGuilhermeSilv/api-golang/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupTestDB creates a test database and returns the DB instance and UserDB
func setupTestDB(t *testing.T) (*gorm.DB, *UserDB) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}

	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	userDb := NewUserDB(db)
	return db, userDb
}

func TestUserDBCreateSuccess(t *testing.T) {
	db, userDb := setupTestDB(t)

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
	_, userDb := setupTestDB(t)

	user, _ := entity.NewUser("John Doe", "john.doe@example.com", "123456")

	err := userDb.Create(user)
	assert.Nil(t, err)

	userFound, err := userDb.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
}
