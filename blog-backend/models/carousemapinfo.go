package models

import (
	"GIN/utils"
)

// updata使用
type MapInfo struct {
	Name string			`gorm:"column:name" json:"name"`
	Url string `gorm:"column:url" json:"url"`
}

type MapInfoList struct {
	Name string			`gorm:"column:name" json:"name"`
	Url string `gorm:"column:url" json:"url"`
	Id int `gorm:"column:id" json:"id"`
}

func (MapInfo) TableName() string  {
	return "carousemapinfo"
}

func (MapInfoList) TableName() string  {
	return "carousemapinfo"
}

//DeleteAll
func (MapInfoList) DeleteAll(mapinfos []MapInfo) int  {
	if err :=  utils.DB.Where("1 = 1").Delete(&MapInfoList{}).Error; err != nil {
		return 1
	}else {
		return 0
	}
}

//AddMap 单个新增
func (MapInfo) AddMap(mapinfo MapInfo) int  {
	if err :=  utils.DB.Model(&MapInfo{}).Create(&mapinfo).Error; err != nil {
		return 1
	}else {
		return 0
	}
}

/**
 * 返回轮播图
 * @author lizibin
 */
func (MapInfoList) SearchCarouseMap() interface{} {
	mapinfos := []MapInfoList{}
	if err :=  utils.DB.Model(&MapInfoList{}).Find(&mapinfos).Error; err != nil {
		return 1
	}else {
		return mapinfos
	}
}
