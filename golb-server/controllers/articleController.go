package controllers

import (
	"github.com/gin-gonic/gin"
	"golb/golb-server/services"
	"golb/golb-server/utils"
)

type ArticleController struct{}

func (con ArticleController) GetArticlesList(c *gin.Context) {
	result := services.ArticleService{}.GetArticlesList()
	if result == nil {
		utils.SearchPageFaild(c)
	}
	utils.SuccessData(c, result)
}
