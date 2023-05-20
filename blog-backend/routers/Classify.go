package routers

import (
	"GIN/controllers/Classify"
	"github.com/gin-gonic/gin"
)

func ClassifyRouterInit(r *gin.Engine)  {
	classifyRouters:= r.Group("/classify")
	{
		// 根据id删除分类
		classifyRouters.DELETE("/delete",Classify.ClassifyController{}.Delete)
		// 添加分类
		classifyRouters.POST("/insert",Classify.ClassifyController{}.Insert)
		// 查询所有文章分类
		classifyRouters.GET("/showAllClassifyInfo",Classify.ClassifyController{}.ShowAllClassifyInfo)
		// 根据分类id查询分类信息
		classifyRouters.GET("/showClassifyInfoByClassifyId",Classify.ClassifyController{}.ShowClassifyInfoByClassifyId)
		// 根据分类id增加文章数量
		classifyRouters.GET("/updateArticleNumberByClassifyId",Classify.ClassifyController{}.UpdateArticleNumberByClassifyId)
		// 根据分类id减少文章数量
		classifyRouters.GET("/updateArticleNumberByClassifyId2",Classify.ClassifyController{}.UpdateArticleNumberByClassifyId2)
		// 根据id修改分类
		classifyRouters.POST("/update",Classify.ClassifyController{}.Updata)
	}
}
