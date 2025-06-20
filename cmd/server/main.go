package main

import (
	"log"
	"net/http"

	"github.com/PedroGuilhermeSilv/api-golang/configs"
	"github.com/PedroGuilhermeSilv/api-golang/internal/entity"
	"github.com/PedroGuilhermeSilv/api-golang/internal/infra/database"
	"github.com/PedroGuilhermeSilv/api-golang/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProductDB(db)
	productHandler := handlers.NewProductHandler(productDB)
	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8080", nil)
}
