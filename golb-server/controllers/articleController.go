package controllers

import (
	"github.com/gin-gonic/gin"
	"golb/golb-server/services"
	"golb/golb-server/utils"
	"strconv"
)

type ArticleController struct{}

// GetArticlesListPage 分页查询文章列表
func (con ArticleController) GetArticlesListPage(c *gin.Context) {
	// 获取当前页数
	pageNum, _ := strconv.ParseInt(c.Query("pageNum"), 10, 64)
	// 页大小
	pageSize, _ := strconv.ParseInt(c.Query("pageSize"), 10, 64)
	// 关键字
	keywords := c.Query("keywords")
	// 查询文章列表
	result := services.ArticleService{}.GetArticlesList(pageNum, pageSize, keywords)
	if result == nil {
		utils.SearchPageFaild(c)
	}
	utils.SuccessData(c, result)
}
