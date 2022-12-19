package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// tokenを作成する
func CreateTokenString(sub any, expirationTime time.Duration, key string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": sub,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(expirationTime).Unix(),
	})
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}

// リクエストからtokenを抜き出す
func ExtractToken(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return authHeader, fmt.Errorf("have no Authorization Header")
	}
	authBody := strings.Split(authHeader, " ")
	if len(authBody) != 2 || authBody[0] != "Bearer" {
		return "", fmt.Errorf("invalid Authorization type")
	}
	return authBody[1], nil
}

// 暗号化されたtokenの復号をおこなう
func DecryptTokenString(tokenString string, key string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(key), nil
	})
	if err != nil {
		return token, err
	}
	if !token.Valid {
		return token, fmt.Errorf("invalid token")
	}
	return token, nil
}
