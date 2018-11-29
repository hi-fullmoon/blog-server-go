package routers

import (
	"zhengbiwen/blog_management_system/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/user/login", controllers.Login)

		api.GET("/categories", controllers.GetCategoryList)
		api.POST("/category", controllers.AddCategory)
		api.DELETE("/category/:cid", controllers.DeleteCategory)
		api.PUT("/category", controllers.UpdateCategory)

		api.GET("/tags", controllers.GetTagList)
		api.POST("/tag", controllers.AddTag)
		api.DELETE("/tag/:tid", controllers.DeleteTag)
		api.PUT("/tag", controllers.UpdateTag)

		api.GET("/articles", controllers.GetArticleList)
		api.GET("/article/:aid", controllers.GetArticleInfo)
		api.POST("/article", controllers.AddArticle)
		api.DELETE("/article/:aid", controllers.DeleteArticle)
		api.PUT("/article", controllers.UpdateArticle)
	}

	return r
}
