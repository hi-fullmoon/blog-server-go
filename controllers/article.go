package controllers

import (
	"log"
	"net/http"
	"strconv"
	"zhengbiwen/blog_management_system/models"

	"github.com/gin-gonic/gin"
)

type articleDataIn struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	CategoryID uint   `json:"category_id"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	TagIds     []uint `json:"tag_ids"`
}

func AddArticle(c *gin.Context) {
	var article articleDataIn
	var err error

	if err = c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    StatusFail,
			"message": "参数有误",
		})
		return
	}

	cid := article.CategoryID
	title := article.Title
	desc := article.Desc
	content := article.Content
	tags := article.TagIds

	err = models.CreateArticle(cid, title, desc, content, tags)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    StatusFail,
			"message": "添加文章失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    StatusSuccess,
		"message": "添加文章成功",
	})
}

func GetArticleList(c *gin.Context) {
	articles, err := models.ReadArticleList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    StatusFail,
			"message": "获取文章列表失败",
		})
		return
	}

	m := make(map[string]interface{})

	out := make([]map[string]interface{}, 0, 10)

	for _, article := range articles {
		var tags []map[string]interface{}
		for _, tag := range article.Tags {
			tags = append(tags, map[string]interface{}{
				"id":   tag.ID,
				"name": tag.Name,
			})
		}
		m = map[string]interface{}{
			"id":            article.ID,
			"created_at":    article.CreatedAt,
			"updated_at":    article.UpdatedAt,
			"desc":          article.Desc,
			"title":         article.Title,
			"category_id":   article.CategoryID,
			"category_name": article.Category.Name,
			"tags":          tags,
		}
		out = append(out, m)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    StatusSuccess,
		"message": "获取文章列表成功",
		"data":    out,
	})
}

func GetArticleInfo(c *gin.Context) {
	aid := c.Param("aid")
	aidUint64, err := strconv.ParseUint(aid, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    StatusFail,
			"message": "参数有误",
		})
		return
	}

	info, err := models.ReadArticleInfo(uint(aidUint64))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    StatusFail,
			"message": "获取文章详情失败",
		})
		return
	}

	var tags []map[string]interface{}
	for _, tag := range info.Tags {
		m := map[string]interface{}{
			"id":   tag.ID,
			"name": tag.Name,
		}
		tags = append(tags, m)
	}

	out := map[string]interface{}{
		"title":         info.Title,
		"desc":          info.Desc,
		"content":       info.Content,
		"category_id":   info.CategoryID,
		"category_name": info.Category.Name,
		"tags":          tags,
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    StatusSuccess,
		"message": "获取文章详情成功",
		"data":    out,
	})
}

func DeleteArticle(c *gin.Context) {
	aid := c.Param("aid")
	aidUint64, err := strconv.ParseUint(aid, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    StatusFail,
			"message": "参数有误",
		})
		return
	}

	err = models.DeleteArticle(uint(aidUint64))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    StatusFail,
			"message": "删除文章失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    StatusSuccess,
		"message": "删除文章成功",
	})
}

func UpdateArticle(c *gin.Context) {
	var article articleDataIn
	var err error

	if err = c.ShouldBindJSON(&article); err != nil {
		log.Fatal("xxxxx", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    StatusFail,
			"message": "参数有误",
		})
		return
	}

	aid := article.ID
	cid := article.CategoryID
	title := article.Title
	desc := article.Desc
	content := article.Content
	tags := article.TagIds

	err = models.UpdateArticle(aid, cid, title, desc, content, tags)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    StatusFail,
			"message": "更新文章失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    StatusSuccess,
		"message": "更新文章成功",
	})
}
