package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"kk-rschain.com/gin-rest-auth/models"
)

func GetUserInfo(c *gin.Context) {
	// ユーザー情報の取得
	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"msg": "failed to get user from Token"})
		return
	}
	user, err := models.GetUserById(userId.(int))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	// ユーザーIDの取得
	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"msg": "failed to get user from Token"})
		return
	}

	if err := models.DeleteUser(userId.(int)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func UpdateUser(c *gin.Context) {
	// ユーザー情報の取得
	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"msg": "failed to get user from Token"})
		return
	}
	user, err := models.GetUserById(userId.(int))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err.Error()})
		return
	}

	// postされたデータの取得
	var newUserInfo struct {
		Name     string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&newUserInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	// TODO: util folder へ切り出し
	// パスワードのハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUserInfo.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	// データベースへ反映
	user.Name = newUserInfo.Name
	user.Password = string(hashedPassword)
	if err := models.UpdateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
