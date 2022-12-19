package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"kk-rschain.com/gin-rest-auth/config"
	"kk-rschain.com/gin-rest-auth/models"
	"kk-rschain.com/gin-rest-auth/router"
)

func init() {
	config.Setup()
	models.Setup()
}

func main() {
	gin.SetMode(config.Server.RunMode)
	router := router.InitRoute()

	service := http.Server{
		Addr:         fmt.Sprintf(":%d", config.Server.HttpPort),
		Handler:      router,
		WriteTimeout: config.Server.WriteTimeout,
		ReadTimeout:  config.Server.ReadTimeout,
	}
	service.ListenAndServe()
}
