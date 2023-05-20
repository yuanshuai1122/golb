package routers

import (
	"GIN/controllers/Email"
	"github.com/gin-gonic/gin"
)

func EmailRoutersInit(r *gin.Engine)  {
	emailRouters:= r.Group("/email")
	{
		// 发送邮件
		emailRouters.GET("/sendEmail",Email.EmailController{}.SendEmail)
	}
}
