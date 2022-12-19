package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// リクエストからint型にしたユーザーidを取得する。
func GetUserId(c *gin.Context) (int, error) {
	val, ok := c.Get("userId")
	if !ok {
		return 0, fmt.Errorf("userId key is not set")
	}
	userId, ok := val.(int)
	if !ok {
		return 0, fmt.Errorf("failed to parse int")
	}
	return userId, nil
}
