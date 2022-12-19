package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"kk-rschain.com/gin-rest-auth/config"
	"kk-rschain.com/gin-rest-auth/utils"
)

func RequireAuth(c *gin.Context) {
	// tokenString 取り出し
	tokenString, err := utils.ExtractToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": err.Error()})
		c.Abort()
		return
	}

	// 復号
	token, err := utils.DecryptTokenString(tokenString, config.Jwt.Key)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": err.Error()})
		c.Abort()
		return
	}

	// tokenからユーザーidを取り出す
	claims := token.Claims.(jwt.MapClaims)
	userId := int(claims["sub"].(float64))

	// NOTE: databaseへの問い合わせは行わない。そこまでするならtokenを採用すべきでない

	// ユーザーidをセット
	c.Set("userId", userId)
	c.Next()
}
