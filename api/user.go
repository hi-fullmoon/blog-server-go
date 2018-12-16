package api

import (
	"net/http"
	"strconv"
	"zhengbiwen/blog-server/models"
	"zhengbiwen/blog-server/session"

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

	res, err := models.GetUserByAccount(user.Account)
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

	token := session.GenerateNewSessionId(res.ID)

	if token == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    StatusFail,
			"message": "生成token失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    StatusSuccess,
		"message": "登录成功",
		"data": map[string]interface{}{
			"token":        token,
			"user_id":      res.ID,
			"nick_name":    res.NickName,
			"account":      res.Account,
			"avatar_image": res.AvatarImage,
		},
	})
	return
}

func GetUser(c *gin.Context) {
	uid := c.Param("uid")

	uidUint64, err := strconv.ParseUint(uid, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    StatusFail,
			"message": "参数有误",
		})
		return
	}

	user, err := models.GetUserById(uint(uidUint64))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    StatusFail,
			"message": "获取用户信息失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    StatusSuccess,
		"message": "获取用户信息成功",
		"data": map[string]interface{}{
			"id":           user.ID,
			"account":      user.Account,
			"nickname":     user.NickName,
			"email":        user.Email,
			"province":     user.Province,
			"city":         user.City,
			"county":       user.County,
			"website":      user.Website,
			"profile":      user.Profile,
			"avatar_image": user.AvatarImage,
		},
	})
}

func UpdateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    StatusFail,
			"message": "参数有误",
		})
		return
	}

	if err := models.UpdateUserInfo(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    StatusFail,
			"message": "更新用户信息失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    StatusSuccess,
		"message": "更新用户信息成功",
	})
}

type userPasswordIn struct {
	ID          uint   `json:"id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
	RePassword  string `json:"re_password"`
}

func UpdateUserPwd(c *gin.Context) {
	var body userPasswordIn

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    StatusFail,
			"message": "参数有误",
		})
		return
	}

	uid := body.ID
	oldPwd := body.OldPassword

	res, err := models.GetUserById(uid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    StatusFail,
			"message": "修改密码失败",
		})
		return
	}

	if oldPwd != res.Password {
		c.JSON(http.StatusOK, gin.H{
			"code":    StatusFail,
			"message": "修改密码失败，原密码不正确",
		})
		return
	}

	newPwd := body.NewPassword
	rePwd := body.RePassword

	if newPwd != rePwd {
		c.JSON(http.StatusOK, gin.H{
			"code":    StatusFail,
			"message": "修改密码失败，新密码和原密码相同",
		})
		return
	}

	if newPwd != rePwd {
		c.JSON(http.StatusOK, gin.H{
			"code":    StatusFail,
			"message": "修改密码失败，两次输入的新密码不一样",
		})
		return
	}

	if err := models.UpdateUserPassword(uid, newPwd); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    StatusFail,
			"message": "修改密码失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    StatusSuccess,
		"message": "修改密码成功",
	})
}
