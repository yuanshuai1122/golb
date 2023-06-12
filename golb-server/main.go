package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golb/golb-server/routers"
	"golb/golb-server/utils"
)

func main() {
	fmt.Println("hello world")
	// 初始化数据库
	utils.InitDB()
	r := gin.Default()

	// 路由组配置
	// 文章相关
	routers.ArticleRoutersInit(r)
	// 归档相关
	routers.ArchiveRoutersInit(r)
	// 后台相关
	routers.AdminRouterInit(r)

	err := r.Run(":8088")
	if err != nil {
		return
	}
}
