package routers

import (
	"GIN/controllers/Message"
	"github.com/gin-gonic/gin"
)

func MessageRoutersInit(r *gin.Engine)  {
	messageRouters:= r.Group("/message")
	{
		// 根据用户id删除留言
		messageRouters.DELETE("/delete",Message.MessageController{}.Delete)
		// 新增留言
		messageRouters.POST("/insert",Message.MessageController{}.Insert)
		// 分页查询留言
		messageRouters.GET("/pageShow",Message.MessageController{}.ShowmessageByPage)
		// 根据留言id 回复留言  更新留言
		messageRouters.POST("/updateMessageReply",Message.MessageController{}.UpdateMessageReply)
	}
}
