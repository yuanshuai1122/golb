package models

import (
	"GIN/utils"
)

// 返回列表使用
type UrlInfo struct {
	UrlId          int    `gorm:"column:urlId" json:"urlId"`
	UrlName        string `gorm:"column:urlName" json:"urlName"`
	UrlAddres      string `gorm:"column:urlAddres" json:"urlAddres"`
	UrlReferral    string `gorm:"column:urlReferral" json:"urlReferral"`
	UrlLitimg      string `gorm:"column:urlLitimg" json:"urlLitimg"`
	WebmasterEmail string `gorm:"column:webmasterEmail" json:"webmasterEmail"`
	UrlPass        int    `gorm:"column:urlPass" json:"urlPass"`
	UrlType        int    `gorm:"column:urlType" json:"urlType"`
}

func (UrlInfo) TableName() string {
	return "urlinfo"
}

// 根据ID查询是否存在
func (UrlInfo) SearchUrlByID(id string) int {
	if err := utils.DB.Where("urlId = ?", id).First(&UrlInfo{}).Error; err != nil {
		return 1
	} else {
		return 0
	}
}

// 根据ID删除
func (UrlInfo) DeleteUrlByID(id string) int {
	if err := utils.DB.Where("urlId = ?", id).Delete(&UrlInfo{}).Error; err != nil {
		return 1
	} else {
		return 0
	}
}

// 新增
func (UrlInfo) Insert(json map[string]interface{}) int {
	if err := utils.DB.Model(&UrlInfo{}).Create(json).Error; err != nil {
		return 1
	} else {
		return 0
	}
}

// 分页查询留言
func (UrlInfo) AuditScreening(currentPage int64, pageSize int64, urlPass string) interface{} {
	offset := (currentPage - 1) * pageSize
	var count int64
	urlInfo := []UrlInfo{}
	// 查总数
	if err := utils.DB.Where("urlPass = ?", urlPass).Find(&urlInfo).Count(&count).Error; err != nil {
		return 1
	}
	if err := utils.DB.Offset(int(offset)).Limit(int(pageSize)).Where("urlPass = ?", urlPass).Find(&urlInfo).Error; err != nil {
		return 1
	} else {
		totalPage := utils.ReturnTotalPage(count, pageSize)
		return map[string]interface{}{
			"currentPage": currentPage,
			"list":        urlInfo,
			"pageSize":    pageSize,
			"total":       count,
			"totalPage":   totalPage,
		}
	}
}

// 分页查询
func (UrlInfo) Findby(currentPage int64, pageSize int64, searchContent string) interface{} {
	offset := (currentPage - 1) * pageSize
	var count int64
	urlInfo := []UrlInfo{}
	// 查总数
	if err := utils.DB.Where("urlName LIKE ?", searchContent+"%").Find(&urlInfo).Count(&count).Error; err != nil {
		return 1
	}
	if err := utils.DB.Offset(int(offset)).Limit(int(pageSize)).Where("urlName LIKE ?", searchContent+"%").Find(&urlInfo).Error; err != nil {
		return 1
	} else {
		totalPage := utils.ReturnTotalPage(count, pageSize)
		return map[string]interface{}{
			"currentPage": currentPage,
			"list":        urlInfo,
			"pageSize":    pageSize,
			"total":       count,
			"totalPage":   totalPage,
		}
	}
}

// 查全部
func (UrlInfo) ShowAllUrlInfo(urlType string) interface{} {
	urlInfo := []UrlInfo{}
	if err := utils.DB.Where("urlType = ?", urlType).Find(&urlInfo).Error; err != nil {
		return 1
	} else {
		return map[string]interface{}{
			"list": urlInfo,
		}
	}
}

// 更新状态
func (UrlInfo) UpdateUrlPass(urlId string, urlPass string) int {
	if err := utils.DB.Model(&UrlInfo{}).Where("urlId = ?", urlId).Update("urlPass", urlPass).Error; err != nil {
		return 1
	} else {
		return 0
	}
}
