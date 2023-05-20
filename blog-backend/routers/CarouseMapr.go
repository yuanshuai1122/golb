package routers

import (
	"GIN/controllers/CarouseMap"
	"github.com/gin-gonic/gin"
)

func CarouseMapRoutersInit(r *gin.Engine)  {
	carouseMapRouters:= r.Group("/carouseMap")
	{
		// 更新轮播图
		carouseMapRouters.POST("/updateCarouseMap",CarouseMap.CarouseMapController{}.UpdateCarouseMap)
		// 查询轮播图
		carouseMapRouters.GET("/searchCarouseMap",CarouseMap.CarouseMapController{}.SearchCarouseMap)
	}
}
