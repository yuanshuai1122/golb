package routers

import (
	"github.com/gin-gonic/gin"
	"golb/golb-server/controllers"
)

func ArticleRoutersInit(r *gin.Engine) {
	articleRouters := r.Group("/articles")
	{
		// 查询所有文章（不分页）
		articleRouters.GET("/list", controllers.ArticleController{}.GetArticlesList)

	}
}
