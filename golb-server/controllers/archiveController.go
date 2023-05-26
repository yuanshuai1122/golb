package controllers

import (
	"github.com/gin-gonic/gin"
	"golb/golb-server/services"
	"golb/golb-server/utils"
	"strconv"
)

type ArchiveController struct{}

// GetArchivesListPage 查询归档列表
func (con ArchiveController) GetArchivesListPage(c *gin.Context) {
	// 获取当前页数
	pageNum, _ := strconv.ParseInt(c.Query("pageNum"), 10, 64)
	// 页大小
	pageSize, _ := strconv.ParseInt(c.Query("pageSize"), 10, 64)
	// 查询归档列表
	result := services.ArchiveService{}.GetArchivesList(pageNum, pageSize)
	if result == nil {
		utils.SearchPageFaild(c)
	}
	utils.SuccessData(c, result)
}
