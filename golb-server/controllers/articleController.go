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

// GetArticleDetail 根据文章id获取文章详情
func (con ArticleController) GetArticleDetail(c *gin.Context) {
	// 获取传入的id
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		utils.MissId(c)
	}

	result := services.ArticleService{}.GetArticleDetail(id)
	if result == nil {
		utils.Failed(c, "文章不存在")
	}
	utils.SuccessData(c, result)

}
