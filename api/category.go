package api

import (
	"net/http"
	"strconv"
	"zhengbiwen/blog-server/models"
	"zhengbiwen/blog-server/utils"

	"github.com/gin-gonic/gin"
)

func AddCategory(c *gin.Context) {
	var category models.Category
	var err error
	if err = c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.StatusFail,
			"message": "参数格式有误",
		})
		return
	}

	name := category.Name
	desc := category.Desc

	if _, isExist := models.CheckCategoryExistByName(name); isExist {
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.StatusFail,
			"message": "添加失败，该分类名称已经存在",
		})
		return
	}

	if err = models.CreateCategory(name, desc); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.StatusFail,
			"message": "添加失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    utils.StatusSuccess,
		"message": "添加成功",
	})
}

func GetCategoryList(c *gin.Context) {
	name := c.Query("name")
	categories, err := models.ReadCategoryList(name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.StatusFail,
			"message": "获取分类列表失败",
		})
	}

	var m map[string]interface{}

	out := make([]map[string]interface{}, 0, 10)
	for _, ca := range categories {
		m = map[string]interface{}{
			"id":            ca.ID,
			"created_at":    ca.CreatedAt.Format("2006-01-02 15:04:05"),
			"updated_at":    ca.UpdatedAt.Format("2006-01-02 15:04:05"),
			"desc":          ca.Desc,
			"name":          ca.Name,
			"article_count": len(ca.Articles),
		}
		out = append(out, m)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    utils.StatusSuccess,
		"message": "获取分类列表成功",
		"data":    out,
	})
}

func DeleteCategory(c *gin.Context) {
	cid := c.Param("cid")

	if cid == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.StatusFail,
			"message": "删除失败",
		})
		return
	}

	cidUint64, err := strconv.ParseUint(cid, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.StatusFail,
			"message": "参数有误",
		})
		return
	}

	err = models.DeleteCategory(uint(cidUint64))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.StatusFail,
			"message": "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    utils.StatusSuccess,
		"message": "删除成功",
	})
}

func UpdateCategory(c *gin.Context) {
	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.StatusFail,
			"message": "参数有误",
		})
		return
	}

	cid := category.ID
	cname := category.Name
	cdesc := category.Desc

	if res, isExist := models.CheckCategoryExistByName(cname); res.ID != cid && isExist {
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.StatusFail,
			"message": "添加失败，该分类名称已经存在",
		})
		return
	}

	err := models.UpdateCategory(cid, cname, cdesc)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.StatusFail,
			"message": "更新失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    utils.StatusSuccess,
		"message": "更新成功",
	})
}
