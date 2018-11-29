package controllers

import (
	"net/http"
	"zhengbiwen/blog_management_system/models"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user models.User
	var err error
	if err = c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    StatusFail,
			"message": "参数有误",
		})
		return
	}

	res, err := models.GetUserInfo(user.Username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    StatusFail,
			"message": "登录失败，账户名不存在",
		})
		return
	}

	if res.Password != user.Password {
		c.JSON(http.StatusOK, gin.H{
			"code":    StatusFail,
			"message": "登录失败，密码不正确",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    StatusSuccess,
		"message": "登录成功",
	})
	return
}
