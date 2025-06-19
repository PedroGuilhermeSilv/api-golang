package entity

import (
	"testing"

	"github.com/PedroGuilhermeSilv/api-golang/pkg/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "john.doe@example.com", "password")
	if err != nil {
		t.Fatalf("Error creating user: %v", err)
	}
	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.NotEmpty(t, user.ID)
	assert.IsType(t, entity.ID{}, user.ID)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john.doe@example.com", user.Email)
	assert.NotEmpty(t, user.Password)

}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe", "john.doe@example.com", "password")
	if err != nil {
		t.Fatalf("Error creating user: %v", err)
	}
	assert.True(t, user.ValidatePassword("password"))
	assert.False(t, user.ValidatePassword("wrong_password"))
}
