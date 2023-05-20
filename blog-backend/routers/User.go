package routers

import (
	"GIN/controllers/User"
	"github.com/gin-gonic/gin"
)

func UserRoutersInit(r *gin.Engine)  {
	userRouters:= r.Group("/user")
	{
		// 根据用户id删除用户    同时删除用户的所有评论以及文章
		userRouters.DELETE("/delete",User.UserController{}.Delete)
		// 根据ID查询用户信息
		userRouters.GET("/showUserByUserId",User.UserController{}.ShowUserByUserId)
		// 查询用户数量
		userRouters.GET("/showUserCount",User.UserController{}.ShowUserCount)
		// 查询所有链接信息，分页，可带查询条件  搜索内容
		userRouters.GET("/page/findby",User.UserController{}.Findby)
		// 修改个人信息
		userRouters.POST("/updateUserBasicsInfo",User.UserController{}.UpdateUserBasicsInfo)
		// 修改密码
		userRouters.GET("/updateUserPass",User.UserController{}.UpdateUserPass)
		// 查询用户名是否重复
		userRouters.GET("/showUserName",User.UserController{}.ShowUserName)
		// 查询用户名是否重复
		userRouters.GET("/searchEmail",User.UserController{}.SearchEmail)
		// 注册
		userRouters.POST("/insertRegisterInfo",User.UserController{}.InsertRegisterInfo)
	}
}
