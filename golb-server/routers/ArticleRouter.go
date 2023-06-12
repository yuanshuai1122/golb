package routers

import (
	"github.com/gin-gonic/gin"
	"golb/golb-server/controllers"
)

func ArticleRoutersInit(r *gin.Engine) {
	// 文章相关接口
	articleRouters := r.Group("/articles")
	{
		// 查询所有文章（分页）
		articleRouters.GET("/list", controllers.ArticleController{}.GetArticlesListPage)
		// 查询文章详情
		articleRouters.GET("/detail", controllers.ArticleController{}.GetArticleDetail)
	}

	// 后台管理相关接口
	adminRouters := r.Group("/admin")
	{
		adminRouters.POST("/articles/create", controllers.ArticleController{}.PostArticle)
	}
}
