package routers

import (
	"GIN/controllers/Ip"

	"github.com/gin-gonic/gin"
)

func IpRoutersInit(r *gin.Engine) {
	ipRouters := r.Group("/ip")
	{
		// 根据用户id删除留言
		ipRouters.DELETE("/delete", Ip.IpController{}.Delete)
		// 添加链接
		ipRouters.POST("/insert", Ip.IpController{}.Insert)
		// 查询所有，分页
		ipRouters.GET("/getall", Ip.IpController{}.GetPage)
		// 批量删除
		ipRouters.DELETE("/deleteall", Ip.IpController{}.DeleteAll)
	}
}
