package routers

import (
	"GIN/controllers/Login"
	"github.com/gin-gonic/gin"
)

func LoginRoutersInit(r *gin.Engine)  {
	loginRouters:= r.Group("/login")
	{
		// 添加登录记录
		loginRouters.GET("/insertLoginInfo",Login.LoginController{}.InsertLoginInfo)
		// 登录
		loginRouters.GET("/tologin",Login.LoginController{}.Login)
	}
}
