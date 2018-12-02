package controllers

import (
	"net/http"
	"strings"
	"zhengbiwen/blog_management_system/session"

	"github.com/gin-gonic/gin"
)

var HeaderToken = "X-Auth-Token"

func ValidateUserSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(HeaderToken)

		if len(token) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"code":    StatusAuthFail,
				"message": "访问失败，重新登录",
			})
			c.Abort()
		}

		tokenSlice := strings.Split(token, "_")

		_, ok := session.IsSessionExpired(tokenSlice[1])

		if ok {
			c.JSON(http.StatusOK, gin.H{
				"code":    StatusAuthFail,
				"message": "token失效，重新登录",
			})
			c.Abort()
		}
	}
}
