package testutils

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SetupTestDB creates a test database and returns the DB instance
// It accepts any number of models to migrate
//
// Example usage:
//
//	db := SetupTestDB(t, &entity.User{}, &entity.Product{})
//	db := SetupTestDB(t, &entity.User{}) // Single model
//	db := SetupTestDB(t) // No models to migrate
func SetupTestDB(t *testing.T, models ...interface{}) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}

	if len(models) > 0 {
		err = db.AutoMigrate(models...)
		if err != nil {
			t.Fatalf("failed to migrate database: %v", err)
		}
	}

	return db
}

// SetupTestDBWithModels is a convenience function that creates a test DB and migrates specific models
// It returns both the DB and a slice of the models for convenience
func SetupTestDBWithModels[T any](t *testing.T, models ...interface{}) (*gorm.DB, T) {
	db := SetupTestDB(t, models...)
	var zero T
	return db, zero
}
