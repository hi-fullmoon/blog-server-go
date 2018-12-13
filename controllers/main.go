package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"page": "home",
	})
}

func CategoryList(c *gin.Context) {
	c.HTML(http.StatusOK, "category.html", gin.H{
		"page": "category",
	})
}

func CategoryArticles(c *gin.Context) {
	c.HTML(http.StatusOK, "category-articles.html", gin.H{
		"page": "category",
	})
}

func TagList(c *gin.Context) {
	c.HTML(http.StatusOK, "tag.html", gin.H{
		"page": "tag",
	})
}

func TagArticles(c *gin.Context) {
	c.HTML(http.StatusOK, "tag-articles.html", gin.H{
		"page": "tag",
	})
}

func Article(c *gin.Context) {
	c.HTML(http.StatusOK, "article.html", gin.H{
		"page": "",
	})
}

func Archive(c *gin.Context) {
	c.HTML(http.StatusOK, "archive.html", gin.H{
		"page": "archive",
	})
}

func About(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", gin.H{
		"page": "about",
	})
}

func MessageBoard(c *gin.Context) {
	c.HTML(http.StatusOK, "message-board.html", gin.H{
		"page": "message-board",
	})
}
