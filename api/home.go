package api

import (
	"net/http"
	"zhengbiwen/blog-server/models"
	"zhengbiwen/blog-server/utils"

	"github.com/gin-gonic/gin"
)

func getCount() (int, int, int) {
	articleCount, err := models.ReadArticleCount()
	if err != nil {
		articleCount = 0
	}

	categoryCount, err := models.ReadCategoryCount()
	if err != nil {
		categoryCount = 0
	}

	tagCount, err := models.ReadTagCount()
	if err != nil {
		tagCount = 0
	}

	return articleCount, categoryCount, tagCount
}

func GetStatisticalData(c *gin.Context) {
	aCount, cCount, tCount := getCount()

	c.JSON(http.StatusOK, gin.H{
		"code":    utils.StatusSuccess,
		"message": "获取数据成功",
		"data": map[string]int{
			"article_total":  aCount,
			"category_total": cCount,
			"tag_total":      tCount,
		},
	})
}
