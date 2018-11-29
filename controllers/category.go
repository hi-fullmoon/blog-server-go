package controllers

import (
	"net/http"
	"strconv"
	"zhengbiwen/blog_management_system/models"

	"github.com/gin-gonic/gin"
)

func AddCategory(c *gin.Context) {
	var category models.Category
	var err error
	if err = c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    StatusFail,
			"message": "添加失败",
		})
		return
	}

	name := category.Name
	desc := category.Desc
	if err = models.CreateCategory(name, desc); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    StatusFail,
			"message": "添加失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    StatusSuccess,
		"message": "添加成功",
	})
}

func GetCategoryList(c *gin.Context) {
	name := c.Query("name")
	categories, err := models.ReadCategoryList(name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    StatusFail,
			"message": "获取分类列表失败",
		})
	}

	var m map[string]interface{}

	out := make([]map[string]interface{}, 0, 10)
	for _, ca := range categories {
		m = map[string]interface{}{
			"id":            ca.ID,
			"created_at":    ca.CreatedAt,
			"updated_at":    ca.UpdatedAt,
			"desc":          ca.Desc,
			"name":          ca.Name,
			"article_count": len(ca.Articles),
		}
		out = append(out, m)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    StatusSuccess,
		"message": "获取分类列表成功",
		"data":    out,
	})
}

func DeleteCategory(c *gin.Context) {
	cid := c.Param("cid")

	if cid == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    StatusFail,
			"message": "删除失败",
		})
		return
	}

	cidUint64, err := strconv.ParseUint(cid, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    StatusFail,
			"message": "参数有误",
		})
		return
	}

	err = models.DeleteCategory(uint(cidUint64))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    StatusFail,
			"message": "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    StatusSuccess,
		"message": "删除成功",
	})
}

func UpdateCategory(c *gin.Context) {
	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    StatusFail,
			"message": "参数有误",
		})
		return
	}

	cid := category.ID
	cname := category.Name
	cdesc := category.Desc

	err := models.UpdateCategory(cid, cname, cdesc)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    StatusFail,
			"message": "更新失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    StatusSuccess,
		"message": "更新成功",
	})
}
