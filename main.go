package main

import (
	"kk-rschain.com/gin-rest-auth/models"
	"kk-rschain.com/gin-rest-auth/router"
)

func init() {
	models.Setup()
}

func main() {
	r := router.InitRoute()
	r.Run("localhost:8080")
}
