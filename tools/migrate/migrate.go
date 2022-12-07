package main

import (
	"kk-rschain.com/gin-rest-auth/models"
)

func main() {
	models.Setup()
	defer models.Close()
	models.DB.AutoMigrate(&models.User{})
}
