package Comment

import (
	"GIN/models"
	"GIN/servise"
	"GIN/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommentController struct{}

/**
 * 根据id删除评论
 * @author lizibin
 */
func (con CommentController) Delete(c *gin.Context){
	id := c.Query("commentId")
	if id == "" {
		utils.NoId(c)
	}
	//result := servise.CommentServise{}.Delete(id)
	//if result == 1 {
	//	utils.DeleteFailed(c)
	//}else {
	//	utils.Success(c)
	//}
	commentinfo := models.CommentInfo{}
	if err :=  utils.DB.Where("parentId = ?",id).First(&commentinfo).Error; err == nil {
		if commentinfo.CommentId != 0 {
			c.JSON(http.StatusOK,gin.H{
				"code": 1001,
				"msg": "存在子评论，不允许删除!",
			})
		}
	}else {
		if err :=  utils.DB.Where("commentId = ?",id).First(&commentinfo).Delete(&commentinfo).Error; err != nil {
			c.JSON(http.StatusOK,gin.H{
				"code": 1001,
				"msg": "删除失败",
			})
		}else {
			c.JSON(http.StatusOK,gin.H{
				"code": 0,
				"msg": "OK",
			})
		}
	}
}

/**
 * 根据id删除评论
 * @author lizibin
 */
func (con CommentController) DeleteByArticleId(c *gin.Context){
	id := c.Query("articleId")
	if id == "" {
		c.JSON(http.StatusOK,gin.H{
			"code": 1001,
			"msg": "ID为空",
		})
	}
	commentinfo := []models.CommentInfo{}
	if err :=  utils.DB.Where("articleId = ?",id).Find(&commentinfo).Delete(&commentinfo).Error; err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code": 1001,
			"msg": "删除失败",})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code": 0,
			"msg": "OK",})
	}
}


/**
 * 添加评论
 * @author lizibin
 */
func (con CommentController) InsertCommentInfo(c *gin.Context){
	commentinfo := models.AddCommentInfo{}
	if err := c.ShouldBindJSON(&commentinfo); err!=nil {
		c.JSON(http.StatusOK,gin.H{
			"code": 1001,
			"msg": "接收失败",
		})
	}else {
		if err :=  utils.DB.Create(&commentinfo).Error; err != nil {
			c.JSON(http.StatusOK,gin.H{
				"code": 1001,
				"msg": "追加失败",})
		}else {
			c.JSON(http.StatusOK,gin.H{
				"code": 0,
				"msg": "OK",
			})
		}
	}


}

/**
 * 查询所有评论
 * @author lizibin
 */
func (con CommentController) ShowAllComment(c *gin.Context){
	commentinfo := []models.CommentInfo{}
	if err :=  utils.DB.Find(&commentinfo).Error; err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code": 1001,
			"msg": "查询失败",
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code": 0,
			"msg": "OK",
			"data": commentinfo,
		})
	}
}

/**
 * 根据文章id查询文章对应的评论
 * @author lizibin
 */
func (con CommentController) ShowCommentByArticleId(c *gin.Context){
	commentinfo := []models.CommentInfo{}
	AllCommentInfoAndUserInfo := []models.AllCommentInfoAndUserInfo{}
	id := c.Query("articleId")
	if id == "" {
		c.JSON(http.StatusOK,gin.H{
			"code": 1001,
			"msg": "ID为空",
		})
		return
	}
	if err :=  utils.DB.Where("articleId = ? AND parentId = ?",id,0).Find(&commentinfo).Error; err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code": 1001,
			"msg": "查询失败",
		})
		return
	}
	for index, _ := range commentinfo {
		fmt.Println(commentinfo[index].UserId)
		letAllCommentInfoAndUserInfo := models.AllCommentInfoAndUserInfo{}
		// 循环结构体 用userId 去查个人信息
		if err :=  utils.DB.Where("userId = ?",commentinfo[index].UserId).First(&letAllCommentInfoAndUserInfo).Error; err != nil {
			c.JSON(http.StatusOK,gin.H{
				"code": 1001,
				"msg": "查询失败",
			})
			return
		}else {
			// 组装child 去查 parentId 非0的
			child := []models.CommentInfo{}
			if err :=  utils.DB.Where("userId = ? AND parentId != 0",commentinfo[index].UserId).Find(&child).Error; err != nil {
				c.JSON(http.StatusOK,gin.H{
					"code": 1001,
					"msg": "查询失败",
				})
				return
			}else {
				letAllCommentInfoAndUserInfo.Child = child
				letAllCommentInfoAndUserInfo.CommentInfo =commentinfo[index]
			}
			AllCommentInfoAndUserInfo = append(AllCommentInfoAndUserInfo, letAllCommentInfoAndUserInfo)
		}
	}
	c.JSON(http.StatusOK,gin.H{
		"code": 0,
		"msg": "OK",
		"data": AllCommentInfoAndUserInfo,
	})
}


/**
 * 查询所有评论数量
 * @author lizibin
 */
func (con CommentController) GetCount(c *gin.Context){
	commentinfo := []models.CommentInfo{}
	var count int64
	if err :=  utils.DB.Find(&commentinfo).Count(&count).Error; err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code": 1001,
			"msg": "查询失败",
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code": 0,
			"msg": "OK",
			"data": count,
		})
	}
}

/**
 * 后台评论管理：分页查询  + 搜索框内容  管理员获取全部  普通用户获取自己的
 * @author lizibin
 */
func (con CommentController) GetCountByUserType(c *gin.Context){
	// 获取参数
	currentPage,_ := strconv.ParseInt(c.Query("currentPage"),10,64) //当前的页码
	pageSize,_ :=  strconv.ParseInt(c.Query("pageSize"), 10, 64)//每页显示的条数
	userType := c.Query("userType")// 用户类型
	userId := c.Query("userId")// 用户id
	searchContent := c.Query("searchContent")// 名称查询
	result := servise.CommentServise{}.Findby(currentPage,pageSize,userType,userId,searchContent)
	if result == 1 {
		utils.SearchPageFaild(c)
	}else {
		utils.SuccessData(c,result)
	}

}

func (con CommentController) GetCountByUserType1(c *gin.Context){
	commentinfo := []models.CommentInfo{}
	// 获取参数
	currentPage,_ := strconv.ParseInt(c.Query("currentPage"),10,64) //当前的页码
	pageSize,_ :=  strconv.ParseInt(c.Query("pageSize"), 10, 64)//每页显示的条数
	userType := c.Query("userType")// 用户类型
	userId := c.Query("userId")// 用户id
	offset := (currentPage - 1) * pageSize
	var count int64

	// 首先判断用户类型 为0 是管理员 查所有
	if userType == "0" {
		// 查总数
		if err := utils.DB.Find(&commentinfo).Count(&count).Error; err != nil {
			c.JSON(http.StatusOK,gin.H{
				"code": 1001,
				"msg" : "查总数错误",
			})
			return
		}
		// 分页查数据
		if err := utils.DB.Offset(int(offset)).Limit(int(pageSize)).Find(&commentinfo).Error; err != nil {
			c.JSON(http.StatusOK,gin.H{
				"code": 1001,
				"msg" : "分页查数据错误",
			})
			return
		}
		var  totalPage int
		if count % pageSize == 0 {
			totalPage = int(count / pageSize)
		}else {
			totalPage = int(count / pageSize) + 1
		}
		c.JSON(http.StatusOK,gin.H{
			"code": 0,
			"currentPage" :currentPage,
			"pageSize": pageSize,
			"total": count,
			"totalPage": totalPage,
			"data": commentinfo,
		})
	}else {
		// 普通用户查userID对应的文章
		// 查总数
		if err := utils.DB.Where("userId = ?", userId).Find(&commentinfo).Count(&count).Error; err != nil {
			c.JSON(http.StatusOK,gin.H{
				"code": 1001,
				"msg" : "查总数错误",
			})
			return
		}
		// 分页查数据
		if err := utils.DB.Offset(int(offset)).Limit(int(pageSize)).Where("userId = ?",userId).Find(&commentinfo).Error; err != nil {
			c.JSON(http.StatusOK,gin.H{
				"code": 1001,
				"msg" : "分页查数据错误",
			})
			return
		}
		var  totalPage int
		if count % pageSize == 0 {
			totalPage = int(count / pageSize)
		}else {
			totalPage = int(count / pageSize) + 1
		}
		c.JSON(http.StatusOK,gin.H{
			"code": 0,
			"currentPage" :currentPage,
			"pageSize": pageSize,
			"total": count,
			"totalPage": totalPage,
			"data": commentinfo,
		})
	}
}