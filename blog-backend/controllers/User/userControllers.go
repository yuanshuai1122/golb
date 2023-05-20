package User

import (
	"GIN/servise"
	"GIN/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

/**
 * 根据用户id删除用户 同时删除用户的所有评论以及文章
 * @author lizibin
 */
func (con UserController) Delete(c *gin.Context) {
	id := c.Query("userId")
	if id == "" {
		utils.MissId(c)
		return
	}
	result := servise.UserServise{}.Delete(id)
	if result == 0 {
		utils.Success(c)
	} else if result == 1 {
		utils.DeleteFailed(c)
	}
}

/**
 * 查询所有链接信息，分页，可带查询条件  审核筛选内容
 * @author lizibin
 */
func (con UserController) Findby(c *gin.Context) {
	currentPage, _ := strconv.ParseInt(c.Query("currentPage"), 10, 64) //当前的页码
	pageSize, _ := strconv.ParseInt(c.Query("pageSize"), 10, 64)       //每页显示的条数
	searchContent := c.Query("searchContent")
	result := servise.UserServise{}.Findby(currentPage, pageSize, searchContent)
	if result == 1 {
		utils.SearchPageFaild(c)
	} else {
		utils.SuccessData(c, result)
	}
}

/**
 * 根据ID获取用户信息
 * @author lizibin
 */
func (con UserController) ShowUserByUserId(c *gin.Context) {
	userID := c.Query("userID")
	result := servise.UserServise{}.ShowUserByUserId(userID)
	if result == 1 {
		utils.SearchPageFaild(c)
	} else {
		utils.SuccessData(c, result)
	}
}

// 查询用户数量
func (con UserController) ShowUserCount(c *gin.Context) {
	result := servise.UserServise{}.ShowUserCount()
	if result == 1 {
		utils.SearchPageFaild(c)
	} else {
		utils.SuccessData(c, result)
	}
}

// UpdateUserBasicsInfo 修改个人信息
func (con UserController) UpdateUserBasicsInfo(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	result := servise.UserServise{}.UpdateUserBasicsInfo(json)
	if result == 1 {
		utils.Failed(c, "修改失败!")
	} else {
		utils.Success(c)
	}
}

// 修改密码
func (con UserController) UpdateUserPass(c *gin.Context) {
	userId := c.Query("userId")
	userPass := utils.BcryptHash(c.Query("userPass"))
	result := servise.UserServise{}.UpdateUserPass(userId, userPass)
	if result == 1 {
		utils.Failed(c, "修改失败!")
	} else {
		utils.Success(c)
	}
}

// ShowUserName 查询用户名是否重复
func (con UserController) ShowUserName(c *gin.Context) {
	userName := c.Query("registerName")
	result := servise.UserServise{}.ShowUserName(userName)
	if result == 0 {
		utils.FailedMsgAndCode(c, "用户名重复!", 1020)
	} else {
		utils.Success(c)
	}
}

// SearchEmail 查询邮箱是否重复
func (con UserController) SearchEmail(c *gin.Context) {
	userEmail := c.Query("registerEmail")
	result := servise.UserServise{}.SearchEmail(userEmail)
	if result == 0 {
		utils.FailedMsgAndCode(c, "邮箱重复!", 1021)
	} else {
		utils.Success(c)
	}
}

//InsertRegisterInfo
func (con UserController) InsertRegisterInfo(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	json["userPass"] = utils.BcryptHash(json["userPass"].(string))
	result := servise.UserServise{}.InsertRegisterInfo(json)
	if result == 0 {
		utils.Success(c)
	} else {
		utils.Failed(c, "注册失败!")
	}
}
