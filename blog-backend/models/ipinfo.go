package models

import (
	"GIN/utils"
)

// 返回列表使用
type IpInfo struct {
	Id   int    `gorm:"column:id" json:"id"`
	Ip   string `gorm:"column:ip" json:"ip"`
	City string `gorm:"column:city" json:"city"`
}

func (IpInfo) TableName() string {
	return "ipinfo"
}

// 根据ID查询是否存在
func (IpInfo) SearchUrlByID(id string) int {
	if err := utils.DB.Where("id = ?", id).First(&IpInfo{}).Error; err != nil {
		return 1
	} else {
		return 0
	}
}

// 根据ID删除
func (IpInfo) DeleteUrlByID(id string) int {
	if err := utils.DB.Where("id = ?", id).Delete(&IpInfo{}).Error; err != nil {
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
func (IpInfo) DeleteAll(id string) int {
	if err := utils.DB.Where("id = ?", id).Delete(&IpInfo{}).Error; err != nil {
		return 1
	} else {
		return 0
	}
}

// 新增
func (IpInfo) Insert(json map[string]interface{}) int {
	if err := utils.DB.Model(&IpInfo{}).Create(json).Error; err != nil {
		return 1
	} else {
		return 0
	}
}

// 分页查询留言
func (IpInfo) GetPage(currentPage int64, pageSize int64) interface{} {
	offset := (currentPage - 1) * pageSize
	var count int64
	ipInfo := []IpInfo{}
	// 查总数
	if err := utils.DB.Find(&ipInfo).Count(&count).Error; err != nil {
		return 1
	}
	if err := utils.DB.Offset(int(offset)).Limit(int(pageSize)).Find(&ipInfo).Error; err != nil {
		return 1
	} else {
		totalPage := utils.ReturnTotalPage(count, pageSize)
		return map[string]interface{}{
			"currentPage": currentPage,
			"list":        ipInfo,
			"pageSize":    pageSize,
			"total":       count,
			"totalPage":   totalPage,
		}
	}
}
