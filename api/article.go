package api

import (
	"net/http"
	"strconv"
	"zhengbiwen/blog-server/models"
	"zhengbiwen/blog-server/utils"

	"github.com/gin-gonic/gin"
)

type articleDataIn struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	CategoryID uint   `json:"category_id"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	TagIds     []uint `json:"tag_ids"`
	CoverImage string `json:"cover_image"`
}

func AddArticle(c *gin.Context) {
	var article articleDataIn
	var err error

	if err = c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.StatusFail,
			"message": "参数有误",
		})
		return
	}

	cid := article.CategoryID
	title := article.Title
	desc := article.Desc
	content := article.Content
	tags := article.TagIds
	image := article.CoverImage

	err = models.CreateArticle(cid, title, desc, content, image, tags)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.StatusFail,
			"message": "添加文章失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    utils.StatusSuccess,
		"message": "添加文章成功",
	})
}

func GetArticleList(c *gin.Context) {
	title := c.Query("title")
	cid := c.Query("category_id")
	tid := c.Query("tag_id")
	cStartAt := c.Query("created_start_at")
	cEndAt := c.Query("created_end_at")
	uStartAt := c.Query("updated_start_at")
	uEndAt := c.Query("updated_end_at")
	pageSize := c.Query("page_size")
	pageNum := c.Query("page_num")

	cidUint64, _ := strconv.ParseUint(cid, 10, 64)
	tidUint64, _ := strconv.ParseUint(tid, 10, 64)

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		pageSizeInt = 10
	}
	pageNumInt, err := strconv.Atoi(pageNum)
	if err != nil {
		pageNumInt = 1
	}

	articles, total, err := models.GetArticleList(title, cStartAt, cEndAt, uStartAt, uEndAt,
		uint(cidUint64), uint(tidUint64),
		pageSizeInt, pageNumInt)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.StatusFail,
			"message": "获取文章列表失败",
		})
		return
	}

	m := make(map[string]interface{})

	out := make([]map[string]interface{}, 0, 10)

	for _, article := range articles {
		tags := make([]map[string]interface{}, 0)
		for _, tag := range article.Tags {
			tags = append(tags, map[string]interface{}{
				"id":   tag.ID,
				"name": tag.Name,
			})
		}
		m = map[string]interface{}{
			"id":            article.ID,
			"created_at":    article.CreatedAt.Format("2006-01-02 15:04:05"),
			"updated_at":    article.UpdatedAt.Format("2006-01-02 15:04:05"),
			"desc":          article.Desc,
			"title":         article.Title,
			"category_id":   article.CategoryID,
			"category_name": article.Category.Name,
			"tags":          tags,
		}
		out = append(out, m)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    utils.StatusSuccess,
		"message": "获取文章列表成功",
		"data": map[string]interface{}{
			"list":      out,
			"total":     total,
			"page_num":  pageNumInt,
			"page_size": pageSizeInt,
		},
	})
}

func GetArticleInfo(c *gin.Context) {
	aid := c.Param("aid")
	aidUint64, err := strconv.ParseUint(aid, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.StatusFail,
			"message": "参数有误",
		})
		return
	}

	info, err := models.ReadArticleInfo(uint(aidUint64))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.StatusFail,
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
		"cover_image":   info.CoverImage,
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    utils.StatusSuccess,
		"message": "获取文章详情成功",
		"data":    out,
	})
}

func DeleteArticle(c *gin.Context) {
	aid := c.Param("aid")
	aidUint64, err := strconv.ParseUint(aid, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.StatusFail,
			"message": "参数有误",
		})
		return
	}

	err = models.DeleteArticle(uint(aidUint64))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.StatusFail,
			"message": "删除文章失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    utils.StatusSuccess,
		"message": "删除文章成功",
	})
}

func UpdateArticle(c *gin.Context) {
	var article articleDataIn
	var err error

	if err = c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.StatusFail,
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
	image := article.CoverImage

	err = models.UpdateArticle(aid, cid, title, desc, content, image, tags)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.StatusFail,
			"message": "更新文章失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    utils.StatusSuccess,
		"message": "更新文章成功",
	})
}
