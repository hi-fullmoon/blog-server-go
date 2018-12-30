package controllers

import (
	"net/http"
	"strconv"
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

func Home(c *gin.Context) {
	aCount, cCount, tCount := getCount()

	pageNo := c.Query("page")
	pageNoInt, err := strconv.Atoi(pageNo)
	if err != nil || pageNoInt == 0 {
		pageNoInt = 1
	}
	if pageNoInt < 0 {
		pageNoInt = 1
	}

	pageSize := c.Query("page_size")
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil || pageSizeInt == 0 {
		pageSizeInt = 10
	}
	if pageSizeInt > 10 {
		pageSizeInt = 10
	}

	articles, total, _ := models.ReadArticleList("", "", "", "", "", 0, 0, pageSizeInt, pageNoInt)

	var prevPageNo, nextPageNo int
	if pageNoInt <= 0 {
		prevPageNo = 0
	} else {
		prevPageNo = pageNoInt - 1
	}

	if pageNoInt*pageSizeInt < total {
		nextPageNo = pageNoInt + 1
	} else {
		nextPageNo = 0
	}

	c.HTML(http.StatusOK, "home.html", gin.H{
		"Page":          "home",
		"ArticleCount":  aCount,
		"CategoryCount": cCount,
		"TagCount":      tCount,
		"Articles":      articles,
		"PageTotal":     total,
		"PrevPageNo":    prevPageNo,
		"NextPageNo":    nextPageNo,
	})
}

func CategoryList(c *gin.Context) {
	aCount, cCount, tCount := getCount()

	categoryList, _ := models.ReadCategoryList("")

	c.HTML(http.StatusOK, "category.html", gin.H{
		"Page":          "category",
		"ArticleCount":  aCount,
		"CategoryCount": cCount,
		"TagCount":      tCount,
		"CategoryList":  categoryList,
	})
}

func CategoryArticles(c *gin.Context) {
	aCount, cCount, tCount := getCount()

	categoryName := c.Param("cName")

	articles, _ := models.ReadArticleListByCategoryName(categoryName)

	c.HTML(http.StatusOK, "category-articles.html", gin.H{
		"Page":          "category",
		"ArticleCount":  aCount,
		"CategoryCount": cCount,
		"TagCount":      tCount,
		"CategoryName":  categoryName,
		"Articles":      articles,
	})
}

func TagList(c *gin.Context) {
	aCount, cCount, tCount := getCount()

	tags, _, _ := models.ReadTagList("", 1000, 0)

	c.HTML(http.StatusOK, "tag.html", gin.H{
		"Page":          "tag",
		"ArticleCount":  aCount,
		"CategoryCount": cCount,
		"TagCount":      tCount,
		"Tags":          tags,
	})
}

func TagArticles(c *gin.Context) {
	aCount, cCount, tCount := getCount()

	tagName := c.Param("tName")

	articles, _ := models.ReadArticleListByTagName(tagName)

	c.HTML(http.StatusOK, "tag-articles.html", gin.H{
		"Page":          "tag",
		"ArticleCount":  aCount,
		"CategoryCount": cCount,
		"TagCount":      tCount,
		"TagName":       tagName,
		"Articles":      articles,
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

	article, _ := models.ReadArticleInfo(uint(aidUint64))

	c.HTML(http.StatusOK, "article.html", gin.H{
		"Page":          "",
		"ArticleCount":  aCount,
		"CategoryCount": cCount,
		"TagCount":      tCount,
		"Article":       article,
	})
}

func Archive(c *gin.Context) {
	aCount, cCount, tCount := getCount()

	m, _ := models.ReadArticleByGroup()

	c.HTML(http.StatusOK, "archive.html", gin.H{
		"Page":          "archive",
		"ArticleCount":  aCount,
		"CategoryCount": cCount,
		"TagCount":      tCount,
		"DateLine":      m,
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

func GetArticlesByTitle(c *gin.Context) {
	title := c.Query("title")

	articles, err := models.ReadArticleListByArticleTitle(title)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.StatusFail,
			"message": "获取结果失败",
		})
		return
	}

	m := make(map[string]interface{})
	out := make([]map[string]interface{}, 0, 10)

	for _, article := range articles {
		m = map[string]interface{}{
			"id":    article.ID,
			"title": article.Title,
		}
		out = append(out, m)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    utils.StatusSuccess,
		"message": "获取结果成功",
		"data":    out,
	})
}
