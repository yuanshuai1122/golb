package models

import (
	"GIN/utils"
)

// 返回列表使用
type TextInfo struct {
	Id   int    `gorm:"column:id" json:"id"`
	Text string `gorm:"column:text" json:"text"`
}

func (TextInfo) TableName() string {
	return "circulatetext"
}

// 根据ID查询是否存在
func (TextInfo) SearchUrlByID(id string) int {
	if err := utils.DB.Where("id = ?", id).First(&TextInfo{}).Error; err != nil {
		return 1
	} else {
		return 0
	}
}

// 根据ID删除
func (TextInfo) DeleteUrlByID(id string) int {
	if err := utils.DB.Where("id = ?", id).Delete(&TextInfo{}).Error; err != nil {
		return 1
	} else {
		return 0
	}
}

/**
* 模块名: DeleteAll
* 代码描述: 批量删除
* 作者:lizibin
* 创建时间:2023/01/02 15:14:56
 */
func (TextInfo) DeleteAll(id string) int {
	if err := utils.DB.Where("id = ?", id).Delete(&TextInfo{}).Error; err != nil {
		return 1
	} else {
		return 0
	}
}

// 新增
func (TextInfo) Insert(json map[string]interface{}) int {
	if err := utils.DB.Model(&TextInfo{}).Create(json).Error; err != nil {
		return 1
	} else {
		return 0
	}
}

// 分页查询留言
func (TextInfo) AuditScreening(currentPage int64, pageSize int64) interface{} {
	offset := (currentPage - 1) * pageSize
	var count int64
	textInfo := []TextInfo{}
	// 查总数
	if err := utils.DB.Find(&textInfo).Count(&count).Error; err != nil {
		return 1
	}
	if err := utils.DB.Offset(int(offset)).Limit(int(pageSize)).Find(&textInfo).Error; err != nil {
		return 1
	} else {
		totalPage := utils.ReturnTotalPage(count, pageSize)
		return map[string]interface{}{
			"currentPage": currentPage,
			"list":        textInfo,
			"pageSize":    pageSize,
			"total":       count,
			"totalPage":   totalPage,
		}
	}
}
