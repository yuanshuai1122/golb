package main

import (
	"GIN/middlewares"
	"GIN/routers"
	"GIN/utils"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// 记录到文件。
	utils.InitConfig()
	utils.InitDB()
	f, _ := os.Create("gin.log")
	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.Default()
	// 解决跨越
	r.Use(middlewares.Cors())
	// 静态资源访问
	r.Static("/static", "./static")
	// 路由组配置
	// 文章相关
	routers.ArticleRoutersInit(r)
	// 类别
	routers.ClassifyRouterInit(r)
	// 评论
	routers.CommentRouterInit(r)
	// 留言
	routers.MessageRoutersInit(r)
	// 链接
	routers.UrlRoutersInit(r)
	routers.EmailRoutersInit(r)
	routers.UploadRoutersInit(r)
	routers.SystemSetupRoutersInit(r)
	routers.LoginRoutersInit(r)
	routers.UserRoutersInit(r)
	routers.CarouseMapRoutersInit(r)
	routers.TextRoutersInit(r)
	routers.IpRoutersInit(r)
	r.Run(":520") // 监听并在 0.0.0.0:8080 上启动服务
}
