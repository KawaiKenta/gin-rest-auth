package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddCorsHeader(c *gin.Context) {
	// TODO:Access-Control-Allow-Originを適切に指定する必要がある
	// header 情報の付加
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Max-Age", "86400")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
	// preflightリクエスト用
	if c.Request.Method == "OPTIONS" {
		c.Status(http.StatusNoContent)
		c.Abort()
		return
	}
	c.Next()
}
