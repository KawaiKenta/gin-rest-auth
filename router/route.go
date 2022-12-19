package router

import (
	"github.com/gin-gonic/gin"
	"kk-rschain.com/gin-rest-auth/controllers"
	"kk-rschain.com/gin-rest-auth/middleware"
)

func InitRoute() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.AddCorsHeader)
	v1 := router.Group("/v1", middleware.RequireAuth)
	{
		// NOTE: requireAuth以下においては userId, ok := c.Get("userId") が使用できる
		v1.GET("/user/info", controllers.GetUserInfo)
		v1.DELETE("/user/delete", controllers.DeleteUser)
		v1.POST("/user/update", controllers.UpdateUser)
	}
	auth := router.Group("/auth")
	{
		auth.POST("/signin", controllers.SignIn)
		auth.POST("/signup", controllers.SignUp)
	}
	return router
}
