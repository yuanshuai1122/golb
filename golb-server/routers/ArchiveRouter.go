package routers

import (
	"github.com/gin-gonic/gin"
	"golb/golb-server/controllers"
)

func ArchiveRoutersInit(r *gin.Engine) {
	archiveRouters := r.Group("/archives")
	{
		// 查询所有文章（不分页）
		archiveRouters.GET("/list", controllers.ArchiveController{}.GetArchivesListPage)

	}
}
