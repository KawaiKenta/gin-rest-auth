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
	models.DB.Migrator().DropTable(
		&models.User{},
	)
	fmt.Printf("All tables in %s is droped\n", config.Database.Name)
}
