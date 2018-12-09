package routers

import (
	"net/http"
	"zhengbiwen/blog_management_system/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.StaticFS("/upload", http.Dir("upload"))
	r.StaticFS("/static", http.Dir("static"))

	r.LoadHTMLGlob("views/*/*")

	r.GET("/", controllers.Home)

	api := r.Group("/v1/api")
	{
		api.POST("/admin/user/login", controllers.Login)

		api.Use(controllers.ValidateUserSession())

		api.GET("/admin/user/:uid", controllers.GetUser)
		api.PUT("/admin/user", controllers.UpdateUser)
		api.POST("/admin/user/password", controllers.UpdateUserPwd)

		api.GET("/admin/categories", controllers.GetCategoryList)
		api.POST("/admin/category", controllers.AddCategory)
		api.DELETE("/admin/category/:cid", controllers.DeleteCategory)
		api.PUT("/admin/category", controllers.UpdateCategory)

		api.GET("/admin/tags", controllers.GetTagList)
		api.POST("/admin/tag", controllers.AddTag)
		api.DELETE("/admin/tag/:tid", controllers.DeleteTag)
		api.PUT("/admin/tag", controllers.UpdateTag)

		api.GET("/admin/articles", controllers.GetArticleList)
		api.GET("/admin/article/:aid", controllers.GetArticleInfo)
		api.POST("/admin/article", controllers.AddArticle)
		api.DELETE("/admin/article/:aid", controllers.DeleteArticle)
		api.PUT("/admin/article", controllers.UpdateArticle)

		api.POST("/admin/upload", controllers.UploadImg)
	}

	return r
}
