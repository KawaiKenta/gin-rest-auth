package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

func RequireAuth(c *gin.Context) {
	// header 取り出し
	authorization := c.GetHeader("Authorization")
	if authorization == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "have no token"})
		c.Abort()
		return
	}
	authBody := strings.Split(authorization, " ")
	if len(authBody) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "have invalid auth header"})
		c.Abort()
		return
	}
	tokenString := authBody[1]

	// JWT_SECRET 取り出し
	if err := godotenv.Load(".env"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		c.Abort()
		return
	}
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "jwt_key is not exist"})
		c.Abort()
		return
	}

	// 復号
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(secretKey), nil
	})
	// tokenの復号方式 & exp チェック
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": err.Error()})
		c.Abort()
		return
	}
	// tokenの検証
	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "invalid token"})
		c.Abort()
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	userId := int(claims["sub"].(float64))
	// データベースへの問い合わせを毎回行うぐらいならセッションを使用したほうが良い
	// オーバーヘッドとして 70µs程度払う必要がある
	// if _, err := models.GetUserById(userId); err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"msg": "invalid user"})
	// 	c.Abort()
	// 	return
	// }

	// ユーザーidをセット
	c.Set("userId", userId)
	c.Next()
}
