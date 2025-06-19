package main

import (
	"fmt"
	"log"

	"github.com/PedroGuilhermeSilv/api-golang/configs"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	fmt.Println(config.DbDriver)
}
