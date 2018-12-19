package api

import (
	"net/http"
	"zhengbiwen/blog-server/utils"

	uuid2 "github.com/satori/go.uuid"

	"github.com/gin-gonic/gin"
)

func UploadImg(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    utils.StatusFail,
			"message": "a bad request",
		})
		return
	}

	uuid, err := uuid2.NewV1()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    utils.StatusFail,
			"message": "a bad request",
		})
		return
	}

	if err := c.SaveUploadedFile(file, "upload/images/"+uuid.String()+".jpg"); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.StatusFail,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    utils.StatusSuccess,
		"message": "上传图片成功",
		"data": map[string]string{
			"img_url": "/upload/images/" + uuid.String() + ".jpg",
		},
	})
}
