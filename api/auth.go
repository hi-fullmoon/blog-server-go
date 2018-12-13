package api

import (
	"net/http"
	"strconv"
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
			return
		}

		tokenSlice := strings.Split(token, "_")

		uidStr, err := strconv.ParseUint(tokenSlice[0], 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    StatusAuthFail,
				"message": "访问失败，重新登录",
			})
			c.Abort()
			return
		}

		_, ok := session.IsSessionExpired(uint(uidStr))

		if ok {
			c.JSON(http.StatusOK, gin.H{
				"code":    StatusAuthFail,
				"message": "token失效，重新登录",
			})
			c.Abort()
		}
	}
}
