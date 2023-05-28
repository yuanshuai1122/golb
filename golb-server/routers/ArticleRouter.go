package routers

import (
	"github.com/gin-gonic/gin"
	"golb/golb-server/controllers"
)

func ArticleRoutersInit(r *gin.Engine) {
	articleRouters := r.Group("/articles")
	{
		// 查询所有文章（分页）
		articleRouters.GET("/list", controllers.ArticleController{}.GetArticlesListPage)
		// 查询文章详情
		articleRouters.GET("/detail", controllers.ArticleController{}.GetArticleDetail)

	}
}
