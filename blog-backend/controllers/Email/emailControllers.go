package Email

import (
	"GIN/servise"
	"GIN/utils"
	"github.com/gin-gonic/gin"
)

type EmailController struct{}

/**
 * 发送邮件
 * @author lizibin
 */
func (con EmailController) SendEmail(c *gin.Context){
	registerEmail := c.Query("registerEmail")
	content := c.Query("content")
	result := servise.EmailServise{}.SendEmail(registerEmail,content)
	if result == 0 {
		utils.Success(c)
	}else if result == 1 {
		utils.Failed(c,"邮件发送失败!")
	}
}

