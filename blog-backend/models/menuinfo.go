package models

import (
	"GIN/utils"
)

// 返回列表使用
type MenuInfo struct {
	Id           int        `gorm:"column:id" json:"id"`
	Name         string     `gorm:"column:name" json:"name"`
	ChineseName  string     `gorm:"column:chineseName" json:"chineseName"`
	Title        string     `gorm:"column:title" json:"title"`
	Path         string     `gorm:"column:path" json:"path"`
	Icon         string     `gorm:"column:icon" json:"icon"`
	ParentMenuId int        `gorm:"column:parentMenuId" json:"parentMenuId"`
	MenuStatus   int        `gorm:"column:menuStatus" json:"menuStatus"`
	Children     []MenuInfo `gorm:"-" json:"childMenu"`
}

func (MenuInfo) TableName() string {
	return "backstagemenuinfo"
}

// 根据用户类型查菜单 有多级
func (MenuInfo) SearchMenuByUserID(userType int) interface{} {
	menuInfo := []MenuInfo{}
	// 类型是0 管理员查全部
	if userType == 0 {
		if err := utils.DB.Where("parentMenuId != ?", 2).Find(&menuInfo).Error; err != nil {
			return 1
		} else {
			return menuInfo
		}
	} else {
		if err := utils.DB.Where("menuStatus = ?  AND parentMenuId != ?", 1, 2).Find(&menuInfo).Error; err != nil {
			return 1
		} else {
			return menuInfo
		}
	}

}

// 通过parentID 去查子级
func (MenuInfo) SearchMenuByParentID(userType int) interface{} {
	menuInfo := []MenuInfo{}
	sql := utils.DB.Where("parentMenuId = ? AND menuStatus = ?", 2, 1).Find(&menuInfo)
	if userType == 0 {
		sql = utils.DB.Where("parentMenuId = ?", 2).Find(&menuInfo)
	}
	if err := sql.Error; err != nil {
		return 1
	} else {
		return menuInfo
	}
}
