package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadImg(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "a bad request",
		})
	}

	filename := file.Filename
	fmt.Println("filename------>", filename)

	if err := c.SaveUploadedFile(file, "upload/images/"+filename); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    StatusSuccess,
		"message": "上传成功",
		"data": map[string]string{
			"img_url": "/upload/images/" + filename,
		},
	})
}
