package Login

import (
	"GIN/servise"
	"GIN/utils"
	"github.com/gin-gonic/gin"
)

type LoginController struct{}

/**
 * 增加登录记录
 * @author lizibin
 */
func (con LoginController) InsertLoginInfo(c *gin.Context){
	id := c.Query("userId")
	if id == "" {
		utils.MissId(c)
		return
	}
	json := make(map[string]interface{})
	json["loginId"] = id
	result := servise.LoginServise{}.InsertLoginInfo(json)
	if result == 0 {
		utils.Success(c)
	}else if result == 1 {
		utils.Failed(c,"添加登录记录失败!")
	}
}

/**
 * 登录 没做token鉴权，密码没加密
 * @author lizibin
 */
func (con LoginController) Login(c *gin.Context){
	userName := c.Query("loginName")
	userPass := c.Query("loginPass")
	result := servise.LoginServise{}.Login(userName,userPass)
	if result == 1 {
		utils.Failed(c,"无此账号!")
	}else if result == 2 {
		utils.Failed(c,"密码不对!")
	}else if result == 3 {
		utils.Failed(c,"查询失败!")
	}else {
		utils.SuccessData(c,result)
	}
}

