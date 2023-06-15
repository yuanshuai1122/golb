package routers

import (
	"github.com/gin-gonic/gin"
	"golb/golb-server/controllers"
)

func CategoryRoutersInit(r *gin.Engine) {
	categoryRouters := r.Group("/category")
	{
		// 查询所有分类（不分页）
		categoryRouters.GET("/list", controllers.CategoryController{}.GetCategoryList)

	}
}
