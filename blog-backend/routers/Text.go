package routers

import (
	"GIN/controllers/Text"

	"github.com/gin-gonic/gin"
)

func TextRoutersInit(r *gin.Engine) {
	textRouters := r.Group("/text")
	{
		// 根据用户id删除留言
		textRouters.DELETE("/delete", Text.TextController{}.Delete)
		// 添加链接
		textRouters.POST("/insert", Text.TextController{}.Insert)
		// 查询所有，分页
		textRouters.GET("/getall", Text.TextController{}.AuditScreening)
		// 批量删除
		textRouters.DELETE("/deleteall", Text.TextController{}.DeleteAll)
	}
}
