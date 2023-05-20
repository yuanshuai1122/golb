package Ip

import (
	"GIN/servise"
	"GIN/utils"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IpController struct{}

/**
 * 根据id删除链接
 * @author lizibin
 */
func (con IpController) Delete(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		utils.MissId(c)
		return
	}
	result := servise.IpServise{}.Delete(id)
	if result == 0 {
		utils.Success(c)
	} else if result == 1 {
		utils.DeleteFailed(c)
	} else {
		utils.NoId(c)
	}
}

/**
* 模块名: DeleteAll
* 代码描述: 批量删
* 作者:lizibin
* 创建时间:2023/01/02 15:12:41
 */
func (con IpController) DeleteAll(c *gin.Context) {
	id := c.Query("id")
	fmt.Println(id)
	result := servise.IpServise{}.DeleteAll(id)
	if result == 0 {
		utils.Success(c)
	} else if result == 1 {
		utils.DeleteFailed(c)
	} else {
		utils.NoId(c)
	}
}

/**
 * 新增
 * @author lizibin
 */
func (con IpController) Insert(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	result := servise.IpServise{}.Insert(json)
	if result == 0 {
		utils.Success(c)
	} else if result == 1 {
		utils.AddFailed(c)
	}
}

/**
 * 查询所有链接信息，分页，可带查询条件  搜索内容
 * @author lizibin
 */
func (con IpController) GetPage(c *gin.Context) {
	currentPage, _ := strconv.ParseInt(c.Query("currentPage"), 10, 64) //当前的页码
	pageSize, _ := strconv.ParseInt(c.Query("pageSize"), 10, 64)       //每页显示的条数
	result := servise.IpServise{}.GetPage(currentPage, pageSize)
	if result == 1 {
		utils.SearchPageFaild(c)
	} else {
		utils.SuccessData(c, result)
	}
}
