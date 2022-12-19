package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"kk-rschain.com/gin-rest-auth/config"
	"kk-rschain.com/gin-rest-auth/models"
	"kk-rschain.com/gin-rest-auth/utils"
)

func SignIn(c *gin.Context) {
	// login情報取得
	var loginInfo struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	// user取得
	user, err := models.GetUserByEmail(loginInfo.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	// password比較
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Wrong Password"})
		return
	}

	// 署名済みtokenを取得
	tokenString, err := utils.CreateTokenString(user.ID, config.Jwt.ExpirationTime, config.Jwt.Key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":     tokenString,
		"user_name": user.Name,
	})
}

func SignUp(c *gin.Context) {
	// ユーザー登録情報
	var newUserInfo struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}
	if err := c.ShouldBindJSON(&newUserInfo); err != nil {
		// TODO: 共通で使うのであれば、どこかへ切り出したほうが良さそう
		type ErrorMessage struct {
			Param   string
			Message string
		}
		var ve validator.ValidationErrors
		errors.As(err, &ve)
		out := make([]ErrorMessage, len(ve))
		for i, fe := range ve {
			out[i] = ErrorMessage{fe.Field(), msgForTag(fe.Tag())}
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": out})
		return
	}

	// パスワードのハッシュ化
	hashedPassword, err := utils.EncryptPassword(newUserInfo.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	// 新規ユーザー作成
	user := models.User{
		Name:       newUserInfo.Name,
		Email:      newUserInfo.Email,
		Password:   hashedPassword,
		IsVerified: false,
	}

	if err := models.CreateNewUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// TODO: token_refresh作成

// TODO: これもどこかへ上とセットで切り出す
func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "必須のフィールドです"
	case "email":
		return "メールの形式を満たしていません"
	case "min":
		return "短すぎます"
	}
	return ""
}
