package controllers

import (
	"net/http"
	"strconv"
	"zhengbiwen/blog_management_system/models"

	"github.com/gin-gonic/gin"
)

func getCount() (int, int, int) {
	articleCount, err := models.GetArticleCount()
	if err != nil {
		articleCount = 0
	}

	categoryCount, err := models.GetCategoryCount()
	if err != nil {
		categoryCount = 0
	}

	tagCount, err := models.GetTagCount()
	if err != nil {
		tagCount = 0
	}

	return articleCount, categoryCount, tagCount
}

func Home(c *gin.Context) {
	aCount, cCount, tCount := getCount()

	articles, _, err := models.ReadArticleList("", "", "", "", "", 0, 0, 0, 0)
	if err != nil {
		articles = nil
	}

	c.HTML(http.StatusOK, "home.html", gin.H{
		"Page":          "home",
		"ArticleCount":  aCount,
		"CategoryCount": cCount,
		"TagCount":      tCount,
		"Articles":      articles,
	})
}

func CategoryList(c *gin.Context) {
	aCount, cCount, tCount := getCount()

	c.HTML(http.StatusOK, "category.html", gin.H{
		"Page":          "category",
		"ArticleCount":  aCount,
		"CategoryCount": cCount,
		"TagCount":      tCount,
	})
}

func CategoryArticles(c *gin.Context) {
	aCount, cCount, tCount := getCount()

	c.HTML(http.StatusOK, "category-articles.html", gin.H{
		"Page":          "category",
		"ArticleCount":  aCount,
		"CategoryCount": cCount,
		"TagCount":      tCount,
	})
}

func TagList(c *gin.Context) {
	aCount, cCount, tCount := getCount()

	c.HTML(http.StatusOK, "tag.html", gin.H{
		"Page":          "tag",
		"ArticleCount":  aCount,
		"CategoryCount": cCount,
		"TagCount":      tCount,
	})
}

func TagArticles(c *gin.Context) {
	aCount, cCount, tCount := getCount()

	c.HTML(http.StatusOK, "tag-articles.html", gin.H{
		"Page":          "tag",
		"ArticleCount":  aCount,
		"CategoryCount": cCount,
		"TagCount":      tCount,
	})
}

func Article(c *gin.Context) {
	aCount, cCount, tCount := getCount()

	aid := c.Param("aid")
	aidUint64, err := strconv.ParseUint(aid, 10, 64)
	if err != nil {
		aidUint64 = 0
	}

	models.UpdateArticleViewCount(uint(aidUint64))

	c.HTML(http.StatusOK, "article.html", gin.H{
		"Page":          "",
		"ArticleCount":  aCount,
		"CategoryCount": cCount,
		"TagCount":      tCount,
	})
}

func Archive(c *gin.Context) {
	aCount, cCount, tCount := getCount()

	c.HTML(http.StatusOK, "archive.html", gin.H{
		"Page":          "archive",
		"ArticleCount":  aCount,
		"CategoryCount": cCount,
		"TagCount":      tCount,
	})
}

func About(c *gin.Context) {
	aCount, cCount, tCount := getCount()

	c.HTML(http.StatusOK, "about.html", gin.H{
		"Page":          "about",
		"ArticleCount":  aCount,
		"CategoryCount": cCount,
		"TagCount":      tCount,
	})
}

func MessageBoard(c *gin.Context) {
	aCount, cCount, tCount := getCount()

	c.HTML(http.StatusOK, "message-board.html", gin.H{
		"Page":          "message-board",
		"ArticleCount":  aCount,
		"CategoryCount": cCount,
		"TagCount":      tCount,
	})
}
