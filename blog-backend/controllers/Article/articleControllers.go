package Article

import (
	"GIN/models"
	"GIN/servise"
	"GIN/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ArticleController struct{}

/**
 * 查询所有文章 不分页
 * @author lizibin
 */
func (con ArticleController) ShowAll(c *gin.Context){
	result := servise.ArticleServise{}.ShowAll()
	if result == 1 {
		utils.SearchPageFaild(c)
	}else {
		utils.SuccessData(c,result)
	}
	//articleinfo := []models.ArticleInfo{}
	//if err := utils.DB.Where("articleState = ?",1).Find(&articleinfo).Error ;err != nil {
	//	c.JSON(http.StatusOK,gin.H{
	//		"code": 0,
	//		"msg": "查询出错",
	//	})
	//}else {
	//	c.JSON(http.StatusOK,gin.H{
	//		"code": 0,
	//		"data": articleinfo,
	//	})
	//}
}

/**
 * 根据文章id获取文章信息
 * @author lizibin
 */
func (con ArticleController) ShowArticleInfo(c *gin.Context){
	articleinfo := models.ArticleInfo{}
	id := c.Query("articleId")
	if err := utils.DB.First(&articleinfo,"articleId = ?",id).Error; err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code": 0,
			"msg": "查询出错",
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code": 0,
			"data": articleinfo,
		})
	}
}

/**
 * 下拉框筛选文章审核状态，查询出相关内容并进行分页
 * @author lizibin
 */
func (con ArticleController) SearchPass(c *gin.Context){
	articleinfo := []models.ArticleInfo{}
	// 获取参数
	currentPage,_ := strconv.ParseInt(c.Query("currentPage"),10,64) //当前的页码
	pageSize,_ :=  strconv.ParseInt(c.Query("pageSize"), 10, 64)//每页显示的条数
	userType := c.Query("userType")// 用户类型
	userId := c.Query("userId")// 用户id
	articlePass := c.Query("articlePass")// 审核状态  1 待审核 2 已通过 3 已拒绝
	offset := (currentPage - 1) * pageSize
	var count int64

	// 首先判断用户类型 为0 是管理员 查所有
	if userType == "0" {
		// articlePass为空返回全部
		if articlePass == "" {
			// 查总数
			if err := utils.DB.Find(&articleinfo).Count(&count).Error; err != nil {
				c.JSON(http.StatusOK,gin.H{
					"code": 1001,
					"msg" : "查总数错误",
				})
				return
			}
			// 分页查数据
			if err := utils.DB.Offset(int(offset)).Limit(int(pageSize)).Find(&articleinfo).Error; err != nil {
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
				"data": articleinfo,
			})
		}else {
			// articlePass不为空 返回对应类型的数据
			// 拿到总数
			if err := utils.DB.Offset(int(offset)).Limit(int(pageSize)).Where("articlePass = ? " , articlePass).Find(&articleinfo).Count(&count).Error; err != nil {
				c.JSON(http.StatusOK,gin.H{
					"code": 1001,
					"msg" : "查总数错误",
				})
				return
			}

			// 分页查数据
			if err := utils.DB.Offset(int(offset)).Limit(int(pageSize)).Where("articlePass = ?",articlePass).Find(&articleinfo).Error; err != nil {
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
				"data": articleinfo,
			})
		}
	}else {
		// 普通用户查userID对应的文章
		// articlePass为空返回全部
		if articlePass == "" {
			// 查总数
			if err := utils.DB.Where("userId = ?", userId).Find(&articleinfo).Count(&count).Error; err != nil {
				c.JSON(http.StatusOK,gin.H{
					"code": 1001,
					"msg" : "查总数错误",
				})
				return
			}
			// 分页查数据
			if err := utils.DB.Offset(int(offset)).Limit(int(pageSize)).Where("userId = ?",userId).Find(&articleinfo).Error; err != nil {
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
				"data": articleinfo,
			})
		}else {
			// articlePass不为空 返回对应类型的数据
			// 拿到总数
			if err := utils.DB.Offset(int(offset)).Limit(int(pageSize)).Where("articlePass = ? AND userId = ? " , articlePass,userId).Find(&articleinfo).Count(&count).Error; err != nil {
				c.JSON(http.StatusOK,gin.H{
					"code": 1001,
					"msg" : "查总数错误",
				})
				return
			}

			// 分页查数据
			if err := utils.DB.Offset(int(offset)).Limit(int(pageSize)).Where("articlePass = ? AND userId = ?",articlePass,userId).Find(&articleinfo).Error; err != nil {
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
				"data": articleinfo,
			})
		}
	}
}

/**
 * 查询所有通过审核并且所有可见的文章  筛选条件  分类id
 * @author lizibin
 */
func (con ArticleController) SearchByClassifyId(c *gin.Context){
	articleinfo := []models.ArticleInfo{}
	// 获取参数
	currentPage,_ := strconv.ParseInt(c.Query("currentPage"),10,64) //当前的页码
	pageSize,_ :=  strconv.ParseInt(c.Query("pageSize"), 10, 64)//每页显示的条数
	classifyId := c.Query("classifyId")

	offset := (currentPage - 1) * pageSize
	var count int64

	if err := utils.DB.Where("articleClassifyId = ? AND articleState = ? AND articlePass = ?",classifyId,1,2 ).Find(&articleinfo).Count(&count).Error; err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code": 1001,
			"msg" : "查总数错误",
		})
		return
	}
	// 分页查数据
	if err := utils.DB.Offset(int(offset)).Limit(int(pageSize)).Where("articleClassifyId = ? AND articleState = ? AND articlePass = ?",classifyId,1,2 ).Find(&articleinfo).Error; err != nil {
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
		"data": articleinfo,
	})
}

/**
 * 根据文章标题进行搜索 并且 进行分页  判断是管理员还是普通用户  和另外两个条件查询重复了 原作者java写的太乱了 我又懒得去适配前端
 * @author lizibin
 */
func (con ArticleController) SearchbyUserType(c *gin.Context){
	articleinfo := []models.ArticleInfo{}
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
			if err := utils.DB.Find(&articleinfo).Count(&count).Error; err != nil {
				c.JSON(http.StatusOK,gin.H{
					"code": 1001,
					"msg" : "查总数错误",
				})
				return
			}
			// 分页查数据
			if err := utils.DB.Offset(int(offset)).Limit(int(pageSize)).Find(&articleinfo).Error; err != nil {
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
				"data": articleinfo,
			})
	}else {
		// 普通用户查userID对应的文章
			// 查总数
			if err := utils.DB.Where("userId = ?", userId).Find(&articleinfo).Count(&count).Error; err != nil {
				c.JSON(http.StatusOK,gin.H{
					"code": 1001,
					"msg" : "查总数错误",
				})
				return
			}
			// 分页查数据
			if err := utils.DB.Offset(int(offset)).Limit(int(pageSize)).Where("userId = ?",userId).Find(&articleinfo).Error; err != nil {
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
				"data": articleinfo,
			})
	}
}

/**
 * 根据文章标题进行搜索 并且 进行分页  判断是管理员还是普通用户
 * @author lizibin
 */
func (con ArticleController) SearchTitle(c *gin.Context){
	articleinfo := []models.ArticleInfo{}
	// 获取参数
	currentPage,_ := strconv.ParseInt(c.Query("currentPage"),10,64) //当前的页码
	pageSize,_ :=  strconv.ParseInt(c.Query("pageSize"), 10, 64)//每页显示的条数
	userType := c.Query("userType")// 用户类型
	userId := c.Query("userId")// 用户id
	articleTitle := c.Query("articleTitle")// 审核状态  1 待审核 2 已通过 3 已拒绝
	offset := (currentPage - 1) * pageSize
	var count int64

	// 首先判断用户类型 为0 是管理员 查所有
	if userType == "0" {
		// articlePass为空返回全部
		if articleTitle == "" {
			// 查总数
			if err := utils.DB.Find(&articleinfo).Count(&count).Error; err != nil {
				c.JSON(http.StatusOK,gin.H{
					"code": 1001,
					"msg" : "查总数错误",
				})
				return
			}
			// 分页查数据
			if err := utils.DB.Offset(int(offset)).Limit(int(pageSize)).Find(&articleinfo).Error; err != nil {
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
				"data": articleinfo,
			})
		}else {
			// articlePass不为空 返回对应类型的数据
			// 拿到总数
			if err := utils.DB.Offset(int(offset)).Limit(int(pageSize)).Where("articleTitle LIKE ? " , articleTitle+"%").Find(&articleinfo).Count(&count).Error; err != nil {
				c.JSON(http.StatusOK,gin.H{
					"code": 1001,
					"msg" : "查总数错误",
				})
				return
			}

			// 分页查数据
			if err := utils.DB.Offset(int(offset)).Limit(int(pageSize)).Where("articleTitle  LIKE ?",articleTitle+"%").Find(&articleinfo).Error; err != nil {
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
				"data": articleinfo,
			})
		}
	}else {
		// 普通用户查userID对应的文章
		// articlePass为空返回全部
		if articleTitle == "" {
			// 查总数
			if err := utils.DB.Where("userId = ?", userId).Find(&articleinfo).Count(&count).Error; err != nil {
				c.JSON(http.StatusOK,gin.H{
					"code": 1001,
					"msg" : "查总数错误",
				})
				return
			}
			// 分页查数据
			if err := utils.DB.Offset(int(offset)).Limit(int(pageSize)).Where("userId = ?",userId).Find(&articleinfo).Error; err != nil {
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
				"data": articleinfo,
			})
		}else {
			// articlePass不为空 返回对应类型的数据
			// 拿到总数
			if err := utils.DB.Offset(int(offset)).Limit(int(pageSize)).Where("articleTitle LIKE ? AND userId = ? " , articleTitle+"%",userId).Find(&articleinfo).Count(&count).Error; err != nil {
				c.JSON(http.StatusOK,gin.H{
					"code": 1001,
					"msg" : "查总数错误",
				})
				return
			}

			// 分页查数据
			if err := utils.DB.Offset(int(offset)).Limit(int(pageSize)).Where("articleTitle LIKE ? AND userId = ?",articleTitle+"%",userId).Find(&articleinfo).Error; err != nil {
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
				"data": articleinfo,
			})
		}
	}
}
/**
 * 通过审核的数量
 * @author lizibin
 */
func (con ArticleController) ShowArticleCount(c *gin.Context){
	// 分层重写
	result := servise.ArticleServise{}.ShowArticleCount()
	if result == 1 {
		utils.SearchPageFaild(c)
	}else {
		utils.SuccessData(c,result)
	}
}


/**
 * 删除文章
 * @author lizibin
 */
func (con ArticleController) Delete(c *gin.Context){
	//articleinfo := models.ArticleInfo{}
	articleId := c.Query("articleId")
	result := servise.ArticleServise{}.Delete(articleId)
	if result == 0 {
		utils.Success(c)
	}else if result == 1 {
		utils.DeleteFailed(c)
	}else {
		utils.Failed(c,"存在首页设置，请取消后删除!")
	}
}

/**
 * 新增文章
 * @author lizibin
 */
func (con ArticleController) Add(c *gin.Context){
	articleinfo := models.AddArticleInfo{}
	err := c.ShouldBindJSON(&articleinfo)
	if err != nil {
		fmt.Println(err)
	}else {
		utils.DB.Create(&articleinfo)
		c.JSON(200,gin.H{
			"code": 0,
			"data": articleinfo,
			"msg": "新增成功",
		})
	}
}

/**
 * 修改文章
 * @author lizibin
 */
func (con ArticleController) Edit(c *gin.Context){
	editarticleinfo := models.EditArticleInfo{}
	err := c.ShouldBindJSON(&editarticleinfo)
	if err != nil {
		fmt.Println(err)
	}else {
		if err := utils.DB.Model(models.AddArticleInfo{}).Where("articleId = ?",editarticleinfo.ArticleId).Updates(&editarticleinfo).Error; err != nil {
			fmt.Println(err)
			c.JSON(200,gin.H{
				"code": 0,
				"data": "",
				"msg": "修改失败",
			})
		}else {
			c.JSON(200,gin.H{
				"code": 0,
				"data": editarticleinfo,
				"msg": "修改成功",
			})
		}
	}
}

/**
 * 根据用户id查询通过审核文章总数
 * @author lizibin
 */
func (con ArticleController) ShowArticleCountByUserId(c *gin.Context){
	articleinfo := []models.ArticleInfo{}
	var count int64
	id := c.Query("userId")
	if id == "" {
		fmt.Println("id为空")
		c.JSON(http.StatusOK,gin.H{
					"code": 1001,
					"data": "ID为空",
		})
		return
	}
	if err := utils.DB.Where("userId = ? AND articlePass = ?",id,2).Find(&articleinfo).Count(&count).Error; err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK,gin.H{
			"code": 1001,
			"data": "查询错误",
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"code": 0,
			"data": count,
		})
	}
}


/**
 * 修改文章点击量，每点击一次+1
 * @author lizibin
 */
func (con ArticleController) UpdateArticleClick(c *gin.Context){
	articleinfo := models.ArticleInfo{}
	id := c.Query("articleId")
	if err := utils.DB.First(&articleinfo,"articleId = ?",id).Update("click",articleinfo.Click + 1).Error; err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code": 0,
			"msg": "修改出错",
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code": 0,
			"msg": "OK",
		})
	}
}
/**
 * 修改文章点击量，每点击一次+1
 * @author lizibin
 */
func (con ArticleController) UpdateArticlePass(c *gin.Context){
	articleinfo := models.ArticleInfo{}
	articleId := c.Query("articleId")
	articlePass := c.Query("articlePass")
	if err := utils.DB.First(&articleinfo,"articleId = ?",articleId).Update("articlePass",articlePass).Error; err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code": 0,
			"msg": "修改出错",
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code": 0,
			"msg": "OK",
		})
	}
}


// --------------------------------------------
// ShowAllPass 查询所有审核通过的文章分页
func (con ArticleController) ShowAllPass(c *gin.Context){
	currentPage,_ := strconv.ParseInt(c.Query("currentPage"),10,64) //当前的页码
	pageSize,_ :=  strconv.ParseInt(c.Query("pageSize"), 10, 64)//每页显示的条数
	keyword := c.Query("keyword")
	result := servise.ArticleServise{}.ShowAllPass(currentPage,pageSize,keyword)
	if result == 1 {
		utils.SearchPageFaild(c)
	}else {
		utils.SuccessData(c,result)
	}
}