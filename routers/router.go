package routers

import (
	"net/http"
	"zhengbiwen/blog_management_system/api"
	"zhengbiwen/blog_management_system/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.StaticFS("/static", http.Dir("static"))
	r.StaticFS("/upload", http.Dir("upload"))

	r.LoadHTMLGlob("views/*/*")

	r.GET("/", controllers.Home)
	r.GET("/categories", controllers.CategoryList)
	r.GET("/categories/:cid", controllers.CategoryArticles)
	r.GET("/tags", controllers.TagList)
	r.GET("/tags/:tid", controllers.TagArticles)
	r.GET("/articles/:aid", controllers.Article)
	r.GET("/archives", controllers.Archive)
	r.GET("/about", controllers.About)
	r.GET("/message-board", controllers.MessageBoard)

	admin := r.Group("/api/admin")
	{
		admin.POST("/user/login", api.Login)

		admin.Use(api.ValidateUserSession())

		admin.GET("/user/:uid", api.GetUser)
		admin.PUT("/user", api.UpdateUser)
		admin.POST("/user/password", api.UpdateUserPwd)

		admin.GET("/categories", api.GetCategoryList)
		admin.POST("/category", api.AddCategory)
		admin.DELETE("/category/:cid", api.DeleteCategory)
		admin.PUT("/category", api.UpdateCategory)

		admin.GET("/tags", api.GetTagList)
		admin.POST("/tag", api.AddTag)
		admin.DELETE("/tag/:tid", api.DeleteTag)
		admin.PUT("/tag", api.UpdateTag)

		admin.GET("/articles", api.GetArticleList)
		admin.GET("/article/:aid", api.GetArticleInfo)
		admin.POST("/article", api.AddArticle)
		admin.DELETE("/article/:aid", api.DeleteArticle)
		admin.PUT("/article", api.UpdateArticle)

		admin.POST("/upload", api.UploadImg)
	}

	return r
}
