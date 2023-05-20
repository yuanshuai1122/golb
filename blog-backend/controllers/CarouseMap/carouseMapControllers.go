package CarouseMap

import (
	"GIN/models"
	"GIN/servise"
	"GIN/utils"
	"github.com/gin-gonic/gin"
)

type CarouseMapController struct{}

/**
 * 更新轮播图
 * @author lizibin
 */
func (con CarouseMapController) UpdateCarouseMap(c *gin.Context){
	var mapinfos []models.MapInfo
	c.ShouldBind(&mapinfos)
	result := servise.MapServise{}.UpdateCarouseMap(mapinfos)
	if result == 0 {
		utils.Success(c)
	}else if result == 1 {
		utils.Failed(c,"修改失败!")
	}
}

// SearchCarouseMap
func (con CarouseMapController) SearchCarouseMap(c *gin.Context){
	result := servise.MapServise{}.SearchCarouseMap()
	if result == 1 {
		utils.Failed(c,"查询失败!")
	}else {
		utils.SuccessData(c,result)
	}
}