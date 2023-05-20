package servise

import (
	"GIN/models"
)

type MapServise struct{}

/**
 *  根据用户id删除用户    同时删除用户的所有评论以及文章
 * @author lizibin
 */
func (MapServise) UpdateCarouseMap(mapinfos []models.MapInfo) int {
	var Flag int
	// 删除所有
	deleteMapAll := models.MapInfoList{}.DeleteAll(mapinfos)
	if deleteMapAll == 1 {
		Flag = 1
	}
	// 新增多个
	for _, v := range mapinfos {
		addResult := models.MapInfo{}.AddMap(v)
		if addResult == 1 {
			Flag = 1
		}
	}
	if Flag == 0 {
		return 0
	} else {
		return 1
	}
}

//SearchCarouseMap
func (MapServise) SearchCarouseMap() interface{} {
	result := models.MapInfoList{}.SearchCarouseMap()
	if result == 1 {
		return 1
	} else {
		return result
	}
}