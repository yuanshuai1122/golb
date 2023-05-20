package Message

import (
	"GIN/servise"
	"GIN/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

type MessageController struct{}

/**
 * 根据id删除留言
 * @author lizibin
 */
func (con MessageController) Delete(c *gin.Context){
	id := c.Query("messageId")
	if id == "" {
		utils.MissId(c)
		return
	}
	result := servise.MessageServise{}.Delete(id)
	if result == 0 {
		utils.Success(c)
	}else if result == 1 {
		utils.DeleteFailed(c)
	}else {
		utils.NoId(c)
	}
}

/**
 * 新增留言
 * @author lizibin
 */
func (con MessageController) Insert(c *gin.Context){
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	result := servise.MessageServise{}.Insert(json)
	if result == 0 {
		utils.Success(c)
	}else if result == 1 {
		utils.AddFailed(c)
	}
}

/**
 * 分页查询留言
 * @author lizibin
 */
func (con MessageController) ShowmessageByPage(c *gin.Context){
	currentPage,_ := strconv.ParseInt(c.Query("currentPage"),10,64) //当前的页码
	pageSize,_ :=  strconv.ParseInt(c.Query("pageSize"), 10, 64)//每页显示的条数
	result := servise.MessageServise{}.ShowmessageByPage(currentPage,pageSize)
	if result == 1 {
		utils.SearchPageFaild(c)
	}else {
		utils.SuccessData(c,result)
	}
}

/**
 * 根据留言id 回复留言  更新留言
 * @author lizibin
 */
func (con MessageController) UpdateMessageReply(c *gin.Context){
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	result := servise.MessageServise{}.UpdateMessageReply(json)
	if result == 1 {
		utils.Failed(c,"更新失败!")
	}else if result == 0{
		utils.Success(c)
	}
}
