package api

import (
	"net/http"
	"strconv"
	"zhengbiwen/blog-server/models"
	"zhengbiwen/blog-server/utils"

	"github.com/gin-gonic/gin"
)

func AddTag(c *gin.Context) {
	var tag models.Tag
	var err error
	if err = c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.StatusFail,
			"message": "参数格式有误",
		})
		return
	}

	name := tag.Name
	if _, isExist := models.CheckTagExistByName(name); isExist {
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.StatusFail,
			"message": "添加标签失败，该标签名称已经存在",
		})
		return
	}

	if err = models.CreateTag(name); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.StatusFail,
			"message": "添加标签失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    utils.StatusSuccess,
		"message": "添加标签成功",
	})
}

func GetTagList(c *gin.Context) {
	name := c.Query("name")
	pageSize := c.Query("page_size")
	pageNum := c.Query("page_num")

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		pageSizeInt = 10
	}
	pageNumInt, err := strconv.Atoi(pageNum)
	if err != nil {
		pageNumInt = 1
	}

	tags, total, err := models.GetTagList(name, pageSizeInt, pageNumInt)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.StatusFail,
			"message": "获取标签列表失败",
		})
	}

	var m map[string]interface{}

	out := make([]map[string]interface{}, 0, 10)
	for _, ca := range tags {
		m = map[string]interface{}{
			"id":            ca.ID,
			"created_at":    ca.CreatedAt.Format("2006-01-02 15:04:05"),
			"updated_at":    ca.UpdatedAt.Format("2006-01-02 15:04:05"),
			"name":          ca.Name,
			"article_count": len(ca.Articles),
		}
		out = append(out, m)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    utils.StatusSuccess,
		"message": "获取标签列表成功",
		"data": map[string]interface{}{
			"list":      out,
			"total":     total,
			"page_num":  pageNumInt,
			"page_size": pageSizeInt,
		},
	})
}

func DeleteTag(c *gin.Context) {
	tid := c.Param("tid")

	if tid == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.StatusFail,
			"message": "删除标签失败",
		})
		return
	}

	cidUint64, err := strconv.ParseUint(tid, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.StatusFail,
			"message": "参数有误",
		})
		return
	}

	err = models.DeleteTag(uint(cidUint64))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.StatusFail,
			"message": "删除标签失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    utils.StatusSuccess,
		"message": "删除标签成功",
	})
}

func UpdateTag(c *gin.Context) {
	var tag models.Tag

	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.StatusFail,
			"message": "参数有误",
		})
		return
	}

	tid := tag.ID
	cname := tag.Name

	if res, isExist := models.CheckTagExistByName(cname); res.ID != tid && isExist {
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.StatusFail,
			"message": "添加标签失败，该标签名称已经存在",
		})
		return
	}

	err := models.UpdateTag(tid, cname)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.StatusFail,
			"message": "更新标签失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    utils.StatusSuccess,
		"message": "更新标签成功",
	})
}
