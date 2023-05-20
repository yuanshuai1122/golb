package Classify

import (
	"GIN/servise"
	"GIN/utils"
	"github.com/gin-gonic/gin"
)

type ClassifyController struct{}

/**
 * 根据id删除分类
 * @author lizibin
 */
func (con ClassifyController) Delete(c *gin.Context){
	id := c.Query("classifyId")
	if id == "" {
		utils.MissId(c)
		return
	}
	result := servise.ClassifyServise{}.Delete(id)
	if result == 0 {
		utils.Success(c)
	}else if result == 1 {
		utils.DeleteFailed(c)
	}else {
		utils.Failed(c,"存在文章不允许删除!")
	}
}


/**
 * 添加分类
 * @author lizibin
 */
func (con ClassifyController) Insert(c *gin.Context){
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	result := servise.ClassifyServise{}.Insert(json)
	if result == 0 {
		utils.Success(c)
	}else if result == 1 {
		utils.AddFailed(c)
	}
}

/**
 *  查询所有文章分类
 * @author lizibin
 */
func (con ClassifyController) ShowAllClassifyInfo(c *gin.Context){
	result := servise.ClassifyServise{}.ShowAllClassifyInfo()
	if result == 1 {
		utils.Failed(c,"查询失败!")
	}else {
		utils.SuccessData(c,result)
	}
}

/**
 *  根据分类id查询分类信息
 * @author lizibin
 */
func (con ClassifyController) ShowClassifyInfoByClassifyId(c *gin.Context){
	id := c.Query("classifyId")
	if id == "" {
		utils.NoId(c)
		return
	}
	result := servise.ClassifyServise{}.ShowClassifyInfoByClassifyId(id)
	if result == 1 {
		utils.Failed(c,"查询失败!")
	}else {
		utils.SuccessData(c,result)
	}
}

/**
 *  根据分类id增加文章数量
 * @author lizibin
 */
func (con ClassifyController) UpdateArticleNumberByClassifyId(c *gin.Context){
	id := c.Query("classifyId")
	if id == "" {
		utils.NoId(c)
		return
	}
	result := servise.ClassifyServise{}.UpdateArticleNumberByClassifyId(id)
	if result == 1 {
		utils.Failed(c,"文章加1出错!")
	}else {
		utils.Success(c)
	}
}

/**
 *  根据分类id减少文章数量
 * @author lizibin
 */
func (con ClassifyController) UpdateArticleNumberByClassifyId2(c *gin.Context){
	id := c.Query("classifyId")
	if id == "" {
		utils.NoId(c)
		return
	}
	result := servise.ClassifyServise{}.UpdateArticleNumberByClassifyId2(id)
	if result == 1 {
		utils.Failed(c,"文章减1出错!")
	}else {
		utils.Success(c)
	}
}


/**
 *  根据id修改分类
 * @author lizibin
 */
func (con ClassifyController) Updata(c *gin.Context){
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	result := servise.ClassifyServise{}.Updata(json)
	if result == 1 {
		utils.Failed(c,"更新失败!")
	}else if result == 0{
		utils.Success(c)
	}
}