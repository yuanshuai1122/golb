package routers

import (
	"GIN/controllers/SystemSetup"
	"github.com/gin-gonic/gin"
)

func SystemSetupRoutersInit(r *gin.Engine)  {
	systemSetupRouters:= r.Group("/systemSetup")
	{
		// 查询全站的系统设置
		systemSetupRouters.GET("/showAllSystemSetup",SystemSetup.SystemSetupController{}.ShowAllSystemSetup)
		// 修改系统设置
		systemSetupRouters.POST("/updateSystemSetup",SystemSetup.SystemSetupController{}.UpdateSystemSetup)
	}
}
