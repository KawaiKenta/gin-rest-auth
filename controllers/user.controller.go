package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"kk-rschain.com/gin-rest-auth/models"
	"kk-rschain.com/gin-rest-auth/utils"
)

func GetUserInfo(c *gin.Context) {
	// ユーザー情報の取得
	userId, err := utils.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	user, err := models.GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	// ユーザーIDの取得
	userId, err := utils.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	if err := models.DeleteUser(userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func UpdateUser(c *gin.Context) {
	// postされたデータの取得
	var newUserInfo struct {
		Name     string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&newUserInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// パスワードのハッシュ化
	hashedPassword, err := utils.EncryptPassword(newUserInfo.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	// ユーザーIdの取得
	userId, err := utils.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	// データベースからユーザーの取得
	user, err := models.GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	// user情報の更新
	user.Name = newUserInfo.Name
	user.Password = hashedPassword

	// データベースへ反映
	if err := models.UpdateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
