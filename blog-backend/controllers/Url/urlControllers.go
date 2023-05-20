package Url

import (
	"GIN/servise"
	"GIN/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UrlController struct{}

/**
 * 根据id删除链接
 * @author lizibin
 */
func (con UrlController) Delete(c *gin.Context){
	id := c.Query("urlId")
	if id == "" {
		utils.MissId(c)
		return
	}
	result := servise.UrlServise{}.Delete(id)
	if result == 0 {
		utils.Success(c)
	}else if result == 1 {
		utils.DeleteFailed(c)
	}else {
		utils.NoId(c)
	}
}

/**
 * 新增
 * @author lizibin
 */
func (con UrlController) Insert(c *gin.Context){
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	result := servise.UrlServise{}.Insert(json)
	if result == 0 {
		utils.Success(c)
	}else if result == 1 {
		utils.AddFailed(c)
	}
}

/**
 * 查询所有链接信息，分页，可带查询条件  审核筛选内容
 * @author lizibin
 */
func (con UrlController) Findby(c *gin.Context){
	currentPage,_ := strconv.ParseInt(c.Query("currentPage"),10,64) //当前的页码
	pageSize,_ :=  strconv.ParseInt(c.Query("pageSize"), 10, 64)//每页显示的条数
	searchContent :=  c.Query("searchContent")
	result := servise.UrlServise{}.Findby(currentPage,pageSize,searchContent)
	if result == 1 {
		utils.SearchPageFaild(c)
	}else {
		utils.SuccessData(c,result)
	}
}


/**
 * 查询所有链接信息，分页，可带查询条件  搜索内容
 * @author lizibin
 */
func (con UrlController) AuditScreening(c *gin.Context){
	currentPage,_ := strconv.ParseInt(c.Query("currentPage"),10,64) //当前的页码
	pageSize,_ :=  strconv.ParseInt(c.Query("pageSize"), 10, 64)//每页显示的条数
	urlPass :=  c.Query("auditContent")
	result := servise.UrlServise{}.AuditScreening(currentPage,pageSize,urlPass)
	if result == 1 {
		utils.SearchPageFaild(c)
	}else {
		utils.SuccessData(c,result)
	}
}

/**
 * 根据链接类型传值
 * @author lizibin
 */
func (con UrlController) ShowAllUrlInfo(c *gin.Context){
	urlType :=  c.Query("urlType")
	result := servise.UrlServise{}.ShowAllUrlInfo(urlType)
	if result == 1 {
		utils.SearchPageFaild(c)
	}else {
		utils.SuccessData(c,result)
	}
}
/**
 * 根据链接类型传值
 * @author lizibin
 */
func (con UrlController) UpdateUrlPass(c *gin.Context){
	urlId :=  c.Query("urlId")
	urlPass :=  c.Query("urlPass")
	result := servise.UrlServise{}.UpdateUrlPass(urlId,urlPass)
	if result == 1 {
		utils.Failed(c,"修改失败")
	}else {
		utils.Success(c)
	}
}
