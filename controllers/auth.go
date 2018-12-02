package controllers

import (
	"net/http"
	"zhengbiwen/blog_management_system/session"

	"github.com/gin-gonic/gin"
)

var HeaderToken = "X-Auth-Token"

func ValidateUserSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(HeaderToken)

		if len(token) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"code":    statusAuthFail,
				"message": "访问失败，重新登录",
			})
			c.Abort()
		}

		_, ok := session.IsSessionExpired(token)

		if ok {
			c.JSON(http.StatusOK, gin.H{
				"code":    statusAuthFail,
				"message": "token失效，重新登录",
			})
			c.Abort()
		}
	}
}
