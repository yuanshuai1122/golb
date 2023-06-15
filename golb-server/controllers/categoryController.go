package controllers

import (
	"github.com/gin-gonic/gin"
	"golb/golb-server/services"
	"golb/golb-server/utils"
)

type CategoryController struct{}

// GetCategoryList 获取分类列表
func (con CategoryController) GetCategoryList(c *gin.Context) {
	// 返回分类列表
	utils.SuccessData(c, services.CategoryService{}.GetCategoryList())
}
