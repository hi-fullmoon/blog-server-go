package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadImg(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    StatusFail,
			"message": "a bad request",
		})
	}

	timeNow := time.Now().UnixNano()
	timeNowStr := strconv.FormatInt(timeNow, 10)

	if err := c.SaveUploadedFile(file, "upload/images/"+timeNowStr+".jpg"); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    StatusFail,
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    StatusSuccess,
		"message": "上传成功",
		"data": map[string]string{
			"img_url": "/upload/images/" + timeNowStr + ".jpg",
		},
	})
}
