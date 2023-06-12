package routers

import (
	"github.com/gin-gonic/gin"
	"golb/golb-server/controllers"
)

func AdminRouterInit(r *gin.Engine) {
	// 后台管理相关接口
	adminRouters := r.Group("/admin")
	{
		adminArticlesRouters := adminRouters.Group("/articles")
		{
			// 新增文章
			adminArticlesRouters.POST("/create", controllers.ArticleController{}.PostArticle)
		}
	}
}
