package main

import (
	"fmt"

	"kk-rschain.com/gin-rest-auth/config"
	"kk-rschain.com/gin-rest-auth/models"
)

func main() {
	config.Setup()
	models.Setup()
	defer models.Close()
	models.DB.AutoMigrate(
		&models.User{},
	)
	fmt.Printf("Tables are created in %s\n", config.Database.Name)
}
