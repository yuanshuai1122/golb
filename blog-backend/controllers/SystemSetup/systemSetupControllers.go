package SystemSetup

import (
	"GIN/servise"
	"GIN/utils"

	"github.com/gin-gonic/gin"
)

type SystemSetupController struct{}

/**
 * 查询
 * @author lizibin
 */
func (con SystemSetupController) ShowAllSystemSetup(c *gin.Context) {
	result := servise.SystemSetupServise{}.ShowAllSystemSetup()
	if result == 1 {
		utils.Failed(c, "系统设置查询失败!")
	} else {
		utils.SuccessData(c, result)
	}
}

/**
 * 修改
 * @author lizibin
 */
func (con SystemSetupController) UpdateSystemSetup(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	result := servise.SystemSetupServise{}.UpdateSystemSetup(json)
	if result == 0 {
		utils.Success(c)
	} else if result == 1 {
		utils.Failed(c, "修改失败!")
	}
}
