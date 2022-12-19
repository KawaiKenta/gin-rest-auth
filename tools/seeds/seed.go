package main

import (
	"strconv"

	"golang.org/x/crypto/bcrypt"
	"kk-rschain.com/gin-rest-auth/config"
	"kk-rschain.com/gin-rest-auth/models"
)

func main() {
	config.Setup()
	models.Setup()
	defer models.Close()
	if err := userSeeds(); err != nil {
		panic(err)
	}
	println("users are created")
}

func userSeeds() error {
	for i := 0; i < 10; i++ {
		hash, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		user := models.User{
			Name:       "user" + strconv.Itoa(i+1),
			Email:      "sample" + strconv.Itoa(i+1) + "@gmail.com",
			Password:   string(hash),
			IsVerified: false,
		}

		if err := models.DB.Create(&user).Error; err != nil {
			return err
		}
	}
	return nil
}
