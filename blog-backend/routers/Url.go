package routers

import (
	"GIN/controllers/Url"
	"github.com/gin-gonic/gin"
)

func UrlRoutersInit(r *gin.Engine)  {
	urlRouters:= r.Group("/url")
	{
		// 根据用户id删除留言
		urlRouters.DELETE("/delete",Url.UrlController{}.Delete)
		// 添加链接
		urlRouters.POST("/insert",Url.UrlController{}.Insert)
		// 查询所有链接信息，分页，可带查询条件  审核筛选内容
		urlRouters.GET("/page/auditScreening",Url.UrlController{}.AuditScreening)
		// 查询所有链接信息，分页，可带查询条件  搜索内容
		urlRouters.GET("/page/findby",Url.UrlController{}.Findby)
		// 根据链接类型传值
		urlRouters.GET("/page/showAllUrlInfo",Url.UrlController{}.ShowAllUrlInfo)
		// 修改链接审核结果
		urlRouters.GET("/updateUrlPass",Url.UrlController{}.UpdateUrlPass)
	}
}
